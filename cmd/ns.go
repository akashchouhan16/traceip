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

// nsCmd represents the trace command
var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Perform look up for Name Servers for a particular host.",
	Long:  `Usage: traceip ns <hostname> command allows to look up the Name Server(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, hostname := range args {
				performNSLookUp(hostname)
			}
		} else {
			Warn.Println(color.HiYellowString("Please provide a valid hostname (traceip.com) with ns command to look up name server(s)."))
		}
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)
	Info = log.New(os.Stdout, color.CyanString("[INFO]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, color.YellowString("[WARNING]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, color.RedString("[ERROR]: "), log.Ldate|log.Ltime|log.Lshortfile)
}

/*
	@about: Perform Name Server (NS) look up for hostname.
	@param: string hostname
	@output: array ns
*/
func performNSLookUp(hostname string) {

	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Name Server"})
	ns, err := net.LookupNS(hostname)
	if err != nil {
		Error.Println(color.HiRedString("Error occured with NS Lookup on %s. Err: %s", hostname, err))
		return
	}

	c := color.New(color.FgBlue).Add(color.Underline).Add(color.Bold)
	c.Println("\n=====================================================================")
	c = color.New(color.FgHiBlue).Add(color.Bold)
	c.Println("Name Server(s) for ", hostname)
	c = color.New(color.FgHiWhite)
	for i := 0; i < len(ns); i++ {
		t.AppendRow(table.Row{i + 1, ns[i].Host})
	}
	c = color.New(color.FgHiWhite)
	c.Println(t.Render())
	c = color.New(color.FgBlue).Add(color.Bold)
	c.Println("=====================================================================")
}
