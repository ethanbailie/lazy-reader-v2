package main

import (
	"fmt"

	"lazy-reader-v2/tools"
)

func main() {
	body, err := tools.GetPapers(tools.QuerySettings{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)
}
