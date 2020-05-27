package redis

import "github.com/go-redis/redis"

// Client defines the Redis client.
var Client *redis.Client

// Nil is used to define the Redis error so we don't have to worry about importing both.
var Nil = redis.Nil

// CreateRedisClient is used to create the Redis client.
func CreateRedisClient(Hostname, Password string) error {
	Client = redis.NewClient(&redis.Options{
		Addr: Hostname, Password: Password,
	})
	return Client.Ping().Err()
}
