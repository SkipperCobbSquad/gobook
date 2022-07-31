package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := WaitForServer("https://exemple.onion"); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}

}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
