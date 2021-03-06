package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zimnx/YamlSchemaToGoStruct/app"
	"os"
)

var (
	packageName,
	output,
	rawSuffix,
	interfaceSuffix string
)

// RootCmd of application
var RootCmd = &cobra.Command{
	Use:   "YamlSchemaToGoStruct [path to config file with schemas]",
	Short: "YamlSchemaToGoStruct generates go structs from yaml schemas",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := app.Run(
			args[0],
			output,
			packageName,
			rawSuffix,
			interfaceSuffix,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.Flags().StringVarP(
		&packageName,
		"package-name",
		"p",
		"esi",
		"package name for implementation and raw structs",
	)
	RootCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"",
		"prefix added to output file",
	)
	RootCmd.Flags().StringVar(
		&rawSuffix,
		"raw-suffix",
		"",
		"suffix added to raw struct names",
	)
	RootCmd.Flags().StringVar(
		&interfaceSuffix,
		"interface-suffix",
		"gen",
		"suffix added to generated interface names",
	)
}

// Execute RootCmd
func Execute() {
	RootCmd.Execute()
}
