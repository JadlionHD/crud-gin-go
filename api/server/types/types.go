package types

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type Post struct {
	ID        int64         `json:"id"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	Tags      []string      `json:"tags"`
	Reactions PostReactions `json:"reactions"`
	Views     int           `json:"views"`
	UserID    int           `json:"userId"`
}

type PostReactions struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

type CreatePostInput struct {
	ID    int64    `json:"id" binding:"required"`
	Title string   `json:"title" binding:"required"`
	Body  string   `json:"body" binding:"required"`
	Tags  []string `json:"tags" binding:"required"`
}
