// +build ignore

package cmdautocomplete

// This file isn't yet in use and might be removed at a later date.
// That said, it would be good to find a way to compile autocompletions.
// I'm just not convined this is the best way to accumplish this.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

func cacheDynamic(p *lang.Process) error {
	fn, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	/*ver, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}*/

	errCallback := func(err error) {
		p.Stderr.Writeln([]byte(err.Error()))
	}

	softTimeout, err := lang.ShellProcess.Config.Get("shell", "autocomplete-soft-timeout", types.Integer)
	if err != nil {
		softTimeout = 100
	}

	var (
		i     int                          // index inside array-based flag completion
		flags []string                     // flags to pass the autocomplete function
		oac   = autocomplete.ExesFlags[fn] // old autocomplete
		//nac   = make([]autocomplete.Flags, len(oac)) // new autocomplete
	)

	//copy(nac, oac)

	if oac == nil {
		return fmt.Errorf("nil oac")
	}

	nac := recursiveCache(fn, oac, errCallback, flags, softTimeout.(int), i)

	autocomplete.ExesFlags[fn+"---"] = nac

	return nil
}

func recursiveCache(fn string, oac []autocomplete.Flags, errCallback func(error), flags []string, softTimeout, i int) []autocomplete.Flags {
	debug.Json("oac", oac)
	debug.Json("flags", flags)

	nac := make([]autocomplete.Flags, len(oac)) // new autocomplete
	copy(nac, oac)

	pt, _ := parser.Parse([]rune(fn+" "+strings.Join(flags, " ")), 0)
	debug.Json("pt", pt)
	softCtx, _ := context.WithTimeout(context.Background(), time.Duration(int64(softTimeout))*time.Millisecond)

	act := &autocomplete.AutoCompleteT{
		Definitions:       make(map[string]string),
		CacheDynamic:      true,
		ParsedTokens:      pt,
		DelayedTabContext: readline.DelayedTabContext{Context: softCtx},
	}

	/*act.Definitions = make(map[string]string)
	act.Items = []string{}
	act.DelayedTabContext = readline.DelayedTabContext{Context: softCtx}
	act.ParsedTokens = pt*/

	getFlags(act)

	nac[i].Flags = act.Items
	nac[i].FlagsDesc = act.Definitions
	/*nac[i].Flags = make([]string, len(act.Items))
	copy(nac[i].Flags, act.Items)

	nac[i].FlagsDesc = make(map[string]string)
	for k, v := range act.Definitions {
		nac[i].FlagsDesc[k] = v
	}*/

	nac[i].Dynamic = ""
	nac[i].DynamicDesc = ""

	dedup := make(map[string]bool)
	for _, flag := range nac[i].Flags {
		dedup[flag] = true
	}
	for flag := range nac[i].FlagsDesc {
		dedup[flag] = true
	}

	for flag := range dedup {
		if len(flag) == 0 || flag[0] == '-' {
			continue
		}

		if nac[i].FlagValues == nil {
			nac[i].FlagValues = make(map[string][]autocomplete.Flags)
		}
		if nac[i].FlagValues[flag] == nil || len(nac[i].FlagValues[flag]) == 0 {
			nac[i].FlagValues[flag] = make([]autocomplete.Flags, 1)
		}

		nac[i].FlagValues[flag] = recursiveCache(fn, nac[i].FlagValues[flag], errCallback, append(flags, flag), softTimeout, i)
	}

	return nac
}

func getFlags(act *autocomplete.AutoCompleteT) {
	pIndex := 0
	autocomplete.MatchFlags(autocomplete.ExesFlags[act.ParsedTokens.FuncName], "", act.ParsedTokens.FuncName, act.ParsedTokens.Parameters, &pIndex, act)
}
