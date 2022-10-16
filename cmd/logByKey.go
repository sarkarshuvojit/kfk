/*
Copyright © 2022 Shuvojit

*/
package cmd

import (
	"log"

	"github.com/sarkarshuvojit/kfk/util/kfkconn"
	"github.com/spf13/cobra"
)

var coorelationId string
var brokerUrl string

func handler(coorelationId string, broker string) {
    messages := kfkconn.GetMessages(broker, "myTopic")

    for _, msg := range messages {
        log.Printf("%s %s\n", msg.Value, msg.Key)
    }

}


var logByKeyCmd = &cobra.Command{
	Use:   "logByKey",
	Short: "Will log messages across multiple topics.",
	Long: `Will log messages accross multiple topics sorted by time having the same key as common. 

Can be utilized to debug or trace how an event flows across different topics for an event triggered by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
        handler(coorelationId, brokerUrl)
	},
}

func init() {
	rootCmd.AddCommand(logByKeyCmd)
    logByKeyCmd.Flags().StringVarP(
        &coorelationId,
        "key",
        "k",
        "",
        "Pass the key which will be used to source the events.",
    )
    logByKeyCmd.MarkFlagRequired("key")

    logByKeyCmd.Flags().StringVarP(
        &brokerUrl,
        "broker",
        "b",
        "localhost:9092",
        "Pass the broker url in <host>:<port> format. Defaults to: localhost:9092",
    )
    logByKeyCmd.MarkFlagRequired("key")
}
