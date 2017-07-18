// +build ignore

package shell

import (
	"github.com/lmorg/murex/lang/proc"
	"io/ioutil"
)

func dynFunctions() func(string) []string {
	return func(line string) (items []string) {
		for name := range proc.GoFunctions {
			if proc.GoFunctions[name].TypeIn == types.Null || proc.GoFunctions[name].TypeIn == types.Generic {
				items = append(items, name)
			}
		}
		return
	}
}

func dynMethods() func(string) []string {
	return func(line string) (items []string) {
		for name := range proc.GoFunctions {
			//if proc.GoFunctions[name].TypeIn == types.Null || proc.GoFunctions[name].TypeIn == types.Generic {
			items = append(items, "-> "+name)
			//}
		}
		return
	}
}

func dynParameters() func(string) []string {
	return func(line string) (items []string) {
		items = listFiles("./")
		//items = append([]string{")"}, items...)
		return
	}
}

func listFiles(path string) (filenames []string) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		filenames = append(filenames, f.Name())
	}
	return
}

func createCompleter() (completer *readline.PrefixCompleter) {
	newFunc := readline.PcItemDynamic(dynFunctions())

	newMethod := []readline.PrefixCompleterInterface{
		readline.PcItem("->", readline.PcItemDynamic(dynMethods())),
		readline.PcItem("|"),
		readline.PcItem("?"),
		readline.PcItem(";"),
	}

	params := readline.PcItemDynamic(
		dynParameters(),
		readline.PcItem(" "),
		readline.PcItem("->", newMethod...),
		readline.PcItem("|"),
		readline.PcItem("?"),
		readline.PcItem(";"),
		//readline.PcItemDynamic(dynMethods()),
	)

	newFunc.SetChildren([]readline.PrefixCompleterInterface{params})
	newMethod[2].SetChildren([]readline.PrefixCompleterInterface{params})
	params.Children[0].SetChildren([]readline.PrefixCompleterInterface{params})
	params.Children[1].SetChildren([]readline.PrefixCompleterInterface{params})
	params.Children[2].SetChildren([]readline.PrefixCompleterInterface{params})

	/*newFunc := readline.PcItemDynamic(dynComplete())
	newFunc.SetChildren([]readline.PrefixCompleterInterface{newFunc})*/

	completer = readline.NewPrefixCompleter(newFunc)

	//completer.Children[0].GetChildren()[0].
	return
}

/*var completer2 = readline.NewPrefixCompleter(
	readline.PcItem("text(",
		readline.PcItemDynamic(listFiles("./"),
			readline.PcItem(","),
			readline.PcItem(")",
				readline.PcItem(".", compChild),
				readline.PcItem("|", compChild),
				readline.PcItem("?", compChild),
				compChild,
			),
		),
	),
)*/
