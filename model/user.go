package model

type User struct {
	ID       int64
	Email    string
	Password string
}

type ChatRoom struct {
	ID    int64
	Title string
}
