package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cfgo",
	Short: "A minimalistic CodeForces helper.",
	Long:  "A minimalistic CodeForces helper.",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	//do something
	// },
}

func check(e error, msg string) {
	if e != nil {
		log.Error(msg)
		panic(e)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(listCmd)
	fetchCmd.PersistentFlags().String("contest", "", "Contest ID.")
	fetchCmd.PersistentFlags().String("problem", "A", "Problem ID. Default: A")
}
