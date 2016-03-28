package linkheader_test

import (
	"fmt"

	"github.com/tomnomnom/linkheader"
)

func ExampleParse() {
	header := "<https://api.github.com/user/58276/repos?page=2>; rel=\"next\"," +
		"<https://api.github.com/user/58276/repos?page=2>; rel=\"last\""
	links := linkheader.Parse(header)

	for _, link := range links {
		fmt.Printf("URL: %s; Rel: %s\n", link.URL, link.Rel)
	}

	// Output:
	// URL: https://api.github.com/user/58276/repos?page=2; Rel: next
	// URL: https://api.github.com/user/58276/repos?page=2; Rel: last
}

func ExampleParseMultiple() {
	headers := []string{
		"<https://api.github.com/user/58276/repos?page=2>; rel=\"next\"",
		"<https://api.github.com/user/58276/repos?page=2>; rel=\"last\"",
	}
	links := linkheader.ParseMultiple(headers)

	for _, link := range links {
		fmt.Printf("URL: %s; Rel: %s\n", link.URL, link.Rel)
	}

	// Output:
	// URL: https://api.github.com/user/58276/repos?page=2; Rel: next
	// URL: https://api.github.com/user/58276/repos?page=2; Rel: last
}

func ExampleLinks_FilterByRel() {
	header := "<https://api.github.com/user/58276/repos?page=2>; rel=\"next\"," +
		"<https://api.github.com/user/58276/repos?page=2>; rel=\"last\""
	links := linkheader.Parse(header)

	for _, link := range links.FilterByRel("last") {
		fmt.Printf("URL: %s; Rel: %s\n", link.URL, link.Rel)
	}

	// Output:
	// URL: https://api.github.com/user/58276/repos?page=2; Rel: last

}
