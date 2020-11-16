package gosrcconv

import (
	"go/constant"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Package struct {
	Package *packages.Package
	Structs map[string]*ObjectStruct
	Consts map[string]*ObjectConst
}

func NewPackage(pkg *packages.Package) *Package {
	return &Package{
		Package: pkg,
		Structs: map[string]*ObjectStruct{},
		Consts: map[string]*ObjectConst{},
	}
}

type Type interface {
	Underlying() Type
}

type Const struct {
	Type types.Type
	Val constant.Value
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

type ObjectConst struct {
	Object types.Object
	Const *Const
}

type ObjectStruct struct {
	Object types.Object
	Struct *Struct
}

// Implementations for Type methods.

func (t *Const) Underlying() Type { return t }

func (t *Struct) Underlying() Type { return t }

func (t *Interface) Underlying() Type { return t }
