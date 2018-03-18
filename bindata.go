// Code generated by go-bindata.
// sources:
// static/index.html
// static/main.css
// static/main.js
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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcd\x8e\xe3\x2c\x10\xbc\xfb\x29\x5a\x9c\xbe\x4f\xab\x98\xcc\x35\xc2\x96\x76\x94\xcb\x4a\x7b\xda\x37\xc0\xd0\x89\xd9\x60\x40\xfc\x69\xac\xc8\xef\xbe\x22\xcc\x68\x1c\x4f\xb2\xda\x9c\x50\xa5\xab\xa8\x2e\xba\xcd\xc6\x38\xe9\xbe\x69\xd8\x88\x5c\xf6\x0d\x00\x00\xd3\xca\x5c\xc0\xa3\xee\x48\x88\xb3\xc6\x30\x22\x46\x02\x71\x76\xd8\x91\x88\x6f\x91\x8a\x10\x08\x8c\x1e\x4f\x1d\xa1\x13\x57\xa6\x2d\x40\xdf\x30\x5a\x45\x1a\x36\x58\x39\xbf\x8b\x49\x95\x41\xc9\x8e\x70\xe7\x48\x85\x6e\xf0\xf8\xd2\xbf\x26\x79\xc6\xc8\xe8\xf8\xb2\xc2\x93\x86\xbc\x13\xda\xf2\xcb\x27\xf8\x6e\xaa\x7f\xe5\x9a\x1b\x81\x07\xb8\x5e\x87\x1b\xb7\x1d\x2a\xb2\x2c\x8c\x6a\xf5\x95\x70\xe4\x4a\xcf\x50\x6b\x57\x2c\x59\xe0\x67\x9c\x5f\x58\x3a\x52\xe6\x0c\x92\xcf\x61\xc5\xf2\x1f\x7f\x1c\xf9\x1c\xb6\x6c\x46\x53\x49\x71\xdd\xde\x0f\xe3\xd2\xb6\xbb\x93\xf5\x13\xe4\x9d\x35\x87\x90\x86\x49\xc5\xd6\x79\xcc\x68\x62\x47\xc8\xbd\x93\x23\x06\xe1\x95\x8b\xca\x9a\x7b\x87\xaa\xa8\x42\xde\x4d\x56\x96\x07\x92\x9f\x75\x1b\x05\x36\x78\xda\x3f\xe6\x5a\x73\x10\x5a\x89\x4b\x47\xb8\x94\xff\xfd\xff\xf1\xb4\xd5\x12\x81\xcc\x75\xc2\x8e\x7c\x23\x3d\x7c\x9f\x6c\x32\xf1\xaf\x0e\xf8\xad\x84\x80\xd3\x5c\xe0\x68\xb5\x44\xdf\x91\x7d\xbb\xdf\x6f\xfd\x7c\xbd\x3d\xa4\xe1\xe9\xed\xbb\xf5\xb0\xd0\x92\xdb\x26\xde\x9f\x2a\x44\xb0\x27\x88\x9e\x9b\xc0\x45\x49\x20\xfc\xfb\x2c\x41\xde\x9d\xac\xef\x48\x04\x65\xee\x24\x36\xae\xcb\xef\x7a\x8d\xed\x2a\xe7\x65\xb9\x21\xb5\xef\x65\xb9\x57\x7e\x30\x14\xf5\x24\x55\x2e\x0b\x52\x17\xa3\x61\x55\x0d\x82\x17\x1d\x19\x63\x74\xe1\x40\xa9\x90\xa6\xfd\x1d\x24\x6a\x95\x7d\x6b\x30\x52\xe3\x26\x9a\x13\x92\x9e\xd1\x5a\xdf\x3f\x26\x26\xe3\x2e\xe7\x56\xd8\x89\xf2\x37\x65\x03\x95\x2a\xc4\x7a\x6c\x27\x55\x34\x9f\x2a\xd4\xed\xbd\x2f\x28\x6b\x5c\xbe\x09\x7f\x02\x00\x00\xff\xff\x42\x55\xeb\x8b\x1a\x04\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 1050, mode: os.FileMode(420), modTime: time.Unix(1521176267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMainCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x2e\xd3\x4d\xce\xc9\x4f\xcc\x8e\x55\xb0\x53\xd0\x52\xa8\x56\x48\xc9\x2c\x2e\xc8\x49\xac\xb4\xca\xcb\xcf\x4b\xb5\x56\xa8\xe5\x02\x04\x00\x00\xff\xff\xf7\x95\xc9\x41\x20\x00\x00\x00")

func staticMainCssBytes() ([]byte, error) {
	return bindataRead(
		_staticMainCss,
		"static/main.css",
	)
}

func staticMainCss() (*asset, error) {
	bytes, err := staticMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/main.css", size: 32, mode: os.FileMode(420), modTime: time.Unix(1521147003, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMainJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xcf\x4e\x84\x30\x10\xc6\xef\x3c\xc5\x04\x0f\x40\x44\xea\x9f\x8b\x61\x83\x07\x0f\x3e\x81\xf1\x62\x3c\x74\xe9\xec\x42\xb2\xdb\x36\x6d\x51\x93\x0d\xef\x6e\xda\x06\xec\xb2\x4d\xcc\x7a\x51\x6e\xed\x4c\xbf\xfe\xbe\xaf\xc3\x3b\x55\x40\xa5\x84\x06\x38\x7e\xc0\xcb\x80\xf9\x21\x01\x00\xc0\x5d\x0d\xd9\x05\x95\x32\x2b\xdd\x9a\x51\x43\x6b\xf0\x35\xfb\x19\x45\xb9\xa6\xad\xe9\x05\xd7\x35\xbc\xbe\x95\x73\x65\x3d\xb0\x2d\x9a\x1a\x0e\x63\x99\xcc\x9b\x84\xc0\x93\x50\x80\xdc\xa0\xea\xf9\x16\xa8\xbb\x2e\x10\xa9\xe6\x56\xba\x17\x03\x37\x35\xa4\xe9\xb7\x26\x43\xdd\xaa\x5e\xda\x46\x5b\x70\xfb\xa3\x2f\xb7\x0a\xa9\x41\x96\x17\x21\x5c\xd7\xeb\x6a\x83\xa6\xed\x9e\x03\xcc\xbc\x58\x45\x3a\x1e\x1d\xee\x54\x9b\x98\xf7\x68\x3a\xc1\x74\xe8\x38\xa2\x17\x54\x1d\xf9\x67\x2f\x74\x65\xd5\x32\x42\x65\x4f\x02\x7b\xe4\xf6\xfa\xe6\x9e\xdc\x65\x45\x65\x3a\xe4\xb9\x42\x2d\x05\xd7\x08\xcd\xc3\x42\x63\x86\x0b\x03\x86\x06\xa6\x13\x95\x7d\x88\xd5\xd1\x91\x31\xf0\x15\x66\x7e\x64\xef\x6c\x54\xe2\xdf\xf1\x0c\x62\x7f\xe0\x77\xac\x94\xb1\x13\x46\x3b\x9a\x1a\x77\x1b\x68\x9c\xfe\x2a\x62\x40\x0a\x1d\x71\x90\x95\x11\xc4\xa3\x11\x72\xbc\xc1\x4e\x79\xd2\x3e\x4d\xa1\xeb\xf4\x8b\x85\x11\x9f\xcb\x66\xe0\xee\x4e\x98\x13\x5a\xda\xb0\x9f\xb5\x11\x1b\xb7\x9f\x62\xd1\xc3\xfa\x7f\xc6\x92\x5e\xa5\x70\xf9\x17\xe1\xf8\xbf\x34\x19\x8b\x24\xf9\x0a\x00\x00\xff\xff\xc1\xc7\x53\xb0\xbb\x04\x00\x00")

func staticMainJsBytes() ([]byte, error) {
	return bindataRead(
		_staticMainJs,
		"static/main.js",
	)
}

func staticMainJs() (*asset, error) {
	bytes, err := staticMainJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/main.js", size: 1211, mode: os.FileMode(420), modTime: time.Unix(1521175990, 0)}
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
	"static/index.html": staticIndexHtml,
	"static/main.css": staticMainCss,
	"static/main.js": staticMainJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
		"main.css": &bintree{staticMainCss, map[string]*bintree{}},
		"main.js": &bintree{staticMainJs, map[string]*bintree{}},
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

