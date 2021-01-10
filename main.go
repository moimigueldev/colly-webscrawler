package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0) // will only be able to hold facts

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li ", func(element *colly.HTMLElement) {
		// fmt.Println("element", element)
		factID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get an ID")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	writeJSON(allFacts)

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", " ")
	// enc.Encode(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
	}

	ioutil.WriteFile("rhinofacts.json", file, 0644) // permissions code
}
