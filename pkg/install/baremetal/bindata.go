// Code generated by go-bindata.
// sources:
// lokomotive-kubernetes-baremetal.tar.gz
// DO NOT EDIT!

package baremetal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _lokomotiveKubernetesBaremetalTarGz = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xe2\x1f\x1d\xe0\x1f\x8b\x08\x00\x06\x6b\x89\x5c\x00\x03\xec\x3d\x6b\x73\xdb\xb6\x96\xfd\xac\x5f\x81\x2a\x9e\xc6\xee\x88\xd4\xc3\xaf\x5c\x6d\xd5\x5d\xc7\x56\x12\x4f\x6c\xc7\xd7\x76\xda\xdd\xf1\xe6\x6a\x20\x12\x92\x70\x4d\x11\x2c\x41\x4a\xd1\xfa\xe6\xbf\xef\x39\x00\x48\x91\x14\xf5\x70\xec\xe6\x4e\x5b\xb3\xd3\xb1\x24\x00\xe7\x1c\x1c\xe0\xbc\x01\xc6\x13\x77\x62\x2c\x22\x3e\x61\xd6\x5d\xdc\x67\xa1\xcf\x22\x26\xeb\x7d\x1a\x32\x6b\xcc\x22\xea\xd5\x1d\xe1\x47\x94\xfb\x2c\xb4\x3c\xee\xc7\x9f\xeb\x99\x5e\xdf\x6d\xf8\x34\xe0\x39\xdc\xdf\xc7\xbf\xcd\xc3\xfd\x46\xf6\x6f\xf2\x7c\xd7\xdc\xdd\xdb\x6b\xb5\x0e\xf6\xf7\x1a\x07\xdf\x35\x5a\xbb\xf8\x87\xec\x6f\x8a\xe0\x31\x4f\x2c\x23\x1a\x12\xf2\x9d\x74\x46\xf1\xaa\x7e\xeb\xda\xff\xa0\x8f\xf7\x88\xf5\x77\xbc\xcd\xb6\xc0\xc3\xd7\x7f\xaf\xd1\xd8\x7b\x5e\xff\x6f\xf1\x3c\x72\xfd\xb1\x35\x14\x9e\xc7\x42\x7b\x46\xc7\x9e\x1d\x8d\x03\x6f\x01\x07\x2e\xf0\xc1\xde\xde\x92\xf5\x6f\x1e\x34\x76\xf7\xf2\xeb\xbf\xdb\x68\xec\xb6\xbe\x23\x8d\x6f\xc1\x80\xbf\xf8\xfa\x5b\x96\x55\x91\x33\x19\xb1\xb1\xdb\xae\x10\x12\xfb\x3c\x92\xf8\x81\x10\x8b\xf8\x74\xcc\xda\x84\x45\x8e\x0b\x7b\x61\x0c\xcb\x6e\x4b\x16\x4e\xb8\xc3\x54\x3b\x21\xcc\xa7\x7d\x0f\x7a\x44\x61\x9c\xfc\xe4\x86\x22\xe0\xbe\x81\x90\x85\xb2\xd7\xb0\x14\x20\xc7\x03\x86\x03\x24\xd8\x38\x83\xb4\x13\x21\xb8\x8f\x98\x0f\xa8\xc9\xbf\x32\xbf\x12\x72\x7b\xad\x31\x7e\xca\xfd\xda\xf5\x27\x3c\x14\xfe\x18\x46\x74\xaa\xdd\x9b\xe3\x93\xde\xe9\xf9\xd1\xdb\x6e\xef\xe6\xe8\x6d\x67\xb2\x6b\xef\xda\xcd\x56\x75\xcd\x80\x8b\xa3\xf3\x6e\x67\xeb\x1e\x69\xea\x21\x85\x5f\xd6\x0d\x38\x3a\xf9\xa5\x7b\x75\x73\x7a\xdd\xed\x1d\x9f\x9d\x76\x2f\x6e\x7a\x1f\xaf\xce\xae\x3b\xa3\x28\x0a\x64\xbb\x5e\xdf\xba\x77\xc5\x18\xc4\x44\xc3\x6a\xb7\x76\x0f\xff\xb6\x0e\xe0\xe9\xc5\xe9\xcd\xe9\xd1\x59\x06\xf0\x65\xb7\x7b\xb5\x06\xec\xab\xc6\x3a\xb0\x67\xa7\xd7\x37\xdd\x8b\x52\x22\x1b\xb6\xfa\x6f\x23\xea\x0c\x98\x45\x92\xe6\x40\x36\xa6\xe5\xbc\x7b\x73\x75\x7a\x7c\x3d\x87\x93\x07\xd3\xdc\x94\x53\xc7\x67\x1f\x01\xde\x55\xb2\x6c\x1c\xf6\x2a\xa7\x5e\xcf\x6c\xa9\xb5\x2b\x78\x8d\x54\xdc\xf4\xae\xba\xc7\x1f\x2e\xde\x9c\xbe\xed\x1d\xbf\xeb\x1e\xbf\xef\xe0\xe6\x5d\x3b\xf2\xfa\xac\x77\x72\x7a\xd5\xa9\x03\xde\xba\x94\x1e\xfe\x75\xd7\x0d\xba\xb9\x42\x6a\x4f\x7a\xc7\x47\xbd\x37\xa7\x67\xdd\xf9\x60\x87\x85\x91\x54\x20\xea\x28\x4f\xa0\x59\x1d\x6a\x3b\x61\xb4\x0e\xe0\x31\x6c\x93\x75\xa0\x36\x81\xf3\xbe\xfb\x3f\x6b\xc1\xdc\xb1\xd9\x5a\x72\xf4\x0e\x53\x54\x1d\x7d\xbc\x79\xb7\x11\x27\xd5\x7e\xda\x84\x33\x01\xdb\x98\x2f\x0a\xe6\x6a\xe6\x20\xb4\x8d\x41\xad\xe4\x8f\x82\xb4\x01\x77\x34\x51\xcb\x59\x94\xe8\x45\x57\x38\x77\xeb\x15\x6b\xd2\xdb\x83\xde\x72\xcc\xa3\x91\x5b\x18\x31\xa6\xf2\xae\xa4\x3f\x1a\x6b\x8f\x45\x76\x40\xa3\xd1\x72\xa5\x5d\xa6\x7c\x6f\x3f\x82\x74\xcd\xb5\xee\x09\x93\x4e\xc8\x83\x88\x0b\xbf\xf3\x2b\x8d\x9c\x11\x19\x88\x50\x81\x47\x45\xce\x87\xf3\x71\x97\x80\x6a\x3e\x0e\xbf\x75\x3f\x73\x19\x49\xcd\xcd\x8c\xf7\x50\x36\xf8\xd4\x07\x6b\xec\x79\xf3\xf1\xbf\x52\x20\xcd\x7d\x3d\xeb\x8c\x63\x2f\xe2\x56\x0c\xb3\xb6\xc1\x5e\x0f\x59\x94\x9b\xe7\x94\xf2\xc8\x02\x8a\x2c\xd7\x97\xeb\x8d\xd4\x83\xe7\xcb\x23\x35\xdd\x93\x8b\x6b\x00\x17\x85\x9c\xc9\x1c\x81\xb2\x63\x2c\xa8\x15\x32\x29\xbc\x09\x2b\xae\x0e\x21\xaf\x19\x00\x60\x9d\x64\x3d\x8a\xcd\x8b\x66\xee\x66\x16\xb0\x8e\xf0\x99\x1c\x89\x28\xfd\xf1\x8a\xa1\x31\x38\x1a\x80\xb6\x03\xa6\x46\x9d\xcc\x9c\x60\x0f\x7e\x66\xce\x35\xf0\x26\xea\xd4\xfb\xdc\xaf\xcb\x11\xb1\x1c\xf2\x72\x3a\xe2\x1e\x23\xdf\x93\x7a\x2c\x43\xf5\xfb\x30\x64\x01\x79\xf9\x8f\xdb\x7f\xbc\xb8\x6d\xcb\x80\x3a\xac\xfd\xe9\xd3\x4b\xa2\x16\x47\x53\xaf\x4c\x33\xf9\x99\xd4\x5d\x36\xa9\xfb\xb1\xe7\xfd\x07\x6c\x52\x22\x3d\x06\xe3\x9a\xf8\xd9\x67\x2f\x57\xac\xd8\x15\xfb\x2d\xe6\xa1\x5a\xb3\x65\xb3\xcd\x74\x59\xe6\x5c\x14\x37\x70\x1e\xc4\x03\x17\xf0\xbd\x06\x42\x26\x9c\x92\x77\xc0\xd7\x10\xa1\x16\x56\x30\x0c\x1c\x0b\x66\x12\x2d\x2e\xdd\xe2\xda\x64\xa4\xfd\x0d\x70\xb7\x74\x67\x23\xd1\xcc\x9f\x94\x8d\xe9\x54\xaf\xde\x83\x19\xfa\x78\xd1\x3b\xba\x7a\x7b\xdd\xb1\xac\x38\xe6\xae\x35\x00\x48\x96\xa4\x13\x00\x37\xa1\x61\xdd\xa1\xce\x88\x25\x90\xac\x40\xb8\x36\xf6\x22\xff\x9b\x51\x3b\x96\x35\x11\x5e\x3c\x66\x1d\xbd\x6e\xb5\x3b\xee\xbb\x9d\x91\x90\x51\x4d\x8a\x38\x74\x0c\x5d\xd9\x45\xcd\x8f\x1e\x8b\xd8\x07\xa6\xe4\x60\x68\xd9\x5a\x37\x52\x8f\x21\x40\x27\x04\x05\x7d\xcb\xf1\x79\x09\x72\x9c\x05\xb4\xd6\xa1\x75\x15\xe2\x2c\x90\x04\xfb\xf2\xa1\x45\xcc\xd4\xe3\x8e\x58\x85\x5c\x75\xd8\x08\xbf\x06\xb5\x40\x42\x19\x00\x43\x85\x08\x22\x24\xdb\x02\xb1\x2a\x21\x01\x5a\x71\x02\x28\x74\xab\xf0\x67\x81\x24\xc8\x97\x0f\xcd\xce\x5f\x0c\x97\x4d\x5c\x0c\xd7\xce\x18\x06\xe7\xa6\xba\x30\xc4\x60\xe2\xd2\x91\x1c\xb7\xc0\x92\xed\xa5\xda\xeb\xab\xd0\xcd\x21\x64\x37\x57\xe9\xb8\x2c\x4e\xea\x8e\x4b\x50\xa2\x16\x93\xa8\xc6\x92\x3e\x6b\x31\x23\x9c\x04\xf1\xaa\x91\x10\xb5\x30\x27\x86\xb0\x57\x28\xa5\x21\x3b\x7c\x4c\x87\x19\x7f\x26\x55\xaf\x97\xa0\xc7\x95\x26\x1d\xdf\xb9\x3c\x24\x56\x40\xb2\xeb\xb5\x51\xff\x82\xbe\x18\x53\x9f\x0f\x18\xd8\xc8\xaf\x19\x8c\x78\xe1\x93\xed\x7e\xd5\xe0\x11\x73\xee\x02\xc1\xfd\xc8\x82\xd9\x87\xec\xeb\x48\xe0\x3e\x75\x54\x06\xe1\x61\x13\xc9\x48\xf9\xc3\xfa\x2b\x91\x7c\xd0\x10\xa3\x47\xeb\x7a\x63\x04\x5e\x3c\x84\xe5\x5e\x02\x21\x31\x94\x7d\xaa\x4d\x68\x55\x5b\x4c\xf4\x01\xf9\x80\x3b\x34\x62\x16\x8d\xa3\x91\x08\x79\x34\xb3\x5c\x1a\xd1\x97\x0b\x0c\x99\xbb\x36\xe4\x5f\x84\x4e\xef\xc8\xcb\xfb\x20\x04\x26\x93\xad\xd6\x97\x97\xf0\x13\x80\x66\x07\x7b\xc4\x72\xd1\xcc\x16\x97\xa4\xe0\xf3\xe6\x88\xb3\x52\xea\xc2\xbb\x88\x84\x63\x92\x31\x1e\xab\xec\x46\x99\x93\x80\x90\x14\x3b\xc1\x31\x11\xa9\xd1\xb2\xa6\x21\x0d\xc0\x44\x16\xa4\x83\xfa\xc2\x9f\x81\x70\x49\x35\xf7\xce\x80\x7a\x92\x15\xbb\x40\x03\x58\x37\xe4\x10\xc8\x8f\x15\x89\x3b\xe6\x5b\x53\xd6\x1f\x09\x71\x57\xd2\x15\xf8\xf7\x7f\xba\xe7\x58\xb8\xac\xf3\x6b\x69\x47\xc7\xe3\x00\x12\x74\xb3\x99\x61\x29\xaf\x16\xc6\xa8\x70\xb0\x07\xae\x20\x44\x8a\x99\x6f\x3d\x63\xd9\x7b\x3c\xf8\xb2\x6c\x8c\x8a\xb6\xb3\xc3\x74\xf8\x2d\xe3\xc1\x80\x7f\x5e\x18\x05\x6a\x1b\x97\xd9\x82\xed\xb6\x48\x5b\x22\x97\x85\x41\x0c\x1c\x37\x0b\x66\x8d\xae\xbc\x65\x1c\x19\xe0\x42\xa1\x17\x2a\x3d\xf4\x80\x2c\x01\xd1\x58\xc8\x81\x41\xf9\x54\x40\xa1\xfb\x7c\xc3\xad\x70\xb3\x0b\x63\x14\x01\xf3\x8d\x13\xc6\x7e\x1d\x7f\x4a\xbd\x17\xfc\x52\x18\x02\x10\xa7\x22\xbc\xb3\xb4\x00\x75\x16\xcd\xb3\x0f\x4b\x69\x79\x14\x00\xc8\x8e\xfa\x1c\x0a\x8f\xd9\x73\x6a\x6c\x2e\x40\xdb\x21\x67\xbf\x62\xe0\x3c\xd7\xd8\xa9\xaa\x40\xaa\x00\x03\x36\x7b\xaa\x81\x2c\x8c\x78\x16\x78\x91\xea\xa7\xc2\xc8\x90\x51\x17\xd6\xc4\x9b\x01\x0c\x10\x8d\xc6\x42\xf3\x90\x23\xcd\xd6\x14\x22\x2f\x0b\xb3\xa1\xd1\xba\xe9\x75\xda\x17\xe2\x1a\x24\xd1\x8d\xbd\xa2\x9c\x68\x1d\x64\x78\xa8\xb7\xce\x43\xd4\x94\x08\x0a\x5a\x40\xc2\x4f\x0f\xd6\x03\x57\x4c\x2a\x2d\x40\xbd\x29\x9d\xc9\xe2\xcf\xd7\xcc\xe9\x34\x1b\x8f\x0f\xcb\xfa\x42\x44\x48\xc2\xe3\xdc\xf7\xd7\x00\x45\x46\xa0\x96\x08\x25\xef\x53\x5e\x13\xb3\x1d\x48\xe0\x51\x9f\x11\x5c\x1a\x68\x87\x28\x0c\xba\x05\xdc\xd2\x79\x8c\x14\xde\xb1\xf0\x5d\x8e\xd0\x32\x01\xe9\xf7\xca\x6e\x27\x34\xd6\x31\x9d\xd4\x4b\x29\xc6\x30\xe7\x49\xc3\xb3\x5f\x41\x70\xb8\x3f\x3c\x81\xc0\xc7\x89\x44\x38\xeb\xe4\xb0\x97\x29\xe8\x1c\x75\xc9\x07\x4b\x2d\x50\x89\x75\x00\x85\xa1\xad\x5f\x24\x62\x08\xcf\xd7\xce\x0d\xb6\x4d\x08\x0e\x0e\x66\x6a\x71\xd7\xa4\x49\x5f\x14\x9c\x76\xa9\x45\xcb\x87\x34\x6a\x90\x0a\x7b\xdb\x24\x14\x29\x17\x50\x99\xb7\x09\xa6\xfb\x8b\x8b\x9d\xd2\xcc\x7d\x8f\xfb\x2c\x9f\xe9\x7d\xff\xf1\x75\xf7\xac\x7b\x63\xf2\xb8\x1f\xaf\xce\x3a\x3a\x31\xd2\xae\xd7\xef\x5e\x49\x7b\xe8\x84\x28\x5b\xa3\x85\xe0\xad\x38\x52\x65\x80\x9b\x76\x73\xd7\xde\x5b\x9c\x4f\xa2\x54\x9f\x6a\x06\x19\x2a\xf2\xda\x79\x11\x35\xe0\x71\x22\xcf\x76\x41\x3d\x7c\x56\xe2\x62\x4d\x31\x8b\x02\x2a\x23\x93\x08\x5f\x42\xd0\x66\x1c\x1c\x80\xf6\xf1\x05\xb8\x28\x33\x1b\x70\xf4\x10\x47\xcf\xe0\xe8\x34\x0f\x9a\xaf\x0a\xfc\x58\xbf\xb9\x56\xb3\x67\x3f\x65\x0f\x62\xca\x90\xe6\xb6\xc9\x7e\x23\xd1\x1d\xc3\x50\xc4\xc1\xb2\xc6\xcd\xe6\xf5\xe2\xfb\xd4\x1b\xcb\xfe\x4a\x7e\x35\x7e\x0a\x26\x66\x12\xfa\x49\x5e\x38\x08\x91\x10\xf5\x5b\x2c\x37\xee\x1c\x2c\x2a\x61\x9f\x61\x28\xc7\x60\x9c\x7a\x64\xd1\x6b\x05\x89\x27\x96\x4f\xaa\x5b\xdb\x9e\x2c\xb0\x8a\x4a\x80\x99\xb1\x24\xd6\x8f\xf5\x1f\x49\xeb\xe7\x34\x59\xb2\x53\x25\x9f\xc8\x0f\x3f\x90\xf1\x64\x93\x81\xab\xbb\x20\x1c\xf4\xf1\xc2\xc1\x5a\x58\x19\xd2\x19\xa8\x04\x92\xf7\x13\xe3\xbc\x7b\x81\x56\x08\x74\x13\x58\xc9\x3b\x36\x93\xd6\x20\x14\x63\x4b\xa5\xf7\x17\x7a\x99\x80\x4c\xe3\x5b\x12\xe3\x16\xa8\x5a\x80\x91\x0b\xc9\x0c\xa4\x24\x20\x5b\x32\xc4\xa0\xed\x27\x8a\x7f\x49\xec\x39\xd7\x4f\xab\x91\xce\xe1\x64\x23\xd0\xa5\xa3\xb7\xb6\x30\x3f\xf3\xe1\xf2\xe6\xba\xd0\xf0\x5b\x4c\x67\xda\x11\x51\x1e\x73\x32\xf1\xf6\xa4\x61\x37\xf7\xec\xc6\x02\x11\x00\x5c\x11\xbd\xd0\x80\xae\x69\x69\x03\xae\x5d\x27\x05\x0c\x3f\xe8\x1d\x8d\x2e\x33\x72\x4a\x7b\x0b\x86\x69\xd5\xad\xff\xaa\x56\x02\xf8\x32\xd5\x95\x3b\x10\xc3\x42\xe5\x0e\xe9\x34\x08\xa4\x1c\xf5\x12\xaf\x9b\xb9\x3d\x5c\xf7\x6c\x8d\x6e\xeb\x7e\xb1\xc3\x97\xca\xbf\xbb\x36\xf9\xfc\xfc\xfe\xcf\x23\xeb\xff\x18\x0e\xac\xac\xfd\xe3\xb3\xba\xfe\xdf\x38\x3c\x3c\x2c\xd4\xff\x5b\x87\xcd\xc3\xdd\xe7\xfa\xff\xb7\x78\xd6\xd7\xff\x9f\x2b\x54\xcf\x15\xaa\xbf\x52\x85\xea\xb9\xf4\xf4\x5c\x7a\x7a\x2e\x3d\x3d\x97\x9e\x9e\x4b\x4f\x4f\x5b\x7a\x7a\x2e\xf7\x3c\x97\x7b\x9e\xcb\x3d\x7f\xb9\x72\x0f\xfe\xfe\xad\x0a\x35\x7f\x90\xd2\xca\xfe\xd7\x85\x13\xcf\x29\xfb\xbf\x5c\xca\xfe\x77\x4d\xef\xa9\xfc\xde\x23\xf3\x3f\x5c\xef\xdf\xd5\x09\xa0\x35\xf9\x9f\xd6\xde\x7e\xb3\x98\xff\xd9\x6d\x3e\xe7\x7f\xbe\xc9\xb3\x3e\xff\x63\x96\x78\x93\xdb\x1f\x1b\x04\xc7\x26\xf8\x06\x7b\x61\x8c\x8b\x50\x82\x91\x4d\x99\xe0\xa3\x32\x08\x6b\xfa\x2c\x49\x49\x48\x0e\xdb\x70\x69\x41\x35\x9d\xcd\xd7\xa7\x74\x5e\x90\xa3\x89\x80\x80\x38\x96\xdc\x1f\x12\xf0\x82\x30\x49\xef\xbb\x34\x74\xc9\xf5\xf5\x3b\x82\x96\x89\x48\x41\x60\x02\x21\x1d\x08\xf0\xda\xc0\xc3\xf2\x66\xc4\xa1\x3e\x08\xbb\xea\x02\x81\x02\xf7\x0c\xac\x00\x34\x9c\x95\x88\x11\x79\x1d\x47\x04\x02\x04\xf8\x4e\xc6\x74\xa6\x3a\x47\x82\xb8\xac\x1f\x0f\x89\xcb\xe5\x5d\xb2\x1a\x24\x08\x05\x70\x7f\x2c\xed\x84\x24\x64\x58\xd2\x5a\x43\x85\xe0\x92\x29\x87\x8e\x40\xbd\x26\xa9\xd5\x22\x40\xa5\x56\x24\xf5\x39\x71\x40\x16\xae\x9b\xcf\x9c\xc8\xce\xad\x3b\x82\xb0\x25\xaa\xf4\x84\xe5\x4b\x2f\xf7\x34\x1b\x16\xf6\x56\x46\x79\xf3\x9b\x3d\x0a\x74\xfe\x62\xcf\x19\x9e\xb3\xf0\xaf\x23\xb0\xf3\xe3\xce\xf2\x96\x16\x3c\xeb\xac\x61\xd9\x62\xaf\x29\x69\x7e\x7d\x65\x92\x80\x9f\x97\x9d\x73\x1c\x7a\xca\x5b\x89\xc2\x19\x30\x87\x54\xb7\xee\xf9\xd0\x57\xe7\x0f\x7a\xcc\x77\xd5\x81\xbb\x2f\xff\x79\x7f\x6f\x87\x20\x0d\xe0\x17\xd8\x21\x9d\xf6\xe0\x53\x38\xfb\xf2\xe5\x07\x01\xf1\x9f\xa1\xdb\xad\x12\x4b\x90\x64\xa8\xfd\x4f\x29\xfc\x0c\x96\xad\x7b\x21\x7b\x03\x8f\x4e\x44\xf8\x25\xd9\x41\xc5\x8a\x92\x0b\xbd\x4c\x53\x0f\xf7\xcf\x97\x62\x87\x63\x0d\xc6\x19\xc1\xee\x64\xde\x42\xf3\x2f\xba\x19\xbc\x52\x09\x14\x2c\x34\x0b\x9c\x5a\x6a\x21\x7a\xca\x42\xf4\x04\x1b\x7f\xa9\x16\x7a\x6e\xdd\x63\x10\x04\x6c\x41\x82\x87\x0b\x70\xf8\xd2\x49\xc6\x2e\x9b\x60\xc4\x2c\x59\x14\x79\x59\x27\x42\xaf\x23\x58\x68\x12\x32\x2c\x98\x3d\x95\xad\xac\x96\x1a\xcb\xea\x9f\xb5\x18\xf6\x18\xfb\x2f\xe2\x28\x88\x23\x69\x47\x83\xd5\x38\xd6\xd8\xff\x46\xb3\x79\x50\xb0\xff\xfb\x7b\x07\xcf\xf7\x3f\xbf\xc9\xa3\xd7\x90\x54\xe7\x71\xa3\xa5\xec\x4f\x95\xdc\x83\x50\x4c\xa8\x17\x33\xd2\x41\xa1\x00\x25\x19\x43\x70\x97\x1e\x22\x2a\x0e\x00\x11\x79\xae\x18\xff\xf1\x9e\xc7\xc8\xff\x55\xf7\xe8\xe4\xbc\x6b\x8f\xdd\x35\x38\xd6\xf9\xff\x87\xad\x56\xf1\xfe\xff\xc1\x61\xe3\x59\xfe\xbf\xc5\xf3\x02\x9d\xe6\x91\x10\x3e\xf9\x89\x8f\x87\x84\x7a\x60\x88\x3b\xd5\x90\x0f\x47\x51\x95\xc8\xd0\xe9\x54\x93\x8b\xbf\xc6\xdf\xb2\x87\x42\x0c\x3d\x46\x03\x8e\x81\xf0\xb8\x0e\xfe\x2b\xe3\xae\xf0\xeb\x91\x86\x83\xf9\x7d\x61\x07\xfe\xb0\xfa\x73\xa5\x92\xc0\xe6\x92\x50\x02\x4a\x82\x8f\xa9\xa7\xbc\xd0\x41\xc8\x58\xf6\x30\x27\xf8\x26\x51\xc8\xfb\xb1\x72\x01\x2a\x95\x1f\xc9\xb9\xee\x5c\x43\xdf\x1a\x1c\x5d\x95\x43\x5d\x36\x00\xba\x9f\x30\xc7\xa3\x21\xc5\x7d\x0c\xbe\xda\x20\xa4\xd0\x1a\x3b\x51\x1c\x32\x85\x4e\x2b\xaa\x38\xa4\xa6\xfb\x1b\x44\xbf\x8d\x44\xb8\x62\x6c\x7a\xc8\x68\x47\x7d\x0a\x42\x3e\xa1\xce\x0c\x2b\xa5\x01\x38\xc5\xe0\xe3\xc3\x80\xcb\x10\xaf\x38\x38\x40\x3d\x96\x59\x3d\xda\x97\x35\x82\xc9\x60\x07\xdc\x44\x70\x36\x6a\x1a\x86\x27\x62\x57\xce\x67\x9d\xd2\x08\xf4\xc6\x81\x54\xbe\x6b\x66\x0e\x30\x28\x74\x46\x3c\x62\x48\x27\x40\x06\x22\x27\x3a\x3b\x98\xc2\x53\x59\x49\x08\x08\x5c\xf5\xdb\x18\x0f\x76\x7a\xfc\x0e\xe6\x44\xde\x5e\x7c\xac\x9f\xa1\x40\xe6\x38\x81\x61\xc1\x84\xbb\x80\x0f\xc3\x12\xdd\x7e\x87\xe8\xbc\xb9\xef\x8f\x15\x55\xc0\x35\x0e\x84\x8f\x3e\x2e\x70\xfb\xc5\x0b\xf2\x86\x51\x64\x96\x24\x3f\x51\x32\x0a\xd9\x60\xbe\xec\xd3\xe9\xd4\x76\x7c\x67\xa0\x0e\x3e\xa5\x49\x71\xc0\x55\x97\x62\x10\x4d\x51\x51\x20\x77\x21\x94\xa0\xbe\xc3\xea\xd5\x9f\x1f\xbb\x8f\x0c\x12\xe6\x66\x14\x92\xde\x4f\x3f\xd5\xe9\xcf\xb8\x37\x32\xdb\xc0\x24\x8e\xc8\x76\xc2\xdf\x9a\x2a\xb5\xde\xce\x87\x82\x67\xec\xc4\x7d\x0a\x48\xd3\x73\x55\x9f\xb6\x13\x5a\x86\x3c\x1a\xc5\x7d\x85\x7f\xe5\x88\x9d\x1d\x40\x7b\x0d\x5b\x01\x76\x22\xac\xbf\x8e\x0b\xf5\x49\xef\x1a\xb9\x3d\x56\x15\x8a\x39\x54\x64\x19\x2c\xc4\x3f\x61\x61\x75\xf1\xc2\x16\xe1\xb0\xbe\x83\x23\x6f\xc1\xfd\x45\x47\xbb\x94\x04\x93\x9a\x37\x5d\x76\x88\x09\x7f\xf5\x0e\xfc\xe0\x27\x69\x6b\xf5\x46\x06\x7d\xe0\xf9\xe6\xec\x1a\xf0\x5f\xbd\x3e\x3a\x9e\x03\xcc\xa7\x5f\x5d\xe1\xc8\xba\x32\xd0\xf5\x5c\x1a\xbe\x1e\xf6\xa9\x53\xdf\xb1\x74\x14\xef\x02\x14\x83\x0d\x82\x45\x20\x79\xb6\x12\x1e\x2c\xb8\xc3\x82\x48\xd6\x4d\x46\x40\x5a\x73\x52\xeb\x69\xd6\x18\xe1\x70\x30\x13\xc8\xbb\x23\x77\x82\xdb\x03\x04\x3f\xd9\x68\x6a\x1f\xdf\x4a\x9f\x07\x01\x8b\xe4\x1c\x9d\x51\x23\x76\x20\x5d\x1f\x11\x52\x33\xb2\xee\xc0\xe4\xc5\x38\x21\xff\x45\xc1\x2e\xed\x90\x5c\x3b\xa0\xbc\x62\xd4\x9d\x29\x71\x3d\xf5\x87\x80\x11\x04\xe8\x32\x14\x60\xd4\x46\x2c\x86\xcf\x6f\x21\xfe\xa5\x3e\xd5\x92\x26\xe0\xc7\x90\xe8\xf2\x1b\x48\xe2\xad\x96\xb8\x55\x44\x61\x7b\x1d\xb3\xf5\x13\xce\xa6\x30\x45\x14\xa2\x13\xe0\x4d\xa5\x72\x09\x9b\x1a\x74\x95\x04\x0d\x83\x42\x78\x2b\x06\x20\x31\x1c\xa0\x22\xeb\x96\x82\xd4\xba\x07\x07\xcc\x6d\x2f\xb9\x8d\x62\xd8\x85\x30\x76\x39\x25\x8e\x97\x35\xd6\x3b\x76\xe5\x81\x3e\xd8\x63\xec\x3f\xc4\x49\x6b\x7d\x7f\x7c\xd6\x9d\xff\xda\x3b\x28\xe6\xff\xf6\x20\x00\x78\xb6\xff\xdf\xe2\x79\x41\xae\x55\xed\x19\x2c\x42\x30\xd3\x9a\x05\x94\x8a\x39\xf1\xab\xf6\x64\xa6\xa2\x14\x89\xe4\xd6\x85\x07\x86\xc4\x26\x47\x78\xed\x8f\xa2\x1e\x2e\x1e\x96\xc1\x03\x1c\x58\x36\x27\x55\x3c\x9f\xd3\x4b\xbe\x56\x49\x15\xf1\x58\x73\x28\xc9\xbd\x43\x1d\x6f\x38\xaa\x86\xae\xe2\x0d\x8f\xf9\xc3\x68\xb4\x3d\xa1\xea\x5d\x31\xa6\xb7\xca\xd5\xcb\x1d\x88\x35\x2a\xea\x0c\x38\xe8\x40\x08\x60\x88\xcb\x02\xe6\xbb\xb2\x27\xfc\x1a\x09\x19\x6c\x67\xa6\x4e\xd2\x22\x38\xcf\x35\x47\x68\x71\x2a\xd3\xe4\x9c\xd4\x98\x82\xe1\xf5\x81\xee\xbe\x3a\xf5\xa4\x80\x8d\x31\xf3\xde\x17\x9f\xf5\xa1\x75\x98\x3b\x30\x65\x1a\xf2\x28\x62\x00\xd5\xa1\x3a\xdb\x47\x01\x17\x75\xb1\x56\x86\xc9\xb2\x39\x5e\xa0\xf9\x56\x25\x11\xaa\x09\x98\x9e\x02\x63\x9b\xe4\x4b\xb5\x56\xda\x3a\x9f\xd9\x92\x0e\xfa\x7c\xa5\x6a\xfc\x54\x51\x0c\x52\x99\x3a\x34\xf5\xf7\x6a\x00\xe8\x03\x86\x7f\x81\x67\x20\x8e\xba\x8e\xac\x8e\x16\xeb\xdf\xb6\xee\x61\x59\xf0\xe4\x4f\x91\x91\xba\xf6\x01\x7a\x50\xb1\x1c\xc8\x74\xd9\xe7\x1d\xf3\x0a\x15\x74\x13\xcc\x78\x34\x49\xfa\xc7\x88\x8f\x19\xf2\x1a\x7e\x3c\x68\x8c\xf1\xb7\x2f\x48\x90\x72\x38\x30\x37\x04\x43\xaa\x98\x5f\xab\x1a\xc2\x4c\x0e\x8d\xa4\x84\xac\x08\x20\xcd\xee\x31\xe8\xc1\x7d\x01\xaf\x4b\xa9\x71\x35\xf2\xdd\x87\xf3\x6e\xa6\xae\xf9\x04\xa8\xd5\xeb\x63\x1c\xda\x43\x57\x63\x15\x52\xf3\xc6\x22\x53\x94\x4e\x0a\xf5\x4f\x82\x5c\x01\x7d\x08\x01\xbf\x03\x76\x9d\xd8\xda\x08\xb9\x79\xe3\xc9\x53\x20\xd7\x37\xb3\x36\x9b\x7a\xfe\xa5\x36\x4f\x88\x7d\x93\xa9\xe7\x5f\x85\xf3\x14\xc8\xf1\xe5\x31\x9b\x4d\x3c\xfb\xc2\x9a\x27\xc3\xbc\xc9\xa4\xb3\xef\xb7\x59\x44\x9c\x51\xae\x09\x7e\x9d\x19\x4f\xf5\x1f\xa8\x30\x19\xbb\x82\xe4\x0f\x0d\x25\xef\x4c\xd2\x2f\x4e\xaa\xe5\x7b\x4e\x48\x66\xa7\xfd\x58\x18\x50\xe8\xec\x94\x01\xcc\x8b\x68\x49\x87\xc2\x7b\x96\xca\xf1\xcf\xf7\xda\x72\x08\xeb\x87\x03\xeb\x56\x0c\x47\xc6\x3e\xc9\x84\xb2\xaf\x47\x2a\xa7\x27\xd9\x41\xcb\x46\xaf\x1b\xba\x64\x22\xe9\xfe\x28\x4c\x63\x24\xa6\x3e\xb1\xae\x14\x80\xb6\x72\x23\xf2\x2f\xcb\x2a\x76\x87\x5d\x8a\xdd\xf7\x1b\x8d\x95\x1d\x81\xa4\xa2\xfe\x5f\x71\xfc\xca\x0c\xfe\xa4\xf6\x2e\xec\xde\xbc\x73\x93\x77\x64\xb0\x58\xa2\xad\xeb\x63\x1c\x19\x0d\x61\x13\x27\x46\xf7\x7c\x76\x60\x1e\xe5\xc0\x18\x26\xfe\xd9\x9c\x97\x07\xa8\xd5\x27\x94\x87\xf4\x9a\x99\x71\xf6\x41\x28\x3e\x5c\x74\x33\x1e\xbe\xda\x72\x7a\xf3\xa5\x7d\xa1\x53\xc0\x42\x4c\xf7\x00\x34\xa0\xde\x42\xd6\x42\xb0\xeb\x0d\xd4\x21\x36\x36\xcf\x5a\xa5\xb7\xee\x02\xd8\x77\xf6\x0a\x41\xca\xdf\x3e\xd5\x0c\x28\x17\x8d\x68\xc4\x65\x4e\x3e\xb0\x46\x6e\x08\x54\x72\xa0\x02\xe8\x0c\x6b\x70\x9e\xb6\x02\x77\x93\x16\xbc\xf1\xf8\x18\xc4\xe2\x20\x3b\x40\x3e\x89\xa8\xbc\x23\x34\xc2\xab\xe3\x30\x91\x1a\x16\xee\x41\xc4\xa6\x4a\xee\xb0\x7a\x0b\xf3\x4d\x27\xa2\x00\xe9\xeb\x0a\xbe\x20\xe6\x68\xbf\x96\xb3\x30\xf6\x7d\x35\xcd\x72\x99\xca\x4d\xd9\x5e\x16\x02\xd5\x96\x76\x2e\xa8\x99\xdf\x31\x1c\x68\x6c\x2a\x47\xcd\xfd\x0d\xe4\xc8\x2c\xf9\x5c\x8e\x10\xad\xda\x6e\x3d\x70\x0f\x56\x89\x8c\xde\x93\x4f\x26\x2e\x66\x8b\xe7\x2e\xac\x16\x4c\xcd\xbc\xa0\x9c\xdf\xf2\x05\x21\xfa\x77\x07\xec\x4f\xfc\x3c\x26\xff\x03\x8b\xc9\x31\x7b\xb8\xae\x02\xbc\xe6\xfd\xbf\x4d\xf8\xbf\x58\xff\x3d\x6c\x3c\xd7\x7f\xbf\xc9\x93\xac\x21\x48\xb9\x39\x03\x8d\x1e\x8a\x16\xab\x44\x99\x24\x0a\x25\x0a\x41\xc5\x55\x95\x8a\x4b\x6f\x38\x61\xc3\x47\x9f\xff\x16\xb3\x54\xf1\x2b\x00\xda\xdc\xcc\xf7\x51\xa5\x32\xc7\x94\xba\x08\x98\xd4\x4c\x0f\xc4\x3c\x0c\xe7\x79\xe2\xec\xbc\xbb\xb9\xb9\x24\xe9\xb9\x60\x92\x40\x23\xdb\xcc\x1e\xda\xc4\xbc\x6d\x36\xc1\x68\xb3\xcf\x14\x4f\x87\x61\xba\xbd\xfd\xaa\xf1\xaa\xb1\xa3\x08\x9d\x93\x36\x3f\x0a\xf3\x30\x72\x8e\xf5\x20\xe5\x9b\x51\x7c\xdb\x88\x96\x19\x53\x82\x71\x19\xd6\x94\x54\x71\x6a\x5b\x27\xf9\x2d\x5d\xd6\xaa\x11\xf3\xb5\x0f\x3c\x4a\xbf\x50\x2f\x18\xc1\xb7\x81\x07\x54\xd3\x30\xed\x9a\x7c\xd7\x7d\x93\x6f\xaa\x73\xc9\x34\xcc\x91\x9d\x87\x4d\xe3\x17\x3d\x68\xfd\x34\xc0\x2c\x5e\xfe\x77\x57\x39\x08\xc9\xc9\xa3\xdf\x73\x66\x2f\x52\x87\xb7\x92\x31\xe5\x2f\x25\x71\xc2\xd8\x85\xf9\xaa\xf9\x69\x05\x5e\xd5\xb6\x1d\x8f\xda\x81\xd1\x00\x5f\x05\x36\x85\x8c\x03\x75\x04\xce\xc3\x6b\x9e\x44\x0c\x00\x1a\xf8\xc7\x60\xea\xa7\x0c\x6f\x0e\xa2\x5b\x61\x67\xd9\x57\x4c\x38\x96\x32\x11\x81\x95\xb0\xf0\x43\x08\x6c\x02\x37\x08\x9b\x11\x55\xc6\x9d\x52\xb0\xcc\xc6\xbc\xc5\x73\xf1\xcd\x4f\xc5\x85\xcb\x60\x86\x19\x3f\x15\x62\xee\x62\x5d\x71\x30\xc3\x00\xe0\xfc\xe8\x18\x0b\x8a\x58\x10\x99\x13\xb3\xdf\x6a\xef\xef\xb5\x1b\x8d\x36\x6d\xb6\xff\xe6\xb4\x29\x5b\x45\x98\xf1\x15\x9e\x88\xb6\x37\x7f\x3f\xb9\xc8\x33\x25\x2b\xa4\x0b\x74\x64\xc3\xa8\xc7\x50\xa0\xe1\x2c\x2e\x49\xab\x46\xf0\xcf\xee\x32\xc4\x8f\x5d\x15\x83\x77\xf3\x15\xe9\xb7\xda\xad\x41\xfb\xd5\x41\x8d\xa4\x3f\x39\xbb\xed\x83\x66\xfb\xf0\x70\x19\x8d\x4f\xb0\x40\x86\xcc\x85\xc5\x69\x65\x17\xc7\xb0\x6a\xe5\x7a\x39\x9e\xd3\x4b\x8a\x7c\xa5\x04\x81\x28\x96\xaa\xf7\x80\xe0\x7b\x49\x12\xb9\x37\x0b\x05\x6a\x27\x95\xe1\xa2\x76\x3a\xd6\x2e\x7f\x8a\x4c\x01\x1d\xd0\xd8\x4b\xc2\xb8\x7b\x13\x03\xe5\x0f\x03\x64\x68\xbd\x7b\x25\x7b\x99\x3b\x01\x9a\xdc\xa2\xa2\x9f\xef\x5c\xbc\x6c\x8d\xfd\x30\x24\x70\x46\xc4\xdc\xac\x56\x44\xd2\x9c\xf4\xa1\x82\xf4\xc1\x95\x57\xef\x80\x02\x0d\x5b\x28\xab\x80\x1a\x53\x27\x75\x81\x7d\xe3\xd8\x57\x57\xce\x74\x75\x57\x85\x3b\x05\x83\x96\xbc\x3f\x3f\xc3\xf3\x9d\xea\x72\x0d\x9f\x5b\x8b\xc5\x73\x8d\x0f\xb3\x0e\xea\x74\x73\xdc\xf7\xb8\x43\x30\x3b\x84\x46\x42\x45\x09\x2f\x51\xbb\xbf\x2c\x20\x4b\xbd\xfc\x52\x36\xe2\xad\x7b\xcd\x28\x37\x79\x21\x16\xb0\x11\xb6\x20\x19\x32\x58\x52\x8a\x81\xa4\xf1\xd9\xe5\x48\x05\x63\x10\x7b\x06\x1e\xc5\x12\xf2\xb6\x71\x49\xa1\x49\xc7\x43\x1b\x33\x60\x5e\xa5\x2e\x5f\xdb\x91\xe0\x10\xad\xc0\xde\x9a\x77\x4c\x4e\x54\x84\x64\xdb\xd4\xe5\xb1\x8c\xaf\xcb\xfa\xab\xf0\x16\x37\x5f\x55\x0f\x29\x27\xa8\x37\x8e\xe2\x72\x8a\x2e\x4e\x61\xfb\xc0\x7a\x0f\xf0\xd0\xc6\xf9\xcd\x47\xb2\x8d\x87\xc9\xb9\xde\x65\xe6\x82\x30\xfa\x3d\x0f\x22\xa5\xb9\x87\xaf\xee\x2f\x25\x84\x07\xb8\x43\x40\xaa\x23\x1d\x56\xf6\xb0\x5e\x2e\xdc\x52\xe2\xce\x55\x93\x5a\xc5\x74\x88\x0a\xc0\x55\xc0\x79\x7a\x39\xd9\x4b\xd4\xda\x93\x50\x3d\xe0\xa1\xc4\x57\x28\xc4\xbe\x5b\x20\x3e\x10\x6e\xcf\xe1\x6e\xf9\x4e\x3b\x3e\x3d\xb9\xd2\xc4\x84\xd4\x1f\x2a\xcf\x05\x36\x16\x1f\xfa\xd9\x63\x24\x00\x41\x3e\x88\x83\x0d\xbb\x85\xff\x7c\x41\xbd\x79\x50\x94\x31\x73\xe9\x70\x19\x3d\x3f\xfd\xd4\xfd\x70\x52\xd9\x88\xa8\xe4\x70\x85\x5d\xb9\x01\xa6\x36\x15\x4f\xb5\xa6\x00\x59\x00\xb6\x62\x4a\xd9\x4d\xdf\x72\xd1\xc3\x93\x34\x2a\xcb\x5c\x53\x8b\xd0\x6c\x80\x84\x2d\x1b\x80\x02\x8b\x2f\xa2\xa8\x20\x31\x95\x7c\xda\x60\x61\xce\x7a\xbe\xbb\xd9\xf9\xbe\x48\xcf\x4a\xe4\x75\x7d\xc9\x2d\xca\xd2\x45\xf9\x7b\xcc\xf0\x1d\x15\x8a\x16\x63\xad\x12\x95\x07\x7e\x9c\x1a\x97\x12\x4e\x7d\x39\x55\xb6\xa9\x3f\x4b\xe9\x26\x27\x86\x36\x50\x90\x89\x46\xf4\x04\x1e\xcf\xd2\x5a\x72\x20\x84\x6d\xc8\xb7\xe5\xc4\xb1\x73\x7d\x76\xc8\x83\xe4\x36\x3b\xb4\xb0\xd8\xae\x98\xfa\x9e\xa0\x6e\x0f\x94\x44\x24\x1c\x91\x0d\x1a\x96\xb1\x52\x1d\xe4\x28\x51\xae\x97\x06\x04\xe1\xe8\x58\x1b\xa5\x87\x97\x37\xf0\xfe\x87\xc1\xa3\xd3\x5b\xf3\xf3\x5c\xf8\x5a\xbf\xd0\x4d\xb9\xa1\xc4\x4b\xc1\xaf\xa5\x36\x49\xdf\xb6\xd1\x40\xf1\xd0\x17\xf7\x98\x39\x3b\xe4\x84\xb3\x00\xfa\x1b\xdf\xd8\x26\x1f\x7d\xc0\x06\x30\xc1\x3b\xc3\x0b\x8e\xf8\xef\x6b\x68\xc7\x1e\x78\x8c\x77\x7c\xec\xa2\x65\xcf\xf5\x5a\x6b\x49\xf2\xb2\x8c\x17\x8a\x4b\x98\xf0\xeb\x88\xa9\x93\x38\x45\xdb\x6e\xd8\x81\x93\xc0\xb4\x4c\x2e\xee\x30\x6e\x82\x09\x07\x93\x4c\x8f\xa2\xce\x26\x17\x22\xc2\xa4\x20\xd5\x8a\x49\x9d\x83\x22\x63\x58\x50\x32\xa2\x13\x96\xf2\x95\x69\xce\xce\xa3\x26\x54\xba\x62\x0e\x54\xc3\x2c\x32\x20\x7b\xaf\xe2\x61\xd3\x57\x2f\x2e\x91\x2e\x2d\xe1\xc0\x09\x5e\xf2\x81\x66\x34\x44\x40\x82\x5e\x46\x24\x2e\x73\xf1\x47\x5d\x63\x49\x78\x92\xfc\x5e\x64\x99\x16\x84\x04\x53\x99\x3b\x5f\xb8\xb3\xf1\xb0\x29\x94\x91\xde\xbd\xbc\xea\x1e\x1f\xdd\x74\x4f\xda\xe4\x3a\x60\x0e\xb8\xb7\xb0\x50\xe4\x43\xf7\x9c\xa8\x17\x1a\x80\xd3\x8b\x53\xc2\x3d\x4d\xa5\x3e\xd0\x89\xf2\x9f\x99\x9c\xce\x01\x6a\xca\xe9\x98\xd7\xc8\x64\x8c\x07\x0c\x7b\x21\x9d\xd6\xc8\x67\xe6\xab\x23\x74\x1e\xc3\xa5\xeb\x83\x31\xbe\x4b\xc7\x27\xc4\x99\x17\x27\xe4\xa6\xaa\xe5\xa5\x47\xc3\xa1\x2c\xd5\x45\x47\xae\x7e\x33\x28\x28\x8f\x44\xb4\xc2\x61\x8c\x49\x52\x25\x4f\xc6\xfe\x63\x8e\x38\xd9\x7e\x76\x89\xfe\x98\xfb\xd5\x59\x36\xdd\x7e\xca\xd3\xa2\x4f\xdb\xf5\x42\x86\x22\x97\xba\x21\x9b\xfa\x5f\x5d\x35\x1a\x18\x88\xcc\xc4\x08\x1d\x68\x9e\x45\xdc\xc1\x9c\xb8\x01\xa8\x18\x6c\x8e\x42\x82\xd9\x3d\x9e\xfb\x29\xe5\x12\xf8\xa7\xcb\x67\x3e\x3f\x0f\x7b\x1e\x93\xff\x3d\x3b\x3d\xee\x5e\x5c\x77\xd7\xe2\x58\x77\xfe\xbf\x59\xfc\xf7\xdf\x5a\x7b\x7b\xcf\xf7\x7f\xbf\xcd\x83\xae\xe5\xf9\xe9\x0d\x18\x0d\x87\xf9\xa0\x91\xb7\xe1\xcb\x4e\xa5\x72\x2c\x82\x99\x3a\xbc\x0d\xd1\xd6\x0e\x69\xc1\x7a\xa5\xf7\x04\x8e\x54\x04\x29\xcb\xba\x9c\x50\x2f\x82\x1e\xef\xe2\x3e\xe8\xa9\x4a\xe5\x92\x85\x63\x2e\xb5\x3d\x95\x04\xa3\x3b\x70\xe1\x86\xa1\xba\x5d\x5b\xd3\x57\x00\x30\x17\x34\xc2\xab\xb5\x35\xe5\x03\xfb\x33\xac\x2e\x4a\x18\x20\xfa\xb8\xeb\x74\xc5\x1a\x4b\x60\x15\xe8\xa9\xe2\xe1\xe4\xd4\xb9\xf2\x01\xc0\x30\x0b\x87\xab\x68\xd1\x15\x8e\x52\xda\xda\x86\x68\x13\xb9\x8d\xb6\xa1\x7a\x6d\x46\x54\x77\x6a\xfa\x32\x2d\x78\xad\xe0\x06\x60\x5b\xd2\xa4\x3c\x22\x2c\x6c\x81\xaf\x04\xda\x57\x85\x3e\x35\xb0\x4a\xe0\xfc\xb9\x48\x43\xd2\xec\xf1\x31\x37\x18\x70\xb8\x9a\xbe\xac\x68\xa3\x56\x53\x74\xd6\xf0\x36\x29\x58\x3e\xf8\xcb\xd4\xb4\x54\xc0\x2c\x47\xb5\xcc\x6d\x80\x1a\x38\x5d\x18\x45\x23\xc3\xd5\xf9\xe3\x3a\x28\x73\xc9\x3c\xaf\x02\x10\xd0\x2d\x56\x73\x9d\x53\xa7\xcf\x28\xeb\xca\x2b\xe0\x37\x2c\x92\xda\x3f\x00\xe7\x27\x37\x13\x2e\x2b\xff\xdf\xde\xb1\xf6\xb6\x6d\x03\xbf\xfb\x57\x10\x6e\x3b\x24\x9d\xad\xc4\x4e\x9c\xa6\x01\xdc\xc2\xb1\x95\x5a\x98\x63\x05\xb2\xd3\x2c\xd8\x8a\x40\xb1\x9d\xc5\xad\x13\x1b\x51\xb2\x35\x1d\xbc\xdf\xbe\x3b\x3e\x24\x52\xa2\x64\xf9\x51\x27\xdd\x74\x28\x50\x58\x21\xef\xc8\x23\x79\x47\x1e\xef\x8e\x57\x0f\x77\xb7\x40\x72\xd0\x67\x7b\x47\x60\x19\xa5\xf8\x99\x9e\xcf\xc6\xb4\xf8\xd5\x78\x34\x1a\xff\x85\x5d\xeb\x89\xcc\xd8\xde\x41\x8e\x1e\x34\xdc\x4b\xcc\x90\xdb\xf3\x47\x17\x33\x14\xf4\x06\xdc\x3d\x19\x06\x60\x12\x8c\x2a\xff\x93\x77\xed\xb2\xad\x3a\x63\x18\x6e\x22\x6f\x73\x2e\xb5\x68\xb0\xee\xdc\x21\x79\x34\x82\xe0\x7b\x6d\x34\x1a\x19\xe9\x85\xbb\x69\x00\xfd\xa6\x49\x3a\xf6\x51\xf7\xac\xe6\x98\xc4\xea\x90\x13\xc7\xfe\x68\x35\xcc\x06\x68\xe9\x0e\xfc\xce\x17\xc8\x99\xd5\x6d\xda\xa7\x5d\x02\x25\x9c\x5a\xbb\x7b\x4e\xec\x23\x52\x6b\x9f\x93\x5f\xac\x76\xa3\x40\xcc\x5f\x61\x0b\xd2\xe9\x10\xdb\xc9\x59\xc7\x27\x2d\xcb\x84\x6f\x56\xbb\xde\x3a\x6d\x58\xed\x0f\xe4\x10\xea\xb5\x6d\x98\xeb\x16\x4c\x72\x40\xda\xb5\x09\x12\xe4\xa8\x2c\xb3\x83\xc8\x8e\x4d\xa7\xde\x84\x9f\xb5\x43\xab\x65\x75\xcf\x0b\xb9\x23\xab\xdb\x46\x9c\x47\xb6\x43\x6a\xe4\xa4\xe6\x74\xad\xfa\x69\xab\xe6\x90\x93\x53\xe7\xc4\xee\x98\x40\xbe\x01\x68\xdb\x56\xfb\xc8\x01\x2a\xe6\xb1\xd9\xee\x1a\x40\x15\xbe\x11\xf3\x23\xfc\x20\x9d\x66\xad\xd5\x42\x52\x39\x7c\x04\xcc\x76\xb0\x7d\xa4\x6e\x9f\x9c\x3b\xd6\x87\x66\x97\x34\xed\x56\xc3\x84\x8f\x87\x26\xb4\xac\x76\xd8\x32\x19\x29\xe8\x54\xbd\x55\xb3\x8e\x0b\xa4\x51\xc3\x44\x19\xb4\x96\x0d\x58\x9c\x1c\x16\x63\xad\x23\x67\x4d\x13\x3f\x21\xbd\x1a\xfc\xab\x77\x2d\xbb\x8d\xdd\xa8\xdb\xed\xae\x03\x3f\x0b\xd0\x4b\xa7\xeb\x57\x3d\xb3\x3a\x66\x81\xd4\x1c\xab\x83\x0c\x39\x72\xec\xe3\x42\x0e\xd9\x09\x35\x6c\x8a\x04\xea\xb5\x4d\x86\x05\x59\x4d\x94\x11\x81\x22\xf8\xfb\xb4\x63\xfa\x08\x49\xc3\xac\xb5\x00\x17\x0c\x4f\x5b\x19\xbe\xb9\x3d\xd1\x33\x78\x0a\x58\x46\xff\x8b\xa3\xc8\xac\x20\x80\x64\xfd\x5f\x2e\x81\xbe\x0f\xdf\xff\x56\xf6\x32\xfd\xbf\x16\xa0\x66\x0d\x8f\xfb\xdc\x28\x37\x78\xa4\xf8\x8e\x7f\x20\x2c\xc9\x80\x1f\x8b\xc7\x6f\x48\x69\x15\xf5\x2a\x0f\xeb\xf0\x2f\xf1\x95\xd8\x1f\x54\x77\x14\x6f\x32\x1a\xde\x6f\xe4\x8b\x20\xe1\xd1\x45\x24\xb8\x86\xdd\xe4\x2e\x29\xe8\xec\xc2\xef\x58\x53\x57\x2c\x6d\xd2\xa0\xe4\xdc\xd6\x56\xe4\x74\x6c\xa9\xa7\x69\x8c\x05\x04\x45\x7a\x37\xa0\xb1\x3b\x06\x76\xbb\x38\xf6\x0c\x98\xe6\x9b\x92\xe3\x92\x7f\x6f\xcd\x6b\xe5\xa5\x73\x34\x5b\x1d\x45\xc5\x1a\x92\x32\x9a\x81\xfc\x4c\x12\xfc\x04\x09\xbb\x63\x60\x68\x68\x98\x1d\x74\xf7\x95\x57\x8c\xa1\x5c\x7c\xe5\x71\x56\xc8\xf7\xf9\x05\x22\xf8\x85\xa1\x5b\x6e\xc4\xfd\x87\x52\x63\xd5\x14\xfa\xaa\x87\x1d\x1f\x07\x7e\x4a\xf6\xfd\x79\x22\x56\xb0\x29\x7d\xb9\x96\xce\x2c\x43\x24\x96\x30\x34\xbc\xdd\x72\x6f\xfa\x7b\xbb\x45\xcc\xb8\xc5\x30\x49\x99\x26\x78\x30\x1c\xe2\xec\x3f\x30\x53\xf4\xe4\xeb\xc0\xf8\xf3\x06\x7b\xfb\x8d\xb6\x83\x59\xbf\x02\x27\xab\xb5\xb7\xe6\x82\x9a\x1b\x8c\xde\x64\x38\x36\xfe\xf8\x16\x78\x65\xa1\xa9\x21\x68\x16\x6b\x66\x35\x2d\x06\xc2\xfc\xad\xc6\x2c\xcd\xd4\xf0\x0f\xe3\xe1\x6e\x54\x65\x0d\xd2\xfb\x4d\x4c\xb7\x44\xee\x8c\xf7\x98\x87\xac\xfa\xf2\xe5\xdf\xf8\xff\xf4\xa7\x1b\xb7\x87\x3f\xe0\xbf\x83\xeb\xc1\xd7\xeb\xc7\xc9\x34\x44\x80\x5a\xd4\xe9\xb3\x20\xd5\xc7\x81\x17\xfc\xf1\xd6\x1b\x8f\x06\xd5\xfb\xfb\xc7\x6d\xcd\xb7\x8e\xff\x91\x35\x4a\x32\xaf\x4c\x15\xc7\x34\xc5\xbe\xc4\xfd\xf1\x94\xa5\x8b\xf1\xb3\x06\x3e\xd8\x02\x02\x63\x70\x81\x0b\xca\x88\x9b\xd4\xac\xba\x67\xbc\x86\x81\xbb\xa5\x37\x95\x11\xdf\x4f\x58\xe7\x88\x90\xe4\x15\x8c\xf1\x4b\x54\xe0\x5c\xe1\x52\x45\x53\x0e\x27\xce\x97\x2b\xb4\x60\x03\xfe\xc7\x94\x34\x06\x73\x21\x9d\x6a\xb3\x66\xe5\x79\x75\xc0\xea\x71\x8f\x36\x3f\xb7\x0b\x09\x80\xb5\x8f\xce\x64\x9e\xf7\x25\x2f\xca\x0a\xe1\xa8\x2b\x2b\x66\xbd\x5f\x58\x58\x39\xd5\xc2\xa1\x39\xcf\x4a\x47\xf2\xd6\x90\x88\x28\xf2\xe7\x1f\x17\x3c\xfa\x69\x2a\xfc\x0a\x65\xab\x69\x84\xbc\x92\xaa\x26\xef\xbb\xe2\x86\x2c\x95\x41\x79\x6d\xea\x19\xe6\x79\x18\xb9\xeb\x94\xc8\x68\x13\xbc\xd0\x6a\x2f\x98\x7b\x28\xb3\x6d\x87\x65\xac\xaf\x2d\xdc\x7e\xdf\x23\xc5\x4b\xc2\x73\xda\xd0\x9a\x72\x7e\x1b\x61\x20\x9d\xce\xa5\x7d\xf4\x36\xec\x4d\x44\x80\x66\xec\x03\x52\x8b\x37\x5f\xcf\x30\x5d\x73\xf9\x65\xcc\x50\x66\xda\x5e\xaf\x53\xa7\x25\x36\x60\xad\xaa\x6d\x4b\x61\xdc\x4a\xf4\xd3\x62\x28\x33\x25\xb3\x72\x25\x93\x3c\xcb\x96\xd2\x35\xa9\x50\x67\x2a\x27\x53\x39\xb2\xca\x11\x6a\xe0\x01\xfd\xbf\x02\xb5\x82\xb6\x3a\xfd\x25\x27\x9d\x66\x5a\xad\x03\xb5\x93\x97\x31\x97\x41\x9a\x31\x15\xca\xea\x88\x9f\xdf\x98\xaa\x1a\x26\x1e\x94\xc4\xe9\x8f\xb2\x22\xc5\x71\x49\x94\x5f\x9f\x4a\x09\x51\x7c\xf6\xc7\xa3\x08\x47\x13\x8f\x25\xbc\xf4\xf7\x3e\x25\xad\xa2\x51\x73\xe9\xb1\xd4\x28\x88\x3f\xa7\xbe\xa3\x26\x13\x14\x9e\xa1\x2a\x5b\x4a\x87\x45\x16\xfb\xfa\xf7\xa5\x82\xb5\xa9\x36\xa6\xca\x34\x5c\xff\xb6\x54\x4b\xfe\x49\x36\xa5\xbc\x25\xab\x11\x08\x8b\x22\xcd\x16\xf4\x73\xdd\x9b\xc2\xc2\x94\xdc\x08\x25\x8f\x65\x61\xb0\x9f\x6d\xd3\xe4\x29\x5e\xe4\xc5\xa5\x6c\xf2\x92\x93\xb3\x88\x85\x24\x57\x09\x99\x2e\x79\xe8\xe3\x8c\xd5\xa3\x5f\x36\xe1\x35\x42\x08\xe6\xcf\x15\xb3\x40\xc3\xe3\xde\x3d\xe7\xbf\x84\xac\x28\xca\xa7\xdb\xe9\xfb\x28\x14\x06\x05\x48\x16\xe6\x94\x1c\x6a\x9c\x4a\xde\x73\xca\x29\xa6\x02\x86\x4d\x0e\x60\xc6\x5e\xb0\xd7\x48\xaa\x84\x3a\xfd\xb0\xc8\xf3\x63\x94\xd6\xe8\x8a\xc5\xfc\xaf\xc9\x68\x3c\xfe\xf2\x30\x31\x48\x9d\x65\xa9\xc6\xbf\xb0\x4f\x1b\x37\xee\xa4\x80\xfe\xde\x9b\xc4\x1b\xde\xc2\x94\x19\xde\x33\x1b\x09\x0a\x2e\xe6\xc1\x89\xcb\x89\x86\xd4\x00\x66\xe1\x7e\x8f\x12\x20\x38\x0d\x8c\x7a\x17\xf0\xf7\xdf\xd2\x0f\xeb\xa7\x69\xfe\x53\xb2\x49\x4f\x65\xc3\x02\xe9\x83\xd2\x9d\x9b\x82\xca\x89\x47\x27\x29\x6c\x40\x3a\x3c\xa5\x0b\xb9\xd5\x04\xb1\xd3\x9c\x1d\x61\x64\xc9\xe8\x34\x4c\x94\x91\xa1\x1c\x1e\xba\xa3\x0b\x11\xa4\xc7\x91\x7d\x06\x09\xba\x91\x2f\xc0\x22\x64\xeb\x13\xdd\xd8\x70\x8d\x56\x45\xbe\xb1\x57\xde\x41\x79\x67\x7f\x5b\xac\x52\xad\x0a\x8b\x76\x6a\x53\xd0\xd6\xbf\xc7\x42\xb4\x31\xf5\x31\x6f\xb7\x84\x10\xc9\xbe\xc5\xd2\xf9\x4a\xff\x82\x4b\xfc\xc1\x8c\xcc\x3a\x9b\x05\x07\x22\x49\x96\xb2\x00\x8e\x94\x72\x94\x67\x97\x48\x21\x19\xe2\x36\x22\x31\xf2\x93\x47\x83\xa7\x90\x9d\x32\xe2\xa5\xe4\x26\x27\xb9\xb8\xcc\x0c\x23\x58\x88\x2b\x73\xc9\x4a\x4e\xf1\x47\x97\x93\xf1\x43\x98\x2c\x23\xd5\xee\xcf\x99\x99\x24\x9d\x6c\x0c\x3f\xfe\xba\x84\x5c\x4c\x91\xd8\xe3\x87\x17\x25\xca\x55\x7b\xd3\xed\x7d\xa1\x0e\x5c\xd0\x71\xf7\x0e\x83\x5a\x88\x26\x23\xe8\xb5\xeb\x5d\x0f\x7b\xe3\xbb\x49\xf0\x60\xc3\xd6\xd0\xf3\x1e\x06\xde\x56\xe9\x4d\xb9\x52\xa2\xa8\x44\x2c\x44\xf8\x96\xa1\xa7\x06\xc4\xe1\xcc\x23\x03\xd8\xd4\x3f\xd2\xb0\x3d\x35\xa2\x6e\x23\xff\xfb\x6d\x1e\x66\xee\x98\x62\xc4\x33\x20\x9b\xd5\x1e\xf1\x1e\x7a\xbd\xc1\x00\x9f\x37\xc2\xd9\xd9\x17\x91\x06\xb4\xc7\xdf\x86\x13\x40\x3a\xff\x19\xe7\xfa\xe1\xf6\x0b\xa5\xaa\xdb\xe2\x8c\x7a\x45\x4e\xa5\x28\x9a\xae\x2c\xde\x92\x38\x17\xbd\x20\xa7\xb7\x28\xb7\xb8\x1b\x99\x70\x70\x16\x89\x6e\x89\x47\x7d\xd0\x87\x3d\x9f\x03\x05\x82\x64\xee\x88\x78\x64\xcb\xc3\x5c\xbf\xe3\x3b\x83\xf7\x0d\xf9\xc3\x66\x13\x7a\xee\x6d\x04\x4b\x52\x74\x5a\x48\xdb\x20\xa2\x31\xd8\x70\x37\xc7\x80\x12\xfd\x1c\xae\xf9\xc0\x8a\x38\x36\xe2\x4a\xb9\x47\x44\x78\x25\x1d\x00\xb6\x1a\xd9\x42\x8c\xdb\xe8\x68\x98\x11\x92\x9c\x4b\x9f\x79\xe5\xb5\x0e\x93\x20\x73\x09\xff\x5f\xc1\x32\xfe\x5f\xbe\xb8\x5d\xca\xff\x6b\xbb\x54\x89\xbc\xff\x54\x01\xc8\xfc\xbf\xd6\x01\x98\x12\x2a\xc8\xdd\x24\x6d\x75\xb9\xad\x6f\x23\x88\x53\x2e\x10\xff\x59\xbe\xcd\x1c\xd3\xb8\x41\xf2\x26\x26\x95\xf8\x6e\x18\x24\x09\x28\xb1\x83\x03\x8d\x46\x0b\xd2\xc5\x0b\xa1\x58\x64\xa2\xbd\xe8\xcf\x26\x28\xfd\x1e\x53\x90\x97\x2a\xe5\xfd\xf2\xde\x5e\xe5\x6d\xe5\xf2\x6d\xa9\xbf\xbb\x7d\xb5\x5f\xde\x75\x77\xca\xfb\xa5\xca\xde\xce\x55\x69\xa7\xdf\xdf\x76\x2b\x97\x7b\xfd\x2b\xe6\x1c\x26\xed\x81\x89\x06\x22\xda\x9d\xbe\xfa\x86\xf2\xcf\x9d\x0c\x79\x56\x48\x4f\x57\x91\x6f\xd0\xa8\xa5\x47\x8d\x0f\xc7\xed\x17\x3f\x5d\x25\xd4\x0f\xaa\x47\x8f\x49\x0c\x83\x1f\x29\xad\x25\x1f\x9b\x38\x49\x8a\x51\x4e\xae\x16\x14\x94\xeb\x61\xcc\x71\xaa\x7a\x58\x50\xa9\x18\x13\x23\x1c\xad\x18\x53\x90\xed\xb8\x79\xcc\x6e\x4c\x9f\x03\x64\xa2\x20\xad\x25\x47\xd7\x26\xd6\x92\x0b\xb2\xa3\x83\x7e\xd3\x17\xad\x19\xbf\xfb\x0b\x07\x6f\xc5\x12\x0f\x17\x7c\xce\x4f\xa3\x2c\x23\xff\x79\x74\xe9\xd2\xef\xff\xec\x6e\x47\xfc\x7f\xb3\xf8\x9f\x35\x81\x9c\x19\x4f\xdc\xd4\xd0\x57\x30\xe8\x0b\xa9\xe2\x93\x97\xcb\x05\x2f\xc6\xa1\xa0\xe7\x43\xdf\xf7\x6f\x77\x60\xf2\xbf\xab\x92\x6d\xa3\x54\x32\x58\x6e\x01\x3f\x6f\x42\x9e\x45\x4e\xb3\xf7\x84\x82\xd2\xff\xbc\x23\xd1\xa2\x98\xfa\x2e\x5d\x49\xb1\x6b\x4d\x59\x7a\xe4\x25\x14\x7c\xea\x31\x78\x4a\x58\x66\xfd\xb3\x3c\xa3\xb3\x9f\x00\x98\x15\xff\xb7\x17\xcd\xff\x56\xae\xbc\xc9\xd6\xff\x3a\x40\x63\xc1\xa4\xc3\x9a\xf7\xe3\xdb\x57\xeb\xad\x44\xf7\x67\x8a\x41\x53\xb9\x48\x5d\xe1\xb5\xa9\xb8\xc4\x8e\x78\x30\x91\x6a\xd5\xbf\x02\xcc\x93\xf7\xec\x50\xaf\xa6\x3c\xc0\x12\x98\xf3\x00\xff\x2c\x9a\x14\xb6\xf0\x1a\x49\xd7\xc2\xc6\x6b\x83\x59\x63\xe5\x86\x91\x83\x78\x64\x21\x3f\x95\xb8\xfa\x4b\xb6\x35\xe6\x5e\x71\xee\xd6\xce\x83\x87\x8d\x86\x07\xb8\x30\xcd\x0f\xb7\x0b\xde\xb8\x3d\xd5\x08\xa8\x1f\x70\xcc\xb6\xa5\x8c\x37\x7e\xd8\xd4\x98\x07\xa9\x81\x2d\x61\x2e\x07\x38\x55\xf3\xc9\x7c\xb7\x99\x61\x43\xfc\x8a\x6e\x2f\x95\x99\x9a\xc8\x73\xf1\xf2\xc6\xe2\x7c\xd6\x32\x58\x63\x6d\x1d\x7b\xb4\xbb\xc1\x4b\xa4\x29\x98\xcc\x73\x3d\x27\x32\x38\xfe\xaa\x63\x01\xe6\xce\xba\xde\x48\xc7\x58\x91\x05\x7c\x39\xa6\x4a\x13\x74\x2e\x86\x3e\xb5\x02\xc8\x20\x83\x0c\x32\xc8\x20\x83\x0c\x32\xc8\x20\x83\x0c\x32\xc8\x20\x83\x0c\x32\xc8\x20\x83\x0c\x32\xc8\xe0\x3f\x09\xff\x02\x2c\xec\x2b\x90\x00\xc8\x00\x00\x01\x00\x00\xff\xff\x77\x3f\x62\x51\xe2\x1f\x00\x00")

func lokomotiveKubernetesBaremetalTarGzBytes() ([]byte, error) {
	return bindataRead(
		_lokomotiveKubernetesBaremetalTarGz,
		"lokomotive-kubernetes-baremetal.tar.gz",
	)
}

func lokomotiveKubernetesBaremetalTarGz() (*asset, error) {
	bytes, err := lokomotiveKubernetesBaremetalTarGzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lokomotive-kubernetes-baremetal.tar.gz", size: 8162, mode: os.FileMode(420), modTime: time.Unix(1552509702, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"lokomotive-kubernetes-baremetal.tar.gz": lokomotiveKubernetesBaremetalTarGz,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"lokomotive-kubernetes-baremetal.tar.gz": {lokomotiveKubernetesBaremetalTarGz, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
