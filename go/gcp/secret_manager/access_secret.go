// Sample quickstart is a basic program that uses Secret Manager.
package main

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {
	// GCP project in which to store secrets in Secret Manager.
	// Replace the <Project ID> with your Project ID
	projectID := "devops-296503"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	// Replace the <Secret Key> with you secret
	accessSecretReq := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/dev_sport_redis/versions/1", projectID),
	}

	secret, err := client.AccessSecretVersion(ctx, accessSecretReq)

	if err != nil {
		log.Fatalf("failed to access secret: %v", err)
	}

	fmt.Println(string(secret.Payload.Data))
}
