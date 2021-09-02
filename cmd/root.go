/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/masterminds/sprig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var cfgFile, tmplFile, valFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "templite",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fncs := sprig.TxtFuncMap()

		tmpl := template.Must(template.New(path.Base(tmplFile)).Funcs(fncs).ParseFiles(tmplFile))

		//check if valFile is readable

		var values []byte
		var err error

		if valFile == "-" {

			values, err = ioutil.ReadAll(os.Stdin)

			if err != nil {

				fmt.Fprintf(os.Stderr, "unable to read from stdin due to error: %v\n", err)

				os.Exit(1)

			}

		} else {

			values, err = os.ReadFile(valFile)
			if err != nil {

				fmt.Fprintf(os.Stderr, "unable to read values file due to error: %v\n", err)

				os.Exit(1)

			}

		}

		v := make(map[string]interface{})
		err = yaml.Unmarshal(values, &v)
		if err != nil {

			fmt.Fprintf(os.Stderr, "unable to Unmarshal values due to error: %v\n", err)

			os.Exit(1)
		}

		err = tmpl.Execute(os.Stdout, v)
		if err != nil {

			fmt.Fprintf(os.Stderr, "unable to execute template due to error: %v\n", err)

			os.Exit(1)
		}

		//template

		//print result to stdout

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.templite.yaml)")
	rootCmd.PersistentFlags().StringVar(&tmplFile, "template", "./template.tmpl", "template file (default is ./template.tmpl)")
	rootCmd.PersistentFlags().StringVar(&valFile, "file", "./values.yaml", "values file (default is ./values.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".templite" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".templite")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
