package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetchProblemIO(contestID int, index string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	c.OnHTML("div.sample-test", func(e *colly.HTMLElement) {
			
		i, err := os.Create("cfin.txt")	
		check(err)
		defer i.Close()

		o, err := os.Create("cfout.txt")	
		check(err)
		defer o.Close()

		// input
		e.ForEach("div.test-example-line", func(_ int, s *colly.HTMLElement) {
			i.Write([]byte(s.Text + "\n"))
		})

		// output
		e.ForEach("div.output pre", func(_ int, s *colly.HTMLElement) {
			o.Write([]byte(s.Text + "\n"))
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.Visit(fmt.Sprintf("https://codeforces.com/contest/%d/problem/%s", contestID, index))
}

func main() {	

	contestID, err := strconv.Atoi(os.Args[1])
	check(err)
	index := os.Args[2]

	fetchProblemIO(contestID, index)
}
