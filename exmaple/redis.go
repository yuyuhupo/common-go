package exmaple

import (
	"fmt"

	"github.com/yuyuhupo/common-go/redis"
)

func Redis() {
	var conf = redis.Config{
		Addr:     "localhost:6379",
		Password: "password",
		DB:       1,
	}
	var r = redis.New(conf)

	if err := r.Set("key", "value"); err != nil {
		fmt.Println(err)
	}

	var value string
	if err := r.Get("key", &value); err != nil {
		fmt.Println(err)
	}
	fmt.Println("value: ", value)
}
