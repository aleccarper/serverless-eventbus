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

func processRequest(c *gin.Context) {
	subscription := eventbus.CreateSubscription(c.PostForm("event_type"), c.PostForm("endpoint"))
	c.JSON(http.StatusCreated, subscription)
}

func main() {
	lambda.Start(Handler)
}
