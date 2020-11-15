package gosrcconv

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Package struct {
	Package *packages.Package
	Structs map[string]*Struct
}

func NewPackage(pkg *packages.Package) *Package {
	return &Package{
		Package: pkg,
		Structs: map[string]*Struct{},
	}
}

//type TopLevelObject struct {
//	Name string
//}

type Type interface {
	Underlying() Type

	String() string
}

type Struct struct {
	Type    types.Type
	Fields  []*types.Var
	Methods []*types.Func
}

// Implementations for Type methods.

func (s *Struct) Underlying() Type { return s }

func (b *Struct) String() string { return "Struct" }
