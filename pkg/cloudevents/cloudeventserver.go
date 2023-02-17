package cloudevents

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/client"
	"strconv"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	dapr "github.com/dapr/go-sdk/client"
)

// Handler is the HTTP handler for the registry.
type Handler struct {
	client client.Client
	dapr   dapr.Client
}

// NewHandler creates a new registry handler.
func NewHandler(ctx context.Context) (*Handler, error) {

	// Create a new client for Dapr using the SDK
	daprClient, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	//	defer daprClient.Close()

	client, err := cloudevents.NewClientHTTP()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &Handler{
		client: client,
		dapr:   daprClient,
	}, nil
}

func (h *Handler) Start(ctx context.Context) error {
	return h.client.StartReceiver(ctx, h.receive)
}

func (h *Handler) receive(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	fmt.Printf("☁️  cloudevents.Event\n%s", event)

	// hook me up
	// Publish events using Dapr pubsub
	for i := 1; i <= 10; i++ {
		order := `{"orderId":` + strconv.Itoa(i) + `}`

		err := h.dapr.PublishEvent(context.Background(), "orderpubsub", "orders", []byte(order))
		if err != nil {
			panic(err)
		}

		fmt.Println("Published data:", order)

		time.Sleep(1000)
	}

	return nil, nil
}
