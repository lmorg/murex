package table

import (
	"bufio"
	"encoding/csv"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["tabulate"] = cmdTabulate

	/*defaults.AppendProfile(`
		autocomplete set tabulate { [{
			"Flags": ({ tabulate --help }),
			"AllowMultiple": true
		}] }
	`)*/
}

func cmdTabulate(p *lang.Process) error {
	p.Stdout.SetDataType("csv")

	const (
		fNoTrim    = "--no-trim"
		fSeparator = "--separator"
		//fJoiner    = "--joiner"
		fHelp = "--help"
	)

	flags := map[string]string{
		fNoTrim:    types.Boolean,
		fSeparator: types.String,
		//fJoiner:    types.String,
	}

	/*help := func() (s []string) {
		for f := range flags {
			s = append(s, f)
		}

		sort.Strings(s)
		return
	}*/

	f, _, err := p.Parameters.ParseFlags(
		&parameters.Arguments{
			Flags:           flags,
			AllowAdditional: false,
		},
	)

	if err != nil {
		return err
	}

	var (
		trim      = true
		separator = `\s[\s]+`
		//joiner    = ","
	)

	for flag, value := range f {
		switch flag {
		case fNoTrim:
			trim = false
		case fSeparator:
			separator = value
		//case fJoiner:
		//	joiner = value
		case fHelp:
			// print help
		}
	}

	rxTableSplit, err := regexp.Compile(separator)
	if err != nil {
		return err
	}

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	/*aw, err := p.Stdout.WriteArray(types.Generic)
	if err != nil {
		return err
	}*/

	var (
		firstRec  bool
		offByOne  bool
		csvWriter = csv.NewWriter(p.Stdout)
	)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		if trim {
			s = strings.TrimSpace(s)
		}

		if !rxTableSplit.MatchString(s) {
			continue
		}

		split := rxTableSplit.Split(scanner.Text(), -1)
		if len(split) == 0 {
			continue
		}

		firstRec = true
		if firstRec && split[0] == "" {
			offByOne = true
		}

		if offByOne && len(split) > 1 && split[0] == "" {
			split = split[1:]
		}

		//aw.WriteString(strings.Join(split, joiner))
		err := csvWriter.Write(split)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	//return aw.Close()
	csvWriter.Flush()
	return csvWriter.Error()
}
