package main

type Post struct {
	ID int
	AuthorID int
	Body string
	Comments []Comment
}

func GetPostByID(id int) (p *Post) {
	return
}

func (p *Post) Create() error {
	return nil
}

func (p *Post) Update() error {
	return nil
}

func (p *Post) Delete() error {
	return nil
}