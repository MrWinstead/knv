package cmd

import (
	"context"
	"fmt"
	"github.com/mrwinstead/knv/configuration"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use: "knv",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			registerSignalHandler()
		},
	}

	configuredViper   *viper.Viper
	rootContext       context.Context
	rootContextCancel context.CancelFunc
	monotonicIDSource io.Reader
)

func init() {
	configuredViper = viper.GetViper()
	configureViperOrFatal(configuredViper)

	rootContext, rootContextCancel = context.WithCancel(context.Background())
	monotonicIDSource = newULIDMonotonicSource()
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
	configuration.PopulateDefaultValues(v)

	rootDbDirFlag := configuration.ConfigKeyToCLIArgument(
		configuration.KeyDirectoryNameRootDatabase)
	rootCmd.PersistentFlags().StringP(rootDbDirFlag, "d", "",
		"the root directory for tables and indexes")
	markFlagRequiredOrFatal(rootCmd, rootDbDirFlag)
	bindPflagOrFatal(v, configuration.KeyDirectoryNameRootDatabase,
		rootCmd.PersistentFlags().Lookup(rootDbDirFlag))
}

func markFlagRequiredOrFatal(cmd *cobra.Command, flagName string) {
	markRequiredErr := cmd.MarkPersistentFlagRequired(flagName)
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
