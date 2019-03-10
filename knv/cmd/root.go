package cmd

import (
	"fmt"
	"github.com/mrwinstead/knv/configuration"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "knv",
}

var configuredViper *viper.Viper

func init() {
	configuredViper = viper.GetViper()
	configureViperOrFatal(configuredViper)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureViperOrFatal(v *viper.Viper) {
	v.AutomaticEnv()

	rootDbDirFlag := configuration.ConfigKeyToCLIArgument(
		configuration.KeyDirectoryNameRootDatabase)
	rootCmd.PersistentFlags().StringP(rootDbDirFlag, "-r", "",
		"the root directory for tables and indexes")
	markFlagRequiredOrFatal(rootCmd, rootDbDirFlag)
	bindPflagOrFatal(v, configuration.KeyDirectoryNameRootDatabase,
		rootCmd.PersistentFlags().Lookup(rootDbDirFlag))
}

func markFlagRequiredOrFatal(cmd *cobra.Command, flagName string) {
	markRequiredErr := cmd.MarkFlagRequired(flagName)
	if nil != markRequiredErr {
		log.Fatal(markRequiredErr, "could not mark CLI flag a required, ",
			flagName)
	}
}

func bindPflagOrFatal(v *viper.Viper, configKey string, flag *pflag.Flag) {
	bindErr := v.BindPFlag(configKey, flag)
	if nil != bindErr {
		log.Fatal(bindErr, "could not bind configuration key to CLI flag")
	}
}
