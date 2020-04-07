package pipeline

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	//测试pipeline
	done := make(chan interface{})
	defer close(done)

	intStream := Generator(done, 1, 2, 3, 4)

	pipelineStream := Multiply(done, Add(done, Multiply(done, intStream, 2), 1), 2)

	for v := range pipelineStream {
		fmt.Println(v)
	}
}
