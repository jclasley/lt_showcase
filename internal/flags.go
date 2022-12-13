package internal

import "flag"

type Options struct {
	AlbumID int
	ID      int
	Title   string
	RawJSON bool
}

// GetFlags parses the CLI flags and loads them onto an Options struct with default values that indicate
// the absence of a flag instead of just the zero values.
func GetFlags() Options {
	var opts Options

	flag.IntVar(&opts.AlbumID, "album", -1, "Lookup images in a specific album by ID")
	flag.IntVar(&opts.ID, "id", -1, "Lookup an image by its ID")
	flag.StringVar(&opts.Title, "title", "", "Lookup all images with a title that contains the given string. Accepts regex")

	flag.BoolVar(&opts.RawJSON, "raw", false, "Print the raw JSON response instead of a formatted table (for piping to jq)")

	flag.Parse()

	return opts
}
