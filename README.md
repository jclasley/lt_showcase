# LT Showcase

## Options

You can pass four flags to the built binary:
1. `-album <int>` -- returns photos in a given album
2. `-id <int>` -- returns the specific photo with that ID
3. `-title <string|regex>` -- returns photos where the title matches the given string or regex
4. `-raw` -- returns the raw JSON (for piping to JQ)

You can always pass the `-h` flag to see these options in your terminal.

All flags can be used in any given combination, but might not return any results.

Example:

```bash
./photos -album 1 -title 'ut.*id'

# Output:

# 1-[48]          ut esse id
# URL:            https://via.placeholder.com/600/68e0a8
# Thumbnail:      https://via.placeholder.com/150/68e0a8
```

## Running locally

#### Install with Go
`go install github.com/jclasley/lt_showcase`

#### Build and run
``` bash
go build -o ./photos .
./photos -album 2 -title 'ut.*ed'
```