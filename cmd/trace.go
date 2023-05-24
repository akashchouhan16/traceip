/*
Copyright © 2023 akashchouhan16/traceip <akash.c1500@gmail.com>

*/
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
	"github.com/spf13/cobra"
)

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

type flag struct {
	forceV4 bool `default:false`
	forceV6 bool `default:false`
}

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	CLIFlags   flag
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

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: color.HiGreenString("Geo locate an IP v4 or IP v6 with the trace command."),
	Long:  color.HiGreenString(`Usage: traceip trace <IP v4/ IP v6> command allows to geo locate the ip address.`),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				displayGeoData(ip)
			}
		} else {
			WarningLog.Println(color.HiYellowString("Please provide an IP v4 or IP v6 with trace command to Geo Locate the IP address."))
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func displayGeoData(ip string) {
	if net.ParseIP(ip) == nil {
		WarningLog.Println(color.HiYellowString("Error: Invalid IP address. Please try again with a valid ip."))
		// fmt.Println("Invalid IP v4 / IP v6 provided. Please try again with a valid IP address")
		return
	} else {
		url := "https://ipinfo.io/" + ip + "/geo"
		responseByte := traceIP(url)

		data := GeoLocation{}

		err := json.Unmarshal(responseByte, &data)
		if err != nil {
			ErrorLog.Println(color.HiRedString("Failed to unmarshal JSON response."))
			// log.Println("Unable to unmarshal the response")
		}

		c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
		c.Println("DATA :")

		c = color.New(color.FgHiYellow).Add(color.Bold)
		c.Printf("IP: %s\nCITY: %s\nREGION: %s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE:%s\nPOSTAL: %s\nOrg: %s\nDoc: %s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal, data.Org, data.Readme)

	}

}

/*
	@about: Geo Locating IP address with an open API @ ipinfo.io.
	@param: string ip
	@output: struct data
*/
func traceIP(url string) []byte {
	fmt.Printf("Locating IP...\n")
	response, err := http.Get(url)

	if err != nil {
		ErrorLog.Println(color.HiRedString("Network call for IP geo locate failed: %s", err))
	}

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ErrorLog.Println(color.HiRedString("IO Util read failed to generate byte response: %s", err))
	}

	return responseInBytes

}
