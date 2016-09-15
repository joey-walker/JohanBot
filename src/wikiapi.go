package main

import "fmt"

const wikiBase string = "https://en.wikipedia.org/w/api.php?"
const action string = "action="
const query string = "query"
const ampersand string = "&"
const searchall string = "srsearch="
const formatJSON string = "format=json"
const wikiUTF string = "utf8=1"
const listsearch string = "list=search"
const searchQuery string = wikiBase + action + query + ampersand + formatJSON + ampersand + listsearch + ampersand + wikiUTF
const searchAllAction string = searchQuery + ampersand + searchall

// Search for an object
func searchAll(searchItem string) {
	strippedItem := convertSpacesToNbsp(searchItem)
	fmt.Println("my strippedItem, ", strippedItem)

}
