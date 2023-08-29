package docgen

import (
	"bytes"
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
	"md":         funcMarkdown,
	"quote":      funcQuote,
	"trim":       strings.TrimSpace,
	"doc":        funcRenderedDocuments,
	"cat":        funcRenderedCategories,
	"link":       funcLink,
	"file":       funcFile,
	"notanindex": funcNotAnIndex,
	"date":       funcDate,
	"time":       funcTime,
	"doct":       funcDocT,
	"othercats":  funcOtherCats,
	"otherdocs":  funcOtherDocs,
	"env":        funcEnv,
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

// Takes: string (path, description)
// Returns: URL to document
func funcLink(path, description string) string {
	split := strings.Split(path, "/")
	if len(split) != 2 {
		panic(fmt.Sprintf("Invalid length of path (%d). Expecting 'cat/doc' instead got '%s'", len(split), path))
	}

	doc := Documents.ByID("", split[0], split[1])
	if doc == nil {
		panic(fmt.Sprintf("nil document (%s)", path))
	}

	return doc.Hierarchy()
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
			s = "undefined"
		}
		return s
	}

	key, value := envvars.Split(env)
	v := make(map[string]interface{})
	envvars.All(v)
	s, _ := v[key].(string)
	return s == value
}
