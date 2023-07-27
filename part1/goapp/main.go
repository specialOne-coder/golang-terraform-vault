// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
)

func config() *vault.Client {
	config := vault.DefaultConfig()
	config.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	return client
}

func auth(token string) (client *vault.Client, err error) {
	client = config()
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	client.SetToken(token)
	return client, err
}

func userpassLogin() (string, error) {
	// to pass the password
	client := config()
	options := map[string]interface{}{
		"password": "user1",
	}
	path := fmt.Sprintf("auth/userpass/login/%s", "user1")

	// PUT call to get a token
	secret, err := client.Logical().Write(path, options)
	if err != nil {
		return "", err
	}

	token := secret.Auth.ClientToken
	return token, nil
}

func rootPlaySecret(client *vault.Client, err error) {
	secretData := map[string]interface{}{
		"username": "ferdo89",
		"password": "Hashi123",
	}

	ctx := context.Background()
	// Write a secret
	_, err = client.KVv2("secret/data/").Put(ctx, "my-secret", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	log.Println("Secret written successfully.")
	secret, err := client.KVv2("secret/data/").Get(ctx, "my-secret")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	log.Println("Secret read successfully.", secret.Data)

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	if value != "Hashi123" {
		log.Fatalf("unexpected password value %q retrieved from vault", value)
	}
}

func userPlaySecret(client *vault.Client, err error) {
	secretData := map[string]interface{}{
		"username": "solo",
		"password": "user123",
	}

	ctx := context.Background()
	// Write a secret
	_, err = client.KVv2("secret/data").Put(ctx, "user1", secretData)
	if err != nil {
		log.Fatalf("Simple user unable to write secret: %v", err)
	}

	log.Println("Secret written successfully.")
	secret, err := client.KVv2("secret/data").Get(ctx, "user1")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	log.Println("Secret read successfully.", secret.Data)

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	if value != "user123" {
		log.Fatalf("unexpected password value %q retrieved from vault", value)
	}
}

func main() {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>> Start root playing >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	rootClient, rerr := auth("hvs.EHHGWWufRijnM97ui6HSw6Sc")
	if rerr != nil {
		log.Fatalf("unable to initialize Vault client: %v", rerr)
	}
	rootPlaySecret(rootClient, rerr)
	log.Println("Access granted for root!")

	log.Println(">>>>>>>>>>>>>>>>>>>>>>>Start user playing >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	upToken, _ := userpassLogin()
	log.Println("Connected Token: ", upToken)

	userClient, uerr := auth(upToken)
	if rerr != nil {
		log.Fatalf("unable to initialize Vault client: %v", uerr)
	}
	userPlaySecret(userClient, uerr)

}
