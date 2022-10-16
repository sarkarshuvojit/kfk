/*
Copyright © 2022 Shuvojit

*/
package cmd

import "testing"

func TestHandler(t *testing.T) {
	type args struct {
		coorelationId string
		broker        string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.coorelationId, tt.args.broker)
		})
	}
}
