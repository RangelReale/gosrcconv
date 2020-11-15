package main

import (
	"fmt"
	"github.com/RangelReale/gosrcconv/pkg/gosrcconv"
	"github.com/davecgh/go-spew/spew"
)

func main() {
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
