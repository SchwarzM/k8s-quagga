package main

import (
  "os"
//  "fmt"
  "path"
  "text/template"
  "github.com/codegangsta/cli"
)

type ospfconfig struct {
  Password string
  Interface string
  RouterId string
  HomeNet string
  ContainerNet string
}

type zebraconfig struct {
  Password string
  PortalNet string
  PortalGw string
}

const zebraTemplate = `password {{.Password}}
log stdout

ip route {{.PortalNet}} {{.PortalGw}}

`

const ospfdTemplate = `
password {{.Password}}
enable password {{.Password}}
log stdout
interface {{.Interface}}
  ip ospf authentication-key {{.Password}}

router ospf
  ospf router-id {{.RouterId}}
  log-adjacency-changes detail
  default-information originate
  network {{.HomeNet}} area 0.0.0.0
  network {{.ContainerNet}} area 0.0.0.0
`

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  app := cli.NewApp()
  app.Version = "0.0.2"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "output",
      Value: "/etc/quagga",
      Usage: "Directory to output config to",
      EnvVar: "K8S_QUAGGA_OUTPUT",
    },
    cli.StringFlag{
      Name: "password",
      Value: "changeme",
      Usage: "Password for authorization",
      EnvVar: "K8S_QUAGGA_PASSWORD",
    },
  }
  app.Commands = []cli.Command{
    {
      Name: "ospfd",
      Usage: "output ospfd config",
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "Interface",
          Value: "eth0",
          Usage: "The interface to announce on",
        },
        cli.StringFlag{
          Name: "RouterId",
          Value: "10.0.0.1",
          Usage: "Router ID to announce",
        },
        cli.StringFlag{
          Name: "ContainerNet",
          Value: "10.2.0.0/24",
          Usage: "Container Network CIDR",
        },
        cli.StringFlag{
          Name: "HomeNet",
          Value: "10.0.1.0/24",
          Usage: "Home Network CIDR",
        },
      },
      Action: func(c *cli.Context) {
        config := ospfconfig{
          Password: c.GlobalString("password"),
          Interface: c.String("Interface"),
          RouterId: c.String("RouterId"),
          HomeNet: c.String("HomeNet"),
          ContainerNet: c.String("ContainerNet"),
        }
        f, err := os.Create(path.Join(c.GlobalString("output"), "ospfd.conf"))
        check(err)
        defer f.Close()
        t, err := template.New("config").Parse(ospfdTemplate)
        check(err)
        err = t.Execute(f, config)
        check(err)
      },
    },
    {
      Name: "zebra",
      Usage: "output zebra config",
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "PortalNet",
          Value: "10.2.0.0/24",
          Usage: "Portal Network CIDR",
        },
        cli.StringFlag{
          Name: "PortalGw",
          Value: "10.0.1.4",
          Usage: "Portal Network Gateway",
        },
      },
      Action: func(c *cli.Context) {
        config := zebraconfig{
          Password: c.GlobalString("password"),
          PortalNet: c.String("PortalNet"),
          PortalGw: c.String("PortalGw"),
        }
        f, err := os.Create(path.Join(c.GlobalString("output"), "zebra.conf"))
        check(err)
        defer f.Close()
        t, err := template.New("config").Parse(zebraTemplate)
        check(err)
        err = t.Execute(f, config)
        check(err)
      },
    },
  }
  app.Run(os.Args)
}

