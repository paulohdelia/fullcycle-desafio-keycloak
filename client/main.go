package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/coreos/go-oidc"
)

var (
	clientID = "app"

	// Temporary secret to test keycloak
	clientSecret = "f493fbb6-34e2-4ee0-82a2-9dd8c0b243e6"
)

func main() {
	ctx := context.Background()

	// the second parameter is the issuer
	// you can find the issuer with the link below
	// http://localhost:8080/auth/realms/fullcycle-demo/.well-known/openid-configuration
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/fullcycle-demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	// key word to validate the request
	// usually is dynamic and set in a ENV variable
	// in this case is hard coded to keep the test simple
	// and focus on learning to use keycloak and oauth
	state := "magica"

	// Redirect to keycloak server - login
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	// Redirect back to app
	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {

		// check if the state word is the same as expected
		// in this case the word should be "magica"
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		// get the access code
		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "failed to exchange token", http.StatusBadRequest)
			return
		}

		// get id_token
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "no_id_token", http.StatusBadRequest)
			return
		}

		resp := struct {
			OAuth2Token *oauth2.Token
			RawIDToken  string
		}{
			oauth2Token, rawIDToken,
		}

		// create json with all the info above
		/* "OAuth2Token": {
			"access_token",
			"token_type",
			"refresh_token",
			"expiry"
		}
		"RawIDToken"
		*/
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
