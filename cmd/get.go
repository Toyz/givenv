package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var secretName string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Create .env file for given project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			secretName = args[0]
		}

		err, res := secretsProvider.Get(secretName)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		datawriter := bufio.NewWriter(file)

		for key, value := range res {
			_, _ = datawriter.WriteString(key + "=" + value + "\r\n")
		}

		datawriter.Flush()
		file.Close()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
