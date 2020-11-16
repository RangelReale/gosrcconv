package main

import (
	"flag"
	"fmt"
	"github.com/RangelReale/gosrcconv/pkg/gosrcconv"
	"github.com/RangelReale/gosrcconv/pkg/pythonsrc"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	outputPath := filepath.Clean(flag.Arg(0))

	fmt.Printf("Output path: %s...\n", outputPath)

	l := gosrcconv.NewLoader()
	err := l.LoadStd("text/template", "text/template/parse")
	if err != nil {
		panic(err)
	}

	c, err := l.Converter()
	if err != nil {
		panic(err)
	}
	for pn, pv := range c.Packages {
		filename := filepath.Join(outputPath, fmt.Sprintf("%s.py", strings.Replace(pn, "/", ".", -1)))
		writer := pythonsrc.NewPythonWriter(c, pv)
		err = writer.OutputFile(filename)
		if err != nil {
			panic(err)
		}
	}
}
