package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckMandatoryArgument(t *testing.T) {

	tests := []struct {
		args    []string
		wantErr bool
	}{
		{[]string{"-cvmtoken testme -cvmaddress test123 -idmaddress idmaddress -idmtoken idmtoken"},
			false},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			var cfg Config
			fmt.Println(strings.Join(tt.args, " "))
			err := ParseFlags("prog", tt.args, &cfg)
			if err != nil {
				fmt.Println("Error")
			}

			//if err := CheckMandatoryArgument(&cfg); (err != nil) != tt.wantErr {
			//	t.Errorf("CheckMandatoryArgument() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}
