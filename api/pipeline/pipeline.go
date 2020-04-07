package pipeline

import "fmt"

//构建一个简易的pipeline

//也就结合了一直说明的channel和goroutine的最佳实践

//标准化生成器
func Generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for _, v := range integers {
			select {
			case <-done:
				fmt.Println("gen is done")
				return
			case intStream <- v:
				fmt.Printf("gen data is %d\n", v)
			}
		}
	}()
	return intStream
}

//保证pipeline中每个组建消耗并且返回同样类型
func Multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)

	go func() {
		defer close(multipliedStream)

		for i := range intStream {
			select {
			case <-done:
				fmt.Println("multi is done")
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()

	return multipliedStream
}

func Add(done <-chan interface{}, intStream <-chan int, addNum int) <-chan int {
	addStream := make(chan int)

	go func() {
		defer close(addStream)

		for i := range intStream {
			select {
			case <-done:
				fmt.Println("add is done")
				return
			case addStream <- i + addNum:
				fmt.Printf("add data is %d\n", i)
			}

		}
	}()
	return addStream
}
