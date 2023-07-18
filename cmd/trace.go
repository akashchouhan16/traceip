/* ***********************************************
 * copyright: github.com/akashchouhan16/traceip
 * author: akashchouhan16
 *************************************************/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

type flag struct {
	v4 bool `default:false`
	v6 bool `default:false`
}

var (
	Warn     *log.Logger
	Info     *log.Logger
	Error    *log.Logger
	CLIFlags flag
)

type GeoLocation struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

/*
	@about: Geo Locating IP address with an open API @ ipinfo.io.
	@param: string ip
	@output: struct data
*/

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Geo locate an IPv4 or IPv6 with the trace command.",
	Long:  `Usage: traceip trace <ip address(es)> command allows to geo locate the ip address(es).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				displayGeoData(ip)
			}
		} else {
			Warn.Println(color.HiYellowString("Please provide an IP v4 or IP v6 with trace command to Geo Locate the IP address."))
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
	Info = log.New(os.Stdout, color.CyanString("[INFO]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, color.YellowString("[WARNING]: "), log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, color.RedString("[ERROR]: "), log.Ldate|log.Ltime|log.Lshortfile)
}

func displayGeoData(ip string) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Parameters", "Value"})

	if net.ParseIP(ip) == nil {
		Error.Println(color.HiRedString("Invalid IP address. Please try again with a valid ip."))

		return
	} else {
		url := "https://ipinfo.io/" + ip + "/geo"
		responseByte := traceIP(url)

		data := GeoLocation{}

		err := json.Unmarshal(responseByte, &data)
		if err != nil {
			Error.Println(color.HiRedString("Failed to unmarshal JSON response."))
			// log.Println("Unable to unmarshal the response")
		}

		c := color.New(color.FgHiCyan).Add(color.Underline).Add(color.Bold)
		c.Println("\n=====================================================================")

		c = color.New(color.FgHiWhite)
		c.Println("IP A")
		t.AppendRow(table.Row{1, "IP Address", data.IP})
		t.AppendRow(table.Row{2, "City", data.City})
		t.AppendRow(table.Row{3, "Region", data.Region})
		t.AppendRow(table.Row{4, "Country", data.Country})
		t.AppendRow(table.Row{5, "Coordinates", data.Loc})
		t.AppendRow(table.Row{6, "Timezone", data.Timezone})
		t.AppendRow(table.Row{7, "Zip/Postal", data.Postal})
		t.AppendRow(table.Row{8, "Organization", data.Org})

		c.Println(t.Render())
		c = color.New(color.FgHiCyan).Add(color.Bold)
		c.Println("=====================================================================")
	}

}

func traceIP(url string) []byte {
	fmt.Printf("Locating IP...\n")
	response, err := http.Get(url)

	if err != nil {
		Error.Println(color.HiRedString("Network call for IP geo locate failed: %s", err))
	}

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Error.Println(color.HiRedString("IO Util read failed to generate byte response: %s", err))
	}

	return responseInBytes

}
