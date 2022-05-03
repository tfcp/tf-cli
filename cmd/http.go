package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Create a basic golang http project template.",
	Long:  `An tools for Go Http Project. Example: tf-cli http demo`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(httpCmd)
	//newCmd.AddCommand()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
