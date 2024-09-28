package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/LeaYeh/h1spec"
	"github.com/LeaYeh/h1spec/config"
)

var (
	VERSION string = "0.0.0"
	COMMIT  string = "(Unknown)"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "h1spec [spec...]",
		Short: "Conformance testing tool for HTTP/1.1 implementation",
		Long:  "Conformance testing tool for HTTP/1.1 implementation.",
		RunE:  run,
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	flags := cmd.Flags()
	flags.StringP("host", "h", "127.0.0.1", "Target host")
	flags.IntP("port", "p", 0, "Target port")
	flags.StringP("path", "P", "/", "Target path")
	flags.IntP("timeout", "o", 2, "Time seconds to test timeout")
	flags.Int("max-header-length", 4000, "Maximum length of HTTP header")
	flags.Int("max-body-length", 4000, "Maximum length of HTTP body")
	flags.BoolP("strict", "S", false, "Run all test cases including strict test cases")
	flags.Bool("dryrun", false, "Display only the title of test cases")
	flags.BoolP("verbose", "v", false, "Output verbose log")
	flags.Bool("version", false, "Display version information and exit")
	flags.Bool("help", false, "Display this help and exit")

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	v, err := flags.GetBool("version")
	if err != nil {
		return err
	}

	if v {
		version()
		return nil
	}

	host, err := flags.GetString("host")
	if err != nil {
		return err
	}

	port, err := flags.GetInt("port")
	if err != nil {
		return err
	}

	path, err := flags.GetString("path")
	if err != nil {
		return err
	}

	timeout, err := flags.GetInt("timeout")
	if err != nil {
		return err
	}

	maxHeaderLen, err := flags.GetInt("max-header-length")
	if err != nil {
		return err
	}

	maxBodyLen, err := flags.GetInt("max-body-length")
	if err != nil {
		return err
	}

	strict, err := flags.GetBool("strict")
	if err != nil {
		return err
	}

	dryRun, err := flags.GetBool("dryrun")
	if err != nil {
		return err
	}

	verbose, err := flags.GetBool("verbose")
	if err != nil {
		return err
	}

	if port == 0 {
		port = 80
	}

	c := &config.Config{
		AgentName:   "tester_h1spec_" + VERSION,
		Host:         host,
		Port:         port,
		Path:         path,
		Timeout:      time.Duration(timeout) * time.Second,
		MaxHeaderLen: maxHeaderLen,
		MaxBodyLen:   maxBodyLen,
		Strict:       strict,
		DryRun:       dryRun,
		Verbose:      verbose,
		Sections:     args,
	}

	success, err := h1spec.Run(c)
	if !success {
		os.Exit(1)
	}

	return err
}

func version() {
	fmt.Printf("Version: %s (%s)\n", VERSION, COMMIT)
}
