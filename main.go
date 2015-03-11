package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "shopiclient"
	app.Usage = "Shopify CLI API client"
	app.Before = SetupClient
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "user",
			EnvVar: "SHOPIFY_USER",
		},
		cli.StringFlag{
			Name:   "password",
			EnvVar: "SHOPIFY_PASSWORD",
		},
		cli.StringFlag{
			Name:   "host",
			EnvVar: "SHOPIFY_HOST",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "webhooks",
			Action: WebhooksDefault,
			Subcommands: []cli.Command{
				cli.Command{
					Name:   "list",
					Usage:  "Lists registered webhooks",
					Action: ListWebhooks,
				},
				cli.Command{
					Name:   "delete",
					Usage:  "Deletes a webhook",
					Action: DeleteWebhook,
					Flags: []cli.Flag{
						cli.IntFlag{
							Name: "id",
						},
					},
				},
				cli.Command{
					Name:   "create",
					Usage:  "Registers a new webhook",
					Action: CreateWebhook,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "topic",
						},
						cli.StringFlag{
							Name: "address",
						},
						cli.StringFlag{
							Name: "format",
						},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

var shopifyClient *ShopifyClient

func SetupClient(context *cli.Context) error {
	shopifyClient = Connect(context.String("host"), context.String("user"), context.String("password"))
	return nil
}

func WebhooksDefault(context *cli.Context) {
	if len(context.Args()) > 0 {
		id, err := strconv.Atoi(context.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
		webhook, _ := shopifyClient.Webhooks().get(id)

		fmt.Println("Got webhook", webhook)
	}
}

func CreateWebhook(context *cli.Context) {
	u, err := url.Parse(context.String("address"))
	if err != nil {
		log.Fatal(err)
	}

	format := strings.Trim(context.String("format"), " ")
	if format != "json" && format != "xml" {
		log.Fatalf("Invalid format %s, expected either \"json\" or \"xml\"!", format)
	}

	topic := context.String("topic")

	shopifyClient.Webhooks().create(topic, u, format)
}

func ListWebhooks(context *cli.Context) {

	webhooks := shopifyClient.Webhooks()
	hooks := webhooks.list()
	format := "%4v  %-20s  %-6s  %-s\n"
	fmt.Printf("Registered webhooks: %d (you only see webhooks registered with the current credentials)\n", len(hooks))
	fmt.Printf(format, "ID", "Topic", "Format", "Address")

	for _, webhook := range hooks {
		fmt.Printf(format, webhook.Id, webhook.Topic, webhook.Format, webhook.Address)
	}
}

func DeleteWebhook(context *cli.Context) {
	webhooks := shopifyClient.Webhooks()
	webhooks.delete(context.Int("id"))
}
