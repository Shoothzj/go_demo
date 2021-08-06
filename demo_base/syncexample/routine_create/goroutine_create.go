package routine_create

var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
