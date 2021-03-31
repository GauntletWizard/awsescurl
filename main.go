package main

/* Copied from https://github.com/awsdocs/amazon-elasticsearch-service-developer-guide/blob/master/doc_source/es-request-signing.md
 * See LICENSE-SAMPLECODE for license.
 */

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"net/http"
	"strings"
	"time"
)

func main() {

	// Basic information for the Amazon Elasticsearch Service domain
	domain := "" // e.g. https://my-domain.region.es.amazonaws.com
	index := "my-index"
	id := "1"
	endpoint := domain + "/" + index + "/" + "_doc" + "/" + id
	region := "" // e.g. us-east-1
	service := "es"

	// Sample JSON document to be included as the request body
	json := `{ "title": "Thor: Ragnarok", "director": "Taika Waititi", "year": "2017" }`
	body := strings.NewReader(json)

	// Get credentials from environment variables and create the AWS Signature Version 4 signer
	credentials := credentials.NewEnvCredentials()
	signer := v4.NewSigner(credentials)

	// An HTTP client for sending the request
	client := &http.Client{}

	// Form the HTTP request
	req, err := http.NewRequest(http.MethodPut, endpoint, body)
	if err != nil {
		fmt.Print(err)
	}

	// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	signer.Sign(req, body, service, region, time.Now())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp.Status + "\n")
}
