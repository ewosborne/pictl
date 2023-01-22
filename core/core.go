package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type CliArgs struct {
	Command string
	Delay   int
	Host    string
	Verbose bool
}

func (c CliArgs) String() string {
	return fmt.Sprintf("Command:%s, Delay:%d, Host:%s, Verbose:%t",
		c.Command, c.Delay, c.Host, c.Verbose)
}

func NewCliArgs(cmd *cobra.Command, command string) CliArgs {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: delay may not be defined.
	delay, _ := cmd.Flags().GetInt("time")

	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		log.Fatal(err)
	}

	return CliArgs{Host: host, Delay: delay, Command: command, Verbose: verbose}
}

func Picmd(cliargs CliArgs) {
	if cliargs.Verbose == true {
		slog.Info("verbose")
		fmt.Println(cliargs)
	}
	reqstr := fmt.Sprintf("http://%s/admin/api.php", cliargs.Host)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, reqstr, nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args

	q := req.URL.Query()
	q.Add("auth", os.Getenv("PIPWHASH"))
	q.Add(cliargs.Command, strconv.Itoa(cliargs.Delay))

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	if cliargs.Verbose == true {
		slog.Info(req.URL.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

}
