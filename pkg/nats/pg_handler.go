package nats

import (
	"github.com/jinzhu/gorm"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

func GetPost(db *gorm.DB, id int) (go_blog.Post) {
	post := go_blog.Post{}

	db.Where("id = ?",id).Find(&post)

	return post
}