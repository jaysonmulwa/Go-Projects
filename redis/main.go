package main

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
    client := redis.NewClient(&redis.Options{
		Addr: "192.168.99.100:6379",
		Password: "",
		DB: 0,
    })

    /*pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)*/

    
    json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
    if err != nil {
        fmt.Println(err)
    }

    err = client.Set(ctx, "id1234", json, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
    val, err := client.Get(ctx, "id1234").Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(val)
}


///Docker 
//$ docker pull redis
//$ docker run --name redis-test-instance -p 6379:6379 -d redis

///https://tutorialedge.net/golang/go-redis-tutorial/
///https://github.com/go-redis/redis

