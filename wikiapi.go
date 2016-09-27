package main

import (
	"fmt"
	"github.com/joey-walker/johanbot/common"
)

const wikiBase string = "https://en.wikipedia.org/w/api.php?"
const action string = "action="
const query string = "query"
const ampersand string = "&"
const searchAllTopics string = "srsearch="
const formatJSON string = "format=json"
const wikiUTF string = "utf8=1"
const listSearch string = "list=search"
const searchQuery string = wikiBase + action + query + ampersand + formatJSON + ampersand + listSearch + ampersand + wikiUTF
const searchAllAction string = searchQuery + ampersand + searchAllTopics

// Search for an object
func searchAll(searchItem string) {
	strippedItem := common.ConvertSpacesToNbsp(searchItem)
	fmt.Println("my strippedItem, ", strippedItem)

}
