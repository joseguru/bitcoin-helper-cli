package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var echoTimes int
var priceCmd = &cobra.Command{
	Use:   "Bitcoin helper",
	Short: "This program will contain a list of bitcoin helper commands",
	Long: `A Fast and Flexible bitcoin helper commands built with
                love by joseguru and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD")
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		fmt.Println(string(body))
		defer resp.Body.Close()

	},
}

func Execute() {
	if err := priceCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
