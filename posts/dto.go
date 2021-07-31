package posts

import "encoding/json"

type Post struct {
	UserID int    `json:"userID"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) ToString() string {
	dBytes, _ := json.Marshal(p)
	return string(dBytes)
}
