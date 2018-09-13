package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/uwuh/namecheck"
	"github.com/uwuh/namecheck/util"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
}

func init() {
	flag.Usage = usage
}

func main() {
	var (
		flagName = flag.String("name", "", "your name idea that you want to check")
		flagHelp = flag.Bool("help", false, "show this message")
	)

	flag.Parse()
	if len(flag.Args()) > 0 || *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *flagName == "" {
		flag.Usage()
		os.Exit(1)
	}

	red := color.New(color.FgRed).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()

	fmt.Printf("\nChecking ")
	yellow("%s\n", *flagName)
	fmt.Printf("Please wait...\n\n")

	if !util.IsConnected() {
		red("No internet available. Please check your connection.\n")
		os.Exit(1)
	}

	channels, duration := namecheck.Check(*flagName, namecheck.DefaultChannels)

	fmt.Printf("Status:\n")
	green("%s Available\n", namecheck.StatusAvailable)
	red("%s Not Available\n", namecheck.StatusNotAvailable)
	yellow("%s Unknown (e.g. Timeout)\n\n", namecheck.StatusUnknown)

	for _, channel := range channels {
		if channel.Status == namecheck.StatusAvailable {
			green("%s %s\n", namecheck.StatusAvailable, channel.Code)
		} else if channel.Status == namecheck.StatusNotAvailable {
			red("%s %s\n", namecheck.StatusNotAvailable, channel.Code)
		} else if channel.Status == namecheck.StatusUnknown {
			yellow("%s %s %s\n", namecheck.StatusUnknown, channel.Code, channel.Error.Error())
		}
	}

	fmt.Printf("\nDuration: %fs\n", duration.Seconds())
}
