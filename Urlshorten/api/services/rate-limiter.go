package services

import (
	"os"
	"strconv"
	"time"

	"github.com/ishansaini194/Projects/database"
	"github.com/redis/go-redis/v9"
)

func CheckRateLimit(ip string) (remaining int, reset int, err error) {
	r := database.CreateClient(1)
	defer r.Close()

	quota, _ := strconv.Atoi(os.Getenv("API_QUOTA"))

	val, err := r.Get(database.Ctx, ip).Result()
	if err == redis.Nil {
		_ = r.Set(database.Ctx, ip, quota, 30*time.Second).Err()
		return quota - 1, 30, nil
	}

	count, _ := strconv.Atoi(val)
	if count <= 0 {
		ttl, _ := r.TTL(database.Ctx, ip).Result()
		return 0, int(ttl.Seconds()), redis.Nil
	}

	r.Decr(database.Ctx, ip)
	ttl, _ := r.TTL(database.Ctx, ip).Result()

	return count - 1, int(ttl.Seconds()), nil
}
