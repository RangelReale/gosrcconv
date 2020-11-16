package main

import (
	"flag"
	"fmt"
	"github.com/RangelReale/gosrcconv/pkg/gosrcconv"
	"github.com/davecgh/go-spew/spew"
	"os"
	"path/filepath"
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
		fmt.Println(pn)
		spew.Dump(pv.Structs)
	}
}
