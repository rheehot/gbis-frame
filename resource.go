// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// _resource/BMDOHYEON_ttf.ttf
// _resource/directions_bus-48px.png
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


func _resourceBmdohyeon_ttfTtfBytes() ([]byte, error) {
	return bindataRead(
		__resourceBmdohyeon_ttfTtf,
		"_resource/BMDOHYEON_ttf.ttf",
	)
}

func _resourceBmdohyeon_ttfTtf() (*asset, error) {
	bytes, err := _resourceBmdohyeon_ttfTtfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_resource/BMDOHYEON_ttf.ttf", size: 808332, mode: os.FileMode(420), modTime: time.Unix(1580640699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __resourceDirections_bus48pxPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xfe\x01\x01\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x30\x00\x00\x00\x30\x08\x06\x00\x00\x00\x57\x02\xf9\x87\x00\x00\x00\x04\x73\x42\x49\x54\x08\x08\x08\x08\x7c\x08\x64\x88\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x1d\x87\x00\x00\x1d\x87\x01\x8f\xe5\xf1\x65\x00\x00\x00\x19\x74\x45\x58\x74\x53\x6f\x66\x74\x77\x61\x72\x65\x00\x77\x77\x77\x2e\x69\x6e\x6b\x73\x63\x61\x70\x65\x2e\x6f\x72\x67\x9b\xee\x3c\x1a\x00\x00\x01\x7b\x49\x44\x41\x54\x68\x81\xed\x98\x3d\x4e\xc3\x40\x10\x85\x3f\x42\x24\x44\x07\x29\x22\x8b\x03\x00\x0d\xe2\x00\x5c\x00\x0e\x41\x4d\x84\xa0\x04\x4b\xd4\x5c\x2d\x14\xb9\x01\x28\x55\xa0\xa0\x02\x44\xe5\x00\x22\x14\xac\x04\x6c\x56\xf1\x78\x3d\xce\x06\x65\x3e\x69\x14\xc5\x1a\x3f\xbf\xe7\xf5\xfa\x67\xc1\x30\x96\x9b\x95\x1a\xfb\xee\x02\x07\xc0\x36\xb0\x03\x64\x40\xc7\x55\x0b\xd8\xf8\xa5\x3f\x01\x5e\x80\x4f\xe0\xc9\xd5\x23\x70\xeb\xaa\xef\x7e\x1b\xa7\x03\x5c\x03\x43\x67\x4a\xb3\x86\x4e\x7b\xb3\x29\xf3\xa7\x7c\x9f\x45\x6d\xe3\x7e\x3d\x03\x27\xda\xe6\xaf\xe6\x60\xdc\xaf\x5c\xcb\xfc\x3e\xf0\x91\x20\xc0\x3b\xb0\x57\x66\xae\x25\x08\x70\x0e\xac\x0a\xfa\xb4\x69\x03\x67\x1a\x42\x4d\x4c\x58\x69\xdd\x95\x99\x93\xdc\x46\x0b\x60\x4d\xd0\xd7\x04\x05\xb0\x3e\xab\x41\x12\x60\x12\xb1\x4f\x1d\x2a\x1d\x4f\x32\x07\x16\x1a\x0b\x90\x1a\x0b\x90\x1a\x0b\x90\x1a\x0b\x90\x9a\x76\xc4\x3e\xfe\x93\x32\x29\xff\x7e\x04\x2c\x40\x6a\x2c\x40\x6a\x62\x02\x8c\x81\x4b\x60\xcb\x55\x0e\xbc\x25\xd4\x29\xc5\xff\xcc\xbb\x08\xf4\xe4\x81\xbe\xb2\x92\xea\xa8\x07\xc8\x02\x3d\x59\x44\x00\xa9\xce\x4c\x62\x2e\xa1\xd0\x27\x5e\xcc\xaa\x85\x8a\x8e\x24\xc0\xab\xf7\xff\x38\xd0\x13\xda\x56\x86\x44\xc7\x3f\x76\x14\x37\xfc\x1d\xd2\x82\xe9\xc9\x37\xa6\xfa\x25\x24\xd1\xe9\x6b\x04\xe8\x45\x98\xd3\xaa\x9e\x46\x80\x36\x30\x48\x60\x7e\x40\xdc\xcb\x66\x90\x6e\x82\x00\x5d\x89\xb1\x2a\x8b\x54\x2a\xf7\xe4\x0a\x88\xbc\x69\x0c\x51\xdd\x95\xba\x5a\x27\x66\xa9\xde\x85\x1e\x02\xdb\xee\x15\x3c\x34\xa5\x3b\xc5\x11\x30\xe2\x67\x92\x8d\x80\xc3\x05\xd6\x35\x0c\xc3\x98\x03\x5f\xdc\x6b\x4f\xd9\x2b\xb8\x8d\xbc\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x51\xc0\xd0\x8f\xfe\x01\x00\x00")

func _resourceDirections_bus48pxPngBytes() ([]byte, error) {
	return bindataRead(
		__resourceDirections_bus48pxPng,
		"_resource/directions_bus-48px.png",
	)
}

func _resourceDirections_bus48pxPng() (*asset, error) {
	bytes, err := _resourceDirections_bus48pxPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_resource/directions_bus-48px.png", size: 510, mode: os.FileMode(420), modTime: time.Unix(1580640699, 0)}
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
	"_resource/BMDOHYEON_ttf.ttf":       _resourceBmdohyeon_ttfTtf,
	"_resource/directions_bus-48px.png": _resourceDirections_bus48pxPng,
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
	"_resource": &bintree{nil, map[string]*bintree{
		"BMDOHYEON_ttf.ttf":       &bintree{_resourceBmdohyeon_ttfTtf, map[string]*bintree{}},
		"directions_bus-48px.png": &bintree{_resourceDirections_bus48pxPng, map[string]*bintree{}},
	}},
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