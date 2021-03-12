package mail

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"gopkg.in/gomail.v2"
)

type AWSCredConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	AwsRegion       string
	SMTPUser        string
	SMTPPass        string
	Host            string
	Port            int
	SenderEmail     string
}

func getAWSCreds() *AWSCredConfig {

	port, err := strconv.Atoi(os.Getenv("AWS_PORT"))

	if err != nil {
		log.Panic(err)
	}
	return &AWSCredConfig{
		// AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsRegion:       os.Getenv("AWS_REGION"),
		SMTPUser:        os.Getenv("AWS_SMTP_USER"),
		SMTPPass:        os.Getenv("AWS_SMTP_PASS"),

		Host:        os.Getenv("AWS_HOST"),
		Port:        port,
		SenderEmail: os.Getenv("AWS_SENDER_EMAIL"),
	}

}

const (
	// Replace AccessKeyID with your AccessKeyID key.
	AccessKeyID = "AKIATY2HYWVHIEPURU5A"
)

func SendOTP(token, email, phone string) error {

	awsCreds := getAWSCreds()

	if len(phone) > 0 {
		sendSMS(awsCreds, phone, "Your one Time Password is "+token)
	}

	sendEmail(awsCreds, "Sajan", email, "One Time Password ", "<p>535 Authentication Credentials Invalid</p>", "Your one Time Password Is  "+token, "aws,ses")

	return nil

}

func sendEmail(awsCreds *AWSCredConfig, senderName, recipient, subject, htmlBody, textBody, tags string) {

	m := gomail.NewMessage()

	// Set the main email part to use HTML.
	m.SetBody("text/html", htmlBody)

	// Set the alternative part to plain text.
	m.AddAlternative("text/plain", textBody)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":      {m.FormatAddress(awsCreds.SenderEmail, senderName)},
		"To":        {recipient},
		"Subject":   {subject},
		"something": {tags},
	})

	// Send the email.
	d := gomail.NewPlainDialer(awsCreds.Host, awsCreds.Port, awsCreds.SMTPUser, awsCreds.SMTPPass)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to ", recipient)
	}

}

func sendSMS(awsCreds *AWSCredConfig, phoneNumber string, message string) error {
	// log.Println("i was here")

	// creds := strings.TrimSpace(awsCreds.AccessKeyID)
	// Create Session and assign AccessKeyID and SecretAccessKey
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsCreds.AwsRegion),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, awsCreds.SecretAccessKey, ""),
	},
	)

	// Create SNS service
	svc := sns.New(sess)

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message:     aws.String(message),
	}

	timer1 := time.NewTimer(2 * time.Second)

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svc.Publish(params)

	if err != nil {
		log.Println(err.Error())
	}
	<-timer1.C
	log.Println(resp) // print the response data.

	log.Println("OTP sent to", phoneNumber)

	return nil
}
