package main

import (
	"net/http"

	"github.com/aleccarper/serverless-eventbus/eventbus"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		ginEngine := eventbus.MountAuthorizedRoute("/subscriptions", "post", processRequest)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}
	return ginLambda.Proxy(req)
}

type Input struct {
	EventType string `form:"event_type" json:"event_type" binding:"required"`
	Endpoint  string `form:"endpoint" json:"endpoint" binding:"required"`
}

func processRequest(c *gin.Context) {
	var input Input
	c.BindJSON(&input)
	subscription := eventbus.CreateSubscription(input.EventType, input.Endpoint)
	c.JSON(http.StatusCreated, subscription)
}

func main() {
	lambda.Start(Handler)
}
