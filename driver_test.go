package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/docker/go-plugins-helpers/volume"
)

func TestNewExampleDriver(t *testing.T) {
	testcases := []struct {
		name string
		want ExampleDriver
	}{

		{
			name: "test-positive",
			want: ExampleDriver{
				volumes:    make(map[string]string),
				m:          &sync.Mutex{},
				mountPoint: "/tmp/exampledriver",
			},
		},
	}
	for _, tt := range testcases {
		if got := NewExampleDriver(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewExampleDriver() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_Create(t *testing.T) {
	// want := volume.Response{}
	testcases := volume.Request{
		Name:    "myexamplevol",
		Options: make(map[string]string),
	}
	var d ExampleDriver
	d.Create(testcases)
	// t.Errorf("%v is the response  is the other", want)
	{
		// TODO
	}

}
