package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/compute/metadata"
)

func main() {
	fmt.Println("Start WLI")

	ctx := context.Background()

	for {
		time.Sleep(3 * time.Second)

		{
			token, err := getMetadata(ctx, "/computeMetadata/v1/instance/service-accounts/default/token")
			if err != nil {
				fmt.Printf("failed getServiceAccountToken : %s\n", err)
			}
			fmt.Printf("SAToken : %s\n", token)
		}
		{
			saEmail, err := getMetadata(ctx, "/computeMetadata/v1/instance/service-accounts/default/email")
			if err != nil {
				fmt.Printf("failed getServiceAccountToken : %s\n", err)
			}
			fmt.Printf("Email from Token : %s\n", saEmail)
		}
		{
			saEmail, err := metadata.Email("")
			if err != nil {
				log.Fatal(ctx, err.Error())
			}
			fmt.Printf("I am %s\n", saEmail)
		}
		fmt.Println()
	}
}

func getMetadata(ctx context.Context, metadata string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://169.254.169.254%s", metadata), nil)
	if err != nil {
		return "", fmt.Errorf("failed getServiceAccountToken http.NewRequest : %w", err)
	}
	req.Header.Set("Metadata-Flavor", "Google")
	req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed getServiceAccountToken http.Do : %w", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed getServiceAccountToken ioutil.ReadAll : %w", err)
	}
	return string(body), nil
}
