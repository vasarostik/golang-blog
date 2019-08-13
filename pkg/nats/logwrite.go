package nats

import (
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
	"os"
)

func WriteInFile(file *os.File, post go_blog.Post, message go_blog.PublishPostMessage) error{

	logger := log.New(file, "", 0)

	if message.Action != "deleted"{
		logger.Printf("User with ID '%d' has %s a post with ID '%d' at '%s': \t Title: '%s' , Content: '%s'", message.UserID, message.Action, message.PostID, message.Timestamp, post.Title, post.Content)
	} else {
		logger.Printf("User with ID '%d' has %s a post with ID '%d' at '%s'", message.UserID, message.Action, message.PostID, message.Timestamp)
	}

	return nil
}
