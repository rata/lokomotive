// Copyright 2020 The Lokomotive Authors
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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// Register platforms by adding an anonymous import.
	_ "github.com/kinvolk/lokoctl/pkg/install/aws"
	_ "github.com/kinvolk/lokoctl/pkg/install/baremetal"
	_ "github.com/kinvolk/lokoctl/pkg/install/packet"

	// Register backends by adding an anonymous import.
	_ "github.com/kinvolk/lokoctl/pkg/backend/local"
	_ "github.com/kinvolk/lokoctl/pkg/backend/s3"
)

var rootCmd = &cobra.Command{
	Use:   "lokoctl",
	Short: "Manage Lokomotive clusters.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(cobraInit)

	// Add kubeconfig flag.
	rootCmd.PersistentFlags().String(
		"kubeconfig",
		"", // Special empty default, use getKubeconfig()
		"Path to kubeconfig file, taken from the asset dir if not given, and finally falls back to ~/.kube/config")
	viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))

	rootCmd.PersistentFlags().String("lokocfg", "./", "Path to lokocfg directory or file")
	viper.BindPFlag("lokocfg", rootCmd.PersistentFlags().Lookup("lokocfg"))
	rootCmd.PersistentFlags().String("lokocfg-vars", "./lokocfg.vars", "Path to lokocfg.vars file")
	viper.BindPFlag("lokocfg-vars", rootCmd.PersistentFlags().Lookup("lokocfg-vars"))
}

func cobraInit() {
	viper.AutomaticEnv()
}