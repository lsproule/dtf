package cmd

import (
	"fmt"

	"github.com/lsproule/dtf/src/rules"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "dtf",
		Short: "dtf is a deterministic testing  framework",
		Long: `dtf is a deterministic testing framework that allows you to write tests that are deterministic and repeatable.
Complete documentation is available at  github.com/lsproule/dtf`,
		Run: func(cmd *cobra.Command, args []string) {
			rules := &rules.RuleSets{} 
			err:= viper.Unmarshal(rules) 
			if err != nil { 
				fmt.Printf("unable to decode into struct, %v", err)  
			}
			fmt.Println(rules) 

		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile , "config", "f", "", "config file (default is ./api.yaml)") 


}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("./api.yaml")
		viper.AddConfigPath(".")         // Look for the file in the current directory.

	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
