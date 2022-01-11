package main

import "testing"

func TestCheckMandatoryArgument(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckMandatoryArgument(); (err != nil) != tt.wantErr {
				t.Errorf("CheckMandatoryArgument() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
