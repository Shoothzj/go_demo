package demo_base

const (
	// iota reset to 0
	// c0 = 0
	c0 = iota
	// c1 = 1
	c1 = iota
	// c2 = 2
	c2 = iota
)

const (
	// iota reset to 0
	// a = 1
	a = 1 << iota
	// b = 2
	b = 1 << iota
	// c = 4
	c = 1 << iota
)

const (
	// u = 0
	u = iota * 42
	// v = 42.0
	v float64 = iota * 42
	// w = 84
	w = iota * 42
)

const x = 0
const y = 0

// 在表达式列表中，每个iota的值都相同，iota仅仅在每个ConstSpec之后递增
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
	bit1, mask1                          // bit1 == 2, mask1 == 1
	_, _                                 // skips iota == 2
	bit3, mask3                          // bit3 == 8, mask3 == 7
)
