package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Manual streming response body stream.

func Stream(body io.ReadCloser) []byte {
	defer body.Close() // Make sure to close the stream to release resources.

	buffer := make([]byte, 1024)

	for {
		n, err := body.Read(buffer)

		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	return buffer
}

func ResponseStream() {
	// url := "https://jsonplaceholder.typicode.com/posts"

	serverURL := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "/posts",
		RawQuery: "",
	}

	res, err := http.Get(serverURL.String())

	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	var posts []Post

	err = json.NewDecoder(res.Body).Decode(&posts)

	if err != nil {
		log.Fatal(err)
		return
	}

}
