package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"
	"os"

	"github.com/spf13/cobra"
)

func NewGroupCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "group [OPTION] [ROBOT/GROUP NAME] [ROBOT/GROUP NAME]",
		Short: "create and mange groups in the mind file",
		Long: "Create and mange groups in the mind file which is used by `mindz multirun`\n" +
			"Execute `mind scan` to search for available robots",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide options!")
				os.Exit(-1)
			}

			if options := args[0]; options == "addr" {
				if len(args) < 3 {
					fmt.Println("Please provide ROBOT NAME and GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.AddRobotToGroup(args[1], args[2])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "addg" {
				if len(args) < 2 {
					fmt.Println("Please provide GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.AddGroup(args[1])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "delr" {
				if len(args) < 3 {
					fmt.Println("Please provide ROBOT NAME and GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.DeleteRobotFromGroup(args[1], args[2])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "delg" {
				if len(args) < 2 {
					fmt.Println("Please provide GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.DeleteGroup(args[1])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "listr" {
				if len(args) < 2 {
					fmt.Println("Please provide GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.ListRobotInGroup(args[1])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "listg" {
				err := cli.ListGroup()
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "runn" {
				if len(args) < 2 {
					fmt.Println("Please provide GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.GroupRun_n(args[1])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else if options == "run" {
				if len(args) < 2 {
					fmt.Println("Please provide GROUP NAME !")
					os.Exit(-1)
				}
				err := cli.GroupRun(args[1])
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			} else {
				fmt.Println("Unrecongized Command!")
				os.Exit(-1)
			}
		},
	}
}
