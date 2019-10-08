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
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yaml-extract",
	Short: "Extract value from a yaml file into another yaml file",
	Long:  `Extract value from a yaml file into another yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		keyConfig := viper.New()
		keyConfig.SetConfigFile(cmd.Flag("keys").Value.String())
		if err := keyConfig.ReadInConfig(); err != nil {
			panic(err)
		}
		keys := keyConfig.GetStringSlice("keys")
		input := viper.New()
		input.SetConfigFile(cmd.Flag("input").Value.String())
		if err := input.ReadInConfig(); err != nil {
			panic(err)
		}
		outM := make(map[string]interface{})
		for _, key := range keys {
			outM[key] = input.Get(key)
		}
		outB, err := yaml.Marshal(&outM)
		if err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(cmd.Flag("output").Value.String(), outB, 0644); err != nil {
			panic(err)
		}
	},
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("input", "i", "./input.yaml", "input yaml file")
	rootCmd.Flags().StringP("output", "o", "./output.yaml", "output yaml file")
	rootCmd.Flags().StringP("keys", "k", "./keys.yaml", "keys to be extracted")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".yaml-extract" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".yaml-extract")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
