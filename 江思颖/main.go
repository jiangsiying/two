//https://blog.lenconda.top/
//https://blog.lenconda.top/page/2/

//标题：<h2 class="post-title"><a href="/posts/
//时间：<time datetime="  "
//标签：<a class="post-meta-tag" href="  "
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误传出给调用者
		return
	}
	defer resp.Body.Close() //结束时关闭

	//循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			//读取结束或者出问题
			fmt.Println("读取网页完成")
			break
		}

		//表示出错
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环读到的buf数据，存入result一次性返回。
		result += string(buf[:n])
	}
	return
}

func working(start, end int) {
	fmt.Printf("正在爬取第%d页到第%d页\n", start, end)
	//爬取每一页的数据
	for i := start; i <= end; i++ {
		url := "https://blog.lenconda.top/page/" + strconv.Itoa(i) + "/"
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err:", err)
			continue
		}
		// fmt.Println("result=", result)读取一个网页
		//将读到的整网页数据，保存成一个文件
		f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
		if err != nil {
			fmt.Println("Creat err:", err)
			continue
		}

		f.WriteString(result) //写内容
		f.Close()             //保存好一个文件，关闭一个文件
	}

	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(doc.Find("title").Text())

	// doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(strings.TrimSpace(s.Find("h3").Text()))
	// })

	//还未写完
	// //(?s:(.*?))
	// ret1 := regexp.MustCompile(`<h2 class="post-title"><a href="/posts/(?s:(.*?))`)
	// //提取需要的信息
	// title := ret.FindAllStringSubmatch(result, -1)
	// for _, name := range title {
	// 	fmt.Println("name", name[1])
	// }

	// ret2 := regexp.MustCompile(`<time datetime="(?s:(.*?))"`)
	// //提取需要的信息
	// time := ret.FindAllStringSubmatch(result, -1)
	// for _, name := range time {
	// 	fmt.Println("name", name[1])
	// }
	// ret3 := regexp.MustCompile(`<a class="post-meta-tag" href="(?s:(.*?))"`)
	// //提取需要的信息
	// label := ret.FindAllStringSubmatch(result, -1)
	// for _, name := range label {
	// 	fmt.Println("name", name[1])
	// }
}
func main() {
	var start, end int
	fmt.Printf("请输入起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=start:）")
	fmt.Scan(&end)

	working(start, end)
}
