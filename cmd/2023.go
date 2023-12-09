/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adventofcode/2023/day1"
	"adventofcode/2023/day2"
	"adventofcode/2023/day3"
	"adventofcode/2023/day4"
	"adventofcode/2023/day5"
	"adventofcode/2023/day6"
	"adventofcode/2023/day7"
	"adventofcode/2023/day8"
	"adventofcode/2023/day9"

	"github.com/spf13/cobra"
)

// 2023Cmd represents the 2023 command
var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		daysFunc := map[string]func(){
			"day1": day1.Main,
			"day2": day2.Main,
			"day3": day3.Main,
			"day4": day4.Main,
			"day5": day5.Main,
			"day6": day6.Main,
			"day7": day7.Main,
			"day8": day8.Main,
			"day9": day9.Main,
		}
		for _, arg := range args {
			if val, ok := daysFunc[arg]; ok {
				val()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(Cmd)
}
