package cmd

import (
	"github.com/spf13/cobra"
)

// microCmd represents the micro command
var microCmd = &cobra.Command{
	Use:   "micro",
	Short: "Create a micro-cmd template.",
	Long:  `An tools for Go micro Project. Example: tf-cli micro demo`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(microCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// microCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// microCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
