package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())

		// Fetch the content of the current URL
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to fetch JavaScript content from %s: %s\n", url, err)
			continue
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			content, err := readResponseBody(response)
			if err != nil {
				fmt.Printf("Failed to read response body for %s: %s\n", url, err)
				continue
			}
			fmt.Printf("JavaScript content fetched for %s:\n%s\n", url, content)
		} else {
			fmt.Printf("Failed to fetch JavaScript content from %s. Status code: %d\n", url, response.StatusCode)
		}
	}
}

func readResponseBody(response *http.Response) (string, error) {
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
