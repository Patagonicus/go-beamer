// Code generated by go-bindata.
// sources:
// admin.tpl
// edit_playlist.tpl
// new_info.tpl
// new_playlist.tpl
// DO NOT EDIT!

package assets

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

var _adminTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcf\x41\x6a\x86\x30\x10\x05\xe0\x7d\x4e\x11\x3c\x80\xc3\xbf\x9f\x0e\x14\xda\x45\x37\xd2\x2b\x8c\x64\x34\x81\x18\x25\x0e\x88\x84\xdc\xbd\x44\x6d\x17\xdd\xe5\xc1\x97\x37\x3c\xf4\xba\x44\x32\x38\xae\xee\x24\x83\xfe\x45\xef\x6e\x09\x09\xc1\xbf\xc8\xa0\xf2\x18\x85\x4c\x29\x99\xd3\x2c\xb6\xff\xe6\x59\xf6\x5a\x0d\x6a\x26\x54\x47\xa5\xf4\x5f\x1f\xb5\x22\xa8\xbb\x32\xb2\xf5\x59\xa6\xb7\xae\x94\xfe\xd3\x05\xad\xb5\x6b\x66\xe0\x45\x9a\x62\xba\x25\x68\x6e\xa5\x92\x5c\xeb\x82\xe7\x8a\xf9\xfb\x0d\x49\x0e\x08\x69\x5a\x3b\x1a\xe4\xb0\xed\x65\x37\x9e\xe5\x6a\x18\xb3\x05\xfa\x67\xb7\xc8\x67\x0c\xbb\xde\xfe\x37\x35\x6e\x10\x9e\x69\x70\x2f\xfd\x09\x00\x00\xff\xff\xc8\x82\x18\xb4\xf1\x00\x00\x00")

func adminTplBytes() ([]byte, error) {
	return bindataRead(
		_adminTpl,
		"admin.tpl",
	)
}

func adminTpl() (*asset, error) {
	bytes, err := adminTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin.tpl", size: 241, mode: os.FileMode(420), modTime: time.Unix(1510598397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _edit_playlistTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xcd\x6a\xeb\x30\x10\x85\xf7\x7a\x8a\x41\x78\x79\xf1\x90\x6d\x18\x0b\xc2\xbd\x59\xdc\x4d\x09\xf4\x09\xe4\x6a\x12\x0b\xfc\x87\x3c\x0e\x0d\x42\xef\x5e\x54\x3b\xa5\x4e\x43\x17\x12\xcc\x41\x73\xce\x27\x66\xa8\x91\xae\x35\x8a\xea\xc1\xdd\x8c\xa2\x66\x67\x8e\xce\x0b\x8c\xad\xbd\xb5\x7e\x12\xd0\x31\x96\xa7\xb5\x28\x5f\x6c\xc7\x29\x69\xc2\x66\x67\x14\x9d\x87\xd0\x81\x7d\x13\x3f\xf4\x95\x46\x76\x5e\xf0\xde\x86\xdf\xbb\xfe\xff\x4b\x09\x27\x7b\x65\x0d\x1d\x4b\x33\xb8\x4a\x8f\xc3\x24\xda\x28\x80\xec\xb8\x07\xf2\xfd\x38\x0b\xc8\x6d\xe4\x4a\x0b\xbf\x8b\x86\xde\x76\x5c\xe9\x7c\x6b\xb8\xda\x76\xe6\xea\x19\x09\xa0\xa1\x3a\x00\x66\xab\x8d\xc9\x34\xd7\x9d\x97\xaf\xd6\xd7\xb5\x44\xa3\x08\x33\xb7\x51\xea\xef\x1c\x02\xf7\x02\xdc\x4b\xf0\x3c\xed\x15\x89\xad\x5b\x36\x2a\xc6\x60\xfb\x0b\x43\xe1\xff\x40\xc1\xb0\xaf\xa0\x3c\x2e\x6f\x52\x52\x24\xc1\x90\x38\x13\x63\xe1\x53\x22\x14\x77\x2f\xf9\xf3\xa3\x5b\x65\xc1\x5c\x34\x94\x90\xbd\xb9\x77\xd9\x06\x97\xb0\x95\x5e\x1d\xae\xd6\xb7\x59\x81\xd1\x5e\x9e\xc1\x94\xa7\xac\x6f\x00\x1e\xf3\x36\x69\xf9\x58\x68\x02\x9f\x7f\xcc\xc6\x3a\x87\x31\x16\x0f\x03\x5a\xfd\xb4\x39\x38\x47\x68\xcd\x2f\xd0\x8a\x70\x5d\x17\x5c\xb6\xe7\x23\x00\x00\xff\xff\xbb\x67\x75\x1b\x45\x02\x00\x00")

func edit_playlistTplBytes() ([]byte, error) {
	return bindataRead(
		_edit_playlistTpl,
		"edit_playlist.tpl",
	)
}

func edit_playlistTpl() (*asset, error) {
	bytes, err := edit_playlistTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "edit_playlist.tpl", size: 581, mode: os.FileMode(420), modTime: time.Unix(1510601588, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _new_infoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\x4d\x6e\x85\x30\x0c\x84\xf7\x39\x85\xe5\x0b\x58\xec\x9d\x1c\xa1\x9b\x9e\xc0\x80\x69\x22\x91\x1f\x05\x53\xca\xed\xab\xbc\xc7\x5b\x8d\x46\xf3\x8d\xf4\x71\xb4\xbc\x07\xc7\x73\x5d\xef\xe0\x38\x4e\xe1\x4b\x2f\x48\x65\xab\xd0\xe4\x47\x99\xe2\x14\x1c\x6f\xb5\x67\x90\xc5\x52\x2d\x1e\xa9\xe8\x45\x83\xa0\xa5\xab\x98\x22\x64\xb5\x58\x57\x8f\xad\x1e\x86\xc1\x01\xb0\xe9\x9f\x49\x57\x81\x22\x59\x3d\xae\x7a\x2c\x3d\xb5\xf1\xc7\xc0\xf4\x59\x5f\x68\x2a\xed\x34\xb0\xbb\xa9\xc7\xe3\x9c\x73\x32\x84\x5f\xd9\x4f\xf5\xf8\xfd\x54\x0a\x8e\x69\x38\x8c\x7c\x4c\xe9\x2d\xfe\x1f\x00\x00\xff\xff\xff\x49\xc9\xe1\xc0\x00\x00\x00")

func new_infoTplBytes() ([]byte, error) {
	return bindataRead(
		_new_infoTpl,
		"new_info.tpl",
	)
}

func new_infoTpl() (*asset, error) {
	bytes, err := new_infoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "new_info.tpl", size: 192, mode: os.FileMode(420), modTime: time.Unix(1510596533, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _new_playlistTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x5d\x0a\xc3\x30\x0c\x83\xdf\x73\x0a\xe3\x0b\x98\xbe\x3b\x3e\xc2\x5e\x76\x82\xb4\xf5\x48\x20\x7f\xb4\xee\x4a\x6f\x3f\xc2\x3a\xd8\x8b\x84\x90\x40\x1f\x47\x2b\x59\x1c\xcf\x6d\xbd\xc4\x71\x9c\xe4\xa1\x27\xf4\x1c\xae\x9c\x76\x63\x8a\x93\x38\x7e\xb5\xad\x40\x58\x2c\xb5\xea\x91\xaa\x9e\xf4\x1b\xd0\xb2\x69\x30\x45\x28\x6a\xb1\xad\x1e\x7b\xdb\x0d\xc5\x01\x70\xaa\xfd\x30\xa8\xa1\xa8\xc7\xa1\x08\x24\x3c\x6f\x40\x7f\xad\x5d\x5d\x3d\xee\xc7\x5c\x92\x21\xbc\x43\x3e\xd4\xe3\xf3\x8e\x24\x8e\x69\x5c\x0f\xbf\xf9\xe8\x8b\xfb\x09\x00\x00\xff\xff\x6b\x89\xc9\x09\xb6\x00\x00\x00")

func new_playlistTplBytes() ([]byte, error) {
	return bindataRead(
		_new_playlistTpl,
		"new_playlist.tpl",
	)
}

func new_playlistTpl() (*asset, error) {
	bytes, err := new_playlistTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "new_playlist.tpl", size: 182, mode: os.FileMode(420), modTime: time.Unix(1510596542, 0)}
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
	"admin.tpl": adminTpl,
	"edit_playlist.tpl": edit_playlistTpl,
	"new_info.tpl": new_infoTpl,
	"new_playlist.tpl": new_playlistTpl,
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
	"admin.tpl": &bintree{adminTpl, map[string]*bintree{}},
	"edit_playlist.tpl": &bintree{edit_playlistTpl, map[string]*bintree{}},
	"new_info.tpl": &bintree{new_infoTpl, map[string]*bintree{}},
	"new_playlist.tpl": &bintree{new_playlistTpl, map[string]*bintree{}},
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

