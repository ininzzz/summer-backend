package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/ininzzz/summer-backend/cache"
)

// ClearDailyOrder 每日0点，清除前一日的订单序列号
func ClearDailyOrder() error {
	yesterday := time.Now().AddDate(0, 0, -1)
	date := fmt.Sprintf("%d-%d-%d", yesterday.Year(), yesterday.Month(), yesterday.Day())
	key := "icr:order:" + date
	return cache.RedisClient.Del(context.Background(), key).Err()
}
