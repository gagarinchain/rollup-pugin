package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gagarinchain/rollup-plugin/plugin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var address string

// startCmd represents the start command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		for k, v := range viper.AllSettings() {
			println(k)
			spew.Dump(v)
		}

		s := &plugin.Settings{}
		if err := viper.Unmarshal(s); err != nil {
			return
		}

		spew.Dump(s)

		plugin.Send(address)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
