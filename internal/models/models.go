package models

// Post представляет пост в системе.
// swagger:model Post
type Post struct {
	ID     int    `json:"id" example:"1"`
	UserID int    `json:"userId" example:"1"`
	Title  string `json:"title" example:"My Title"`
	Body   string `json:"body" example:"Post content"`
}
