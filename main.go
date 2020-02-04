package main

import (
	"fmt"
	//"github.com/Shopify/sarama"
	//"os"
	//"os/signal"
	"github.com/landoop/schema-registry"
	//"github.com/linkedin/goavro/v2"
)

//func main() {
//
//	config := sarama.NewConfig()
//	config.Consumer.Return.Errors = true
//
//	// Specify brokers address. This is default one
//	brokers := []string{"localhost:9092"}
//
//	// Create new consumer
//	master, err := sarama.NewConsumer(brokers, config)
//	if err != nil {
//		panic(err)
//	}
//
//	defer func() {
//		if err := master.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//	topic := "new-topic"
//	// How to decide partition, is it fixed value...?
//	consumer, err := master.ConsumePartition(topic, 2, sarama.OffsetOldest)
//	if err != nil {
//		panic(err)
//	}
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, os.Interrupt)
//
//	// Count how many message processed
//	msgCount := 0
//
//	// Get signnal for finish
//	doneCh := make(chan struct{})
//	go func() {
//		for {
//			select {
//			case err := <-consumer.Errors():
//				fmt.Println(err)
//			case msg := <-consumer.Messages():
//				msgCount++
//				fmt.Println("Received messages", string(msg.Key), string(msg.Value), msg.Headers)
//			case <-signals:
//				fmt.Println("Interrupt is detected")
//				doneCh <- struct{}{}
//			}
//		}
//	}()
//
//	<-doneCh
//	fmt.Println("Processed", msgCount, "messages")
//}

//type Key struct {
//	Id int64 `json:"id"`

func main() {
	//client, _ := schemaregistry.NewClient("http://capp-schemaregistry.dev-mytaxi.com:8081")
	client, _ := schemaregistry.NewClient("localhost:8081")
	//schema, _ := client.Subjects()
	//schema, _ := client.GetLatestSchema("trackdriver")
	//schema, _ := client.GetLatestSchema("com.pickme.events.Key")
	//schema, _ := client.GetSchemaByID(5)
	//schema, _ := client.GetSchemaById(207) deprecated
	//	schema, _ := client.IsLatestSchemaCompatible("com.pickme.events.Key", `{
	//  "type": "record",
	//  "name": "Key",
	//  "namespace": "com.pickme.events",
	//  "fields": [
	//    {
	//      "name": "key",
	//      "type": "long"
	//    }
	//  ]
	//}`)
	//	bo, schema, _ := client.IsRegistered("com.pickme.events.Key", `{
	//  "type": "record",
	//  "name": "Key",
	//  "namespace": "com.pickme.events",
	//  "fields": [
	//    {
	//      "name": "key",
	//      "type": "long"
	//    }
	//  ]
	//}`)

	schema, _ := client.Versions("telecom_italia_data-value")
	//schema, _ := client.GetConfig("com.pickme.events.Key")

	fmt.Println(schema)
	//fmt.Println(bo)
	//codec,_:= goavro.NewCodec(schema.Schema)
	//val := `{"key":2}`
	//valByteArr := []byte(val)
	//native,_,_:= codec.NativeFromTextual(valByteArr)
	//fmt.Println(native.(map[string]interface{})["key"])
}
