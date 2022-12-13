package internal

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryBuilder(t *testing.T) {
	tt := []struct {
		albumID   int
		id        int
		title     string
		wantParts []string
	}{
		{
			albumID:   1,
			id:        1,
			title:     "foo",
			wantParts: []string{"albumId=1", "id=1", "title_like=foo"},
		},
		{
			albumID:   1,
			id:        -1,
			title:     "foo",
			wantParts: []string{"albumId=1", "title_like=foo"},
		},
		{
			albumID:   -1,
			id:        1,
			title:     "foo",
			wantParts: []string{"id=1", "title_like=foo"},
		},
		{
			albumID:   -1,
			id:        -1,
			title:     "foo",
			wantParts: []string{"title_like=foo"},
		},
		{
			albumID:   -1,
			id:        -1,
			title:     "",
			wantParts: []string{},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.wantParts), func(t *testing.T) {
			opts := Options{
				AlbumID: tc.albumID,
				ID:      tc.id,
				Title:   tc.title,
			}
			req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/photos", nil)
			require.NoError(t, err)

			opts.EncodeQuery(req)

			got := req.URL.Query().Encode()

			for _, wantPart := range tc.wantParts {
				require.Contains(t, got, wantPart)
			}
		})
	}
}
