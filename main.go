package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
)

func main() {
	client := youtube.Client{}
	var url string

	isError := true
	for isError {
		fmt.Print("Enter the url: ")
		fmt.Scan(&url)
		video, err := client.GetVideo(url)
		if err != nil {
			fmt.Println("There was an error: ", err)
			continue
		}
		isError = false
		fmt.Println(video.Title)
	}

	fmt.Println("Hello, World!")
}
