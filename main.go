package main

import "github.com/wtrep/shopify-backend-challenge-auth/auth"

func main() {
	// TODO check environment variables
	auth.SetupAndServeRoutes()
}
