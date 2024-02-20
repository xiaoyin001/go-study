package tests

import (
	"fmt"
	"testing"
)

/*
	对闭包来说，函数在该语言中得是一等公民。
	一般来说，一个函数返回另外一个函数，这个被返回的函数可以引用外层函数的局部变量，这形成了一个闭包。
	通常，闭包通过一个结构体来实现，它存储一个函数和一个关联的上下文环境。
	但 Go 语言中，匿名函数就是一个闭包，它可以直接引用外部函数的局部变量
*/

func app() func(string) string {
	t := "Hi"

	// 比如这里的 c 和上面的 t 就是一个闭包，可以理解为是一整个结构体
	c := func(b string) string {
		t = t + " " + b
		fmt.Println("t 地址:", &t)
		return t
	}

	// 这里的返回值 c 可以理解为返回了一个结构体，里面包含 c和t 两个成员变量
	return c
}

func TestBiBao(t *testing.T) {
	// 这里的a可以理解为 一个结构体 里面有上面的c和t
	a := app()
	// 这里的a可以理解为 一个结构体 里面也有上面的c和t，但是 a和b 这里已经不相关联了
	b := app()

	// 这里操作一次a 相当于是修改了a这个“结构体”里面t的值   t = Hi go
	a("go")
	// 这里操作的是b这个“结构体”里面t的值   t = Hi All
	fmt.Println(b("All"))
	// 同理这里操作的是“结构体”a 里面的 t = Hi go All
	fmt.Println(a("All"))
	// 同理这里操作的是“结构体”a 里面的 t = Hi go All xiaoyin
	fmt.Println(a("xiaoyin"))
}
