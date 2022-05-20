package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Req struct {
	Name    string    `json:"name" form:"name" binding:"required" `
	Address time.Time `json:"address" form:"address"`
}

const (
	TimeLayout         = "2006/01/02/15:04:05"
	TimeId             = "20060102150405"
	pictureChanMaxsize = 2
)

func main() {
	r := gin.Default()
	r.GET("/time_test", func(c *gin.Context) {
		var req Req
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("参数错误,err=%+v", err),
			})
			return
		}
		c.JSON(200, gin.H{
			"data": req,
		})
	})

	// 测试下载文件
	r.GET("/upload_some_pic", func(c *gin.Context) {
		var err error
		var wg sync.WaitGroup
		//url := "https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/undefined/ingame/images/202009/20200917160412-968279.jpg,"+
		//	"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/undefined/ingame/images/202009/20200917154019-160772.jpg," +
		//	"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029114835-384084.jpg," +
		//	"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029114843-776638.png"
		url := "https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150102-177900.jpg," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150108-534646.jpg," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150115-246488.jpg," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150120-187433.jpg," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150126-920047.jpg," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150130-667574.png," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150140-512908.png," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150153-616065.png," +
			"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202010/20201029150159-327623.jpg"
		urls := strings.Split(url, ",")

		type File struct {
			Content []byte
			Name    string
		}

		fileChan := make(chan *File, pictureChanMaxsize)
		//isDone := make(chan struct{}) //是否完成下载
		var sum = 0
		for index, v := range urls {
			wg.Add(1)
			sum++
			go func(v string, index int) {

				defer wg.Done()
				res, err := http.Get(v)
				if err != nil {
					fmt.Printf("err11=%+v\n", err)
					c.JSON(400, gin.H{
						"err": err,
					})
					return
				}

				defer res.Body.Close()
				buff, _ := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Printf("err22=%+v\n", err)
					c.JSON(400, gin.H{
						"err": err,
					})
					return
				}
				f := &File{
					Content: buff,
				}

				pos := strings.LastIndex(v, ".")
				f.Name = fmt.Sprintf("%d%s", index, v[pos:])
				//c.Header("Content-Disposition", "attachment; filename="+f.Name)
				//c.Header("Content-Transfer-Encoding", "chunked")
				//c.Header("Content-Type", "application/octet-stream")
				//c.Header("Cache-Control", "no-cache")
				//
				//w := bufio.NewWriter(c.Writer)
				//_, err = w.Write(f.Content)
				//if err != nil {
				//	fmt_test.Println("writer err=",err)
				//	return
				//}
				//w.Flush()
				fileChan <- f
				//fmt_test.Println("下载图片成功 b=",string(buff))
				fmt.Println("下载成功")
			}(v, index)

		}

		//压缩打包
		t := time.Now().Format(TimeId)
		fmt.Println("t=", t)
		zipFileName := fmt.Sprintf("smoba_%v.zip", t)
		buf := bytes.NewBuffer(nil)

		zipwriter := zip.NewWriter(buf)
		defer  func(err error) {
			zipwriter.Close()
			if err != nil {
				fmt.Println("处理图片失败...")
				return
			}
			bs := buf.Bytes()
			_, err = c.Writer.Write(bs)
			if err != nil {
				fmt.Println("写入失败err=", err)
				return
			}

			fmt.Println("下载完毕")
		}(err)
		wg.Add(1)
		go func() {
			defer wg.Done()
			over := 0
			for v := range fileChan {
				//buf.Reset()
				over++
				iowriter, err := zipwriter.Create(v.Name)
				if err != nil {
					fmt.Println("err false=", err)

				}
				fmt.Println("len=", len(v.Content))
				n, err := iowriter.Write(v.Content)
				fmt.Println("n=", n)
				if err != nil {
					fmt.Println("err=", err)
					return
				}
				//zipwriter.Flush()
				fmt.Println("压缩文件 name=",v.Name)
				if over == sum {
					break
				}
			}
		}()
		//for  {
		//	flag := false
		//	select {
		//	case p := <-fileChan:
		//		fmt_test.Println("收到工作的信号")
		//		wg.Add(1)
		//		go func(v *File) {
		//			defer wg.Done()
		//			mw.Lock()
		//			defer mw.Unlock()
		//			iowriter, err := zipwriter.Create(v.Name)
		//			if err != nil {
		//				fmt_test.Println("err false=", err)
		//
		//			}
		//			fmt_test.Println("len=",len(v.Content))
		//			n, err := iowriter.Write(v.Content)
		//			fmt_test.Println("n=",n)
		//			if err != nil {
		//				fmt_test.Println("err=", err)
		//				return
		//			}
		//			fmt_test.Println("1328466461")
		//		}(p)
		//	case <-isDone:
		//		// 收到结束信号
		//		l := len(fileChan)
		//		fmt_test.Println("收到结束信号,workChanLen=",l)
		//		for l<0 {
		//			l--
		//			v := <-fileChan
		//			fmt_test.Println("收到结束信号,但继续工作")
		//			wg.Add(1)
		//			go func(v *File) {
		//				defer wg.Done()
		//				mw.Lock()
		//				defer mw.Unlock()
		//				iowriter, err := zipwriter.Create(v.Name)
		//				if err != nil {
		//					fmt_test.Println("err false=", err)
		//
		//				}
		//				fmt_test.Println("len=",len(v.Content))
		//				n, err := iowriter.Write(v.Content)
		//				fmt_test.Println("n=",n)
		//				if err != nil {
		//					fmt_test.Println("err=", err)
		//					return
		//				}
		//				fmt_test.Println("1328466461")
		//			}(v)
		//		}
		//		fmt_test.Println("结束for循环")
		//		flag = true
		//	default:
		//		fmt_test.Println("睡一会")
		//		time.Sleep(300*time.Millisecond)
		//
		//	}
		//	if flag == true {
		//		break
		//	}
		//
		//}
		fmt.Println("等待线程结束.....")
		wg.Wait()
		fmt.Println("所有线程都结束")
		c.Header("Content-Disposition", "attachment; filename="+zipFileName)
		c.Header("Content-Transfer-Encoding", "chunked")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Cache-Control", "no-cache")


	})

	r.Run(":8080")
}
