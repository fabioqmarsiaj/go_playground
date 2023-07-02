package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the application cli to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search for IPs and Server names on internet"

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search for IPs on internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context){
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}