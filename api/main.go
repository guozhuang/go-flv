package main

import (
	"api/pipeline"
	"fmt"
	"time"
)

/*import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/user", CreateUser)

	router.GET("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()

	http.ListenAndServe(":80", r)
}*/

func main() {
	testPiipeline()

	for {
		time.Sleep(time.Duration(5) * time.Second)
		break
	}
}

func testPiipeline() {
	done := make(chan interface{})
	defer close(done)

	intStream := pipeline.Generator(done, 1, 2, 3, 4)

	pipelineStream := pipeline.Multiply(done, pipeline.Add(done, pipeline.Multiply(done, intStream, 2), 1), 2)

	for v := range pipelineStream {
		fmt.Println(v)
	}
}
