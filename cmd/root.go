package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/toyz/givenv/providers"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var provider string

var rootCmd = &cobra.Command{
	Use:   "givenv",
}

var secretsProvider providers.ProviderInterface

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.givenv.yaml)")
	rootCmd.PersistentFlags().StringVar(&provider, "provider", "aws", "Secrets Manager provider")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".givenv")
	}

	secretsProvider = providers.InitProvider(provider)
	if secretsProvider == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Secrets provider does not exist")
		_ = rootCmd.Help()
		os.Exit(1)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
