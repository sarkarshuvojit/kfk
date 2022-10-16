package kfkconn

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sarkarshuvojit/kfk/util"
	"github.com/segmentio/kafka-go"
)

func GetMessages(broker string, topic string) ([]kafka.Message) {
    currentTimestamp := util.GetCurrentTimeStamp()
    config := kafka.ReaderConfig{
		Brokers:         []string{broker},
		GroupID:         fmt.Sprintf("kfk-group-%s", currentTimestamp),
        GroupTopics:     []string{topic},
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

    processedMessages := 0
    maxOffSet := 0

    messages := []kafka.Message{}


	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}

        processedMessages++
        messages = append(messages, m)

		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}

        maxOffSet = int(reader.Stats().Offset)

        if maxOffSet == processedMessages {
            break
        }
	}

    return messages
}

