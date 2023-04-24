package logic

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
	osgameError "osgame/model/error"
	model "osgame/model/osserver/proto/resdata"
	"path/filepath"
	"reflect"
	"strings"
)

// parse 解析osgame数据档
func parse(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

// ParseFile 解析数据档文件
func ParseFile(file *os.File, v interface{}) error {
	return parse(file, v)
}

// GetDeepestFileContent
func GetDeepestFileContent(rc *zip.Reader, filename string, categories ...string) (io.ReadCloser, error) {
	fileList := GetZipFileList(rc, filename)
	// 贪婪匹配
	for n := len(categories); n > 0; n-- {
		prepath := filepath.Join(categories[:n]...)
		for _, fd := range fileList {
			if strings.Contains(fd.FileHeader.Name, prepath) {
				return fd.Open()
			}
		}
	}
	return nil, nil
}

// GetZipFileList
func GetZipFileList(rc *zip.Reader, filename string) []*zip.File {
	var list []*zip.File
	for _, f := range rc.File {
		if f.FileInfo().Name() == filename {
			list = append(list, f)
		}
	}
	return list
}

// GetZipFileContent
func GetZipFileContent(rc *zip.Reader, filename string) (io.ReadCloser, error) {
	for _, f := range rc.File {
		log.Println(f.FileInfo().Name())
		if f.FileInfo().Name() == filename {
			return f.Open()
		}
	}
	return nil, fmt.Errorf(osgameError.NotFindFile)
}

// 获取zip中的目录
func GetZipFileInDirectory(rc *zip.Reader, filename string) (map[os.FileInfo]io.ReadCloser, error) {
	var dirname string
	// 找到文件夹名字
	for _, f := range rc.File {
		fi := f.FileInfo()
		if fi.Name() == filename && fi.IsDir() {
			dirname = f.Name[:len(f.Name)-1]
			break
		}
	}
	if dirname == "" {
		return nil, fmt.Errorf("%s not found", filename)
	}
	rcs := make(map[os.FileInfo]io.ReadCloser)
	for _, f := range rc.File {
		fi := f.FileInfo()
		d := filepath.Dir(f.Name)
		if d == dirname {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			rcs[fi] = rc
		}
	}
	return rcs, nil
}
func GetZipAllFile(rc *zip.Reader) (map[string]io.ReadCloser, error) {
	list := make(map[string]io.ReadCloser, 0)
	for _, f := range rc.File {
		if f.FileInfo().IsDir() {
			continue
		}
		//fmt.Println("in GetZipAllFile f.Name=", f.Name)
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		list[f.Name] = rc
	}
	return list, nil
}

type databindecoder model.DataBinBytes

func (d *databindecoder) Decode(ptr interface{}) error {
	ptrt := reflect.TypeOf(ptr)
	ptrv := reflect.ValueOf(ptr)
	if ptrt.Kind() != reflect.Ptr {
		return fmt.Errorf("ptr should be pointer")
	}
	t := ptrt.Elem().Elem()
	isElPtr := true
	if t.Kind() != reflect.Ptr {
		isElPtr = false
	}
	if !t.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		return fmt.Errorf("切片元素未实现proto.Message")
	}
	ss := reflect.New(ptrt.Elem()).Elem()
	for _, item := range d.Records {
		var ev reflect.Value
		// 确保ev为元素指针
		if !isElPtr {
			ev = reflect.New(t)
		} else {
			ev = reflect.New(t.Elem())
		}
		el := ev.Interface()
		nes, ok := el.(proto.Message)
		if !ok {
			return fmt.Errorf("切片元素未实现proto.Message")
		}

		err := proto.Unmarshal(item, nes)
		if err != nil {
			return err
		}
		if isElPtr {
			ss = reflect.Append(ss, ev)
		} else {
			ss = reflect.Append(ss, ev.Elem())
		}
	}
	ptrv.Elem().Set(ss)
	return nil
}

func NewDatabinDecoder(rc io.Reader) (*databindecoder, error) {
	d := &model.DataBinBytes{}
	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("读取失败，err:%s", err.Error())
	}
	err = proto.Unmarshal(buf, d)
	if err != nil {
		return nil, fmt.Errorf("解析数据档失败, err:%s", err.Error())
	}
	return (*databindecoder)(d), nil
}

type i18NDecoder model.I18N

func I18NDecode(rc io.Reader) (*i18NDecoder, error) {
	d := &model.I18N{}
	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("读取lang文件失败，err:%s", err.Error())
	}
	err = proto.Unmarshal(buf, d)
	if err != nil {
		return nil, fmt.Errorf("解析langw文件失败, err:%s", err.Error())
	}
	return (*i18NDecoder)(d), nil
}
