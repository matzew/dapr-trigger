package main

import (
	"context"
	"github.com/matzew/dapr-trigger/pkg/cloudevents"
	"go.uber.org/zap"
	"log"
)

func main() {
	ctx := context.Background()
	handler, err := cloudevents.NewHandler(ctx)
	err = handler.Start(ctx)
	if err != nil {
		log.Fatal("handler.Start() returned an error", zap.Error(err))
	}
}
