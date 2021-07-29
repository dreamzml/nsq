// Code generated by go-bindata. DO NOT EDIT.
// sources:
// static/build/base.css
// static/build/bootstrap.min.css
// static/build/favicon.png
// static/build/glyphicons-halflings-regular.eot
// static/build/glyphicons-halflings-regular.svg
// static/build/glyphicons-halflings-regular.ttf
// static/build/glyphicons-halflings-regular.woff
// static/build/glyphicons-halflings-regular.woff2
// static/build/index.html
// static/build/main.js
// static/build/main.js.map
// static/build/nsq_blue.png
// static/build/vendor.js

package nsqadmin

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// bindataBaseCss reads file data from disk. It returns an error on failure.
func bindataBaseCssBytes() ([]byte, error) {
	asset, err := bindataBaseCss()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataBaseCss() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/base.css"
	name := "base.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataBootstrapMinCss reads file data from disk. It returns an error on failure.
func bindataBootstrapMinCssBytes() ([]byte, error) {
	asset, err := bindataBootstrapMinCss()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataBootstrapMinCss() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/bootstrap.min.css"
	name := "bootstrap.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataFaviconPng reads file data from disk. It returns an error on failure.
func bindataFaviconPngBytes() ([]byte, error) {
	asset, err := bindataFaviconPng()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataFaviconPng() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/favicon.png"
	name := "favicon.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataGlyphiconshalflingsregularEot reads file data from disk. It returns an error on failure.
func bindataGlyphiconshalflingsregularEotBytes() ([]byte, error) {
	asset, err := bindataGlyphiconshalflingsregularEot()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataGlyphiconshalflingsregularEot() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/glyphicons-halflings-regular.eot"
	name := "glyphicons-halflings-regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataGlyphiconshalflingsregularSvg reads file data from disk. It returns an error on failure.
func bindataGlyphiconshalflingsregularSvgBytes() ([]byte, error) {
	asset, err := bindataGlyphiconshalflingsregularSvg()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataGlyphiconshalflingsregularSvg() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/glyphicons-halflings-regular.svg"
	name := "glyphicons-halflings-regular.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataGlyphiconshalflingsregularTtf reads file data from disk. It returns an error on failure.
func bindataGlyphiconshalflingsregularTtfBytes() ([]byte, error) {
	asset, err := bindataGlyphiconshalflingsregularTtf()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataGlyphiconshalflingsregularTtf() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/glyphicons-halflings-regular.ttf"
	name := "glyphicons-halflings-regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataGlyphiconshalflingsregularWoff reads file data from disk. It returns an error on failure.
func bindataGlyphiconshalflingsregularWoffBytes() ([]byte, error) {
	asset, err := bindataGlyphiconshalflingsregularWoff()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataGlyphiconshalflingsregularWoff() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/glyphicons-halflings-regular.woff"
	name := "glyphicons-halflings-regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataGlyphiconshalflingsregularWoff2 reads file data from disk. It returns an error on failure.
func bindataGlyphiconshalflingsregularWoff2Bytes() ([]byte, error) {
	asset, err := bindataGlyphiconshalflingsregularWoff2()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataGlyphiconshalflingsregularWoff2() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/glyphicons-halflings-regular.woff2"
	name := "glyphicons-halflings-regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataIndexHtml reads file data from disk. It returns an error on failure.
func bindataIndexHtmlBytes() ([]byte, error) {
	asset, err := bindataIndexHtml()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataIndexHtml() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/index.html"
	name := "index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataMainJs reads file data from disk. It returns an error on failure.
func bindataMainJsBytes() ([]byte, error) {
	asset, err := bindataMainJs()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataMainJs() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/main.js"
	name := "main.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataMainJsMap reads file data from disk. It returns an error on failure.
func bindataMainJsMapBytes() ([]byte, error) {
	asset, err := bindataMainJsMap()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataMainJsMap() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/main.js.map"
	name := "main.js.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataNsqbluePng reads file data from disk. It returns an error on failure.
func bindataNsqbluePngBytes() ([]byte, error) {
	asset, err := bindataNsqbluePng()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataNsqbluePng() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/nsq_blue.png"
	name := "nsq_blue.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataVendorJs reads file data from disk. It returns an error on failure.
func bindataVendorJsBytes() ([]byte, error) {
	asset, err := bindataVendorJs()
	if asset == nil {
		return nil, err
	}
	return asset.bytes, err
}

func bindataVendorJs() (*asset, error) {
	path := "/mnt/e/siteProgram/golangProject/src/nsq/nsqadmin/static/build/vendor.js"
	name := "vendor.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"base.css":                           bindataBaseCss,
	"bootstrap.min.css":                  bindataBootstrapMinCss,
	"favicon.png":                        bindataFaviconPng,
	"glyphicons-halflings-regular.eot":   bindataGlyphiconshalflingsregularEot,
	"glyphicons-halflings-regular.svg":   bindataGlyphiconshalflingsregularSvg,
	"glyphicons-halflings-regular.ttf":   bindataGlyphiconshalflingsregularTtf,
	"glyphicons-halflings-regular.woff":  bindataGlyphiconshalflingsregularWoff,
	"glyphicons-halflings-regular.woff2": bindataGlyphiconshalflingsregularWoff2,
	"index.html":                         bindataIndexHtml,
	"main.js":                            bindataMainJs,
	"main.js.map":                        bindataMainJsMap,
	"nsq_blue.png":                       bindataNsqbluePng,
	"vendor.js":                          bindataVendorJs,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"base.css": {Func: bindataBaseCss, Children: map[string]*bintree{}},
	"bootstrap.min.css": {Func: bindataBootstrapMinCss, Children: map[string]*bintree{}},
	"favicon.png": {Func: bindataFaviconPng, Children: map[string]*bintree{}},
	"glyphicons-halflings-regular.eot": {Func: bindataGlyphiconshalflingsregularEot, Children: map[string]*bintree{}},
	"glyphicons-halflings-regular.svg": {Func: bindataGlyphiconshalflingsregularSvg, Children: map[string]*bintree{}},
	"glyphicons-halflings-regular.ttf": {Func: bindataGlyphiconshalflingsregularTtf, Children: map[string]*bintree{}},
	"glyphicons-halflings-regular.woff": {Func: bindataGlyphiconshalflingsregularWoff, Children: map[string]*bintree{}},
	"glyphicons-halflings-regular.woff2": {Func: bindataGlyphiconshalflingsregularWoff2, Children: map[string]*bintree{}},
	"index.html": {Func: bindataIndexHtml, Children: map[string]*bintree{}},
	"main.js": {Func: bindataMainJs, Children: map[string]*bintree{}},
	"main.js.map": {Func: bindataMainJsMap, Children: map[string]*bintree{}},
	"nsq_blue.png": {Func: bindataNsqbluePng, Children: map[string]*bintree{}},
	"vendor.js": {Func: bindataVendorJs, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
