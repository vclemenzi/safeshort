package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func CheckURL(url string) bool {
	filePath := "./banned_urls.list"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to open the file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if IsURLIn(line, url) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Unable to read the file:", err)
	}

	return false
}

func IsURLIn(s, urlStr string) bool {
    parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error during url parsing:", err)
		return false
	}

	parsedDomain := parsedURL.Hostname()
	if parsedDomain == "" {
		return false
	}

	return s == parsedDomain
}
