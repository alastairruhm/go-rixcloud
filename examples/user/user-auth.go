package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alastairruhm/go-rixcloud/rixcloud"
)

func main() {
	httpClient := &http.Client{}
	client := rixcloud.NewClient(httpClient, "username", "password")
	profile, _, err := client.Profiles.VerifyCredentials()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("userId: ", profile.UserID)
}
