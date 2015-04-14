package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

const WEBHOOK_FORMAT = "%4v  %-20s  %-6s  %-s"

func main() {
	app := cli.NewApp()
	app.Name = "shopiclient"
	app.Usage = "Shopify CLI API client"
	app.Before = SetupClient
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Be verbose",
		},
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
			Name:   "channels",
			Action: ChannelsDefault,
		},
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
				cli.Command{
					Name:   "auto-test",
					Usage:  "Automatically set up a webhook for the given topic and start a server to listen",
					Action: AutoTestWebhook,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "topic",
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
	shopifyClient.Verbose = context.IsSet("verbose")
	return nil
}

func WebhooksDefault(context *cli.Context) {
	if len(context.Args()) > 0 {
		id, err := strconv.Atoi(context.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
		webhook, _ := shopifyClient.Webhooks().get(id)

		prettyListWebhooks(webhook)
	}
}

func prettyListWebhooks(hooks ...*Webhook) {
	fmt.Printf(WEBHOOK_FORMAT, "ID", "Topic", "Format", "Address")
	fmt.Println()
	for _, webhook := range hooks {
		fmt.Printf(WEBHOOK_FORMAT, webhook.Id, webhook.Topic, webhook.Format, webhook.Address)
		fmt.Println()
	}

}

func AutoTestWebhook(context *cli.Context) {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	format := "json"
	topic := context.String("topic")

	u, err := url.Parse(fmt.Sprintf("http://%s:8080", hostname))
	if err != nil {
		log.Fatal(err)
	}

	webhook, _ := shopifyClient.Webhooks().create(topic, u, format)
	fmt.Println("Created new webhook for automatic testing:")
	prettyListWebhooks(webhook)

	log.Println("Now listening for webhooks, press ^C to exit...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go interruptHandler(c, webhook)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func interruptHandler(c chan os.Signal, webhook *Webhook) {
	for sig := range c {
		log.Println("Caught ^C ", sig.String())
		shopifyClient.Webhooks().delete(webhook.Id)

		os.Exit(0)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println()
	log.Println("--< Incoming Request >---------------------------------------------------------")
	fmt.Printf("%s", b)
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

	webhook, _ := shopifyClient.Webhooks().create(topic, u, format)
	fmt.Println("Created new webhook:")
	prettyListWebhooks(webhook)
}

func ListWebhooks(context *cli.Context) {

	webhooks := shopifyClient.Webhooks()
	hooks := webhooks.list()
	fmt.Printf("Registered webhooks: %d (you only see webhooks registered with the current credentials)\n", len(hooks))

	prettyListWebhooks(hooks...)
}

func DeleteWebhook(context *cli.Context) {
	webhooks := shopifyClient.Webhooks()

	id, err := strconv.Atoi(context.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	webhooks.delete(id)
}

func ChannelsDefault(context *cli.Context) {
	channels := shopifyClient.Channels()
	ch := channels.List()

	fmt.Printf("Found %d channels.\n", len(ch))
	fmt.Printf("%-5s %-10s\n", "ID", "Provider ID")
	for _, c := range ch {
		fmt.Printf("%-5d %-10d\n", c.Id, c.ProviderId)
	}
}
