package gosrcconv

type Converter struct {
	Loader   *Loader
	Packages map[string]*Package
}

func NewConverter(Loader *Loader) (*Converter, error) {
	ret := &Converter{
		Loader:   Loader,
		Packages: make(map[string]*Package),
	}
	err := ret.process()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
