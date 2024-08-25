package requests

import (
	"fmt"
	"log"
	"net/url"
)

func LearnURLPackage() {
	serverUrl := "https://example.com:8080/path?query=1#fragment"

	output, err := url.Parse(serverUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output.Scheme)
	fmt.Println(output.Host)
	fmt.Println(output.RawQuery)
	fmt.Println(output.Fragment)

	// Building url.
	u := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/path",
		RawQuery: "query-1=1",
		Fragment: "fragment",
	}

	fmt.Println(u.String())
}
