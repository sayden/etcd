// Copyright 2015 CoreOS, Inc.
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
	"fmt"
	"os"
	"sync"

	"github.com/coreos/etcd/Godeps/_workspace/src/github.com/cheggaaa/pb"
	"github.com/coreos/etcd/Godeps/_workspace/src/github.com/spf13/cobra"
)

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "A low-level benchmark tool for etcd3",
	Long: `benchmark is a low-level benchmakr tool for etcd3.
It uses gRPC client directly and does not depend on 
etcd client libray.
	`,
}

var (
	endpoints    string
	totalConns   uint
	totalClients uint

	bar     *pb.ProgressBar
	results chan result
	wg      sync.WaitGroup
)

func init() {
	RootCmd.PersistentFlags().StringVar(&endpoints, "endpoint", "127.0.0.1:2378", "comma-separated gRPC endpoints")
	RootCmd.PersistentFlags().UintVar(&totalConns, "conns", 1, "Total number of gRPC connections")
	RootCmd.PersistentFlags().UintVar(&totalClients, "clients", 1, "Total number of gRPC clients")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
