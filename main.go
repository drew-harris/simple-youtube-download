package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
)

func main() {
	client := youtube.Client{}
	var url string

	fmt.Print("Enter the url: ")
	fmt.Scan(&url)
	video, err := client.GetVideo(url)
	if err != nil {
		fmt.Println("There was an error: ", err)
	}

	fmt.Println("Title: ", video.Title)
	fmt.Println("Description: ", video.Description)
	fmt.Println("Duration: ", video.Duration)

	formats := video.Formats.WithAudioChannels()

	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, World!")
}
