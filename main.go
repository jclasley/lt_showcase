package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/jclasley/lt_showcase/internal"
)

func main() {
	opts := internal.GetFlags()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/photos", nil)
	if err != nil {
		log.Printf("failed to construct request: %v", err)
		return
	}

	opts.EncodeQuery(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to execute request: %v", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %d", resp.StatusCode)
		return
	}

	var data []response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("failed to decode response: %v", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("No results found. Try a different query.")
		return
	}

	if opts.RawJSON {
		if err := json.NewEncoder(os.Stdout).Encode(data); err != nil {
			log.Printf("failed to encode response: %v", err)
			return
		}
		return
	}

	tw := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)

	for _, d := range data {
		msg := fmt.Sprintf("%d-[%d]\t%s\nURL:\t%s\nThumbnail:\t%s\n\n", d.AlbumID, d.ID, d.Title, d.URL, d.Thumb)
		if _, err := tw.Write([]byte(msg)); err != nil {
			log.Printf("failed to write to tabwriter: %v", err)
			continue
		}
	}
	if err := tw.Flush(); err != nil {
		log.Printf("failed to flush tabwriter: %v", err)
	}

}

type response struct {
	AlbumID int    `json:"albumId"`
	ID      int    `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Thumb   string `json:"thumbnailUrl"`
}
