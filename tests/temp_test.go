package tests

import (
	"fmt"
	"testing"
)

type AAA struct {
	Name string
}

func (a *AAA) ShowName1() {
	fmt.Println("Name=", a.Name)
}

func (a AAA) ShowName2() {
	fmt.Println("Name=", a.Name)
}

func Test_Go(t *testing.T) {
	mA := AAA{
		Name: "aaa",
	}

	mA.ShowName1()
	mA.ShowName2()
}
