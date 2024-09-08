package cmd

import (
	"fmt"
	"os"

	"github.com/NeGat1FF/github-activity/activity"
	"github.com/spf13/cobra"
)

var TypeFlag string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-activity [username]",
	Short: "Github User Activity is a CLI tool for fetching user activity",
	Long: `Github User Activity is a CLI tool for fetching activity of specific user. 

Example:
	github-activity NeGat1FF

You can also use the --type flag to filter and print specific activity types

Example:
	github-activity NeGat1FF --type push
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		acts, err := activity.GetActivity(args[0])
		if err != nil {
			return err
		}
		switch TypeFlag {
		case "all":
		case "push":
			acts, err = activity.FilterActivities("push", acts)
		case "issue":
			acts, err = activity.FilterActivities("issue", acts)
		case "fork":
			acts, err = activity.FilterActivities("fork", acts)
		case "star":
			acts, err = activity.FilterActivities("star", acts)
		case "create":
			acts, err = activity.FilterActivities("create", acts)
		case "delete":
			acts, err = activity.FilterActivities("delete", acts)
		default:
			return fmt.Errorf("invalid flag")
		}

		if err != nil {
			return err
		}

		activity.PrintActivities(acts)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-activity.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVar(&TypeFlag, "type", "all", "Specify type to list")
}
