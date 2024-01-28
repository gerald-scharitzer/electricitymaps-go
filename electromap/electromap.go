package main

import (
	"encoding/csv"
	"flag"
	"os"

	em "gopkg.in/gerald-scharitzer/electromap.v0"
)

const (
	usageHead = `Usage: electromap [-h] [{ health | help | zones }] [-]`
	usageTail = `  -     process standard input
  health
        print the API health
  help  print the command help
  zones
        print the zones`
)

func usage() {
	println(usageHead)
	flag.PrintDefaults()
	println(usageTail)
}

func main() {
	flag.Usage = usage
	var format = flag.String("f", "", `format output as: csv`)
	var hFlag = flag.Bool("h", false, "print the command help")
	flag.Parse()
	if *hFlag {
		usage()
	}

	// TODO var session em.Session
	switch flag.Arg(0) {
	case "":
		// TODO session = em.GetSession()
	case "-": // TODO
		// sessionp, err := em.GetSessionFromStdin()
		// if err != nil {
		// 	panic(err.Error())
		// }
		// session = *sessionp
	case "health":
		health, err := em.GetHealth(nil)
		if err != nil {
			panic(err.Error())
		}
		println(health.Monitors.State, health.Status)
	case "help":
		usage()
	case "zones":
		zones, err := em.GetZones(nil)
		if err != nil {
			panic(err.Error())
		}
		var csvWriter *csv.Writer
		if *format == "csv" {
			csvWriter = csv.NewWriter(os.Stdout)
			err = csvWriter.Write([]string{"Key", "Country", "Name"})
			if err != nil {
				panic(err.Error())
			}
		}
		for key, zone := range *zones {
			if *format == "csv" {
				err = csvWriter.Write([]string{key, zone.Country, zone.Name})
				if err != nil {
					panic(err.Error())
				}
				csvWriter.Flush()
				err = csvWriter.Error()
				if err != nil {
					panic(err.Error())
				}
			} else {
				println(key, zone.Country, zone.Name)
			}
		}
	default:
		usage()
		return
	}
	// TODO println("apiRoot", session.ApiRoot)
	// TODO println(len(session.AuthToken) > 0)
}
