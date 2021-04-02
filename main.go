package main

/* Copied from https://github.com/awsdocs/amazon-elasticsearch-service-developer-guide/blob/master/doc_source/es-request-signing.md
 * See LICENSE-SAMPLECODE for license.
 */

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

var (
	method  = flag.String("X", "GET", "Method (from curl's -X)")
	bodyloc = flag.String("f", "-", "File to read (defaults to stdin)")
)

func main() {

	flag.Parse()

	// Basic information for the Amazon Elasticsearch Service domain
	endpoint := flag.Arg(0)
	service := "es"

	// Read in the request body
	var body io.ReadSeeker
	var err error
	if *bodyloc == "-" {
		log.Fatal("Unimplemented, you must specify a file")
	} else {
		body, err = os.Open(*bodyloc)
	}
	if err != nil {
		log.Fatalf("Failed to read request body %s", err)
	}

	// Get credentials from environment variables and create the AWS Signature Version 4 signer
	config, err := session.NewSession()
	if err != nil {
		log.Fatalf("Error getting AWS Credentials %s", err)
	}
	creds := config.Config.Credentials
	region := config.Config.Region

	signer := v4.NewSigner(creds)

	// An HTTP client for sending the request
	client := &http.Client{}

	// Form the HTTP request
	req, err := http.NewRequest(*method, endpoint, body)
	if err != nil {
		fmt.Print(err)
	}

	// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	signer.Sign(req, body, service, *region, time.Now())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp.Status + "\n")
	io.Copy(os.Stdout, resp.Body)
}
