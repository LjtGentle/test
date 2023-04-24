package logic

import (
	"archive/zip"
	"fmt"
	logger "git.woa.com/iegm-open/gin-logger"
	"git.woa.com/imkd-ingame/community_admin/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"path/filepath"
)

const mediaPath = "/osgame/mediaPath"

// 上传美术资源
func LoadMedia(c *gin.Context, f multipart.File, size int64) error {
	zr, err := zip.NewReader(f, size)
	if err != nil {
		return err
	}
	err = LoadAllIcon(c, zr)
	if err != nil {
		log.Printf("上传图片失败,err=%+v\n", err)
		return err
	}
	//// 上传棋子图标
	//err = LoadIcon(c, "IconHeroHead", zr)
	//if err != nil {
	//	log.Printf("上传棋子图标失败，err=%s", err.Error())
	//}
	//// 上传装备图标
	//err = LoadIcon(c, "IconEquip", zr)
	//if err != nil {
	//	log.Printf("上传装备图标失败，err=%s", err.Error())
	//}
	//// 上传领主图标
	//err = LoadIcon(c, "IconLordHead", zr)
	//if err != nil {
	//	log.Printf("上传领主图标失败，err=%s", err.Error())
	//}
	//// 上传技能图标
	//err = LoadIcon(c, "IconSkill", zr)
	//if err != nil {
	//	log.Printf("上传技能图标失败，err=%s", err.Error())
	//}
	//// 上传羁绊图标
	//err = LoadIcon(c, "IconFetter", zr)
	//if err != nil {
	//	log.Printf("上传羁绊图标失败，err=%s", err.Error())
	//}
	return nil
}

// 上传图标
func LoadIcon(c *gin.Context, dir string, zr *zip.Reader) error {
	fmt.Printf("dir = %s\n", dir)
	rcs, err := GetZipFileInDirectory(zr, dir)
	if err != nil {
		return err
	}
	for fi, rc := range rcs {
		ext := filepath.Ext(fi.Name())
		log.Printf("filename = %s, ext = %s", fi.Name(), ext)
		if ext != ".png" {
			continue
		}
		b, err := ioutil.ReadAll(rc)
		rc.Close()
		if err != nil {
			return err
		}
		link, hash, err := services.PutCDN(c, &b, "/osgame_static/"+dir, fi.Name())
		if err != nil {
			return err
		}
		log.Printf("link=%s, hash=%s", link, hash)
	}
	return nil
}

func LoadAllIcon(c *gin.Context, zr *zip.Reader) error {
	fileMap, err := GetZipAllFile(zr)
	if err != nil {
		return err
	}
	logger.Debug(c, "in LoadAllIcon len=%d", len(fileMap))
	for fileName, reader := range fileMap {
		b, err := ioutil.ReadAll(reader)
		reader.Close()
		if err != nil {
			log.Printf("in LoadAllIcon call ioutil.ReadAll err=%+v", err)
			continue
		}
		//fileName = strings.Join(strings.Split(fileName, "/")[1:], "/")
		logger.Debug(c, "in LoadAllIcon fileName=%s", fileName)
		link, hash, err := services.PutCDN(c, &b, mediaPath, fileName)
		if err != nil {
			log.Printf("in LoadAllIcon PutCDN err=%+v", err)
			continue
		}
		log.Printf("link=%s, hash=%s", link, hash)

	}
	return nil
}
