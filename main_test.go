package main

import "testing"

func TestMyFunc(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MyFunc()
			if (err != nil) != tt.wantErr {
				t.Errorf("MyFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
