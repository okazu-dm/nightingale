package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/okazu-dm/nightingale/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "nightingale",
	Short: "Notify something to somewhere",
	Long: `Notify something to somewhere.
Use config.json or command line options.`,
	Run: func(c *cobra.Command, args []string) {
		stdin := bufio.NewReader(os.Stdin)
		inp, err := ioutil.ReadAll(stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		rendered, err := cmd.Render(&cfg, inp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
			os.Exit(1)
		}
		if err := cmd.Notice(&cfg, rendered); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func main() {
	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string
var cfg cmd.Config

const EnvPrefix string = "NIGHTINGALE"

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default: ./config.yml)")
	rootCmd.PersistentFlags().StringP("template_path", "t", "default", "Template file path")
	rootCmd.PersistentFlags().StringP("slack_url", "", "", "Slack Channel to notify something.")
	viper.BindPFlag("template_path", rootCmd.PersistentFlags().Lookup("template_path"))
	viper.BindPFlag("slack.url", rootCmd.PersistentFlags().Lookup("slack_url"))
	viper.BindEnv("slack.url", EnvPrefix+"_SLACK_URL")
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(EnvPrefix)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		os.Exit(1)
	}
	viper.Unmarshal(&cfg)
}
