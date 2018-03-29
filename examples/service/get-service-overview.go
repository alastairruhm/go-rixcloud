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
	v, _, err := client.Profiles.GetServiceOverview()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("total   service: ", v.TotalService)
	fmt.Println("actived service: ", v.ActiveCount)
}
