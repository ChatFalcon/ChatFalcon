package config

import (
	"context"
	"encoding/json"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/redis"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// GetConfig is used to get the config from the cache or MongoDB.
func GetConfig() (*ServerConfig, error) {
	// Try and get the config from Redis if it is initialised.
	if redis.Client != nil {
		resp := redis.Client.Get("config")
		err := resp.Err()
		switch err {
		case nil:
			b, _ := resp.Bytes()
			var cfg ServerConfig
			err := json.Unmarshal(b, &cfg)
			if err != nil {
				logrus.Error("Failed to de-serialize the config from the cache: ", err)
			}
			return &cfg, nil
		case redis.Nil:
			break
		default:
			logrus.Error("Failed to get config from cache: ", err)
		}
	}

	// Attempt to get the config from MongoDB.
	res := mongo.DB.Collection("serverConfig").FindOne(context.TODO(), bson.M{})
	err := res.Err()
	if err != nil {
		return nil, err
	}
	var cfg ServerConfig
	err = res.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
