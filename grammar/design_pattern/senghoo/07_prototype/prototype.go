package _7_prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prorotypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prorotypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prorotypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prorotypes[name] = prototype
}
