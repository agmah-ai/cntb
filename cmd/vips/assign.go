package cmd

import (
	"context"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var assignVipCmd = &cobra.Command{
	Use:     "vip [ip] [resourceId]",
	Short:   "Assign a VIP to a machine resource",
	Long:    `Assign a VIP to a VPS/VDS/Bare Metal using the machine id.`,
	Example: `cntb assign vip 127.0.0.1 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		_, httpResp, err := client.ApiClient().VIPApi.
			AssignIp(context.Background(), assignVipResourceId, assignVipIp, assignVipResourceType).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while adding vip to resource")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide an ip and resourceId.")
		}

		assignVipIp = args[0]

		resourceIdInt64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		assignVipResourceId = resourceIdInt64

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		assignVipResourceType = viper.GetString("type")

		validTypes := map[string]bool{"instances": true, "bare-metal": true}
		if !validTypes[assignVipResourceType] {
			log.Fatal("Invalid resourceType: %s. Allowed values: instances, bare-metal", assignVipResourceType)
		}

		return nil
	},
}

func init() {
	contaboCmd.AssignInstanceCmd.AddCommand(assignVipCmd)

	assignVipCmd.Flags().StringVarP(&assignVipResourceType, "type", "t", "",
		`Resource type (instances | bare-metal) (required)`)

	assignVipCmd.MarkFlagRequired("type")
	viper.BindPFlag("type", assignVipCmd.Flags().Lookup("type"))

}
