package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

// imports as package "cli"

func main() {
	app := cli.NewApp()
	app.Name = "Simple DNS Lookup Cli"
	app.Usage = "Used for querying IPS, MX records, and Nameservers!"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "cloudkite.io",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Look up nameservers for a particular host",
			Flags: flags,

			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},

		{
			Name:  "ip",
			Usage: "Returns the IP address for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},

		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},

		{
			Name:  "mx",
			Usage: "Returns the MX records for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
