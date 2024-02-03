// command line interface for the electromap package
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	em "gopkg.in/gerald-scharitzer/electromap.v0"
)

const (
	usageHead = `Usage: electromap [-f format] [-h] [-v] [{ health | help | version | zones }] [-]`
	usageTail = `  -     process standard input
  health
        print the API health
  help  print the command help
  version
        print the module version
  zones
        print the zones`
)

func usage() {
	_, err := fmt.Println(usageHead)
	if err != nil {
		panic(err.Error())
	}
	flag.PrintDefaults()
	_, err = fmt.Println(usageTail)
	if err != nil {
		panic(err.Error())
	}
}

func version() {
	_, err := fmt.Println(em.VersionString())
	if err != nil {
		panic(err.Error())
	}
}

func main() { // TODO set exit code
	flag.Usage = usage
	var format = flag.String("f", "", `format output as: csv`)
	var hFlag = flag.Bool("h", false, "print the command help")
	var vFlag = flag.Bool("v", false, "print the module version")
	flag.Parse()
	if *hFlag {
		usage()
	}
	if *vFlag {
		version()
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
		_, err = fmt.Println(health.Monitors.State, health.Status)
		if err != nil {
			panic(err.Error())
		}
	case "help":
		usage()
	case "version":
		version()
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
				_, err = fmt.Println(key, zone.Country, zone.Name)
				if err != nil {
					panic(err.Error())
				}
			}
		}
	default:
		usage()
		return
	}
	// TODO fmt.Println("apiRoot", session.ApiRoot)
	// TODO fmt.Println(len(session.AuthToken) > 0)
}
