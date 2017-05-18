package main

import (
	"compress/gzip"
	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
)

func main() {
	readFlags()

	if fTrace != "" {
		file, err := os.Create(fTrace)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = trace.Start(file)
		if err != nil {
			panic(err)
		}
		//defer trace.Stop()

		ctrlC := make(chan os.Signal)
		signal.Notify(ctrlC, os.Interrupt, syscall.SIGTERM)
		go func() {
			//os.Stderr.WriteString("1")
			<-ctrlC
			//os.Stderr.WriteString("2")
			trace.Stop()
			//os.Stderr.WriteString("3")
			file.Close()
			os.Stderr.WriteString("[SIGTERM]")
			os.Exit(1)
		}()
	}

	switch {
	case fCommand != "":
		execSource([]rune(fCommand))

	case fStdin:
		os.Stderr.WriteString("Not implimented yet.\n")
		os.Exit(1)

	case len(fSource) > 0:
		for _, filename := range fSource {
			execSource(diskSource(filename))
		}

	default:
		shell.Start()
	}

	debug.Log("[FIN]")
}

func diskSource(filename string) []rune {
	var b []byte

	file, err := os.Open(filename)
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		os.Exit(1)
	}

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(1)
		}
		b, err = ioutil.ReadAll(gz)

		file.Close()
		gz.Close()

		if err != nil {
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(1)
		}

	} else {
		b, err = ioutil.ReadAll(file)
		file.Close()
		if err != nil {
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(1)
		}
	}

	return []rune(string(b))
}

func execSource(source []rune) {
	exitNum, err := lang.ProcessNewBlock(
		source,
		nil,
		nil,
		nil,
		types.Null,
	)

	if err != nil {
		if exitNum == 0 {
			exitNum = 1
		}
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		os.Exit(exitNum)
	}

	if exitNum != 0 {
		os.Exit(exitNum)
	}
}
