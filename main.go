package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TestResult struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusCodes        map[int]int
	TotalTime          time.Duration
}

func main() {
	// Parse command line arguments
	url := flag.String("url", "", "URL of the service to test")
	requests := flag.Int("requests", 0, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Number of concurrent requests")
	flag.Parse()

	if *url == "" || *requests == 0 {
		fmt.Println("Error: --url and --requests are required parameters")
		flag.Usage()
		return
	}

	// Initialize test result
	result := TestResult{
		StatusCodes: make(map[int]int),
	}

	// Create a wait group to wait for all goroutines
	var wg sync.WaitGroup
	// Create a channel to control concurrency
	semaphore := make(chan struct{}, *concurrency)
	// Create a channel to collect results
	results := make(chan int, *requests)

	startTime := time.Now()

	// Launch goroutines for each request
	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			resp, err := http.Get(*url)
			if err != nil {
				results <- 0
				return
			}
			defer resp.Body.Close()
			results <- resp.StatusCode
		}()
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for statusCode := range results {
		result.TotalRequests++
		if statusCode == http.StatusOK {
			result.SuccessfulRequests++
		}
		result.StatusCodes[statusCode]++
	}

	result.TotalTime = time.Since(startTime)

	// Print report
	printReport(result)
}

func printReport(result TestResult) {
	fmt.Println("\n=== Load Test Report ===")
	fmt.Printf("Total Time: %v\n", result.TotalTime)
	fmt.Printf("Total Requests: %d\n", result.TotalRequests)
	fmt.Printf("Successful Requests (200): %d\n", result.SuccessfulRequests)
	fmt.Println("\nStatus Code Distribution:")
	for code, count := range result.StatusCodes {
		fmt.Printf("Status %d: %d requests\n", code, count)
	}
}
