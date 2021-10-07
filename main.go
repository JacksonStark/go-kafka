package main

import (
	"context"
	"fmt"
	"strconv"
	// "log"
	// "os"
	"time"

	"github.com/segmentio/kafka-go"
)

// init topic and broker address constants
const (
	topic          = "my-kafka-topic"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)

func main() {
	// create context
	ctx := context.Background()

	// produce messages in new goroutine, since both produce & consume are blocking
	go produce(ctx)
	consume(ctx)
}

func produce(ctx context.Context) {
	// init a counter
	i := 0

	// logger for more granular clarity
	// l := log.New(os.Stdout, "kafka writer: ", 0)

	// init writer w/ broker addresses & topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
		// Logger: l,
	})

	for {
		// each kafka message has key:val
		err := w.WriteMessages(ctx, kafka.Message{
			// key decides which partition (therefore which broker) message gets published on
			Key: []byte(strconv.Itoa(i)),
			// arbitrary message payload for value
			Value: []byte("this is message " + strconv.Itoa(i)),
		})

		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log successful write
		fmt.Println("writes:", i)

		i++
		// sleep for quick sec
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	// logger for more granular clarity
	// l := log.New(os.Stdout, "kafka reader: ", 0)

	// init reader w/ broker addresses, topic, etc
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker1Address, broker2Address, broker3Address},
		Topic:    topic,
		// group consumers to prevent duplicate messages being received
		GroupID:  "my-kafka-group",
		// min batch size required for data to be received
		MinBytes: 50, 
		// MaxBytes required if MinBytes set
		MaxBytes: 1e6,
		// wait for 3s max before receiving new data, regardless of MinBytes threshold
		MaxWait: 3 * time.Second,
		// Logger: l,
	})

	for {
		// blocks until next event received
		msg, err := r.ReadMessage(ctx)

		if err != nil {
			panic("could not read message " + err.Error())
		}

		// log successfully received message
		fmt.Println("received: ", string(msg.Value))
	}
}