package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/joho/godotenv"
)

const (
	// Sender ...
	Sender = "sender address"
	// Recipient ...
	Recipient = "some email address"
	// Subject ...
	Subject = "Amazon SES with golang tutorial"
	// HtmlBody ...
	HtmlBody = "<b>Holaaa</b>"
	// CharSet ...s
	CharSet = "UTF-8"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AWSAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWSSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWSSessionToken := os.Getenv("AWS_SESSION_TOKEN")
	AWSRegion := os.Getenv("AWS_REGION")

	conf := &aws.Config{
		Region:      aws.String(AWSRegion),
		Credentials: credentials.NewStaticCredentials(AWSAccessKeyID, AWSSecretAccessKey, AWSSessionToken)}
	sess, err := session.NewSession(conf)
	if err != nil {
		log.Fatal("Failed create session to aws")
	}

	// Create an SES session.
	svc := ses.New(sess)

	// result, err := svc.ListIdentities(&ses.ListIdentitiesInput{IdentityType: aws.String("EmailAddress")})

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// assamble email
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	fmt.Println(AWSAccessKeyID)
	fmt.Println(AWSSecretAccessKey)
}
