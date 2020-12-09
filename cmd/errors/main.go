package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"todoapp/lib/errors/generate"

	"gopkg.in/yaml.v2"
)

func generateErrors(errorTags map[string]generate.ErrorMap) {
	err := generate.Validate(errorTags)
	if err != nil {
		panic(err)
	}

	output, err := os.OpenFile("pkg/errors/errors.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	err = generate.Generate(errorTags, output)
	if err != nil {
		panic(err)
	}

	err = output.Close()
	if err != nil {
		panic(err)
	}
}

func generateCmd(errorTags map[string]generate.ErrorMap) *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "generate pkg/errors/errors.go from errors.yml",
		Run: func(cmd *cobra.Command, args []string) {
			generateErrors(errorTags)
		},
	}
}

func nextErrorCodeCmd(errorTags map[string]generate.ErrorMap) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-code",
		Short: "find the next error code for rpc-status",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing rpc-status")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			n, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				panic(err)
			}

			nextCode, err := generate.NextErrorCodeForRPCStatus(errorTags, uint32(n))
			if err != nil {
				panic(err)
			}

			fmt.Println(nextCode)
		},
	}
	return cmd
}

func main() {
	file, err := os.Open("errors.yml")
	if err != nil {
		panic(err)
	}

	errorTags := make(map[string]generate.ErrorMap)
	err = yaml.NewDecoder(file).Decode(&errorTags)
	if err != nil {
		panic(err)
	}

	rootCmd := &cobra.Command{
		Use:   "errors",
		Short: "errors utility for generating pkg/errors/errors.go",
	}

	rootCmd.AddCommand(
		generateCmd(errorTags),
		nextErrorCodeCmd(errorTags),
	)

	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
