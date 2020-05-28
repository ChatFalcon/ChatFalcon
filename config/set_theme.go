package config

import (
	"context"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/redis"
	"go.mongodb.org/mongo-driver/bson"
)

// SetTheme is used to set the current theme which is being used.
func (c *ServerConfig) SetTheme(Theme string) (err error) {
	_, err = mongo.DB.Collection("serverConfig").UpdateOne(context.TODO(), bson.M{}, bson.M{"$set": bson.M{
		"currentTheme": Theme,
	}})
	if err != nil {
		return
	}
	if redis.Client != nil {
		err = redis.Client.Del("config", "theme_html").Err()
		if err != nil {
			return
		}
	}
	c.CurrentTheme = Theme
	return
}
