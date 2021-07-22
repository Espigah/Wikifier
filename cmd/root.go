package cmd

import (
	"fmt"
	"os"

	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose *bool //nolint: gochecknoglobals

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{ //nolint: gochecknoglobals
	Use:   "wikifier",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// ctx := context.Background()
	// if err := rootCmd.ExecuteContext(ctx); err != nil {
	// 	logrus.Errorf("Failed to run command %s", err)
	// 	os.Exit(1)
	// }
}

func init() { //nolint: gochecknoinits
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	verbose = rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Run in verbose mode")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// rootCmd.AddCommand(aws.RootCmdAws)
}

func AddCommand(cmds *cobra.Command) {
	rootCmd.AddCommand(cmds)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName("")    // name of config file (without extension)
	viper.SetConfigType("wkf") // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("/etc/wikifier/")   // path to look for the config file in
	// viper.AddConfigPath("$HOME/.wikifier/") // call multiple times to add many search paths
	viper.AddConfigPath(".") // optionally look for config in the working directory

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		err = viper.Unmarshal(&app.Config)
		if err != nil { // Handle errors reading the config file
			panic(err.Error())
		}

	}

	// fmt.Printf("%+v\n", Config)
	// fmt.Printf("%+v\n", viper.GetString("teste"))
	// fmt.Printf("%+v\n", viper.GetString("DBDriver"))
}
