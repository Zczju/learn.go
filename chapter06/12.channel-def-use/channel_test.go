package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{}
	fmt.Println("sampleMap:", sampleMap)
	var intCh chan int
	fmt.Println("intCh", intCh) // 引用类型，不实例化就是nil
	intCh = make(chan int, 1)   // 空的chan和map一样，只能读不能写。没有buffer就写入会deadlock
	fmt.Println("intCh", intCh)

	fmt.Println("装入数据")
	intCh <- 3
	fmt.Println("取出数据")
	out := <-intCh
	fmt.Println("数据是", out)
}

func TestChanPutGet(t *testing.T) {
	intCh := make(chan int) // 创建一个不带size的channel，不带buffer
	workerCount := 10       // 多个worker，看是否是线程安全的
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i
		}(i)
	}

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d拿到%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果 : 如果没有out(取数据的worker)	，那么in部分会等待，直到有out开始
// size为0即手把手的交换信息
func TestChanPutGet2_oWait(t *testing.T) {
	intCh := make(chan int) // 创建一个不带size的channel，不带buffer
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作:", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作:", time.Now())
		}(i)
	}

	fmt.Println(time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果 : 带有buffer的会直接允许in，不影响out
func TestChanPutGet2_oWait_withBuffer(t *testing.T) {
	intCh := make(chan int, 10) // 创建一个带size的channel，带buffer
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作:", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作:", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果 : buffer里面只能装指定size的obj，buffer装满以后就不能再装了，要等待；带buffer即有一定的冗余空间
func TestChanPutGet2_oWait_withSmallBuffer(t *testing.T) {
	intCh := make(chan int, 2) // 创建一个带size的channel，带buffer
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作:", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作:", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// out 先启动， in 后续启动
// 结果 :没有in的时候，out也要等待
func TestChanPutGet2_oFirst_withBuffer(t *testing.T) {
	intCh := make(chan int, 2) // 创建一个带size的channel，带buffer，解决等待问题，提升程序性能
	workerCount := 10
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}

	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作:", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作:", time.Now())
		}(i)
	}

	time.Sleep(1 * time.Second)
}

// range没有close的话，在取出所有数据之后panic :deadlock
func TestRangeChannel(t *testing.T) {
	intCh := make(chan int, 10) // 创建一个带size的channel，带buffer
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	for o := range intCh {
		fmt.Println(o)
	}
}

func TestRangeClosedChannel(t *testing.T) {
	intCh := make(chan int, 10) // 创建一个带size的channel，带buffer
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5
	close(intCh)
	{
		o1, ok := <-intCh
		fmt.Println("直接取数", o1, ok)
	}

	{
		for o := range intCh {
			fmt.Println("range取数", o)
		}
	}

	o1, ok := <-intCh   //close了的channel，读取会读到假数据
	fmt.Println(o1, ok) // 取到的是假数据
}

// 当所有的channel都不ready的时候，select会等待，直到某个channel ready
func TestSelectChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", o)
	}
	fmt.Println("DONE")
}

// default 直接运行，不去等待channel ready
func TestSelectChannelWithDefault(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("DONE")
}

// case的优先级高于default， 只要有case中的channel ready， default就不走了
func TestSelectChannelWithDefaultAndChannelReady(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())

	ch1 <- 1
	//go func() {
	//  time.Sleep(1 * time.Second)
	//  ch1 <- 1
	//}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("DONE")
}

// closed channel 在select的时候总是可用
func TestSelectChannelWithDefaultAndClosedChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())

	fmt.Println("关闭 ch1")
	close(ch1)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch1 <- 1
	//}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", o)
	}
	fmt.Println("DONE")
}

func TestMultipleSelect(t *testing.T) {
	ch1 := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-ch1:
				fmt.Println(time.Now(), i)
			}
		}(i)
	}

	fmt.Println("关闭channel: ", time.Now())
	close(ch1) // 关闭channel之后，向所有的协程发送信号，并走select里面case为ch1的部分

	time.Sleep(1 * time.Second)
}

func TestDualChannel(t *testing.T) {
	c := make(chan struct{})
	close(c)
	// close(c) // panic: close of closed channel 只能关闭一次
}

func TestPut2ClosedChannel(t *testing.T) {
	intCh := make(chan int, 10) // 创建一个带size的channel，带buffer
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5
	close(intCh)

	//intCh <- 6 // panic: send on closed channel 已经关闭的channel不能再装数据
}

func TestOutOnly(t *testing.T) {
	c := make(chan int, 10)
	<-c // fatal error: all goroutines are asleep - deadlock!
}

func TestSingleGoroutineApp(t *testing.T) {
	intCh := make(chan int)
	intCh <- 1 // fatal error: all goroutines are asleep - deadlock!
	<-intCh
}

func TestMultipleChannelReadySelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	close(ch1)
	close(ch2)

	ch1Counter, ch2Counter := 0, 0

	for i := 0; i < 10000; i++ {
		select {
		case <-ch1:
			ch1Counter++
		default:
			select {
			case <-ch2:
				ch2Counter++ // 修改优先级，由随机变为必然
			}

		}
	}
	fmt.Println("ch1Counter: ", ch1Counter)
	fmt.Println("ch2Counter: ", ch2Counter)
	// 两个都ready的时候，会随机走一个
}
