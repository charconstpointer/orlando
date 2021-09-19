package orlando

type Hasher func(s string) uint32

type Filter struct {
	hashers []Hasher
	bits    []*struct{}
}

func NewFilter(size int, hashers ...Hasher) (*Filter, error) {
	return &Filter{
		hashers: hashers,
		bits:    make([]*struct{}, size),
	}, nil
}

func (f *Filter) Insert(s string) error {
	for _, h := range f.hashers {
		i := h(s) % uint32(len(f.bits))
		f.bits[i] = &struct{}{}
	}
	return nil
}

func (f *Filter) Contains(s string) bool {
	for _, h := range f.hashers {
		i := h(s) % uint32(len(f.bits))
		if f.bits[i] == nil {
			return false
		}
	}
	return true
}
