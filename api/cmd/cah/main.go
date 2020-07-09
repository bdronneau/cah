package main

import (
	"flag"
	"log"
	"os"

	"cah/pkg/handlers"

	"github.com/peterbourgon/ff"
	"github.com/sirupsen/logrus"
)

func main() {
	fs := flag.NewFlagSet("cah", flag.ExitOnError)

	var (
		loglevel    = fs.String("log-level", "DEBUG", "Log level")
		environment = fs.String("env", "prod", "run level for app")
	)

	cahConfig := handlers.Flags(fs)

	err := ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
		ff.WithEnvVarPrefix("CAH"),
	)

	if err != nil {
		log.Fatalf("error while parsing flags: %v", err)
	}

	if err := configLogger(*loglevel); err != nil {
		log.Fatalf("unable to configure logger: %v", err)
	}

	app, _ := handlers.New(cahConfig)

	app.NewHTTP(*environment)
}

func configLogger(logLevel string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	logrusLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrusLevel)
	logrus.SetReportCaller(true)

	return nil
}
