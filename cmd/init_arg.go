package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	DB_NAME  string
	DB_USER  string
	DB_PWD   string
	DB_URL   string
	Duration int64
	Port     string
	GrpcPort string
	rootCmd  = &cobra.Command{
		Use:   "traffic controller",
		Short: "traffic controller",
		Long:  "traffic controller to manager server ,chain",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&DB_NAME, "name", "applier", "db name")
	rootCmd.PersistentFlags().StringVar(&DB_USER, "user", "root", "db user")
	rootCmd.PersistentFlags().StringVar(&DB_PWD, "pwd", "Admin@123", "password")
	rootCmd.PersistentFlags().StringVar(&DB_URL, "url", "host.docker.internal:3306", "db url")
	rootCmd.PersistentFlags().Int64Var(&Duration, "duration", 5, "cache duration")
	rootCmd.PersistentFlags().StringVar(&Port, "port", "8088", "listener ports")
	rootCmd.PersistentFlags().StringVar(&GrpcPort, "gport", ":10080", "listener ports of grpc")
}
