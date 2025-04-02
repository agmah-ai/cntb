package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vipsGetCmd = &cobra.Command{
	Use:     "vips",
	Short:   "All about your Virtual IPs (vips).",
	Long:    `Retrieves information about one or multiple virtual IPs. Filter by name.`,
	Example: `cntb get vips`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveVipListRequest := client.ApiClient().
			VIPApi.RetrieveVipList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listVipNameFilter != "" {
			ApiRetrieveVipListRequest = ApiRetrieveVipListRequest.ResourceName(listVipNameFilter)
		}

		resp, httpResp, err := ApiRetrieveVipListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving vips")

		arr := make([]jmap, 0)
		for _, entry := range resp.Data {
			entryModified, _ := util.StructToMap(entry)
			entryModified["ipv4"] = entry.V4.Ip
			entryModified["net"] = entry.V4.Net
			entryModified["name"] = entry.ResourceName
			entryModified["displayName"] = entry.ResourceDisplayName
			arr = append(arr, entryModified)
		}

		responseJson, _ := json.Marshal(arr)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"ipv4", "type", "net", "resourceId", "name", "displayName", "region", "dataCenter", "vipId",

			},
			WideFilter: []string{
				"ipv4", "type", "net", "resourceId", "name", "displayName", "region", "dataCenter", "vipId", "tenantId", "customerId",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listVipNameFilter = viper.GetString("name")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(vipsGetCmd)

	vipsGetCmd.Flags().StringVarP(&listVipNameFilter, "name", "n", "",
		`Filter by instance name`)
}
