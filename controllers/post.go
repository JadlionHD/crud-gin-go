package controllers

type Post struct {
	ID        int           `json:"id"`
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
}
