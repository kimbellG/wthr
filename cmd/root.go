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
	"log"

	"wthr/weather"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

type Config struct {
	APIKey string `mapstructure:"API_Key"`
}

var conf Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wthr",
	Short: "app for get weather.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: mainController,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func assetsInputArgs(isCurrentArgs bool) {
	if !isCurrentArgs {
		log.Fatalln("Input city.")
	}
}

func mainController(cmd *cobra.Command, args []string) {
	assetsInputArgs(isCurrentArgs(cmd, args))
	fmt.Println("----------------------------------------------")

	if fstatus, _ := cmd.Flags().GetBool("geolocation"); fstatus {
		getWeatherByIPCoordinates()
	}

	for _, arg := range args {
		request := weather.WeatherRequestByCityName{arg, conf.APIKey}
		if fstatus, _ := cmd.Flags().GetBool("current"); fstatus {
			fmt.Printf(weather.GetCurrentWeather(request))
			fmt.Println("----------------------------------------------")
		}
	}
}

func getWeatherByIPCoordinates() {
	var coord weather.AnswerIpGeolocationServer = *weather.GetGeolocationCoordinates()
	fmt.Println(weather.GetCurrentWeatherForGeolocation(weather.WeatherRequestByGeoCoord{Coordinate: coord, APIKey: weather.GEOAPIKEY}))
}

func isCurrentArgs(cmd *cobra.Command, args []string) bool {
	geoFlagStatus, _ := cmd.Flags().GetBool("geolocation")
	return geoFlagStatus || len(args) > 0
}

func init() {
	cobra.OnInitialize(initConfig)
	log.SetFlags(0)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wthr.yaml)")
	rootCmd.Flags().BoolP("current", "c", false, "Get current weather in given city")
	rootCmd.Flags().BoolP("geolocation", "g", false, "Get weather by geolocation")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".wthr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wthr")
		viper.SetConfigType("env")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	assetsConfigFile(viper.ReadInConfig())

	viper.Unmarshal(&conf)

	if conf.APIKey == "" {
		fmt.Print("Incorrect config file:")
		log.Fatalln(formatConfigErrorMessage)
	}
}

func assetsConfigFile(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatalln(formatConfigErrorMessage())
	}
}

func formatConfigErrorMessage() string {
	return fmt.Sprintf(`Create config file $HOME/.wthr.yaml or input him with option --config
Required info: API_KEY for %s`, weather.APIURL)
}
