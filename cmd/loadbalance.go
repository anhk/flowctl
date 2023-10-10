package main

import (
	"fmt"

	"github.com/anhk/flowctl/app"
	"github.com/spf13/cobra"
)

var LbOption struct {
	Service string
	Target  string
}

var lbAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "ad"},
	Short:   "Add",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("add loadbalance %+v", LbOption)
		return nil
	},
}

var lbDelCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"d", "de"},
	Short:   "Del",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("del loadbalance %+v", LbOption)
		return nil
	},
}

var lbClrCmd = &cobra.Command{
	Use:     "clear",
	Aliases: []string{"c", "cl", "cle", "clea"},
	Short:   "Clear",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Clear loadbalance ")
		return nil
	},
}

var lbSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.NewLoadbalance().Setup()
	},
}

var lbCmd = &cobra.Command{
	Use:     "loadbalance",
	Aliases: []string{"lb"},
	Short:   "LoadBalance",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("loadbalance show ")
		return nil
	},
}

func init() {
	// add
	lbAddCmd.PersistentFlags().StringVarP(&LbOption.Service, "service", "s", "", "Servcie")
	lbAddCmd.PersistentFlags().StringVarP(&LbOption.Target, "target", "t", "", "Target")
	lbCmd.AddCommand(lbAddCmd)

	// del
	lbDelCmd.PersistentFlags().StringVarP(&LbOption.Service, "service", "s", "", "Servcie")
	lbDelCmd.PersistentFlags().StringVarP(&LbOption.Target, "target", "t", "", "Target")
	lbCmd.AddCommand(lbDelCmd)

	// clear
	lbCmd.AddCommand(lbClrCmd)

	// setup
	lbCmd.AddCommand(lbSetupCmd)

	// loadbalance
	rootCmd.AddCommand(lbCmd)
}
