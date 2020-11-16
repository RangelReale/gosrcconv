package gen

import (
	"fmt"
	"io"
	"strings"
)

type GenFile struct {
	pb     strings.Builder
	indent int
}

func NewGenFile() *GenFile {
	return &GenFile{}
}

func (g *GenFile) Append(format string, args ...interface{}) {
	_, err := fmt.Fprintf(&g.pb, format, args...)
	if err != nil {
		panic(err)
	}
}

func (g *GenFile) StartLine() {
	if g.indent > 0 {
		g.Append(strings.Repeat("    ", g.indent))
	}
}

func (g *GenFile) Line(format string, args ...interface{}) {
	g.StartLine()
	g.Append(format, args...)
	g.Append("\n")
}

func (g *GenFile) NL() {
	g.Append("\n")
}

func (g *GenFile) I() {
	g.indent++
}

func (g *GenFile) D() {
	g.indent--
}

func (g *GenFile) Output(out io.Writer) error {
	_, err := out.Write([]byte(g.pb.String()))
	return err
}
