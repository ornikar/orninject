package cmd

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	author  = "Julien BREUX <julien.breux@ornikar.com>"
	website = "https://ornikar.github.io/orninject/"
)

var (
	version = "dev"
	commit  = "dev"
	date    = "n/a"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version/build info",
	Long:  "Show the version and build information",
	Run:   versionRun,
}

// versionRun represents the run command
func versionRun(cmd *cobra.Command, args []string) {
	printKeyValue("Version:  ", version)
	printKeyValue("Commit:   ", commit)
	printKeyValue("Date:     ", date)
	printKeyValue("Author:   ", author)
	printKeyValue("Website:  ", website)
}

// printKeyValue
func printKeyValue(key, value string) {
	fmt.Println(aurora.Cyan(key), aurora.White(value))
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
