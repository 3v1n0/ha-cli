package cmd

import (
	"fmt"

	helper "github.com/home-assistant/hassio-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var hassosImportCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"im", "sync"},
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("hassos import")

		section := "hassos"
		command := "config/sync"
		base := viper.GetString("endpoint")

		resp, err := helper.GenericJSONPost(base, section, command, nil)
		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	hassosCmd.AddCommand(hassosImportCmd)
}
