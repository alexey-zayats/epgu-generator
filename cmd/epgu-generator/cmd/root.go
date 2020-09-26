package cmd

import (
	"epgu-generator/internal/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "epgu-generator",
	Short: "epgu-generator",
	Long:  "epgu-generator",
	Run:   rootMain,
}

// Execute entry point to cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Show cmd usage & exit
func usage(cmd *cobra.Command) {
	if err := cmd.Help(); err != nil {
		logrus.Fatal(err)
	}
	os.Exit(0)
}

func init() {

	cfgParams := []config.Param{
		{Name: "log-level", Value: "info", Usage: "log level", ViperBind: "Log.Level"},
		{Name: "workers", Value: 8, Usage: "number of workers", ViperBind: "Workers"},
	}

	viper.SetEnvPrefix(config.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config.Apply(rootCmd, cfgParams)

	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "./config/epgu-generator.yaml", "Config file")
	cobra.OnInitialize(config.Init)
}

func rootMain(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
