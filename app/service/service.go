package service

import (
	"Grpc-crud/app/models"
	"Grpc-crud/app/pb"
	er "Grpc-crud/app/utils/errors"
	res "Grpc-crud/app/utils/response"
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
) 

func connDynamo()(*session.Session, error){
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		log.Fatalf("Error in aws connection  %v", err)
		return nil, err
	}
	return sess, nil
}



func WriteRequest(b models.Binder)(*res.ResRespo, *er.ResErr) {
	se, err := connDynamo()
	if err != nil {
		log.Fatalf("Error in aws session err %v", err)
		return nil, er.InternalServerError("Error in aws session err")
	}
    svc := dynamodb.New(se)
	av, err := dynamodbattribute.MarshalMap(b)
	if err != nil {
        log.Fatalf("Got error marshalling map: %v", err)
       return nil, er.InternalServerError("Got error marshalling map:")
   
    } 
	input := &dynamodb.PutItemInput{
        Item: av,
        TableName: aws.String("blocks-tables"),
    }
	_, err = svc.PutItem(input)

    if err != nil {
		log.Fatalf("Got error calling PutItem: %v", err)
		return nil, er.InternalServerError("Got error calling inserting: ")
    }
	return res.ReuqstAccepted("insert the data into the Dynanmo table"), nil
}