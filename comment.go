package main

type Comment struct {
	ID int
	AuthorID int
	Body string
	PostID int
}

func GetCommentByID() (c *Comment) {
	return
}

func (c *Comment) Create() error {
	return nil
}

func (c *Comment) Update() error {
	return nil
}

func (c *Comment) Delete() error {
	return nil
}