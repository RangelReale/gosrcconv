package pythonsrc

import (
	"fmt"
	"github.com/RangelReale/gosrcconv/pkg/gen"
	"github.com/RangelReale/gosrcconv/pkg/gosrcconv"
	"go/types"
	"io"
	"os"
	"strings"
)

type PythonWriter struct {
	converter *gosrcconv.Converter
	pkg       *gosrcconv.Package
	gf        *gen.GenFile
}

func NewPythonWriter(converter *gosrcconv.Converter, pkg *gosrcconv.Package) *PythonWriter {
	return &PythonWriter{
		converter: converter,
		pkg:       pkg,
	}
}

func (w *PythonWriter) Output(out io.Writer) error {
	w.gf = gen.NewGenFile()
	defer func() {
		w.gf = nil
	}()

	qual := types.RelativeTo(w.pkg.Package.Types)

	var err error
	w.gf.Line("# package: %s", w.pkg.Package.PkgPath)

	for sn, st := range w.pkg.Structs {
		err = w.writeStruct(sn, st, qual)
		if err != nil {
			return err
		}
	}

	return w.gf.Output(out)
}

func (w *PythonWriter) OutputFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err = w.Output(f); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

func (w *PythonWriter) writeStruct(name string, ostruct *gosrcconv.ObjectStruct, qf types.Qualifier) error {
	fieldCount := 0

	w.gf.Line("# %s", w.converter.Loader.FileSet.Position(ostruct.Object.Pos()))

	w.gf.Line("class %s:", ostruct.Object.Name())
	w.gf.I()
	for _, field := range ostruct.Struct.Fields {
		w.gf.StartLine()
		w.gf.Append("%s: ", w.pythonIdent(field.Name()))
		w.gf.Append(w.returnType(field.Type(), qf))
		//p.AppendLineComment(gf, f)
		w.gf.NL()
		fieldCount++
	}

	if fieldCount == 0 {
		w.gf.Line("pass")
	}
	w.gf.D()
	w.gf.NL()

	return nil
}

func (w *PythonWriter) returnType(typ types.Type, qf types.Qualifier) string {
	return w.internalReturnType(typ, qf, make([]types.Type, 0, 8))
}

func (w *PythonWriter) internalReturnType(typ types.Type, qf types.Qualifier, visited []types.Type) string {
	visited = append(visited, typ)

	var tb strings.Builder

	switch t := typ.(type) {
	case nil:
	case *types.Basic:
		tb.WriteString(w.pythonType(t))
	//case *types.Array:
	//case *types.Slice:
	case *types.Struct:
		panic("Cannot process internal struct")
	//case *types.Pointer:
	//case *types.Tuple:
	//case *types.Signature:
	//case *types.Interface:
	//case *types.Map:
	//case *types.Chan:
	//case *types.Named:
	default:
		tb.WriteString("Any")
		//panic("Cannot determine type to return")
	}

	return tb.String()
}

func (w *PythonWriter) pythonType(ptype *types.Basic) string {
	switch ptype.Kind() {
	case types.UntypedBool, types.Bool:
		return "bool"
	case types.UntypedInt, types.Int, types.Int8, types.Int16, types.Int32, types.Int64, types.Uint, types.Uint8,
		types.Uint16, types.Uint32, types.Uint64:
		return "int"
	case types.Uintptr:
		return "Optional[int]"
	case types.UntypedFloat, types.Float32, types.Float64:
		return "float"
	case types.UntypedComplex, types.Complex64, types.Complex128:
		return "complex"
	case types.UntypedString, types.UntypedRune, types.String:
		return "str"
	case types.UntypedNil:
		return "None"
	case types.Invalid, types.UnsafePointer:
		return "Optional[Any]"
	}

	return ptype.Name()
}

func (w *PythonWriter) pythonIdent(ident string) string {
	if ident == "True" || ident == "False" {
		return fmt.Sprintf("%s_", ident)
	}
	return ident
}
