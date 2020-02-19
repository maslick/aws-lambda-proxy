package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"net/http"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return httpadapter.New(initRouter()).ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

func initRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/path1", handler1)
	router.HandleFunc("/path2", handler2)
	return router
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("one"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("two"))
}
