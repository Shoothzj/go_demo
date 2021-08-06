package ut

import "fmt"

type Base struct {
}

func (receiver *Base) xxx() {
	fmt.Println("base xxx")
	receiver.yyy()
}

func (receiver *Base) yyy() {
	fmt.Println("base yyy")
}

type BaseMock struct {
	Base
}

func (receiver *BaseMock) yyy() {
	fmt.Println("base mock yyy")
}
