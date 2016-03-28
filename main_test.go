package linkheader

import (
	"testing"
)

func TestSimple(t *testing.T) {
	// Test case stolen from https://github.com/thlorenz/parse-link-header :)
	header := "<https://api.github.com/user/9287/repos?page=3&per_page=100>; rel=\"next\", " +
		"<https://api.github.com/user/9287/repos?page=1&per_page=100>; rel=\"prev\"; pet=\"cat\", " +
		"<https://api.github.com/user/9287/repos?page=5&per_page=100>; rel=\"last\""

	links := Parse(header)

	if len(links) != 3 {
		t.Errorf("Should have been 3 links returned, got %d", len(links))
	}

	if links[0].URL != "https://api.github.com/user/9287/repos?page=3&per_page=100" {
		t.Errorf("First link should have URL 'https://api.github.com/user/9287/repos?page=3&per_page=100'")
	}

	if links[0].Rel != "next" {
		t.Errorf("First link should have rel=\"next\"")
	}

	if len(links[0].Params) != 1 {
		t.Errorf("First link should have exactly 1 params, but has %d", len(links[0].Params))
	}

	if len(links[1].Params) != 2 {
		t.Errorf("Second link should have exactly 2 params, but has %d", len(links[1].Params))
	}

	if links[1].Params["pet"] != "cat" {
		t.Errorf("Second link's 'pet' param should be 'cat', but was %s", links[1].Params["pet"])
	}

}

func TestLinkMethods(t *testing.T) {
	header := "<https://api.github.com/user/9287/repos?page=1&per_page=100>; rel=\"prev\"; pet=\"cat\""
	links := Parse(header)
	link := links[0]

	if !link.HasParam("rel") {
		t.Errorf("Link should have param 'rel'")
	}

	if link.HasParam("foo") {
		t.Errorf("Link should not have param 'foo'")
	}

	val, err := link.Param("pet")
	if err != nil {
		t.Errorf("Error value should be nil")
	}
	if val != "cat" {
		t.Errorf("Link should have param pet=\"cat\"")
	}

	_, err = link.Param("foo")
	if err == nil {
		t.Errorf("Error value should not be nil")
	}

}

func testLinksMethods(t *testing.T) {
	header := "<https://api.github.com/user/9287/repos?page=3&per_page=100>; rel=\"next\", " +
		"<https://api.github.com/user/9287/repos?page=1&per_page=100>; rel=\"stylesheet\"; pet=\"cat\", " +
		"<https://api.github.com/user/9287/repos?page=5&per_page=100>; rel=\"stylesheet\""

	links := Parse(header)

	filtered := links.FilterByRel("next")

	if filtered[0].URL != "https://api.github.com/user/9287/repos?page=3&per_page=100" {
		t.Errorf("URL did not match expected")
	}

	filtered = links.FilterByRel("stylesheet")
	if len(filtered) != 2 {
		t.Errorf("Filter for stylesheet should yield 2 results but got %d", len(filtered))
	}

	filtered = links.FilterByRel("notarel")
	if len(filtered) != 0 {
		t.Errorf("Filter by non-existant rel should yeild no results")
	}

}

func testParseMultiple(t *testing.T) {
	headers := []string{
		"<https://api.github.com/user/58276/repos?page=2>; rel=\"next\"",
		"<https://api.github.com/user/58276/repos?page=2>; rel=\"last\"",
	}

	links := ParseMultiple(headers)

	if len(links) != 2 {
		t.Errorf("Should have returned 2 links")
	}
}
