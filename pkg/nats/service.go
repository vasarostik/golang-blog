package nats

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/nats-io/go-nats"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
	"os"
)

func Start(cfg *config.Configuration, natsClient *nats.Conn, db *gorm.DB, logfile *os.File) error {

	var message = new(go_blog.PublishPostMessage)

	if _, err := natsClient.Subscribe(cfg.NATS_Subscriber.Subject, func (m *nats.Msg) {

		if err := json.Unmarshal(m.Data, &message); err != nil {
			panic(err)
		}

		post := GetPost(db,message.PostID)
		log.Printf("Action: Post with id '%d' was %s now", message.PostID , message.Action)

		err := WriteInFile(logfile, post, *message)

		if err != nil {
			panic(err)
		}

	}); err != nil {
		log.Fatalf("Failed to start subscription on '%s': %v", cfg.NATS_Subscriber.Subject, err)
	} else {

	}

	return nil
}
