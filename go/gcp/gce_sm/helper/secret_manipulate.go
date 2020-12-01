package helper

import (
	"context"
	"fmt"
	"log"
	"math"
	"runtime"
	"sync"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// CreateSecret is to create individual secret
func CreateSecret(ctx context.Context, client *secretmanager.Client, secret Secret, project string) {
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", project),
		SecretId: secret.Name,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
			Labels: secret.Labels,
		},
	}

	createSecretResp, err := client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		log.Printf("Failed to create secret: %v\n", err)
	}

	payload := []byte(secret.Value)

	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: createSecretResp.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}

	// Call the API.
	_, err = client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Printf("failed to add secret version: %v\n", err)
	}
}

// BulkCreateSecrets can create multiple secretes in one shot
func BulkCreateSecrets(ctx context.Context, client *secretmanager.Client, wg *sync.WaitGroup, secretFile SecretFile) {
	cpuNo := runtime.NumCPU()
	secretNo := len(secretFile.Secrets)
	gru := int(math.Min(float64(cpuNo), float64(secretNo)))
	secretChan := make(chan Secret, secretNo)

	for _, s := range secretFile.Secrets {
		secretChan <- s
	}

	close(secretChan)

	wg.Add(gru)

	for g := 0; g < gru; g++ {
		go func() {
			defer wg.Done()

			for secret := range secretChan {
				CreateSecret(ctx, client, secret, secretFile.Project)
			}
		}()
	}
}

// DeleteSecret is to create individual secret
func DeleteSecret(ctx context.Context, client *secretmanager.Client, secret Secret, project string) {
	deleteSecretReq := &secretmanagerpb.DeleteSecretRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s", project, secret.Name),
	}

	err := client.DeleteSecret(ctx, deleteSecretReq)
	if err != nil {
		log.Printf("Failed to create secret: %v\n", err)
	}
}

// BulkDeleteSecrets can Delete multiple secretes in one shot
func BulkDeleteSecrets(ctx context.Context, client *secretmanager.Client, wg *sync.WaitGroup, secretFile SecretFile) {
	cpuNo := runtime.NumCPU()
	secretNo := len(secretFile.Secrets)
	gru := int(math.Min(float64(cpuNo), float64(secretNo)))
	secretChan := make(chan Secret, secretNo)

	for _, s := range secretFile.Secrets {
		secretChan <- s
	}

	close(secretChan)

	wg.Add(gru)

	for g := 0; g < gru; g++ {
		go func() {
			defer wg.Done()

			for secret := range secretChan {
				DeleteSecret(ctx, client, secret, secretFile.Project)
			}
		}()
	}
}
