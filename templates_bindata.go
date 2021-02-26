// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package main generated by go-bindata.// sources:
// templates/dot.tmpl
// templates/dot_relations.tmpl
// templates/dot_subgraphs.tmpl
// templates/dot_tables.tmpl
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _templatesDotTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x41\x6b\xe3\x30\x10\x85\xef\xfe\x15\xc2\x67\x47\xeb\x78\x77\xa1\xa5\x51\xa0\x87\x16\x72\x49\x0a\xcd\xa9\xa5\x14\xd9\x9a\xd8\x6a\x95\x91\x91\xc6\x04\x2a\xfc\xdf\x8b\x9d\x04\x3b\x69\x48\x7d\x30\xe2\x9b\xa7\x79\x8f\xd1\x84\x30\x61\x0a\x36\x1a\x81\xc5\xca\x52\xcc\x26\x6d\x1b\x95\x4e\xd6\x15\x0b\x11\x63\x8c\xed\xcf\xaf\xfd\xb9\xfb\xba\x0b\x7a\xc3\xf8\x5a\x93\x81\xfd\xff\x9e\xc8\xe9\xbc\x21\xf0\xdc\xc8\x1c\x4c\xdf\xe3\xa8\xef\x89\x98\xcd\x1e\x57\xcb\x35\x7b\x5a\x2d\x96\xeb\xc9\xf3\xe2\xe5\x41\xc4\x59\x1a\xcf\x43\xb8\xd6\xa7\x6d\x67\x7f\xba\x6b\xf3\x79\x72\xda\xee\xa3\xf1\x24\xcc\x19\x34\xb6\x10\x94\x9c\xe4\x04\x54\x27\x59\xd0\x2a\xf0\x50\x8b\x94\xff\x1f\x84\x4e\xe2\xe7\x0f\x58\x4b\x25\xe2\x94\x67\x49\xca\xb3\x78\xc0\x5b\xe9\x4a\x8d\x5d\x25\x1d\xd1\x02\xb0\x00\x24\x27\x09\x04\xb9\x06\x86\x8a\xaf\x8d\x46\xf0\xc2\x3a\xaa\x6c\x4f\xdf\xee\xa2\x63\x92\xd1\x50\x37\x16\xc9\xeb\x2f\x10\xd3\x7f\x17\xcc\x46\x56\x35\xe0\x4e\x2b\xaa\xc4\x94\xa7\x23\x9b\x4a\xd6\x20\xd0\x22\x8c\x3d\x40\x95\x63\x0f\xa5\x9d\xc8\x2d\x55\xc9\x05\xd7\x6c\x80\xd2\x39\xbb\xeb\x69\xca\x6f\x7f\x33\xee\x07\x2f\xb1\x34\x20\xfe\x66\x67\x58\x69\x4f\x12\x0b\x10\x53\x7e\x33\x8e\x15\x02\xc1\xb6\x36\x92\xf6\x1b\xf7\xee\x9b\xbc\xdf\x31\x1f\x33\x7e\x78\xab\x73\x89\x03\x23\x49\x5b\xbc\x22\x21\x99\x1b\x38\xd4\xdb\x28\x04\x40\xd5\xb6\xdf\x01\x00\x00\xff\xff\x55\x2a\xfd\x0d\xdd\x02\x00\x00")

func templatesDotTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDotTmpl,
		"templates/dot.tmpl",
	)
}

func templatesDotTmpl() (*asset, error) {
	bytes, err := templatesDotTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot.tmpl", size: 733, mode: os.FileMode(436), modTime: time.Unix(1592765204, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDot_relationsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xd1\x4b\xc3\x30\x10\xc6\xdf\xfb\x57\x1c\x7d\x72\xba\x54\xf7\xec\x3a\x10\xc1\x27\x99\x30\xf6\x26\x22\x37\x73\xdb\x02\x31\xc1\x34\x32\x47\xc8\xff\x2e\xa9\x5d\xac\x6b\x5a\x9d\x6f\xc7\xf1\xf5\xf7\x5d\xee\xbe\x3a\xc7\x69\x2d\x14\x41\xce\xb5\x7d\x36\x24\xd1\x0a\xad\xaa\xdc\xfb\xcc\x39\x83\x6a\x43\x50\x2c\x0e\x5d\xef\x33\x00\xe7\x8a\x7b\x5a\xdb\x25\xae\x24\xcd\xf1\x95\xbc\x07\xc6\x42\x77\x21\x36\xdb\x1f\xed\xc7\x0c\x20\xe8\x19\x88\x35\x9c\xd1\x1b\x7c\x49\x6e\xd1\x70\xa1\x50\x0a\xbb\x87\xfc\x3c\x1f\x01\xab\xb9\x00\x68\x8c\xde\x6d\x09\x79\xa9\x5f\x8c\xde\x8d\x43\x29\x71\x45\xb2\x9c\x4e\xef\x1e\xe6\xcb\xd9\x55\x51\xcc\xa7\x97\x75\x39\x1b\x47\x38\xc9\x8a\x06\x1c\x2e\xf2\x51\x9a\x6f\x89\xba\x16\x93\xff\x58\xb0\x84\x85\xd2\x8a\x8e\x00\xdd\x77\x06\x91\x4e\xce\x71\xd8\x67\xcb\xc7\xfb\xc4\x60\x8a\x47\x6c\xb3\xe8\x78\xae\x1b\x6b\x8d\x58\xbd\x5b\xaa\x0a\x49\x2a\xca\x24\xa9\x32\xd0\x93\x2a\xef\x4f\x67\x87\xb1\xa3\xf0\xa3\xf3\x8a\x9e\x2f\xfe\xf8\x9a\x7a\xe3\x21\x6f\xc3\xa9\xb1\x28\x64\x93\x9a\x50\xf6\xa6\x26\x7d\xd1\x0e\xff\x28\x33\xdf\xf4\x70\xab\x8e\xc1\xe4\x74\x03\x96\x30\xf8\x2d\x31\x51\xa4\x93\x53\x34\xff\x65\x32\x30\x43\x1b\x46\xc5\xfb\x87\x1c\x48\x7c\xe4\x54\x76\x2f\xa9\xe4\x58\x6d\x89\xa7\x6e\xf9\x74\x9d\xb5\x1b\xad\xfa\x33\x00\x00\xff\xff\x6c\xee\x75\x57\x7d\x04\x00\x00")

func templatesDot_relationsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDot_relationsTmpl,
		"templates/dot_relations.tmpl",
	)
}

func templatesDot_relationsTmpl() (*asset, error) {
	bytes, err := templatesDot_relationsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot_relations.tmpl", size: 1149, mode: os.FileMode(436), modTime: time.Unix(1614335786, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDot_subgraphsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\xc1\xf8\xd8\x98\xa6\xc7\x1a\x1d\x4a\x8f\xfd\x84\x52\x82\x1c\xad\x5d\x11\x45\x36\xd2\xba\x50\x96\xfd\xf7\x52\x39\x4d\x31\x24\x17\xed\xac\xf4\x76\x46\x48\x22\x8e\x06\x1f\x09\xb5\x9b\xf8\x90\x97\x7e\x4c\x76\xfe\xcc\xb5\x6a\x25\x92\x6c\x1c\x09\x0d\x9f\x1e\xd0\x30\x9e\x0d\xda\xd7\xb0\x64\xa6\x94\x55\x2b\xe0\x8f\xc6\x71\xdd\x3d\x88\x34\x7c\x52\x85\x54\x00\xf0\x66\xf6\x8f\x5d\x51\x22\x7e\xb8\xce\xbe\x30\x27\xdf\x2f\x4c\xb9\x0d\xb6\xa7\xa0\x5a\x0a\x0c\x6a\x91\xfb\x50\xdd\x89\x50\x74\x25\xf8\xbe\x63\x3f\x1e\xa7\x30\x25\xd5\xcc\xdf\x81\xcc\xe0\x43\x20\xb7\x5e\xa2\x1c\x98\xdb\x19\xd7\xb1\x6d\xca\x6a\xe2\xe3\x97\xcf\xab\x47\x59\xe2\xe4\x08\xef\x45\xae\x4d\xa6\xd9\xec\x4b\xff\xd1\x55\xa5\x92\x1b\xff\x91\xb3\x8f\x81\xa2\x79\xda\x10\x22\x4c\xe7\x39\x58\xbe\xbc\x3c\xdb\x3e\x50\xae\xd1\xaa\xde\x06\x12\x05\xcb\x7e\x8a\x17\x06\xf8\xfd\xa0\x1d\x28\x3a\xec\x74\xa3\x7f\x02\x00\x00\xff\xff\xa6\xfd\x06\x47\xd3\x01\x00\x00")

func templatesDot_subgraphsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDot_subgraphsTmpl,
		"templates/dot_subgraphs.tmpl",
	)
}

func templatesDot_subgraphsTmpl() (*asset, error) {
	bytes, err := templatesDot_subgraphsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot_subgraphs.tmpl", size: 467, mode: os.FileMode(436), modTime: time.Unix(1592765656, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDot_tablesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\xdf\x6f\x9b\x30\x10\x7e\xe7\xaf\x38\x9d\xaa\x3d\x25\x90\x6e\xdd\x1e\x56\x1b\x89\x5f\x69\x90\x18\x44\xd4\xda\xa4\x4d\xd3\x04\xc1\x89\xd0\x5c\xa8\xc0\x9d\x54\x21\xfe\xf7\xc9\x2e\x24\x50\x85\xb7\xf3\xf9\xbe\xfb\xbe\xfb\xce\xee\xba\x82\x1f\xcb\x8a\x03\x16\xb5\xfc\x23\xb3\x5c\xf0\x16\xfb\xde\xe8\xba\x26\xab\x4e\x1c\x6e\xe4\xdf\x15\xdc\x48\xf8\x4a\xc1\x64\xfa\xb6\xef\x0d\x80\xae\x93\xfc\xe9\x59\x64\x72\x0a\x44\x30\x35\x72\x0d\xbc\x2a\x60\xfd\x2e\x36\xae\x50\xe1\xd0\xcc\x64\xa5\x14\xbc\xef\xe1\x97\xc8\x72\x2e\x28\x31\x00\x00\x08\x73\xdc\x28\xd0\x21\x80\x13\x85\x0f\x31\xc5\x28\xd8\x32\x1c\x52\x5e\x10\x45\x7b\xc7\xf7\xc3\xf8\x81\xe2\x66\x9a\x7d\xdc\x3b\x9e\xce\xde\x8d\xd9\x1f\xa1\xcf\x76\x14\x6f\x3f\x0d\x19\x7b\xc8\x13\x96\x8e\xa1\x3a\xf8\x23\x8f\x17\xc4\x2c\x48\x11\xdc\x24\xf5\x83\x54\xb5\x9f\x37\xde\x98\x9f\x11\xbe\x0f\xc5\x6e\xc2\x58\xf2\x0d\xa7\x24\x36\xd9\x26\x31\x83\x7d\x12\xc6\x6c\xfd\x18\xfe\x0c\x28\xde\xde\x21\x6c\x1d\x2f\xa0\xb8\xe3\xe2\x1f\x97\xe5\x21\x83\xbc\x16\x05\xda\xc4\xb5\x2f\x1e\x10\xcb\xb5\x89\xa5\xd0\x36\xb1\x98\x7f\x16\x6a\x5d\x94\x92\x5d\x6a\x8d\xb1\xb2\xb8\x3c\x82\xe9\xd5\xe2\xe5\xa9\x6a\xb5\xd5\x97\x9b\x61\x89\x6a\x87\x07\xbd\xc3\xa1\xec\x5c\xb4\x34\xbf\xf6\x79\x32\xfd\xb5\x79\x3e\xe2\x4c\xb6\x96\x7c\xee\x35\xd7\xe5\x48\xd9\x94\xf9\x8b\xe4\xad\xa9\x37\x3c\x91\xa9\x89\x75\xf3\x37\x73\x9c\xa6\xcc\x04\x84\x32\x13\xe5\x01\xe7\x84\x6a\x09\x49\x94\xa4\x14\x4f\x0d\x7f\xfd\xb2\x41\xfb\x43\x95\xb7\xcf\xf7\x5d\xb7\x40\x73\x5d\xd6\xf8\x22\xcf\xec\x0b\x36\x0f\xb5\x33\x43\xa7\x58\x62\xe9\x07\xaa\xca\x6d\xe3\x32\xb1\xfe\x26\x13\x25\xf9\xe9\x50\x8b\xba\xd1\xa0\xd5\xb1\x14\x42\x1f\x29\x2a\xef\x96\x4a\x71\x65\x00\xb4\xf2\x55\x70\xaa\x10\xbc\x30\xde\xd3\xff\xbe\x9f\xfe\xae\xff\x01\x00\x00\xff\xff\xca\xed\x4a\x5f\xc7\x03\x00\x00")

func templatesDot_tablesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDot_tablesTmpl,
		"templates/dot_tables.tmpl",
	)
}

func templatesDot_tablesTmpl() (*asset, error) {
	bytes, err := templatesDot_tablesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot_tables.tmpl", size: 967, mode: os.FileMode(436), modTime: time.Unix(1592671482, 0)}
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
	"templates/dot.tmpl":           templatesDotTmpl,
	"templates/dot_relations.tmpl": templatesDot_relationsTmpl,
	"templates/dot_subgraphs.tmpl": templatesDot_subgraphsTmpl,
	"templates/dot_tables.tmpl":    templatesDot_tablesTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"dot.tmpl":           &bintree{templatesDotTmpl, map[string]*bintree{}},
		"dot_relations.tmpl": &bintree{templatesDot_relationsTmpl, map[string]*bintree{}},
		"dot_subgraphs.tmpl": &bintree{templatesDot_subgraphsTmpl, map[string]*bintree{}},
		"dot_tables.tmpl":    &bintree{templatesDot_tablesTmpl, map[string]*bintree{}},
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
