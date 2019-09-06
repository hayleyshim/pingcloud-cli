/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/reoim/pingcloud-cli/ping"
	"github.com/reoim/pingcloud-cli/ping/gcp"
	"github.com/spf13/cobra"
)

// used for local flags
var list bool

// gcpCmd represents the gcp command
var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Check latencies of GCP regions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		if list {

			// Init tabwriter
			tr := tabwriter.NewWriter(os.Stdout, 40, 8, 2, '\t', 0)
			fmt.Fprintf(tr, "GCP Region Code\tGCP Region Name")
			fmt.Fprintln(tr)
			fmt.Fprintf(tr, "------------------------------\t------------------------------")
			fmt.Fprintln(tr)
			for r, n := range gcp.GCPRegions {
				fmt.Fprintf(tr, "[%v]\t[%v]", r, n)
				fmt.Fprintln(tr)
			}
			// Flush tabwriter
			tr.Flush()
		} else if len(args) == 0 {

			// Init tabwriter
			tr := tabwriter.NewWriter(os.Stdout, 40, 8, 2, '\t', 0)
			fmt.Fprintf(tr, "GCP Region Code\tGCP Region Name\tLatency")
			fmt.Fprintln(tr)
			fmt.Fprintf(tr, "------------------------------\t------------------------------\t------------------------------")
			fmt.Fprintln(tr)

			// Flush tabwriter
			tr.Flush()

			for r, i := range gcp.GCPEndpoints {
				p := ping.PingDto{
					Region:  r,
					Name:    gcp.GCPRegions[r],
					Address: i,
				}
				p.Ping()
			}
			fmt.Println("")
			fmt.Println("You can also add region after command if you want http trace information of the specific region")
			fmt.Println("ex> pingcloud-cli gcp us-central1")
		} else {
			for _, r := range args {
				if i, ok := gcp.GCPEndpoints[r]; ok {
					p := ping.PingDto{
						Region:  r,
						Name:    gcp.GCPRegions[r],
						Address: i,
					}
					p.VerbosePing()
				} else {
					fmt.Printf("Region code [%v] is wrong.  To check available region codes run the command with -l or --list flag\n", r)
					fmt.Println("Usage: pingcloud-cli gcp -l")
					fmt.Println("Usage: pingcloud-cli gcp --list")

				}
			}

		}
		fmt.Println("")

	},
}

func init() {
	rootCmd.AddCommand(gcpCmd)

	gcpCmd.Flags().BoolVarP(&list, "list", "l", false, "List all available regions")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gcpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gcpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
