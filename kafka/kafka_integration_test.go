// +build integration

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kafka

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/intelsdi-x/snap/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/Shopify/sarama.v1"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// integration test
func TestPublish(t *testing.T) {
	// This integration test requires KAFKA_BROKERS
	if os.Getenv("SNAP_TEST_KAFKA") == "" {
		fmt.Println("Skipping integration")
		return
	}
	brokers := os.Getenv("SNAP_TEST_KAFKA")

	// Pick a random topic
	topic := fmt.Sprintf("%d", time.Now().Nanosecond())
	fmt.Printf("Topic: %s\n", topic)
	config := make(map[string]ctypes.ConfigValue)
	Convey("Publish to Kafka", t, func() {
		Convey("publish and consume", func() {
			k := NewKafkaPublisher()

			// Build some config
			config["brokers"] = ctypes.ConfigValueStr{Value: brokers}
			config["topic"] = ctypes.ConfigValueStr{Value: topic}

			// Get validated policy
			cp, _ := k.GetConfigPolicy()
			cfg, _ := cp.Get([]string{""}).Process(config)

			t := time.Now().String()

			// Send data to create topic. There is a weird bug where first message won't be consumed
			// ref: http://mail-archives.apache.org/mod_mbox/kafka-users/201411.mbox/%3CCAHwHRrVmwyJg-1eyULkzwCUOXALuRA6BqcDV-ffSjEQ+tmT7dw@mail.gmail.com%3E
			k.Publish("", []byte(t), *cfg)
			// Send the same message. This will be consumed.
			k.Publish("", []byte(t), *cfg)

			// start timer and wait
			m, _ := consumer(topic, brokers)

			Convey("So we should receive what we published to Kafka", func() {
				So(string(m.Value), ShouldEqual, t)
			})
		})
	})
}

func consumer(topic, brokers string) (*sarama.ConsumerMessage, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true // Handle errors manually instead of letting Sarama log them.

	master, err := sarama.NewConsumer([]string{brokers}, config)
	if err != nil {
		return nil, err
	}
	defer func() {
		master.Close()
	}()

	consumer, err := master.ConsumePartition(topic, 0, 0)
	if err != nil {
		return nil, err
	}
	defer func() {
		consumer.Close()
	}()

	msgCount := 0
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Millisecond * 500)
		timeout <- true
	}()

	select {
	case err := <-consumer.Errors():
		return nil, err
	case m := <-consumer.Messages():
		msgCount++
		return m, nil
	case <-timeout:
		return nil, errors.New("timed out waiting for produced message")
	}
}
