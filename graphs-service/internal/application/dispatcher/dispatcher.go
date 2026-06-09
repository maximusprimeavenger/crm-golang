package dispatcher

import (
	"context"
	"fmt"
	"graphs-service/internal/application/aggregator"
	"graphs-service/internal/transport/kafka-reader/models"

	"github.com/segmentio/kafka-go"
)

func Dispatcher(ctx context.Context, agg *aggregator.Aggregator, reader *kafka.Reader, ch <-chan models.Envelope) {
	for {
		select {
		case <-ctx.Done():
			return

		case env := <-ch:
			if err := agg.Apply(env.Event); err != nil {
				continue
			}
			fmt.Println("Dispatcher GOT THE MESSAGE: ", env.Event.Payload)
			_ = reader.CommitMessages(ctx, env.Msg)
			fmt.Println("Message commited")
		}
	}
}
