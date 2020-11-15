package gosrcconv

import (
	"github.com/hashicorp/go-multierror"
	"go/token"
	"golang.org/x/tools/go/packages"
)

type Loader struct {
	Packages []*packages.Package
	FileSet  *token.FileSet
}

func NewLoader() *Loader {
	return &Loader{
		Packages: []*packages.Package{},
		FileSet:  token.NewFileSet(),
	}
}

func (l *Loader) LoadStd(packageNames ...string) error {
	cfg := &packages.Config{Fset: l.FileSet, Mode: PackagesLoadMode}
	pkgs, err := packages.Load(cfg, packageNames...)
	if err != nil {
		return err
	}

	if err = l.checkPackageErrors(pkgs); err != nil {
		return err
	}

	l.Packages = append(l.Packages, pkgs...)

	return nil
}

func (l *Loader) LoadDir(directoryName string) error {
	cfg := &packages.Config{Fset: l.FileSet, Dir: directoryName, Mode: PackagesLoadMode}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		l.Packages = append(l.Packages, pkgs...)
	}
	return err
}

func (l *Loader) Converter() (*Converter, error) {
	return NewConverter(l)
}

func (l *Loader) checkPackageErrors(pkgs []*packages.Package) error {
	var errorList error
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		for _, err := range pkg.Errors {
			errorList = multierror.Append(errorList, err)
		}
	})
	return errorList
}
