# viewjs

JavaScript Content Fetcher

A lightweight command-line tool written in Go for fetching JavaScript content from URLs and displaying it directly in the terminal. This tool is designed to quickly check the availability of JavaScript resources, providing information about successful fetches, failures, and HTTP status codes.

Usage:
1. Pipe URLs from a file to the tool: cat urljs.txt | go run viewjs.go

This tool does not save the fetched content to files; it focuses on efficiently checking and presenting JavaScript resource status.
