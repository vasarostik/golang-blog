package nats

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nats-io/go-nats"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
)

func Start(cfg *config.Configuration, natsClient *nats.Conn, db *gorm.DB) error {

	var message = new(go_blog.PublishPostMessage)

	if _, err := natsClient.Subscribe(cfg.NATS_Subscriber.Subject, func (m *nats.Msg) {

		if err := json.Unmarshal(m.Data, &message); err != nil {
			panic(err)
		}

		post := GetPost(db,message.PostID)

		fmt.Printf("Received a message - (Title:%s, PostID:%d, UserID:%d, Content:%s, Action:%s, TimeOfAction: %s)\n", post.Title, message.PostID, post.UserID, post.Content, message.Action, message.Timestamp)

	}); err != nil {
		log.Fatalf("Failed to start subscription on '%s': %v", cfg.NATS_Subscriber.Subject, err)
	}

	return nil
}
