// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	//"fmt"
        "os"
	"os/exec"

	"github.com/spf13/cobra"
        "github.com/fatih/color"
)

var verbose bool // used for flags

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "List all pods in all namespaces",
        Long: `Running in verbose mode:
ktl pods -v

gives more detailed info.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("pods called")
		pods()
	},
}

func pods() {
    if verbose {
        //fmt.Println("Listing Pods")
        color.Yellow("Verbose List Of All Pods")
	cmdStr := "kubectl get pods --all-namespaces -o wide"
        cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
    } else {
        color.Yellow("List Of All Pods")
        cmdStr := "kubectl get pods --all-namespaces"
        cmd := exec.Command("bash", "-c", cmdStr)
        cmd.Stdout = os.Stdout
        cmd.Stdin = os.Stdin
        cmd.Stderr = os.Stderr
        cmd.Run()
    }
}

func init() {
	rootCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//podsCmd.Flags().BoolP("list", "l", false, "List all pods of all namespaces")
        podsCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose list all pods")
}
