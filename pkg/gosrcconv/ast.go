package gosrcconv

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Poser interface {
	Pos() token.Pos
}

func (c *Converter) FileOf(poser Poser) *ast.File {
	for _, pkg := range c.Packages {
		for _, file := range pkg.Package.Syntax {
			if file.Pos() <= poser.Pos() && file.End() > poser.Pos() {
				return file
			}
		}
	}
	return nil
}

func (c *Converter) AstOf(typeObj types.Object) (typeAst ast.Node) {
	ast.Inspect(c.FileOf(typeObj), func(node ast.Node) bool {
		if node == nil {
			return true
		}

		switch node.(type) {
		case *ast.File:
			// ignore
		default:
			if node.Pos() == typeObj.Pos() {
				typeAst = node
				return false
			}
		}
		return true
	})
	return
}
