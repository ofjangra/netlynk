package helpers

import (
	"context"
	"time"

	"github.com/ofjangra/netlynk_server/models"
)

func CreateALink(link models.Links) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	_, err := linksCollection.InsertOne(ctx, link)

	if err != nil {
		return err
	}

	return nil
}
