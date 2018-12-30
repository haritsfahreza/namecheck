package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/haritsfahreza/namecheck"
	"github.com/haritsfahreza/namecheck/util"
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
	ctx := context.Background()
	var (
		flagName        = flag.String("name", "", "your name idea")
		flagChannelCode = flag.String("code", "", "specify channel code")
		flagChannelType = flag.String("type", "", "specify channel type ('domain' & 'social')")
		flagHelp        = flag.Bool("help", false, "show this message")
		flagChannelList = flag.Bool("list", false, "show available channel list")
	)

	flag.Parse()
	if len(flag.Args()) > 0 || *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	if len(flag.Args()) > 0 || *flagChannelList {
		fmt.Printf("\nList of available channel codes:\n\n")
		for _, channel := range namecheck.DefaultChannels {
			fmt.Printf("   %s\n", channel.Code)
		}
		os.Exit(0)
	}

	if *flagName == "" {
		flag.Usage()
		os.Exit(1)
	}

	red := color.New(color.FgRed).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()

	channels := namecheck.DefaultChannels
	if *flagChannelCode != "" {
		ch, err := namecheck.FindChannelByCode(ctx, *flagChannelCode)
		if err != nil {
			red(err.Error())
			os.Exit(1)
		}
		channels = []*namecheck.Channel{ch}
	}

	if *flagChannelType != "" {
		if *flagChannelType == "domain" {
			channels = namecheck.FindChannelsByType(ctx, namecheck.TypeDomain)
		} else if *flagChannelType == "social" {
			channels = namecheck.FindChannelsByType(ctx, namecheck.TypeSocial)
		} else {
			red("No channel type found. Please use the available channel type")
			os.Exit(1)
		}
	}

	fmt.Printf("\nChecking ")
	yellow("%s\n", *flagName)
	fmt.Printf("Please wait...\n\n")

	if !util.IsConnected() {
		red("No internet available. Please check your connection.\n")
		os.Exit(1)
	}

	channels, duration := namecheck.Check(ctx, *flagName, channels)

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
