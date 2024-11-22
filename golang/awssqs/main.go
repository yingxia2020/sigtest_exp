package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// Create a new AWS session using default credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new SQS service client
	svc := sqs.New(sess, aws.NewConfig().WithRegion("us-east-2"))

	// Create a new queue
	queueName := "my-queue"
	createParams := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}
	start := time.Now()
	createResp, err := svc.CreateQueueWithContext(context.Background(), createParams)
	end := time.Now()
	fmt.Printf("Create Queue Elapsed time: %s\n", end.Sub(start))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*createResp.QueueUrl)
	//URL := "https://sqs.us-east-2.amazonaws.com/200891666809/xyqueue.fifo"
	queueUrl := createResp.QueueUrl

	// Send a message to the queue
	for i:= 0; i<5; i++ {
		msg := fmt.Sprintf("Hello world %d!", i)
		sendParams := &sqs.SendMessageInput{
			QueueUrl:    queueUrl,
			MessageBody: aws.String(msg),
//			MessageGroupId: aws.String("xygroup"),
		}
		start = time.Now()
		_, err = svc.SendMessageWithContext(context.Background(), sendParams)
		end = time.Now()
		fmt.Printf("Send Message Elapsed time: %s\n", end.Sub(start))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for i:=0; i<5; i++ {
		go func() {
			start = time.Now()
			// Receive messages from the queue
			receiveParams := &sqs.ReceiveMessageInput{
				QueueUrl:            queueUrl,
				MaxNumberOfMessages: aws.Int64(1),
				WaitTimeSeconds:     aws.Int64(10),
			}

			receiveResp, err := svc.ReceiveMessageWithContext(context.Background(), receiveParams)
			end = time.Now()
			fmt.Printf("Receive Message Elapsed time: %s\n", end.Sub(start))
			if err != nil {
				fmt.Println(err)
				return
			}

			// Print the received messages
			for _, msg := range receiveResp.Messages {
				fmt.Println(*msg.Body)
			}

			deleteParams := &sqs.DeleteMessageInput{
				QueueUrl:            queueUrl,
				ReceiptHandle:       receiveResp.Messages[0].ReceiptHandle,
			}
			start = time.Now()
			_, _ = svc.DeleteMessageWithContext(context.Background(), deleteParams)
			end = time.Now()
			fmt.Printf("Delete Message Elapsed time: %s\n", end.Sub(start))
		}()
	}

	time.Sleep(1 * time.Second)

	// Delete the queue
	deleteParams := &sqs.DeleteQueueInput{
		QueueUrl: queueUrl,
	}

	start = time.Now()
	_, err = svc.DeleteQueueWithContext(context.Background(), deleteParams)
	end = time.Now()
	fmt.Printf("Delete Queue Elapsed time: %s\n", end.Sub(start))
	if err != nil {
		fmt.Println(err)
		return
	}
}
