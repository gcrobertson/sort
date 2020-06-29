/*
 *	clear && go run main.go -sorts=bubble,counting,heap,insertion,merge,quick,radix,selection,shell -size=100000
 *
 *	clear && go run main.go -sorts=counting,radix -size=20 -range=99999 -natural=true
 *
 */

package main

import (
	"reflect"
	"testing"
)

func Test_validateCLISorts(t *testing.T) {

	*sorts = "bubble,merge,radix"

	tests := []struct {
		name   string
		sorts  string
		expect map[string]bool
	}{
		{
			name:  "test1",
			sorts: "bubble,merge,radix",
			expect: map[string]bool{
				"bubble":    true,
				"counting":  false,
				"heap":      false,
				"insertion": false,
				"merge":     true,
				"quick":     false,
				"radix":     true,
				"selection": false,
				"shell":     false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*sorts = tt.sorts
			validateCLISorts()

			res := reflect.DeepEqual(tt.expect, sortMap)

			if res != true {
				t.Errorf("got %v, want %v\n", sortMap, tt.expect)
			}
		})
	}
}
