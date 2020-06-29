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

func Test_validateCLISize(t *testing.T) {

	tests := []struct {
		name     string
		size     int
		xrange   int
		expsize  int
		exprange int
	}{
		{
			name:     "test1",
			size:     1000000,
			xrange:   9999,
			expsize:  1000000,
			exprange: 9999,
		},
		{
			name:     "test2",
			size:     -20,
			xrange:   1000000,
			expsize:  10,
			exprange: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			*size = tt.size
			*xrange = tt.xrange

			validateCLISize()

			if *size != tt.expsize {
				t.Errorf("got %v, want %v\n", *size, tt.expsize)
			}

			if *xrange != tt.exprange {
				t.Errorf("got %v, want %v\n", *xrange, tt.exprange)
			}
		})
	}
}

func Test_initializeIntSlice(t *testing.T) {

	tests := []struct {
		name     string
		size     int
		xrange   int
		xnatural bool
		want     []int
	}{
		{
			name:     "test1",
			size:     10,
			xrange:   5,
			xnatural: true,
			want:     make([]int, 10),
		},
		{
			name:     "test2",
			size:     100,
			xrange:   9990,
			xnatural: false,
			want:     make([]int, 100),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			*size = tt.size
			*xrange = tt.xrange
			*xnatural = tt.xnatural

			initializeIntSlice()

			min, max := PresortSlice[0], PresortSlice[0]
			for i := 0; i < len(PresortSlice); i++ {
				if min < PresortSlice[i] {
					min = PresortSlice[i]
				} else if max > PresortSlice[i] {
					max = PresortSlice[i]
				}
			}

			if len(PresortSlice) != len(tt.want) {
				t.Errorf("initializeIntSlice() = %v, want %v", PresortSlice, tt.want)
			}
			if max > tt.xrange {
				t.Errorf("range out of bound = %v, want <= %v", max, tt.xrange)
			}
			if tt.xnatural == true && min < 0 {
				t.Errorf("range covers negative numbers = %v, want only natural numbers %v", min, tt.xnatural)
			}
		})
	}
}

func Test_setupSortInfoSlice(t *testing.T) {

	tests := []struct {
		name     string
		sorts    string
		size     int
		xrange   int
		xnatural bool
	}{
		{
			name:     "test1",
			sorts:    "bubble,insertion,merge,selection,radix",
			size:     100,
			xrange:   999,
			xnatural: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*sorts = tt.sorts
			*size = tt.size
			*xrange = tt.xrange
			*xnatural = tt.xnatural
			validateCLISorts()
			initializeIntSlice()
			setupSortInfoSlice()

			if len(SortInfoSlice) != 5 {
				t.Errorf("got SortInfoSlice length %v, want %v", len(SortInfoSlice), 5)
			}
			for i := 0; i < len(SortInfoSlice); i++ {
				if !reflect.DeepEqual(SortInfoSlice[i].ArgumentToFunc, PresortSlice) {
					t.Errorf("got presort slice = %v, want <= %v", SortInfoSlice[i-1].ArgumentToFunc, PresortSlice)
				}
			}
		})
	}
}
