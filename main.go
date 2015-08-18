package main

import (
  "os"
//  "fmt"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "output",
      Value: "/etc/quagga",
      Usage: "Directory to output config to",
      EnvVar: "K8S_QUAGGA_OUTPUT",
    },
  }
  app.Commands = []cli.Command{
    {
      Name: "ospfd",
      Usage: "output ospfd config",
      Action: func(c *cli.Context) {
        println("Printing ospfd config to ", c.GlobalString("output") )
      },
    },
    {
      Name: "zebra",
      Usage: "output zebra config",
      Action: func(c *cli.Context) {
        println("Printing zebra config to ", c.GlobalString("output") )
      },
    },
  }
  app.Run(os.Args)
}

