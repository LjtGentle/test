package logic

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
	"time"
)

const (
	CDNPath = "/osgame/data/%s" // 大玩法/小玩法
)

// 更新数据档
func UpdateDataBin(f multipart.File, size int64, language string) ([]string, error) {
	// 记录启动时间
	now := time.Now()

	// 执行数据解析录入
	links, err := LoadDataBin(f, size, language)
	if err != nil {
		log.Printf("加载数据档失败,err=%v", err)
		return nil, err
	}
	t := time.Now().Sub(now)
	fmt.Println("耗时:", t)
	//fmt.Printf("links=%+v\n", links)
	return links, nil
}

// 加载数据档
func LoadDataBin(f multipart.File, size int64, language string) ([]string, error) {
	zr, err := zip.NewReader(f, size)
	if err != nil {
		return nil, err
	}
	outlinks := make([]string, 0)
	osGameTypes := [][]string{{"Standard_S1", "Basic"}}
	basic := [][]string{{"Basic", "Public"}}
	// 处理玩法的
	for fileName, fileModel := range DataBinBytesFileStruct {
		convert := &FileConvert{
			SFileName: fileName,
			TFileName: strings.Split(fileName, ".")[0] + ".json",
			ModelPtr:  fileModel,
		}
		links, err := convert.LoadResDataBinPutCDN(language, zr, osGameTypes)
		if err != nil {
			log.Printf("读取comm并上传%s数据失败，中止加载，err=%v\n", convert.SFileName, err)
			return nil, err
		}
		outlinks = append(outlinks, links...)
	}
	// 处理基本的
	for fileName, fileModel := range DataBinBytesFileStructBasic {
		convert := &FileConvert{
			SFileName: fileName,
			TFileName: strings.Split(fileName, ".")[0] + ".json",
			ModelPtr:  fileModel,
		}
		links, err := convert.LoadResDataBinPutCDN(language, zr, basic)
		if err != nil {
			log.Printf("读取Basic并上传%s数据失败，中止加载，err=%v\n", convert.SFileName, err)
			return nil, err
		}
		outlinks = append(outlinks, links...)
	}
	return outlinks, nil
}

// FileConvert 文件类型转换
type FileConvert struct {
	SFileName string
	TFileName string
	ModelPtr  interface{}
}

// LoadResDataBinPutCDN 读取.bytes 文件 转换.json文件 上传cdn
func (convert *FileConvert) LoadResDataBinPutCDN(language string, zr *zip.Reader, OsGameTypes [][]string) ([]string, error) {
	fmt.Printf("OsGameTypes=%+v\n", OsGameTypes)
	links := make([]string, 0)
	for _, gt := range OsGameTypes {
		// 读取基本信息
		rc, err := GetDeepestFileContent(zr, convert.SFileName, gt...)
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		decoder, err := NewDatabinDecoder(rc)
		if err != nil {
			return nil, err
		}
		err = decoder.Decode(convert.ModelPtr)
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(convert.ModelPtr)
		if err != nil {
			fmt.Println("in LoadResDataBinPutCDN call fist json Marshal err=", err)
			return nil, err
		}
		dataBinStr := string(b)
		fmt.Printf("dataBinStr=%s\n", dataBinStr)
		// 读取对应的lang文件
		langFileName := strings.Split(convert.SFileName, ".")[0]
		langFileName += "." + language + ".lang"
		fmt.Printf("LoadResDataBinPutCDN lang file %s\n", langFileName)
		langRc, err := GetDeepestFileContent(zr, langFileName, gt...)
		//_, err = GetDeepestFileContent(zr, langFileName, gt...)
		if err != nil {
			return nil, err
		}
		if langRc != nil {
			// 有翻译文件存在
			defer langRc.Close()
			lang, err := I18NDecode(langRc)
			if err != nil {
				return nil, err
			}
			fmt.Printf("LoadResDataBinPutCDN 准备替换字符串\n")
			for index, v := range lang.Records {
				fmt.Printf("index=%d\tv=%+v\n", index, v)
				//if index > 10 {
				//	break
				//}
				// 去换行
				//v.Translation = strings.ReplaceAll(v.Translation, "\"", "\\\"")
				v.Translation = strings.ReplaceAll(v.Translation, "\"", "'")
				//v.Translation = strings.ReplaceAll(v.Translation, "\\", "\\\\")
				v.Translation = strings.ReplaceAll(v.Translation, "\n", "\\n")
				dataBinStr = strings.ReplaceAll(dataBinStr, v.Key, v.Translation)
			}
			//dataBinStr = strings.ReplaceAll(dataBinStr, "\n", "\\n")
		}
		//logger.Info(c, "dataBinStr=%s", dataBinStr)
		b = []byte(dataBinStr)
		err = json.Unmarshal(b, convert.ModelPtr)
		if err != nil {
			fmt.Printf("json Unmarshal err=%+v\n", err)
			return nil, err
		}
		var data = map[string]interface{}{
			"dataTable": convert.ModelPtr,
		}
		dataBin, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("json marshal err=%+v\n", err)
			return nil, err
		}
		link := string(dataBin)
		links = append(links, link)
	}

	return links, nil
}
