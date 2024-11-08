package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

// Define the rabbit trigger event struct to be received
type StructEvent struct {
	Data            *RabbitMQData `json:"data"`
	Id              *string       `json:"id"`
	Source          *string       `json:"source"`
	SpecVersion     *string       `json:"specversion"`
	Type            *string       `json:"type"`
	DataContentType *string       `json:"datacontenttype"`
	Time            *string       `json:"time"`
	Subject         *string       `json:"subject"`
	AliyunAccountId *string       `json:"aliyunaccountid"`
}

type RabbitMQData struct {
	Envelope *RabbitEnvelope `json:"envelope"`
	Props    *RabbitProps    `json:"props"`
	Body     *string         `json:"body"`
}

type RabbitEnvelope struct {
	DeliveryTag *int    `json:"deliveryTag"`
	Redeliver   *bool   `json:"redeliver"`
	Exchange    *string `json:"exchange"`
	RoutingKey  *string `json:"routingKey"`
}

type RabbitProps struct {
	MessageId *string `json:"messageId"`
}

func HandleRequest(ctx context.Context, events []StructEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "Get fccontext fail.", nil
	}
	flog := fctx.GetLogger()

	for _, evt := range events {
		// The trigger event data is in the `Data` json object from the json array
		flog.Info("event id:", *evt.Id)
		flog.Info("body of rabbitmq messageid :", *evt.Data.Body)
	}

	return fmt.Sprintf("Receive RabbitMQ Trigger Event: %v", events), nil
}

func main() {
	fc.Start(HandleRequest)
}
