package _6_builder

// TODO: 泛型优化
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

// Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	num string
}

func (b *Builder1) Part1() {
	b.num += "1"
}

func (b *Builder1) Part2() {
	b.num += "2"
}

func (b *Builder1) Part3() {
	b.num += "3"
}

func (b *Builder1) GetResult() string {
	return b.num
}

type Builder2 struct {
	num int
}

func (b *Builder2) Part1() {
	b.num += 1
}

func (b *Builder2) Part2() {
	b.num += 2
}

func (b *Builder2) Part3() {
	b.num += 3
}

func (b *Builder2) GetResult() int {
	return b.num
}
