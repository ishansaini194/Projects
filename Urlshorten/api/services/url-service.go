package services

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ishansaini194/Projects/database"
)

func GenerateID(custom string) string {
	if custom != "" {
		return custom
	}
	return uuid.New().String()[:6]
}

func Exists(id string) bool {
	r := database.CreateClient(0)
	defer r.Close()

	val, _ := r.Get(database.Ctx, id).Result()
	return val != ""
}

func Save(id, url string, expiry time.Duration) error {
	if expiry == 0 {
		expiry = 24
	}

	r := database.CreateClient(0)
	defer r.Close()

	return r.Set(
		database.Ctx,
		id,
		url,
		expiry*3600*time.Second,
	).Err()
}

func BuildShortURL(id string) string {
	return os.Getenv("DOMAIN") + "/" + id
}
