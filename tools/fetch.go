package tools

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type QuerySettings struct {
	Categories []string
	MaxResults int
}

func GetPapers(opts QuerySettings) (string, error) {
	baseURL := "http://export.arxiv.org/api/query"
	params := url.Values{}

	finalURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	var categories []string
	if len(opts.Categories) == 0 {
		categories = []string{
			"cs.CL",
			"cs.AI",
			"cs.LG",
		}
	} else {
		categories = opts.Categories
	}

	var maxResults string
	if opts.MaxResults == 0 {
		maxResults = "1"
	} else {
		maxResults = strconv.Itoa(opts.MaxResults)
	}

	search_query := ""
	for i, c := range categories {
		if i > 0 {
			search_query += "+OR+"
		}
		category := "cat:" + c
		search_query += category
	}

	params.Add("search_query", search_query)
	params.Add("start", "0")
	params.Add("max_results", maxResults)
	params.Add("sortBy", "submittedDate")
	params.Add("sortOrder", "descending")

	finalURL.RawQuery = params.Encode()

	resp, err := http.Get(finalURL.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
