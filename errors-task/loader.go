package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	urlArgPtr := flag.String("url", "", "File download URL")
	outputArgPtr := flag.String("output", "", "Output file name")

	flag.Parse()

	if *urlArgPtr == "" {
		printErrorAndExit("URL is not defined")
	}
	if *outputArgPtr == "" {
		printErrorAndExit("Output file name is not defined")
	}

	err := downloadFile(*urlArgPtr, *outputArgPtr)

	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUrl):
			fmt.Printf("Error: %v\nEnter correct URL\n", err)
		case errors.Is(err, ErrConnectionFailed):
			fmt.Printf("Error: %v\nCheck your connection and try again\n", err)
		case errors.Is(err, ErrDownloadFailed):
			fmt.Printf("Error: %v\nCheck file availability for download\n", err)
		case errors.Is(err, ErrFileNotFound):
			fmt.Printf("Error: %v\nCheck if the URL is correct or if the file is on the server\n", err)
		default:
			fmt.Println(err)
		}
	}
}

func printErrorAndExit(message string) {
	fmt.Println(message)
	fmt.Println("Usage of program:")
	flag.PrintDefaults()
	os.Exit(1)
}
