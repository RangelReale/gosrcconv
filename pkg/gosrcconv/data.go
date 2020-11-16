package gosrcconv

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Package struct {
	Package *packages.Package
	Structs map[string]*ObjectStruct
}

func NewPackage(pkg *packages.Package) *Package {
	return &Package{
		Package: pkg,
		Structs: map[string]*ObjectStruct{},
	}
}

type Type interface {
	Underlying() Type
}

type Struct struct {
	Type    types.Type
	Fields  map[string]*types.Var
	Methods map[string]*types.Func
}

type Interface struct {
	Type    types.Type
	Fields  []*types.Var
	Methods []*types.Func
}

// Object

type ObjectStruct struct {
	Object types.Object
	Struct *Struct
}

// Implementations for Type methods.

func (s *Struct) Underlying() Type { return s }
