package models

type Error struct {
	Status     int
	StatusText string
	Message    string
	Back       string
}

type Info struct {
	User
	Post
	Comment []Comment
}

type InfoPosts struct {
	User
	Posts    []Post
	Category []string
}
