package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all the problems for a particular contest",
	Long: "List all the problems for a particular contest.",
	Run: func(cmd *cobra.Command, args[] string) {
		contestId := args[0]
		listContestProblems(contestId)
	},
}

func listContestProblems(contestId string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Default().Print("Visiting ", r.URL)
	})

	c.OnHTML("table.problems", func (e *colly.HTMLElement) {
		e.ForEach("td a", func (_ int, s *colly.HTMLElement) {
			if s.Attr("title") != "Participants solved the problem" {
				fmt.Print(strings.TrimSpace(s.Text))
			}
			
			if len(s.Text) > 1 {
				fmt.Println()
			}
		})
	})

	c.OnError(func (_ *colly.Response, err error) {
		log.Error("Something went wrong: ", err)
	})

	c.Visit("https://codeforces.com/contest/" + contestId)
}