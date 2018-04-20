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
	// need id of acvtived service
	v, _, err := client.Profiles.GetTrafficData("xxxx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	upload := v.Upload
	fmt.Printf("upload  : %5d M\n", upload>>20)
	dowload := v.Download
	fmt.Printf("download: %5d M\n", dowload>>20)
	fmt.Printf("used    : %5d M\n", (upload+dowload)>>20)
	total := v.Total
	fmt.Printf("total   : %5d M\n", total>>20)
}
