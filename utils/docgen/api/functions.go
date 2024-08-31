package docgen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/lmorg/murex/utils/envvars"
)

var funcMap = template.FuncMap{
	"md":           funcMarkdown,
	"quote":        funcQuote,
	"trim":         strings.TrimSpace,
	"doc":          funcRenderedDocuments,
	"cat":          funcRenderedCategories,
	"link":         funcLink,
	"bookmark":     funcLinkBookmark,
	"section":      funcSection,
	"file":         funcFile,
	"notanindex":   funcNotAnIndex,
	"date":         funcDate,
	"time":         funcTime,
	"doct":         funcDocT,
	"othercats":    funcOtherCats,
	"otherdocs":    funcOtherDocs,
	"env":          funcEnv,
	"fn":           funcFunctions,
	"vuepressmenu": funcVuePressMenu,
	"dump":         funcDump,
}

var funcMap__fn = template.FuncMap{}

func init() {
	for k, v := range funcMap {
		funcMap__fn[k] = v
	}
	delete(funcMap__fn, "fn")
}

/************
 * Markdown *
 ************/

// Takes: string (contents as read from YAML in a machine readable subset of markdown)
// Returns: markdown contents cleaned up for printing
func funcMarkdown(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.TrimSuffix(s, "\n")
	return s
}

/************
 *   Quote  *
 ************/

// Takes: strings (contents)
// Returns: contents with some characters escaped for printing in source code (eg \")
func funcQuote(s ...string) string {
	return strconv.Quote(funcMarkdown(strings.Join(s, "")))
}

/************
 *    Doc   *
 ************/

// Takes: string (category name), string (document name), int (index of rendered document template)
// Returns: contents of rendered document template
func funcRenderedDocuments(cat, doc string, index int) (string, error) {
	if len(Config.renderedDocuments[cat]) == 0 {
		return "", errors.New("category does not exist, doesn't have any documents, or hasn't yet been parsed")
	}

	if len(Config.renderedDocuments[cat][doc]) <= index {
		return "", errors.New("index greater than number of templates currently parsed for that document")
	}

	return Config.renderedDocuments[cat][doc][index], nil
}

/************
 *    Cat   *
 ************/

// Takes: string (category name) and int (index of rendered category template)
// Returns: contents of rendered category template
func funcRenderedCategories(cat string, index int) (string, error) {
	if len(Config.renderedCategories[cat]) == 0 {
		return "", errors.New("category does not exist or hasn't yet been parsed")
	}

	if len(Config.renderedCategories[cat]) <= index {
		return "", errors.New("index greater than number of templates currently parsed for that category")
	}

	return Config.renderedCategories[cat][index], nil
}

/************
 *   Link   *
 ************/

// Takes: string (description, doc, relativePath prefix...)
// Returns: URL to document
func funcLink(description, docId string, relativePath ...string) string {
	return funcLinkBookmark(description, docId, "", relativePath...)
}

/************
 *  LinkBM  *
 ************/

// Takes: string (description, doc, bookmark (w/o hash), relativePath prefix...)
// Returns: URL to document
func funcLinkBookmark(description, docId, bookmark string, relativePath ...string) string {
	doc := Documents.ByID("", "???", docId)
	if doc == nil {
		panic(fmt.Sprintf("nil document (%s)", docId))
	}

	if doc.CategoryID == "" {

		return description
	}

	cat := Config.Categories.ByID(doc.CategoryID)
	path := cat.Templates[0].DocumentFilePath(doc)

	if bookmark != "" {
		path += "#" + bookmark
	}

	path = strings.Join(relativePath, "/") + "/" + path

	if os.Getenv("DOCGEN_TARGET") == "vuepress" {
		path = strings.Replace(path, "/docs", "", 1)
	}

	return fmt.Sprintf("[%s](%s)", description, path)
}

/************
 * Section  *
 ************/

// Takes: string (description, cat, bookmark (w/o hash), relativePath prefix...)
// Returns: URL to document
func funcSection(description, catId, bookmark string, relativePath ...string) string {
	cat := Config.Categories.ByID(catId)
	if cat == nil {
		panic(fmt.Sprintf("nil category (%s)", catId))
	}

	if bookmark != "" {
		bookmark = "#" + bookmark
	}

	path := fmt.Sprintf("%s/%s%s%s",
		strings.Join(relativePath, "/"),
		cat.Templates[0].OutputPath,
		cat.Templates[0].CategoryFile,
		bookmark)

	if os.Getenv("DOCGEN_TARGET") == "vuepress" {
		path = strings.Replace(path, "/docs", "", 1)
	}

	return fmt.Sprintf("[%s](%s)", description, path)
}

/************
 *  Include *
 ************/

// Takes: string (contents)
// Looks for: {{ include "filename" }}
// Returns: document with search string replaced with contents of filename
func funcInclude(s string) string {
	rx := regexp.MustCompile(`\{\{ include "([-_/.a-zA-Z0-9]+)" \}\}`)
	if !rx.MatchString(s) {
		return s
	}

	match := rx.FindAllStringSubmatch(s, -1)
	for i := range match {
		path := match[i][1]
		f := fileReader(path)
		b := bytes.TrimSpace(readAll(f))

		w := bytes.NewBuffer([]byte{})
		t := template.Must(template.New(path).Funcs(funcMap).Parse(string(b)))
		if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}

		s = strings.Replace(s, match[i][0], w.String(), -1)
	}

	return s
}

/************
 *   File   *
 ************/

// Takes: slice of strings (file path)
// Returns: contents of file based on a concatenation of the slice
func funcFile(path ...string) string {
	f := fileReader(strings.Join(path, ""))
	b := readAll(f)
	return string(b)
}

func init() {
	funcMap["include"] = funcInclude
}

/************
 *NotAnIndex*
 ************/

// Takes: number (eg variable)
// Returns: number incremented by 1
func funcNotAnIndex(i int) int {
	return i + 1
}

/************
 *   Date   *
 ************/

// Takes: DateTime (time.Time)
// Returns: Date as string
func funcDate(dt time.Time) string {
	return dt.Format("02.01.2006")
}

/************
 *   Time   *
 ************/

// Takes: DateTime (time.Time)
// Returns: Time as string
func funcTime(dt time.Time) string {
	return dt.Format("15:04")
}

/************
 *   doct   *
 ************/

// Takes: string (category, document ID)
// Returns: document type
func funcDocT(cat, doc string) *document {
	return Documents.ByID("!!!", cat, doc)
}

/************
 * OtherCats*
 ************/

// Returns: list of documents in that category
func funcOtherCats() []category {
	return Config.Categories
}

/************
 * OtherDocs*
 ************/

// Takes: string (category ID)
// Returns: list of documents in that category
func funcOtherDocs(id string) (d documents) {
	for i := range Documents {
		if Documents[i].CategoryID == id {
			d = append(d, Documents[i])
		}
	}

	sort.Sort(d)
	return
}

/************
 *    Env   *
 ************/

// Takes: string (env, formatted: `key=val`)
// Returns: true or false if the env matches systems env
func funcEnv(env string) any {
	if !strings.Contains(env, "=") {
		s := os.Getenv(env)
		if s == "" {
			s = "undef"
		}
		return s
	}

	key, value := envvars.Split(env)
	v := make(map[string]interface{})
	envvars.All(v)
	s, _ := v[key].(string)
	return s == value
}

/************
 *    fn    *
 ************/

// Takes: string, to use as template
// Returns: parsed string
func funcFunctions(s string) string {
	w := bytes.NewBuffer([]byte{})
	t, err := template.New("__fn").Funcs(funcMap__fn).Parse(s)
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
	return w.String()
}

/************
 * VuePress *
 ************/

// Takes: string, category ID
// Returns: JSON string
func funcVuePressMenu(catID string) string {
	//var menu vuePressMenuItem
	cat := Config.Categories.ByID(catID)
	if cat == nil {
		panic(fmt.Sprintf("cannot find category with ID '%s'", catID))
	}

	menu := vuePressSubMenu(cat)

	for i := range Documents {
		if Documents[i].CategoryID != cat.ID || len(Documents[i].SubCategoryIDs) > 0 {
			continue
		}
		menu = append(menu, map[string]any{
			"text": vueTitle(Documents[i].Title),
			"link": Documents[i].Hierarchy() + ".html",
		})
	}

	b, err := json.MarshalIndent(menu, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("cannot marshal JSON: %s", err.Error()))
	}

	return string(b)
}

func vuePressSubMenu(cat *category) []map[string]any {
	var menu []map[string]any
	for _, sub := range cat.SubCategories {
		var subMenu []map[string]any
		for i := range Documents {
			if Documents[i].IsInSubCategory(sub.ID) {
				subMenu = append(subMenu, map[string]any{
					"text": vueTitle(Documents[i].Title),
					"link": Documents[i].Hierarchy() + ".html",
				})
			}
		}
		menu = append(menu, map[string]any{
			"text":        vueTitle(sub.Title),
			"icon":        sub.VueIcon,
			"children":    subMenu,
			"collapsible": true,
		})
	}

	return menu
}

var rxVueTitle = regexp.MustCompile("[\\r\\n`]")

func vueTitle(s string) string {
	return rxVueTitle.ReplaceAllString(s, "")
}

/************
 *   Dump   *
 ************/

// Returns: A JSON dump of something (this is an internal tool for debugging)
func funcDump() string {
	b, _ := json.MarshalIndent(Config.Categories, "", "    ")
	return string(b)
}
