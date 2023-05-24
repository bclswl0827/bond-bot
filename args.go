package main

import "flag"

type Args struct {
	Path    string
	Version bool
}

func (args *Args) ReadFlags() {
	flag.StringVar(&args.Path, "config", "./config.json", "Path to config file")
	flag.BoolVar(&args.Version, "version", false, "Show version information")
	flag.Parse()
}
