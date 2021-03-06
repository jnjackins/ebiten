package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _text_png = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\xc0\x00\x00\x00\x80\b\x03\x00\x00\x00jZ\x96Y\x00\x00\x00\x06PLTE\x00\x00\x00\xff\xff\xff\xa5ٟ\xdd\x00\x00\x00\x01tRNS\x00@\xe6\xd8f\x00\x00\a\xb2IDATx\x9c\xec\x1c\x89v#!\b\xf2\xe6\xff\u007f\x99MT\xe4\xf4\x984m\x8e\r\xbb\xb5\x13\x83\x1c\xe2\x01\x8c\xf6\x02o\x0e_\x05n\x80\x83\x9aX?j\x1b\xf1w\xda\xde\xe0\x02x\x05iu\xfd\x87\xf5\xa1Է\xb2}\x8b\x86<\xb6\xf2\xb9`,\x80t\xfb\xd7>\xd0\x15J)ߒ\xd5\xe0/`\xd0Gҗ\x87\xae.\xf2\xa3\xb2g\x11\x9e|c\xaa\xa8\xfb\x12D\x1a\xbb\x90\xf0\xba}\xb8\xfep\xdda\xeb\xd9\x06\xa8x2\xeaU5E\xa8Yc\xa9ǽ\xa2k\xeaYoq\r+\x807\xe4f\x01C\xa2\x97T\xdb\xe0\xed\u007f\xed\x17~ڑ\xc40\xdd\x05\x19\xceM\xacR\xa2\xebZ3Yqcp\x8f֜X\xf3\xe3U\b\xd5\a\x94\xaa\xb6\ue532Y\x80ʘQ\x16P\x1asi\xd8R\xa7O\x9dy\xaf#%\x11\x05a)\xd3\x00\x9am\x1b\xf4QK\u00ad\xe3P\xab&e\x01QR\x9e\x9d\xea)\u05f8r\x9f]\xcbG\xad\x10\x94\x00\x00v\xd5V\xe8}\x12\xb7\xe9.\xebh\xedJ\xb7\"=\x02N\xcd\abTLF\xc0\r\x0e[\xad\xf7\x81V\xde*ȱ\xff\x912\xfb\x8d\x93\x01\xcc\xca\xf4\xe5#\xd9\a\x1c\x03Wa\xec\xf3\x03ر\x03ْT}\xdepo\x15\x9a\x8b\xf4\xe09\xf0\xf9\xf0u\xa7\x9f\r\x1f\xa6\xc0b\x0e\x87\xcd\xe5\x05\xc0x\xa3\xfd\xb1|\x90\xc2:mg\xdc\xe9\xb2b\x83P\xeb{#U\xa2\xc93\xb2+\x12\x19\xf4=\xb6\xfeƸA\xa1\x0e\xb5ؽ\x13/%:^+CH\xfb\xc4EAcS\b\x02DbN\xd8\xdbO2\a\x9aWz\xf5\xed\x8a{\xe7\tu\x1f\x1d\x9cߔ\xc9O\xfeA>\x1a\x93\xe7h\xae\xb3\xe4K\xe5\xf3[\x05\xa8nr\xa47\xbd\x12'`\x1a\x9b\x11\xae\x03\x1aP\xa6N\xe4T\x8eqB\xd3\x18\x84*bA\xa58\xc4P?*\xa2\\\xa0G4\xde9\x8a<I诬\xd4)\xeb\xf1\x93y\xc0YM\x90\x17\x9d/d\xa0\xbaEm&\x0f:\xba\x8e4\xd5e\tZ\bh\x81\x92\xb1\xa2Q\x8a\xcd5=Y5\xbac\xa7V\x98H\xa1G\b\b\xbd\xab\xdc\x04\x82V9p\xd2\xc1\x18Op\xf4d\xed\t\x9b\xb8P\xf8~\x9f\xf9K\xc9$\xaes\x00{\xef:\xe3\xa9h\x9e\x9d\xf5\xb6\xd4j%\xd3y\xad⸒\xaf\x99Ο\x9a[\x10\r\xf5\xccE;\xa7\xbc~h\xba\xab\xadx\xde\x04QB\xdfC\xdc\x1e\xfdl\xe1i\"3M\x82B\xa9Ix\x8d~\x9f\x85|\xe5\x9c,\xdb?\x84\x87\x11zI\xf80g\xee\r\xe1\xab\xc0\xb3a\xbc\x13\xbf\x0eh\xbf\a\xd5F\\\xe0\xd0U\xc5MjH(\x19B\xce\xcc \xa7\x89\xc4\xcbg7\x03U\x12\x90k:\x1dM\xd3<\x03i|\x95\x1b\xe7\xb2\xf2\xd5\"Ss\xecL\xb6\xd1\xfd\xf4\x12\xe3V\xee\x1d\xaf\x80\x89\xed?\xa0\xa7\x06\xde\u05cf\x1c1\xe0X\x17CW\xf0\xe7K\xdb\xf4\x9c\xc7@\xd2\xf3\x10\xea\x03\x18O\x9d\xf2\x96;t\x1a\xd3\xea̹\xbc\xae&I\x8a\xc2eJm\x0f\x8cSH\x81\xdd\uf09dĘ\x85\t\x12\xa1\xca\xe0]\x02aF\x13k\x12\xdfs1\tM\xd21w\xb0\xedH\x01\x0e\"\xc1\xfb\xd8\xc6\xed3S[\xc8\xcf\x03|\xb2\x92\xc2V<`\x19\xcc\xe0\x02m\xf1\xa1\x1f\f\xa4\xfd\x013\xc3L\x12\xcb\x1bm/\xb9\xecu\xe2o\xebtJ\x83\x04\xd7v:9\vO\xe1\xa0\xd1\x18E\x1e\xf6\xf2\xbeG\xb5\xe3\xb5´\xb5C[\u007f\xbb\U000dc025\x99\x1a\xe8\xef|\xed}{\x86\x96\xb2i\xea%\x81\xb7\xb9?\x03\xfc\xec\xc0\xe4^x{o\xf4\xab\xc0\xb3\xe1r\xef:\x84\xe6\xd7/\x83\xe4\xb6\x13~\x98>\x86\xafƹ\x18t\x18\xbbB\x9dEU\xa5\xcaS]\xf6vQ\\{\x1a\u007f\xb8\"\xab\xe3Bo\x11R\x06\xd0Iɣg6\xc5m\xd6.g\xe2,\xa08\x1a\xfd\x17\x19\x14$L\xc2\xc8z\xeaG\xea\xb1\a\xa8\xab\x90r\x13&A\xa0\xce$ې\xae\a\x8f\xa0\x03\xc3QH\xa9\xea%P\xb4!e\x8fQ;\x02\xda\xf5\xc2͵\x03̷\f\xa6\xd7\xcd\x1b\x80ɻ\xb7\xa4y\xac\xdf\fh\x14i\xc3=0;\xa4:\x9d\xcc;\x06\xdc>\x00(47[l\xa0\r&q\xa6\xf3x\xb9\x9aj\xd0]\xf4H\xbb\xbft)\x04\xf4\x04\xa0\xed\x80\xe40©x@N\x9b\x89qMrCS\xa7\xec\x8dOkK\xbe-\xd7c\xe5\xe2O\x19\xd6W\xa3\xd4\xe7\xc9\xe3\x96g\x9d\x8d1\x95\xcbV\x9b\xa8c^\x93\x9d\xf8$U\x9aWd\x8d\xceD\xe8\xe7H\u007f\x06|\x807\xfa\xe6\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0U\xe0\xd9\xf0\xbf(\xc0I\xed;\x00Cv\x01͙\xcaGq49xԵ\xea\xa7f*\xfb\xddK9\xb4\xedr;\xfd~\xa6dA\x01\xd43\xba6\xe6ĵ\x88`\x9a\x9b\xd7\x03\xc3\x1e@\xc3'(\x00^\xb1(\v8l\xb4(ir\aѿ\xbf7)`e\x8c\xbc\xc3,%\x937\x06}\x8a\x17#\xa6\x9c\xb2N\xe8\xf0\xb3b\x99\xf3\x8f\n\x18\x9bk\xeec2,\x88\x98Y\xcb\x16\xc4L\x944Ģ\x88\x88\xd1\xc0L\x17\xe3\x9d\r\xdd\xed\xde\xe0\xfaCH\xee\xf2\xab\ahw4P]\xa6\xf4I\\N\xc1\xf6\xdcv\xa9HҤ\xfc*\x038\x1f깲ڡYc\xab\xef\U000d42d4\x18Rܮ\xbb3\xcb\xc5\x13\xe0\x00\xa6\x1eC\r\xa4\r<WH,P\xbe\xb5\xa3@\x9b\xa3\x9bq\xf2\x8e\xac\x0f\x109Z\xa2\x93\xee\xfd\xe4\xe1\xf2\xa0Om<\xfd\x12S\x1c>\xe2\xc1\xb70z]?\xe1\xa7/\xb3\x94\xa2-\x8e\xe0מ\xd8\xc5\x10f\x84\xfd8Z'\x02f\xbe\fxr\xdd\x1c\xd0\x17\x8f\xda\xceZ@N|R;\x1b\xa6\xe9P j\xda\r\xbf\xe5\x8a\xcc\b\xe1\xfd\x80g\x04\xea\xc4h}\x1f·^\x16g\x8d\x12\x01~\xe1\xbcLF\xf1\xf1\\~\x11\x12a\xdfJ\xfe\a\xc0\xff⍾.|\x82\x02z\x8d\xe7+P\xe8\xf6?\r\x11\xdf\xf8z\xa1\x1e\x00\xfc^;\xe0\xa59\f\xa9a\xb2c\xeb\x82\xcf:\x80\xf2h\x93\xfb_\xda/C\x98Գ\x17\x97\xb7\xc5L\xd5)\x97\xfdM\x12@\xb3\xb9\xa3\xa9\xa6p\x0e}\x85\xef\xfe8\x8c>'\x1eJ\xe7:\xf6\x9b\x17\xa1\x04\xcc\x0e\xec\xf4c8\xbc\xc3\xe78\br7\xa48:\xf6\x18\x8e/A=\xaf\x0f<\xe5\x1eVR\xaac\xf3\xe6\b}\xd7`!֖O\x18\xe0\xd2<\xbbEy\xbf\x1b\x11or\x0eaǧ5ץo\xf7\xb4\xf7\x8e\x9c\x91u[\xc7Ce\xdctC\x83}\v \x1f#\xb9\x0e\xa1\x1as\xcdK+Ѡ\\\xc1R\x83vNeEŹ\xc9\xc7dVI\xa9\x99\xfc\x1a4\xc9N\xb2\xd8=\xb5X\xfe0\x94\x1c\xdfK\xcb\x16\xf6ڒ#\xec\"\"\xda\xf3H\xb6d\x1c\\\xe2t\x99Ξh\xf9\vOw\x9f\xc7yi^\xccS\u007f1q\xee\x85O\xf0F\xdf\x1b>A\x81_\x8e\a\x9as|&\x1e\x98qю\xb7EjE\x16\x0fd\xb4w\xe3\x01h\xa9\xdb\xedx`\xc1\x05!*0\x85\x9f&\x83«\x8b\x19\xee\x83\xf1\x1e\x05\xdb\xfcp\xb3\x81\xfe\xfeX;B\xe9\x1f\x8f\xcaʚ\xf9\x8b%\x80\xb9Q0\xf6g\xa9I\xe7\xef\x15\xf8\x12L\x00\xd45\x9a\x95\x12V\x8e\xcaqK\x80ɍ\x02]r\xca\x16\x97\x98\xaa<6\xa4W3\x99Ƙs\xd8\xf0\xba$\xfc<3Əԅt\xa5\xf1\xd3\xef\x8d\a6\xe0\x1eg\xfd\xe4%\xa0i\x02\xe01\xc0\xa4RK\xfbk\x02<\x94r\xff\xbe\x97\x9b\x93x\x9dqXLb\x11\x0e\xb0\xfd\x815_\x02\x9c\xbd\xb3n:\xe7dg\xde\x03s\x1ez\xd7\xc3%\xf69\xda\u007f\x03\xa8\xff\xbc\xdeG\xdcP\xfe\x17\x00\x00\xff\xff!o0\xc6\x16:4\xa9\x00\x00\x00\x00IEND\xaeB`\x82")

func text_png_bytes() ([]byte, error) {
	return _text_png, nil
}

func text_png() (*asset, error) {
	bytes, err := text_png_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "text.png", size: 2058, mode: os.FileMode(420), modTime: time.Unix(1420820510, 0)}
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
	"text.png": text_png,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"text.png": &_bintree_t{text_png, map[string]*_bintree_t{}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
