package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleGet(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	cname, err := net.LookupCNAME("examplesvc.sector7.internal")
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	if strings.HasSuffix(cname, ".") {
		cname = strings.TrimSuffix(cname, ".")
	}

	endpoint := fmt.Sprintf("http://%s", cname)
	fmt.Println("get", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func makeHandler() func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return handleGet(request)
	}
}

func main() {
	handler := makeHandler()
	lambda.Start(handler)
}
