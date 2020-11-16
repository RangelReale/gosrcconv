package main

import "github.com/RangelReale/gosrcconv/pkg/gosrcconv"

type PythonWriter struct {
	converter *gosrcconv.Converter
	pkg       *gosrcconv.Package
}

func NewPythonWriter(converter *gosrcconv.Converter, pkg *gosrcconv.Package) *PythonWriter {
	return &PythonWriter{
		converter: converter,
		pkg:       pkg,
	}
}
