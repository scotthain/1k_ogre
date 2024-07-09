/*
Copyright © 2024 Scott Hain <elejia@gmail.com>
*/
package cmd

import (
	"fmt"
	v "scotthain/ogre/internal"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's.`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Println(v.Version())
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
