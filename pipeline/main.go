package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
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

func Pipeline() {
	pipe := client.Pipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(incr.Val())
}

func Pipelined() {
	var incr *redis.IntCmd

	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipelined_counter")
		pipe.Expire(ctx, "pipelined_counter", time.Hour)
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(incr.Val())
}

func Cmd() {
	client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Set(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("key%d", i), 0)
		}

		return nil
	})

	cmds, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
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
func main() {
	//Pipeline()
	//Pipelined()

	Cmd()
}
