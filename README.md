![](http://i.imgur.com/GQOYw.gif)

`$ go get github.com/bsdf/lastfm`

```go
package main

import (
	"github.com/bsdf/lastfm"
	"fmt"
)

func main() {
  	lastfm := lastfm.LastFM{"Y0UR_4P1_K3Y_H3R3"}
	artists, err := lastfm.GetTopArtists("bsdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	topArtist := artists[0]
    fmt.Println("my favorite artist is:", topArtist.Name)
}
```
