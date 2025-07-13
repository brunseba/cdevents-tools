package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cdevents-cli",
	Short: "A CLI tool for generating and sending CDEvents",
	Long: `CDEvents CLI is a command-line tool for generating and sending CDEvents 
into CI/CD toolchains using CloudEvents as transport.

CDEvents is a common specification for Continuous Delivery events, enabling
interoperability in the complete software production ecosystem.

This tool supports:
- Generating various CDEvents (pipeline, task, build, deployment, etc.)
- Sending events via HTTP, Kafka, or other transports
- Loading event templates from YAML files
- Integration with CI/CD systems`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cdevents-cli.yaml)")
	rootCmd.PersistentFlags().StringP("output", "o", "json", "output format (json, yaml, cloudevent)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	
	// Bind flags to viper
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cdevents-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cdevents-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Fprintf(os.Stderr, "Using config file: %s\n", viper.ConfigFileUsed())
		}
	}
}
