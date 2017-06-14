package main

import (
	"github.com/docker/go-plugins-helpers/volume"
	"reflect"
	"sync"
	"testing"
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

func TestExampleDriver_Create(t *testing.T) {
	type fields struct {
		volumes    map[string]string
		m          *sync.Mutex
		mountPoint string
	}
	type args struct {
		r volume.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   volume.Response
	}{
		{
			name: "test-1",
			fields: fields{
				volumes:    make(map[string]string),
				m:          &sync.Mutex{},
				mountPoint: "/tmp/exampledriver/",
			},
			args: args{
				r: volume.Request{
					Name:    "myvolumename",
					Options: make(map[string]string),
				},
			},
			want: volume.Response{},
		},
		{
			name: "test-2",
			fields: fields{
				volumes:    make(map[string]string),
				m:          &sync.Mutex{},
				mountPoint: "/tmp/exampledriver/",
			},
			args: args{
				r: volume.Request{
					Name:    "newvolumename",
					Options: make(map[string]string),
				},
			},
			want: volume.Response{},
		},
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Create(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Create() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
