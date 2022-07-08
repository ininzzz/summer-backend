package tasks

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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

//清除当前文件夹中的./imgs文件夹下的本地图片
//因为上传到COS的方法目前要先转存到本地的imgs文件夹下，若以后改进上传到COS的方法，可以不必定期清空imgs文件夹
func ClearDailyImgs() error {
	dir, _ := ioutil.ReadDir("./imgs")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"./imgs", d.Name()}...))
	}
	return nil
}
