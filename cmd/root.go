/* ***********************************************
 * copyright: 2023 @github.com/akashchouhan16/traceip
 * author: @akashchouhan16
 *************************************************/
package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "traceip",
	Short: color.HiGreenString("traceip is a command line tool for geolocation lookups on IP addresses. Visit https://github.com/akashchouhan16/traceip for more details."),
	Long: color.HiGreenString(`
 ======================================================	
| ████████╗██████╗  █████╗  ██████╗███████╗ ██╗██████╗  |
| ╚══██╔══╝██╔══██╗██╔══██╗██╔════╝██╔════╝ ██║██╔══██╗ |
|    ██║   ██████╔╝███████║██║     █████╗   ██║██████╔╝ |
|    ██║   ██╔══██╗██╔══██║██║     ██╔══╝   ██║██╔═══╝  |
|    ██║   ██║  ██║██║  ██║╚██████╗███████╗ ██║██║      |
|    ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚══════╝ ╚═╝╚═╝      |
 ====================================================== 
A command line tool for geolocation lookups on IP addresses.
 `),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	if len(args) == 0 {
	// 		c := color.New(color.FgGreen)
	// 		c.Println("Use traceip trace [...ip] to geo locate a collection of IPs or visit https://github.com/akashchouhan16/traceip for more details.")
	// 	}
	// },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.traceip.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("ipv4", "4", false, "Enforce IPv4 geo tracing only {wip}")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Enforce IPv6 geo tracing only {wip}")
}
