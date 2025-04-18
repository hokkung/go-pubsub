package main

import (
	"context"
	"encoding/json"
	"fmt"
	go_pubsub "github.com/hokkung/go-pubsub"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ProductTopic string `envconfig:"product_topic" required:"true"`
	ProductSubID string `envconfig:"product_sub_id" required:"true"`
}

func NewConfig() *Config {
	var c Config
	if err := envconfig.Process("app", &c); err != nil {
		panic(err)
	}
	return &c
}

type Product struct {
	Name string `json:"name"`
}

type ProductPublisher struct {
	topicName string
	publisher go_pubsub.Publisher
}

func NewProductPublisher(publisher go_pubsub.Publisher, topic string) *ProductPublisher {
	return &ProductPublisher{
		topicName: topic,
		publisher: publisher,
	}
}

func (e *ProductPublisher) Publish(ctx context.Context, data *Product, attrs map[string]string) (*go_pubsub.PublishResponse, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return e.publisher.Publish(ctx, &go_pubsub.PublishRequest{
		Topic:      e.topicName,
		Data:       bytes,
		Attributes: attrs,
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	ctx := context.Background()
	appConfig := NewConfig()
	pubsubConfig := go_pubsub.NewConfig()
	client, err := go_pubsub.NewClient(ctx, pubsubConfig)
	if err != nil {
		panic(err)
	}

	publisher := go_pubsub.NewPublisher(client)
	productPub := NewProductPublisher(publisher, appConfig.ProductTopic)
	productSub := go_pubsub.NewSubscriber(client)

	go func() {
		ctx := context.Background()
		res, err := productPub.Publish(ctx, &Product{
			Name: "Product2",
		}, map[string]string{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.ID)
	}()
	
	productSub.Subscribe(ctx, appConfig.ProductSubID, func(ctx context.Context, data []byte, attributes map[string]string) error {
		var product Product
		err := json.Unmarshal(data, &product)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(product)
		fmt.Println(attributes)

		return nil
	})
}
