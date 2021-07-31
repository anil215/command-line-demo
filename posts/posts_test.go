package posts

import (
	"encoding/json"
	"fmt"
	"testing"

	"example.com/command/http"
)

func TestGetPostForAnId(t *testing.T) {
	p := &Posts{
		PostGetter: http.NewHTTP("https://jsonplaceholder.typicode.com"),
	}
	post, err := p.GetPostForAnId(1)
	if err != nil {
		fmt.Printf("failed to get posts with err: %s", err.Error())
		return
	}
	dBytes, _ := json.Marshal(post)
	fmt.Println(string(dBytes))
}
