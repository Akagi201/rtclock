// Code generated by go-bindata.
// sources:
// data/template/index.html
// DO NOT EDIT!

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

var _dataTemplateIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x5f\x73\xe2\x36\x10\x7f\xe7\x53\x6c\x75\x6d\x03\x13\xfc\x8f\xe4\xae\x8d\xcf\xa6\x0f\x69\x67\xfa\x74\x97\xb9\xd0\xe9\xb3\x90\xd6\x58\x41\x96\x7c\x92\x88\xa1\x29\xdf\xbd\x23\x9b\x80\x71\x72\x97\xa9\x67\x0c\xde\xfd\x49\xfb\xe7\xb7\xab\x55\x56\xba\x4a\xce\x47\x59\x89\x94\xcf\x47\x00\x00\x59\x85\x8e\x02\x2b\xa9\xb1\xe8\x72\xb2\x71\x45\xf0\x2b\xe9\x43\xa5\x73\x75\x80\x5f\x37\xe2\x31\x27\xb7\x5a\x39\x54\x2e\x58\xec\x6a\x24\xc0\x3a\x29\x27\x0e\xb7\x2e\xf2\x96\xcf\x36\x2a\x5a\x61\x4e\x1e\x05\x36\xb5\x36\xae\xb7\xbc\x11\xdc\x95\x39\xc7\x47\xc1\x30\x68\x85\x29\x08\x25\x9c\xa0\x32\xb0\x8c\x4a\xcc\x93\x30\x7e\x36\xe5\x84\x93\x38\xff\xb2\xb8\x95\x9a\xad\xb3\xa8\x13\x3b\xc8\x32\x23\x6a\x07\xd6\xb0\x9c\xf8\x28\x6d\x1a\x45\x8c\xab\x07\x1b\x32\xa9\x37\xbc\x90\xd4\x60\xc8\x74\x15\xd1\x07\xba\x8d\xa4\x58\xda\xe8\xe1\xeb\x06\xcd\x2e\xba\x0a\x67\x61\x72\x10\xc2\x4a\xa8\xf0\xc1\x92\x79\x16\x75\x06\x0f\xd6\xa5\x50\x6b\x28\x0d\x16\x27\xe3\x85\x56\xce\x86\x2b\xad\x57\x12\x69\x2d\x6c\x6b\x9c\x59\xfb\x5b\x41\x2b\x21\x77\xf9\xdd\xe2\xf2\x9e\x2a\xfb\xaf\xff\x47\x23\x8a\xf4\x3a\x8e\xa7\xd7\x71\x2c\xa6\xbf\xc4\xb1\x7f\xc5\xcf\x76\xb3\xf4\x3c\xb3\x9d\x11\x52\x0a\x46\x5a\x5f\xdd\x63\x50\xe6\xc4\xba\x9d\x44\x5b\x22\x3a\xd2\x8f\x63\x88\x0d\x02\xdb\xa8\x7a\xbd\xea\x52\x6d\x6c\x15\x32\x6b\x23\x2e\xac\xeb\x24\x9f\x1f\xb3\xf6\xd9\x5e\x6b\x05\xdc\xae\xc6\x43\xe1\x4e\x98\x7f\x96\x9a\xef\xe0\xa9\x17\x16\xc0\x92\xb2\xf5\xca\xe8\x8d\xe2\x01\xd3\x52\x9b\x14\x96\x92\xb2\xf5\xc7\xb3\x45\x07\xa4\x29\x85\xc3\x73\xc4\xb3\x96\x42\x72\x55\x6f\xa3\x24\xbc\x86\x4a\x2b\xca\xf4\x14\xc8\xad\xde\x18\x81\x06\x3e\x61\x43\xa6\x70\x90\xa6\x1e\xd7\xb6\xa6\x6c\x60\xa5\xa2\x66\x25\x54\x0a\x71\xbd\x3d\x07\x6a\xca\xb9\x50\xab\x14\x92\xb8\xde\xc2\xec\x0c\xdf\x8f\x8e\x9f\x65\xf2\x66\x56\xef\x66\xb3\xd9\xab\x49\xad\x0c\xa2\xda\xa1\x94\xba\x79\x99\x5a\x60\xc5\x3f\x98\x42\x12\xce\xb0\x7a\x05\x6d\x50\xac\x4a\x97\xc2\x87\x38\x1e\xc4\xad\xad\x70\x42\xab\x14\x0a\xb1\x45\x7e\x0e\xb6\xa7\xc2\xa7\x14\xff\xf4\x0d\x1a\xce\xd5\x4e\xd7\x2f\x74\x12\x0b\xf7\x42\x79\x64\xeb\xfd\xf7\xc8\x7a\x57\xa1\xb5\x74\x85\x76\xc0\xd9\xb3\xf7\x2b\x4f\xb5\x7f\x93\xc3\xc7\xf7\x4a\xf2\x02\x7e\xa6\xfb\xe6\xe6\xe6\x0d\xef\xf5\xc0\x7f\xdb\x5d\x41\xdb\x1d\x29\xd4\x06\x83\xc6\xd0\x7a\xc0\x9c\x36\xbc\x55\xa7\xb0\x34\x48\xd7\x81\x57\xbc\xee\x9f\xf3\x01\x20\x85\xc2\xa0\x3c\x14\x2c\x79\x3f\xe4\xbe\x57\xed\x0f\x03\xe6\xfc\x6f\x16\xb5\x47\x6b\x3e\xca\xa2\x6e\xbc\x66\xfe\x2c\xf9\x61\x9b\x9c\xe6\x57\x99\xcc\x47\x19\x17\x8f\x20\x78\x4e\x9e\xf3\xf4\xb3\x87\x8b\xc7\xf9\x28\xeb\x4f\xa0\x1f\xc7\x5c\xb3\x4d\x85\xca\x4d\x42\x83\x94\xef\xc6\xc5\x46\x31\xdf\x32\x30\x9e\xf4\x88\x11\x05\x8c\x1b\xa1\xb8\x6e\xc2\xbf\x71\x79\xaf\xd9\x1a\xdd\x64\x48\x9c\x85\x1c\x14\x36\x70\x5c\x31\x26\x8d\x9f\x1c\x4f\x4f\xe1\x9f\xda\xba\xfd\x3e\x62\x3e\x40\x32\x19\xb0\x69\x43\xad\x0e\x61\x42\x0e\xa7\x00\x70\xe8\xa1\x8b\x98\x1c\x6b\x47\x26\xa1\xbf\x11\xc6\x24\xab\xe7\x04\x2e\x01\x43\x4e\x1d\x85\x4b\x20\x59\x54\xcf\xbd\x9f\x28\x7a\x63\x3f\xad\x6b\x54\xfc\x7f\x59\x60\x5a\x59\x2d\x31\x94\x7a\x35\x26\x0d\x2e\x6d\x9b\x6b\x0a\xa7\xfd\xdf\x70\x7c\xe1\x83\x9d\xb6\xd3\xef\x62\x12\x52\x25\x2a\xea\x70\xfc\x64\x99\xd1\x52\x2e\xfc\xe9\xea\x97\xa3\xeb\x91\xf1\x64\x3f\x85\x8b\x82\x5a\x77\x31\xa0\x6d\xdf\xeb\x0d\x40\x69\x71\x40\x16\x95\x68\xdc\x98\x1c\x8b\x01\x9f\x3e\x2f\xe0\xfe\xaf\xbb\xbb\xcf\x5f\x16\x7f\xfc\xfe\x43\xbf\x0a\x5d\x6f\xed\x27\x1f\x47\xa7\xfb\x29\x8b\x0e\x9d\xd5\x5e\xba\xf3\xff\x02\x00\x00\xff\xff\x5f\x46\x13\xb1\xd5\x07\x00\x00")

func dataTemplateIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataTemplateIndexHtml,
		"data/template/index.html",
	)
}

func dataTemplateIndexHtml() (*asset, error) {
	bytes, err := dataTemplateIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/template/index.html", size: 2005, mode: os.FileMode(420), modTime: time.Unix(1490156972, 0)}
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
	"data/template/index.html": dataTemplateIndexHtml,
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
	"data": &bintree{nil, map[string]*bintree{
		"template": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{dataTemplateIndexHtml, map[string]*bintree{}},
		}},
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
