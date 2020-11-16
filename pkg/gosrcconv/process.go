package gosrcconv

import (
	"errors"
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

func (c *Converter) process() error {
	for _, pkg := range c.Loader.Packages {
		err := c.processPackage(pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Converter) processPackage(pkg *packages.Package) error {
	c.Packages[pkg.PkgPath] = NewPackage(pkg)

	if pkg.Types != nil {
		qual := types.RelativeTo(pkg.Types)
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			object := scope.Lookup(name)
			err := c.processObject(pkg, object, qual)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Converter) processObject(pkg *packages.Package, object types.Object, qual types.Qualifier) error {
	var tname *types.TypeName
	typ := object.Type()

	switch tobj := object.(type) {
	case *types.PkgName:
	case *types.Const:
	case *types.TypeName:
		tname = tobj
	case *types.Var:
	case *types.Func:
	case *types.Label:
	case *types.Builtin:
	case *types.Nil:
	default:
		return errors.New(fmt.Sprintf("Unknown object type: %T", tobj))
	}

	if typ == nil {
		return nil
	}

	if tname != nil {
		// We have a type object: Don't print anything more for
		// basic types since there's no more information (names
		// are the same; see also comment in TypeName.IsAlias).
		//if _, ok := typ.(*types.Basic); ok {
		//	return
		//}
		if tname.IsAlias() {
			//buf.WriteString(" =")
		} else {
			typ = typ.Underlying()
		}
	}

	retType := c.returnType(pkg, object, typ, qual)
	switch rt := retType.(type) {
	case *Struct:
		c.Packages[pkg.PkgPath].Structs[object.Name()] = &ObjectStruct{
			Object: object,
			Struct: rt,
		}
	}

	return nil
}

func (c *Converter) returnType(pkg *packages.Package, object types.Object, typ types.Type, qual types.Qualifier) Type {
	return c.internalReturnType(pkg, object, typ, qual, make([]types.Type, 0, 8))
}

func (c *Converter) internalReturnType(pkg *packages.Package, object types.Object, typ types.Type,
	qual types.Qualifier, visited []types.Type) Type {
	visited = append(visited, typ)

	switch t := typ.(type) {
	case nil:
	case *types.Basic:
	case *types.Array:
	case *types.Slice:
	case *types.Struct:
		ret := &Struct{
			Type:    typ,
			Fields:  map[string]*types.Var{},
			Methods: map[string]*types.Func{},
		}
		for i := 0; i < t.NumFields(); i++ {
			f := t.Field(i)
			if f.Embedded() {
				continue
			}
			ret.Fields[f.Name()] = f
		}

		for _, meth := range typeutil.IntuitiveMethodSet(t, nil) {
			switch meth.Kind() {
			case types.MethodVal:
				ret.Methods[meth.Obj().Name()] = meth.Obj().(*types.Func)
			case types.MethodExpr:
				ret.Methods[meth.Obj().Name()] = meth.Obj().(*types.Func)
			default:
				panic(fmt.Sprintf("unsupported selector(%T)", meth))
			}
		}

		return ret
	case *types.Pointer:
	case *types.Tuple:
	case *types.Signature:
	case *types.Interface:
	case *types.Map:
	case *types.Chan:
	case *types.Named:
	default:

	}

	return nil
}
