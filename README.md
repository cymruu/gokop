package main

import (
	"fmt"

	"github.com/cymruu/gokop/v1"
)

func main() {
	client := v1.CreateWykopV1API("apikey", "secret", "connection")
	entry, err := client.Index(int64(38199789))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(entry)
	resp, err := client.AddComment(entry, "czesc")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

}
