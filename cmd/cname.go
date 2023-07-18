/* ***********************************************
 * copyright: github.com/akashchouhan16/traceip
 * author: akashchouhan16
 *************************************************/
package cmd

import (
	"log"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var cnameCmd = &cobra.Command{
	Use:   "cname",
	Short: "Perform look up for canonical name (cname) for a particular host.",
	Long:  `Usage: traceip cname <hostname> command allows to look up the cname.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, hostname := range args {
				performCNAMELookUp(hostname)
			}
		} else {
			Warn.Println(color.HiYellowString("Please provide a valid hostname (traceip.com) with cname command to look up canonical name(s)."))
		}
	},
}

func init() {
	rootCmd.AddCommand(cnameCmd)
	Info = log.New(os.Stdout, color.CyanString("[INFO]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, color.YellowString("[WARNING]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, color.RedString("[ERROR]: "), log.Ldate|log.Ltime|log.Lshortfile)
}

/*
	@about: Perform CNAME look up for hostname.
	@param: string hostname
	@output: array CNAME
*/
func performCNAMELookUp(hostname string) {

	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Canonical Name"})
	cname, err := net.LookupCNAME(hostname)
	if err != nil {
		Error.Println(color.HiRedString("Error occured with CNAME Lookup on %s. Err: %s", hostname, err))
		return
	}

	c := color.New(color.FgBlue).Add(color.Underline).Add(color.Bold)
	c.Println("\n=====================================================================")
	c = color.New(color.FgHiBlue).Add(color.Bold)
	c.Println("canonical name (CNAME) for ", hostname)
	c = color.New(color.FgHiWhite)
	t.AppendRow(table.Row{1, cname})
	c.Println(t.Render())
	c = color.New(color.FgBlue).Add(color.Bold)
	c.Println("=====================================================================")
}
