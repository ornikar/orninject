package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ornikar/orninject/internal/template"
	"github.com/spf13/cobra"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace [FILE]",
	Short: "Replace environment variables by values in file",
	Args:  replaceArgs,
	Run:   replaceRun,
}

// replaceRun represents the run command
func replaceRun(cmd *cobra.Command, args []string) {
	if err := template.Replace(args[0]); err != nil {
		fmt.Println("Unable to replace:")
		fmt.Printf(" %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
	os.Exit(0)
}

// replaceArgs checks the arguments
func replaceArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a file argument")
	}

	if _, err := os.Stat(args[0]); !os.IsNotExist(err) {
		return nil
	}

	return fmt.Errorf("invalid file specified: %s", args[0])
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
