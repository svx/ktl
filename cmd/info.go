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

// clusterInfoCmd represents the clusterInfo command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows common info about the cluster",
	Long: `Display addresses of the master and services.

This is the short version of the kubectl command:

kubectl cluster-info`,

	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("clusterInfo called")
                clusterINFO()
	},
}

func clusterINFO() {
        color.Yellow("Cluster Information")
	cmdStr := "kubectl cluster-info"
        cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
