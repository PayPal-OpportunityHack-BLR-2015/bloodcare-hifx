package services

import "gopkg.in/redis.v3"

type Redis struct {
	Client *redis.Client
}

func NewRedis(conString string) (*Redis, error) {

	var final *Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	final = &Redis{client}

	_, err := client.Ping().Result()
	//  fmt.Print("\nRedis Ping result: ", pong)
	return final, err
}
