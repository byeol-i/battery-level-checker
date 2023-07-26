package consumer

// returns device's id list
// func (c *Consumer) GetUserDevice(uid string) ([]string, error) {
// 	saramaConfig := c.kafkaConf
// 	saramaConfig.Consumer.Return.Errors = true

// 	client, err := sarama.NewClient(c.brokerList, saramaConfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer client.Close()

// 	consumer, err := sarama.NewConsumerFromClient(client)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer consumer.Close()

// 	// No idea
// 	partitions, err := client.Partitions(uid)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var recentPartition int32
// 	var recentOffset int64

// 	for _, partition := range partitions {
// 		offset, err := client.GetOffset(uid, partition, sarama.OffsetNewest)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if offset > recentOffset {
// 			recentPartition = partition
// 			recentOffset = offset
// 		}
// 	}

// 	pc, err := consumer.ConsumePartition(uid, recentPartition, recentOffset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer pc.Close()

// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, os.Interrupt)

// 	result := []string{}
// 	ConsumerLoop:
// 	for {
// 		select {
// 		case msg := <-pc.Messages():
// 			result = append(result, string(msg.Value))
// 			fmt.Println("Received message:", string(msg.Value))
// 			break ConsumerLoop
// 		case err := <-pc.Errors():
// 			return nil, err
// 		case <-signals:
// 			break ConsumerLoop
// 		}
// 	}

// 	return result, nil
// }