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
		ginEngine := eventbus.MountAuthorizedRoute("/events", "post", processRequest)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}
	return ginLambda.Proxy(req)
}

func processRequest(c *gin.Context) {
	event := eventbus.CreateEvent(c.PostForm("event_type"), c.PostForm("payload"))
	c.JSON(http.StatusCreated, event)
}

func main() {
	lambda.Start(Handler)
}
