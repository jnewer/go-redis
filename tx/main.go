package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func TxPipelined() {
	cmds, err := client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}
}

const maxRetries = 1000

func Increment(key string) error {
	txf := func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		n++

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})

		return err
	}

	for i := 0; i < maxRetries; i++ {
		err := client.Watch(ctx, txf, key)
		if err != nil {
			return nil
		}

		if err == redis.TxFailedErr {
			continue
		}

		return err
	}

	return errors.New("达到最大次数")

}

func main() {
	TxPipelined()
}
