package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// middlewares ------------------------------------------------------------------
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}

// ------------------------------------------------------------------------------

// structs-----------------------------------------------------------------------
type URLHolder struct {
	URLs       []string `json:"urls"`
	RetryLimit int      `json:"retry_limit"`
}
type CheckResult struct {
	URL     string `json:"url"`
	Latency string `json:"latency"`
	Err     string `json:"error"`
}

// ------------------------------------------------------------------------------

// interfaces--------------------------------------------------------------------
type URLChecker interface {
	Check(url string, retry_limit int) (latency time.Duration, err error)
}

// ------------------------------------------------------------------------------

// implementations --------------------------------------------------------------
func (uh URLHolder) Check(url string, retry_limit int) (latency time.Duration, err error) {
	var startTime time.Time
	var response *http.Response

	for i := 0; i <= retry_limit; i++ {
		startTime = time.Now()
		response, err = http.Get(url)
		if err == nil {
			response.Body.Close()
			return time.Since(startTime), nil
		}
	}

	return 0, err
}

// ------------------------------------------------------------------------------

// workers-----------------------------------------------------------------------
func checkerHelper(uh URLHolder, url string, rate_limit int, ch chan<- CheckResult) {
	latency, err := uh.Check(url, rate_limit)

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	ch <- CheckResult{URL: url, Latency: latency.String(), Err: errMsg}
}

// ------------------------------------------------------------------------------

// conductor---------------------------------------------------------------------
func recieve_urls_and_check(w http.ResponseWriter, r *http.Request) {
	var urls URLHolder
	var results []CheckResult
	err := json.NewDecoder(r.Body).Decode(&urls)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(urls)
	var channel = make(chan CheckResult, len(urls.URLs))

	for _, item := range urls.URLs {
		go checkerHelper(urls, item, urls.RetryLimit, channel)
	}

	for i := 0; i < len(urls.URLs); i++ {
		res := <-channel
		results = append(results, res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

// ------------------------------------------------------------------------------

// main execution point
func main() {
	http.HandleFunc("POST /check", loggingMiddleware(recieve_urls_and_check))

	// server stuff
	fmt.Println("Server starting on: 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
