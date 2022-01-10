package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func ReturnValueFromPointer(pointer *string) string {

	if pointer != nil {
		return *pointer
	}
	return ""
}

func CheckMandatoryArgument() error {
	//Arguments can be provided by command line or environment variable
	requiredItems := [...]string{"cvmtoken", "cvmaddress", "idmtoken", "idmaddress"}
	envPrefix := "CVM_"

	for s := range requiredItems {
		found := false
		name := requiredItems[s]
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})
		if !found {
			key := fmt.Sprintf("%s%s", envPrefix, strings.ToUpper(name))
			val, present := os.LookupEnv(key)
			fmt.Println(key)
			if present {
				flag.Set(name, val)
			} else {
				return errors.New(fmt.Sprintf("Mandatory argument %s not found", name))
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("Staring process of syncronization")
	strnill := ""
	cvmtoken := flag.String("cvmtoken", "", "CVM token")
	cvmaddress := flag.String("cvmaddress", "", "CVM portal Address")
	cvmmode := flag.Bool("dryrun", true, "Dry run")
	idmtoken := flag.String("idmtoken", "", "IDM token")
	idmaddress := flag.String("idmaddress", "", "IDM address")

	flag.Parse()

	if ReturnValueFromPointer(cvmtoken) == strnill {
		fmt.Println("CVM Token not defined")
		flag.PrintDefaults()

	}
	fmt.Println(CheckMandatoryArgument())
	fmt.Println(*cvmtoken)
	fmt.Println(*cvmmode)
	fmt.Println(*cvmaddress)
	fmt.Println(*idmtoken)
	fmt.Println(*idmaddress)
	fmt.Println(CheckMandatoryArgument())
}
