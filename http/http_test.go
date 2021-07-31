package http

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	h := &HTTP{
		GetURL: "https://jsonplaceholder.typicode.com",
	}
	dBytes, err := h.Get("/posts/1")
	if err != nil {
		t.Log(fmt.Sprintf("err occured: %s", err.Error()))
		return
	}
	t.Log(string(dBytes))
}

func TestPrintSomething(t *testing.T) {
	fmt.Println("Say hi")
	t.Log("Say bye")
}
