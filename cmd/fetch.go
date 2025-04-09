package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch the input and output for given contest and problem.",
	Long:  "Fetch the input and output for given contest and problem.",
	Run: func(cmd *cobra.Command, args []string) {
		contestIDStr, _ := cmd.Flags().GetString("contest")
		contestID, err := strconv.Atoi(contestIDStr)
		check(err)
		problem, _ := cmd.Flags().GetString("problem")

		fetchProblemIO(contestID, problem)
	},
}

func fetchProblemIO(contestID int, problem string) {
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

	c.Visit(fmt.Sprintf("https://codeforces.com/contest/%d/problem/%s", contestID, problem))
}
