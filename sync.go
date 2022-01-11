package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func CheckMandatoryArgument(conf *Config) error {
	//Arguments can be provided by command line or environment variable
	requiredItems := [...]string{"Cvmtoken", "Cvmaddress", "Idmtoken", "Idmaddress"}
	envPrefix := "CVM_"

	e := reflect.ValueOf(conf).Elem()

	for s := range requiredItems {
		va := e.FieldByName(requiredItems[s])
		if va.IsZero() {
			key := fmt.Sprintf("%s%s", envPrefix, strings.ToUpper(requiredItems[s]))
			val, present := os.LookupEnv(key)
			if present {
				e.FieldByName(requiredItems[s]).Set(reflect.ValueOf(val))
			} else {
				return errors.New(fmt.Sprintf("Mandatory argument %s not found", strings.ToLower(requiredItems[s])))

			}
		}
	}
	return nil
}

type Config struct {
	Cvmtoken   string
	Cvmaddress string
	dryrun     bool
	Idmtoken   string
	Idmaddress string
}

func ParseFlags(progname string, args []string, cfg *Config) (err error) {
	fmt.Println("Staring process of syncronization")
	flagSet := flag.NewFlagSet("ArgFlagset", flag.ExitOnError)
	var config = cfg

	flagSet.StringVar(&config.Cvmtoken, "cvmtoken", "", "CVM token")
	flagSet.StringVar(&config.Cvmaddress, "cvmaddress", "", "CVM portal Address")
	flagSet.BoolVar(&config.dryrun, "dryrun", true, "Dry run")
	flagSet.StringVar(&config.Idmtoken, "idmtoken", "", "IDM token")
	flagSet.StringVar(&config.Idmaddress, "idmaddress", "", "IDM address")

	err = flagSet.Parse(args)
	//	fmt.Println("test")
	//	fmt.Println(args)
	if err != nil {
		flagSet.PrintDefaults()
		return fmt.Errorf("Please try again")
	}

	return nil

}

func main() {
	var cfg Config
	ParseFlags(os.Args[0], os.Args[1:], &cfg)
	err := CheckMandatoryArgument(&cfg)
	if err == nil {
		fmt.Println("ok")
	}

}
