package main

func main() {
	a := &Iterator{}

	for a.Next() {

	}
}

type Iterator struct {
	val   Value
	idx   *runtime.Runtime
	ctx   *adt.OpContext
	arcs  []field
	p     int
	cur   Value
	f     adt.Feature
	isOpt bool
}

func (i *Iterator) Next() bool {
	if i.p >= len(i.arcs) {
		i.cur = Value{}
		return false
	}
	f := i.arcs[i.p]
	f.arc.Finalize(i.ctx)
	p := linkParent(i.val.parent_, i.val.v, f.arc)
	i.cur = makeValue(i.val.idx, f.arc, p)
	i.f = f.arc.Label
	i.isOpt = f.isOptional
	i.p++
	return true
}
