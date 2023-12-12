package tests

import (
	"fmt"
	"math"
	"testing"
)

func TestWeiYunSuan(t *testing.T) {
	var id int64 = math.MaxInt64
	fmt.Printf("%b \n", id) // 这个数打印出来只有63位，还有一位隐藏的符号位
	// 111111111111111111111111111111111111111111111111111111111111111

	fmt.Printf("%b \n", int64(id>>40)) // 右移40位，得到前23位数据
	// 000000000000000000000000000000000000000011111111111111111111111
	// 11111111111111111111111

	fmt.Printf("%b \n", int64(id^(id>>40<<40))) // 后40
	// 000000000000000000000000000000000000000011111111111111111111111   【id>>40】
	// 111111111111111111111110000000000000000000000000000000000000000   【id>>40<<40】
	// 111111111111111111111111111111111111111111111111111111111111111
	// 000000000000000000000001111111111111111111111111111111111111111   【id^(id>>40<<40)】
	// 1111111111111111111111111111111111111111

	fmt.Printf("%b \n", int64(id<<24>>24)) // 24？ 这种直接先左移是不正确的
}
