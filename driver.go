package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
)

// ExampleDriver is a initialsied with different arguments.
type ExampleDriver struct {
	volumes    map[string]string
	m          *sync.Mutex
	mountPoint string
}

// NewExampleDriver creates a plugin handler from an existing volume
// driver. This could be used, for instance, by the `local` volume driver built-in
// to Docker Engine and it would create a plugin from it that maps plugin API calls
// directly to any volume driver that satisfies the volume.Driver interface from
// Docker Engine.
func NewExampleDriver() ExampleDriver {
	return ExampleDriver{
		volumes:    make(map[string]string),
		m:          &sync.Mutex{},
		mountPoint: "/tmp/mntdir1",
	}
}

// Create request results in API call to volume registry
// that creates requested volume if possible.
func (d ExampleDriver) Create(r volume.Request) volume.Response {
	logrus.Infof("Create volume: %s", r.Name)
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.volumes[r.Name]; ok {
		return volume.Response{}
	}

	volumePath := filepath.Join(d.mountPoint, r.Name)
	_, err := os.Lstat(volumePath)

	if err != nil {
		fmt.Printf("Creating new directory %s", volumePath)
		//cmd := exec.Command("sudo mkdir", "-p", volumePath)
		os.Mkdir(volumePath, os.ModePerm)
	}
	d.volumes[r.Name] = volumePath
	return volume.Response{}
}

// List all volumes and their respective mount points.
func (d ExampleDriver) List(r volume.Request) volume.Response {
	logrus.Info("Volumes list ", r)
	volumes := []*volume.Volume{}
	for name, path := range d.volumes {
		volumes = append(volumes, &volume.Volume{
			Name:       name,
			Mountpoint: path,
		})
	}
	return volume.Response{Volumes: volumes}
}

// Get request the information about specified volume
// and returns the name, mountpoint & status.
func (d ExampleDriver) Get(r volume.Request) volume.Response {
	logrus.Info("Get volume ", r)
	if path, ok := d.volumes[r.Name]; ok {
		return volume.Response{
			Volume: &volume.Volume{
				Name:       r.Name,
				Mountpoint: path,
			},
		}
	}
	return volume.Response{
		Err: fmt.Sprintf("volume named %s not found", r.Name),
	}
}

// Remove is called to delete a volume.
func (d ExampleDriver) Remove(r volume.Request) volume.Response {
	logrus.Info("Remove volume ", r)
	d.m.Lock()
	defer d.m.Unlock()

	if _, ok := d.volumes[r.Name]; ok {
		delete(d.volumes, r.Name)
	}
	return volume.Response{}
}

// Path is called to get the path of a volume mounted on the host.
func (d ExampleDriver) Path(r volume.Request) volume.Response {
	logrus.Info("Get volume path", r)
	if path, ok := d.volumes[r.Name]; ok {
		return volume.Response{
			Mountpoint: path,
		}
	}
	return volume.Response{}
}

// Mount bind the volume to a container specified by the Path.
func (d ExampleDriver) Mount(r volume.MountRequest) volume.Response {
	logrus.Info("Mount volume ", r)
	if path, ok := d.volumes[r.Name]; ok {
		return volume.Response{
			Mountpoint: path,
		}
	}
	return volume.Response{}
}

// Unmount is called to stop the container from using the volume
// and it is probably safe to unmount.
func (d ExampleDriver) Unmount(r volume.UnmountRequest) volume.Response {
	logrus.Info("Unmount ", r)
	return volume.Response{}
}

// Capabilities indicates if a volume has to be created multiple times or only once.
func (d ExampleDriver) Capabilities(r volume.Request) volume.Response {
	return volume.Response{Capabilities: volume.Capability{Scope: "global"}}
}
