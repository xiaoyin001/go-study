package tests

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestNeiCun(t *testing.T) {
	Consume()
}

type AAAAAAA struct {
	A    int64
	B    int64
	C    int64
	D    int64
	AAAA []*AAAAAAA
}

func Consume() {
	var c = make(chan *AAAAAAA, 10)

	go func() {
		for {
			select {
			case d := <-c:
				fmt.Println(d)
			default:
				fmt.Println("11")
			}

			// 这个携程取数据 1ms 一个（太慢了），给的地方一直在给，这里处理不过来，导致数据一直堆积，内存就在暴涨
			time.Sleep(time.Millisecond)
		}
	}()

	gob.Register(&AAAAAAA{})
	i := 0
	for {
		aaaa := &AAAAAAA{}
		for i := 0; i < 1000; i++ {
			aaaa.AAAA = append(aaaa.AAAA, &AAAAAAA{})
		}

		var bf bytes.Buffer
		enc := gob.NewEncoder(&bf)
		err := enc.Encode(&aaaa)
		if err != nil {
			fmt.Println("enc:", err)
		}

		dec := gob.NewDecoder(&bf)
		dd := &AAAAAAA{}
		err = dec.Decode(&dd)
		if err != nil {
			fmt.Println("dec:", err)
		}

		go func(d *AAAAAAA) {
			// 这行的休眠是迷惑的代码，感觉是已经在控制速度了，但是外面的速度是很快的
			time.Sleep(time.Millisecond)
			// 可能 1ms 就生成了100份，然后这100份数据，都延迟 1ms 总的只要了 1ms ，然后就将这100份数据放入管道，如果管道放不下，还得继续阻塞
			// 这100份数据这里总时间可能只用了 2ms，但是上面处理这100份数据至少需要100ms，这样存取速度差就出来了，就导致内存中的数据越来越多
			c <- d
		}(dd)

		i++
		if i == math.MaxInt {
			break
		}

		// time.Sleep(time.Millisecond)
	}
}
