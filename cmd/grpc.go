package cmd

import (
	"github.com/spf13/cobra"
)

// microCmd represents the micro command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Create a basic golang gRpc project template.",
	Long:  `An tools for Go Grpc Project. Example: tf-cli grpc demo`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// microCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// microCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
