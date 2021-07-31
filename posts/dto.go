package posts

type Post struct {
	UserID int    `json:"userID"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
