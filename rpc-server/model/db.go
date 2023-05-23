package model

import (
	"context"
	"encoding/json"
	"github.com/pbthang/go_messaging_app/rpc-server/kitex_gen/rpc"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var db *redis.Client
var ctx = context.Background()
var redisAddr string
var redisPassword string

func init() {
	if os.Getenv("ENV") == "PROD" {
		redisAddr = "redis:6379"
		redisPassword = ""
	} else {
		redisAddr = "localhost:6379"
		redisPassword = ""
	}
	db = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})
	_, err := db.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to redis")
	}
}

func InsertMessage(message *rpc.Message) error {
	jsonObj, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = db.RPush(ctx, message.Chat, jsonObj).Result()
	if err != nil {
		return err
	}
	return err
}

func GetMessages(chat string, cursor, limit int64, reverse bool) ([]*rpc.Message, bool, int64, error) {
	// get from db
	messageJsons, err := db.LRange(ctx, chat, 0, -1).Result()
	length := len(messageJsons)
	messages := make([]*rpc.Message, 0)
	if err != nil {
		return nil, false, 0, err
	}

	// unmarshal messages
	for _, messageInterface := range messageJsons {
		var message rpc.Message
		err = json.Unmarshal([]byte(messageInterface), &message)
		if err != nil {
			return nil, false, 0, err
		}
		messages = append(messages, &message)
	}

	// filter messages
	hasMore := true
	if length == 0 {
		hasMore = false
	}
	nextCursor := int64(0)
	if cursor >= 0 && limit >= 0 {
		var filteredMessages []*rpc.Message
		count := 0
		for i, message := range messages {
			if count >= int(limit) {
				break
			}
			if i == length-1 {
				hasMore = false
			}
			if hasMore {
				nextCursor = messages[i+1].SendTime
			}
			if message.SendTime >= cursor {
				filteredMessages = append(filteredMessages, message)
				count++
			}
		}
		messages = filteredMessages
	}

	// reverse messages
	if reverse {
		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages[i], messages[j] = messages[j], messages[i]
		}
	}

	if !hasMore {
		nextCursor = 0
	}

	return messages, hasMore, nextCursor, nil
}

// for testing purpose only

func FlushAll() {
	db.FlushAll(ctx)
}

func DeleteChat(chat string) {
	db.Del(ctx, chat)
}
