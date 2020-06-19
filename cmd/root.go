/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var stdoutLogFormat = logging.MustStringFormatter(
	`%{time:15:04:05.000} [%{shortfile}] [%{level}] %{message}`,
)

var log = logging.MustGetLogger("main")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rollups",
	Short: "A brief description of your application",
	Long:  `A longer description `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/settings.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	initLogger("DEBUG")
	envCfg, envFound := os.LookupEnv("GN_SETTINGS")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else if envFound {
		viper.SetConfigFile(envCfg)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra_test" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("settings.yaml")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initLogger(logLevel string) {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	errBackend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, stdoutLogFormat)
	errBackendFormatter := logging.NewBackendFormatter(errBackend, stdoutLogFormat)
	backendLeveled := logging.AddModuleLevel(errBackendFormatter)
	errBackendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	errBackendLeveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backendLeveled, backendFormatter)
	logging.SetBackend(errBackendLeveled, errBackendFormatter)
}
