package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	session := session.New()
	getIamIdentity(session)
}

func getIamIdentity(session *session.Session) {
	svc := sts.New(session)
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Caller identity: (string) ", result.String())

	data, err := os.ReadFile("/var/run/secrets/eks.amazonaws.com/serviceaccount/token")
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
	}
	fmt.Println("File content: ", string(data))
}
