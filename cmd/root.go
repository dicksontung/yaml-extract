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
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	"github.com/adam-hanna/arrayOperations"
	"github.com/dicksontung/viper"
	"github.com/mitchellh/go-homedir"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yaml-extract",
	Short: "Extract value from a yaml file into another yaml file",
	Long:  ``,
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
		difference, _ := arrayOperations.Difference(keys, input.AllKeys())
		diff := difference.Interface().([]string)
		for _, unwantedKey := range diff {
			input.Unset(unwantedKey)
		}
		outB, err := yaml.Marshal(input.AllSettings())
		if err != nil {
			panic(err)
		}
		if cmd.Flag("output").Changed {
			if err := ioutil.WriteFile(cmd.Flag("output").Value.String(), outB, 0644); err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("%s", outB)
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
	rootCmd.Flags().StringP("input", "i", "", "input yaml file")
	rootCmd.MarkFlagRequired("input")
	rootCmd.Flags().StringP("output", "o", "", "output yaml file")
	rootCmd.Flags().StringP("keys", "k", "", "keys to be extracted")
	rootCmd.MarkFlagRequired("keys")
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
