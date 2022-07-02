package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

//kp = key prefix
const (
	TestSetKey  = "TestSetKey"
	TestHSetKey = "TestHSetKey"
	BlogKP      = "blog:"       //blog:id  //HMSet，存的是一个blog的全部信息
	JwtKP       = "Jwt:"        //Jwt:jwt  //HMSet，存的是使用当前jwt的用户的信息
	CodePhoneKP = "code:phone:" //code:phone:phoneNumber  //set，存的是手机号的验证码
	IncrKP      = "icr:"        //icr:业务类型:date  //生成全局唯一id的序列号，每日更新序列号从0开始
	BlogLikesKP = "likes:blog:"
)

type TestStuct struct {
	Id       int    `json:"id"`
	Password string `json:"pwd"`
}

var RedisClient *redis.Client

//结构体转map
func struct2map(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

//结构体转为map存入redis
func RedisHMSetStruct(ctx context.Context, key string, value interface{}) error { //exp为0表示永不过期
	res := struct2map(value)
	err := RedisClient.HMSet(ctx, key, res).Err()
	if err != nil {
		return err
	}
	return nil

	//example:
	// data := TestStuct{
	// 	1,
	// 	"中文",
	// }
	// err_set := RedisHMSetStruct(ctx, TestSetKey, data)
	// if err_set != nil {
	// 	fmt.Printf("err.Error(): %v\n", err_set.Error())
	// }
}

const (
	epoch          = int64(1656726415) // 设置起始时间(时间戳) 2022-07-02 09:47:34
	timestampBits  = uint(32)          // 时间戳占用位数
	sequenceBits   = uint(32)          // 序列所占的位数
	timestampShift = sequenceBits      // 时间戳左移位数
)

func getId(ctx context.Context, kind string) int64 {
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	timestamp := now.Unix()
	t := timestamp - epoch
	key := IncrKP + kind + ":" + date
	sequence, _ := RedisClient.Incr(ctx, key).Result()
	r := int64((t)<<timestampShift | (sequence))
	fmt.Printf("r: %v\n", r)
	return r

	//example:
	// for i := 0; i < 1000; i++ {
	// 	getId(ctx, "blog")
	// }
}
func main() {
	//初始化
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisClient = rdb
	data := TestStuct{
		1,
		"中文",
	}
	err_HMset := RedisHMSetStruct(ctx, TestHSetKey, data)
	if err_HMset != nil {
		fmt.Printf("err.Error(): %v\n", err_HMset.Error())
	}

	// hgetall参数：ctx,key
	hmap, err := rdb.HGetAll(ctx, TestHSetKey).Result()
	if err != nil {
		fmt.Printf("hmgetall error: %v\n", err)
		return
	}
	fmt.Printf("hmgetall:  %v\n", hmap) //map[string]string
	//构造map[string]interface{}
	mmp := make(map[string]interface{})
	for k, v := range hmap {
		mmp[k] = v
		fmt.Printf("k: %v\n", k)

	}
	fmt.Printf("mmp[id] type: %T\n", mmp["id"])
	mmp["id"], _ = strconv.ParseUint(mmp["id"].(string), 10, 64)
	fmt.Printf("mmp[id] type: %T\n", mmp["id"])
	fmt.Printf("mmp: %v\n", mmp)
	//手动构造结构体
	data3 := TestStuct{}
	data3.Id, _ = strconv.Atoi(hmap["id"])
	data3.Password = hmap["pwd"]
	fmt.Printf("data3: %v\n", data3)
}

// // 延长过期时间
// func ExpireRedis(key string, t int) bool {
// 	expire := time.Duration(t) * time.Second
// 	if err := Redis.Expire(ctx, key, expire).Err(); err != nil {
// 		fmt.Println(err)
// 		return false
// 	}
// 	return true
// }
