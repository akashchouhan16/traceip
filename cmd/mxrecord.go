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

// mxrecordCmd represents the trace command
var mxrecordCmd = &cobra.Command{
	Use:   "mx",
	Short: "Perform MX Records look up for a particular host.",
	Long:  `Usage: traceip mx <hostname> command allows to look up the MX Record(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, hostname := range args {
				performMXLookUp(hostname)
			}
		} else {
			Warn.Println(color.HiYellowString("Please provide a valid hostname (traceip.com) with mx command to look up mx record(s)."))
		}
	},
}

func init() {
	rootCmd.AddCommand(mxrecordCmd)
	Info = log.New(os.Stdout, color.CyanString("[INFO]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, color.YellowString("[WARNING]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, color.RedString("[ERROR]: "), log.Ldate|log.Ltime|log.Lshortfile)
}

/*
	@about: Perform MX Records look up for hostname.
	@param: string hostname
	@output: array mxrecords
*/
func performMXLookUp(hostname string) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "MX Record Host", "MX Record Pref"})

	mxrecords, err := net.LookupMX(hostname)
	if err != nil {
		Error.Println(color.HiRedString("Error occured with MX records Lookup on %s. Err: %s", hostname, err))
		return
	}

	c := color.New(color.FgBlue).Add(color.Underline).Add(color.Bold)
	c.Println("\n=====================================================================")
	c = color.New(color.FgHiBlue).Add(color.Bold)
	c.Println("MS Records for", hostname)
	c = color.New(color.FgHiWhite)
	for i := 0; i < len(mxrecords); i++ {
		t.AppendRow(table.Row{i + 1, mxrecords[i].Host, mxrecords[i].Pref})
	}
	c = color.New(color.FgHiWhite)
	c.Println(t.Render())
	c = color.New(color.FgBlue).Add(color.Bold)
	c.Println("=====================================================================")
}
