/* ***********************************************
 * copyright: 2023 @github.com/akashchouhan16/traceip
 * author: @akashchouhan16
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
	"github.com/spf13/cobra"
)

type flag struct {
	forceV4 bool `default:false`
	forceV6 bool `default:false`
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

// traceCmd represents the trace command
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
	Info = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func displayGeoData(ip string) {
	if net.ParseIP(ip) == nil {
		Error.Println(color.HiRedString("Invalid IP address. Please try again with a valid ip."))
		// fmt.Println("Invalid IP v4 / IP v6 provided. Please try again with a valid IP address")
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

		c := color.New(color.FgGreen).Add(color.Underline).Add(color.Bold)
		c.Println("==============================")

		c = color.New(color.FgHiGreen).Add(color.Bold)
		c.Printf("	- IP: %s\n	- CITY: %s\n	- REGION: %s\n	- COUNTRY :%s\n	- LOCATION :%s\n	- TIMEZONE:%s\n	- POSTAL: %s\n	- Org: %s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal, data.Org)
		c.Println("==============================")
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
		Error.Println(color.HiRedString("Network call for IP geo locate failed: %s", err))
	}

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Error.Println(color.HiRedString("IO Util read failed to generate byte response: %s", err))
	}

	return responseInBytes

}
