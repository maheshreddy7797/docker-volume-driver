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
	// TODO: Add test cases.
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

func TestExampleDriver_List(t *testing.T) {
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.List(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.List() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestExampleDriver_Get(t *testing.T) {
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Get(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestExampleDriver_Remove(t *testing.T) {
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Remove(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Remove() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestExampleDriver_Path(t *testing.T) {
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Path(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Path() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestExampleDriver_Mount(t *testing.T) {
	type fields struct {
		volumes    map[string]string
		m          *sync.Mutex
		mountPoint string
	}
	type args struct {
		r volume.MountRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   volume.Response
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Mount(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Mount() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestExampleDriver_Unmount(t *testing.T) {
	type fields struct {
		volumes    map[string]string
		m          *sync.Mutex
		mountPoint string
	}
	type args struct {
		r volume.UnmountRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   volume.Response
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Unmount(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Unmount() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestExampleDriver_Capabilities(t *testing.T) {
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := ExampleDriver{
			volumes:    tt.fields.volumes,
			m:          tt.fields.m,
			mountPoint: tt.fields.mountPoint,
		}
		if got := d.Capabilities(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ExampleDriver.Capabilities() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
