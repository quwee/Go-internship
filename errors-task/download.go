package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	ErrInvalidUrl       = errors.New("invalid URL")
	ErrConnectionFailed = errors.New("connection failed")
	ErrDownloadFailed   = errors.New("download failed")
	ErrFileNotFound     = errors.New("file not found")
)

func downloadFile(fileUrl, fileName string) error {
	// validate url
	parsedUrl, err := url.Parse(fileUrl)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidUrl, err)
	}

	if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return fmt.Errorf("%w: incorrect host or scheme", ErrInvalidUrl)
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileUrl, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	// perform request
	var httpClient http.Client

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConnectionFailed, err)
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("%w: %w", ErrFileNotFound, err)
		}
		return fmt.Errorf("%w: %w", ErrDownloadFailed, err)
	}

	// load file
	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("could not copy file: %w", err)
	}

	return nil
}
