// Package commands is the command line boot loader"
package commands

import (
	"github.com/spf13/cobra"
	"go-mongo/commands/actions"

	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// CommandEngine is the structure of cli
type CommandEngine struct {
	rootCmd *cobra.Command
}

// NewCommandEngine the command line boot loader
func NewCommandEngine() *CommandEngine {
	var rootCmd = &cobra.Command{
		Use:   config.GetString("app.command"),
		Short: config.GetString("app.name") + " command line",
		Long:  config.GetString("app.name") + " command line",
	}
	defer func() {
		r := recover()
		if r != nil {
			log.Error(r)
		}
	}()

	rootCmd.PersistentFlags().StringP("config", "c", "configurations", "the config path location")

	//rootCmd.Execute()
	return &CommandEngine{
		rootCmd: rootCmd,
	}
}

// GetRoot the command line service
func (c *CommandEngine) GetRoot() *cobra.Command {
	return c.rootCmd
}

// Run the all command line
func (c *CommandEngine) Run() {

	var commands = []*cobra.Command{
		// this for run server by commands
		{
			Use:   "serve",
			Short: config.GetString("app.name") + " Listening HTTP server",
			Long:  config.GetString("app.name") + " Listening HTTP server",
			Run: func(cmd *cobra.Command, args []string) {
				actions.StartServer()
			},
		},
	}
	for _, command := range commands {
		c.rootCmd.AddCommand(command)
	}
	c.rootCmd.Execute()
}
