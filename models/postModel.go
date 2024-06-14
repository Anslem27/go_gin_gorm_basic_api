package models

import "gorm.io/gorm"

/* refer to https://gorm.io/docs/models.html */

type Post struct {
	gorm.Model
	Title string
	Body  string
	// Comments []Comment
}

type Comment struct {
	gorm.Model
	PostID  uint
	Content string
	Post    Post
}
