package internal

import (
	"fmt"
	"net/http"
)

// EncodeQuery takes the options received from the CLI flags and loads them onto the request
// as URL encoded query parameters.
func (o *Options) EncodeQuery(r *http.Request) {
	q := r.URL.Query()
	if o.AlbumID != -1 {
		q.Add("albumId", fmt.Sprintf("%d", o.AlbumID))
	}

	if o.ID != -1 {
		q.Add("id", fmt.Sprintf("%d", o.ID))
	}

	if o.Title != "" {
		q.Add("title_like", o.Title)
	}

	r.URL.RawQuery = q.Encode()
}
