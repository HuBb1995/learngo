package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//江苏省人民政府
//const baseUrl = "http://www.jiangsu.gov.cn"
//var listRe = regexp.MustCompile(`<li><a title="([^"]+)" target="_blank" href="(/art/[0-9]+/[0-9]+/[0-9]+/[^.]+.html)">[^<]+</a><b>([^<]+)</b></li>`)
//var listRe2 = regexp.MustCompile(`<li><a title="([^"]+)" target="_blank" href="(http://www.jiangsu.gov.cn/art/[0-9]+/[0-9]+/[0-9]+/[^.]+.html)">[^<]+</a><b>([^<]+)</b></li>`)
//var contentRe = regexp.MustCompile(`<p style="text-indent: ([0-9]em);[^"]*">([^<]+)</p>`)
//var pictureRe = regexp.MustCompile(`<img title="([^"]+)" src="([^"]+)"/>`)

//江苏省发展和改革委员会
//var listRe2 = regexp.MustCompile(`<a href='(http://fzggw.jiangsu.gov.cn/art/[^.]+.html)' target='_blank' title='([^']+)'[^>]+>[^<]+</a>`)
//var dateRe = regexp.MustCompile(`<td align='center' width='80'>([^<]+)</td>`)
//var contentRe = regexp.MustCompile(`<p style="text-indent: ([0-9]em); text-align: left;">([^<]+)</p>`)
//江苏省教育厅
//const baseUrl = "http://jyt.jiangsu.gov.cn"
//
//var listRe = regexp.MustCompile(`<a title="([^"]+)" href="(/art/[^.]+.html)"[^t]+target="_blank">[^<]+</a><span>([^<]+)</span>`)
//var contentRe = regexp.MustCompile(`<p style="text-indent: ([0-9]em);[^"]*">([^<]+)</p>`)
//var contentRe2 = regexp.MustCompile(`<p style="TEXT-ALIGN: left; TEXT-INDENT: ([0-9]em)">([^<]+)</p>`)

//江苏省科学技术厅
//const baseUrl = "http://kxjst.jiangsu.gov.cn"

//政策文件
//var listRe = regexp.MustCompile(`<a href="(/art/[^.]+.html)" target="blank" title="([^"]+)">[^<]+</a>`)
//var contentRe = regexp.MustCompile(`<p style="text-indent: ([0-9]em);[^"]*">([^<]+)</p>`)
//var dateRe = regexp.MustCompile(`<td width="180" align="center" style="vertical-align:middle;color:#999999; border-bottom:1px #CCC dashed;">\(([^)]+)\)</td>`)
//政策解读
//var listRe = regexp.MustCompile(`<a href="(/art/[^.]+.html)" target="blank" title="([^"]+)">[^<]+</a>`)
//var contentRe = regexp.MustCompile(`<p style="text-indent: ([0-9]em); text-align: left;">([^<]+)</p>`)
//var dateRe = regexp.MustCompile(`<td width="180" align="center" style="vertical-align:middle;color:#999999; border-bottom:1px #CCC dashed;">\(([^)]+)\)</td>`)

//江苏省工业和信息化厅
//const baseUrl = "http://gxt.jiangsu.gov.cn"

//政策文件
//var listRe = regexp.MustCompile(`<a href="(/art/[^.]+.html)" target="_blank" title="([^"]+)">[^<]+</a>([^<]+)</li>`)
//var contentRe = regexp.MustCompile(`<p style="background: white; text-align: justify; line-height: ([0-9]em); text-indent: 37px; margin-bottom: 0px; -ms-text-justify: inter-ideograph;"><span style="font-family: 宋体, SimSun; font-size: 16px;">([^<]+)</span></p>`)
//政策解读
//var listRe = regexp.MustCompile(`<a href="(/art/[^.]+.html)" target="_blank" title="([^"]+)">[^<]+</a>([^<]+)</li>`)
//var contentRe = regexp.MustCompile(`<p style="text-align: left; line-height: 2em; text-indent: ([0-9]em);"><span style="font-family: 微软雅黑,Microsoft YaHei;"><strong><span style="font-size: 16px;">([^<]+)</span></strong></span></p>`)

//江苏省民族宗教事务委员会
//const baseUrl = "http://mzw.jiangsu.gov.cn"

//政策文件
//var listRe = regexp.MustCompile(`<a title="([^"]+)" href="(/art/[^.]+.html)"[^t]+target="_blank">[^<]+</a><span>([^<]+)</span>`)
//var contentRe = regexp.MustCompile(`<p style="text-align: left; text-indent: ([0-9]em);">([^<]+)</p>`)
//政策解读
//var listRe = regexp.MustCompile(`<a title="([^"]+)" href="(/art/[^.]+.html)"[^t]+target="_blank">[^<]+</a><span>([^<]+)</span>`)
//var contentRe = regexp.MustCompile(`<p style="text-align: left; text-indent: ([0-9]em);">([^<]+)</p>`)

//江苏省公安厅

//const baseUrl = "http://gat.jiangsu.gov.cn"

//政策文件
//政策解读
//var listRe = regexp.MustCompile(`<a target="_blank" href="(/art/[^.]+.html)" title="([^"]+)">[^<]+</a>`)

//江苏省民政厅
//const baseUrl = "http://mzt.jiangsu.gov.cn"

//政策文件
//var listRe = regexp.MustCompile(`<a title="([^"]+)" target="_blank" href="(/art/[^.]+.html)">[^<]+</a><b>([^<]+)</b>`)

//政策图解
var pictureRe = regexp.MustCompile(`<img title="([^"]*)" src="([^"]+)"[^>]+>`)

//江苏省司法局
//const baseUrl = "http://sfj.jiangsu.gov.cn"

//政策文件
//政策解读
//var listRe2 = regexp.MustCompile(`<a style="font-size: 15px;color: #333333;" href='([^']+)' target='_blank' title='([^']+)' class='bt_link' style='line-height:32px;' >([^<]+)</a></td>
//<td style="font-size: 15px;color: #333333;" align='center' width='80'>([^<]+)</td>`)

//江苏省财政厅
//const baseUrl = "http://czt.jiangsu.gov.cn"

//政策文件
//政策解读
//var listRe = regexp.MustCompile(`<a href="(/art/[^.]+.html)" title="([^"]+)" target="_blank">[^<]+</a>`)

//江苏省人力资源和社会保障厅
//const baseUrl = "http://jshrss.jiangsu.gov.cn"

//政策文件
//政策解读
//var listRe = regexp.MustCompile(`<a target="_blank" href="(/art/[^.]+.html)" title="([^"]+)" id="maodian">[^<]+</a>`)

//江苏省自然资源厅
//const baseUrl = "http://zrzy.jiangsu.gov.cn"

//政策文件
//政策解读
//var listRe = regexp.MustCompile(`<a title="([^"]+)" href="(/gtxxgk/[^"]+)" >([^<]+)</a>`)

//江苏省生态环境厅
//const baseUrl = "http://hbt.jiangsu.gov.cn"

//政策文件
//var listRe2 = regexp.MustCompile(`<a href='(http://hbt.jiangsu.gov.cn/art/[^.]+.html)' target='_blank' title='([^']+)' class='bt_link' style='line-height:25px; padding-left:10px;' >[^<]+</a>`)
//政策解读
//var listRe = regexp.MustCompile(`<a href="(/art/[^.]+.html)" target="_blank" title="([^"]+)">[^<]+<span>([^<]+)</span></a>`)

//江苏省住房和城乡建设厅
//const baseUrl = "http://jsszfhcxjst.jiangsu.gov.cn"

//政策解读
//var listRe2 = regexp.MustCompile(`<a href='([^']+)' target='_blank' title='([^']+)' class='bt_link' style='line-height:25px;' >[^<]+</a>`)

//江苏省交通运输厅
const baseUrl = "http://jtyst.jiangsu.gov.cn"

//政策文件
//政策解读
var listRe = regexp.MustCompile(`<a target="_blank" href="([^"]+)" title="([^"]+)"[^>]*>[^<]+</a>`)

//江苏省水利厅
//const baseUrl = "http://jssslt.jiangsu.gov.cn"

//政策文件
//政策解读
//var listRe2 = regexp.MustCompile(`<a href='([^']+)' target='_blank' title='([^']+)'[^>]+>[^<]+</a>`)

func main() {
	request, _ := http.NewRequest(http.MethodGet, "http://jtyst.jiangsu.gov.cn/col/col77126/index.html", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	client := &http.Client{}
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	//body, _ = simplifiedchinese.GBK.NewDecoder().Bytes(body)

	//fmt.Printf("%s", body)
	matches := listRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		m[1] = append([]byte(baseUrl), m[1]...)
	}
	//matches2 := listRe2.FindAllSubmatch(body, -1)
	//matches := append(matches, matches2...)

	//dateMatches := dateRe.FindAllSubmatch(body, -1)

	itemCount := 0
	for _, m := range matches {
		itemCount++
		url := string(m[1])
		url = strings.ReplaceAll(url, "amp;", "")
		title := string(m[2])
		//date := string(m[3])
		fmt.Printf("#%d, Url: %s, Title: %s\n", itemCount, url, title)
		printContent(url)
		//savePicture(url, title)
	}
}

func printContent(url string) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	client := &http.Client{}
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	//body, _ = simplifiedchinese.GBK.NewDecoder().Bytes(body)
	fmt.Printf("%s", body)
	//matches := contentRe.FindAllSubmatch(body, -1)
	////matches2 := contentRe2.FindAllSubmatch(body, -1)
	////matches = append(matches, matches2...)
	//var text string
	//for _, m := range matches {
	//	text = strings.Join([]string{text, string(m[1]), string(m[2])}, "")
	//}
	//text = strings.ReplaceAll(text, "2em", "\n  ")
	//text = strings.ReplaceAll(text, "0em", "\n")
	//text = strings.ReplaceAll(text, "&nbsp;", " ")
	//
	//fmt.Println(text)
}

func savePicture(url string, title string) {
	//获取网页
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	//读取图片链接
	body, _ := io.ReadAll(resp.Body)

	match := pictureRe.FindSubmatch(body)
	//获取图片
	resp, _ = http.Get(baseUrl + string(match[2]))
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	body, _ = simplifiedchinese.GBK.NewDecoder().Bytes(body)
	//图片存储
	file, _ := os.Create("政策图解/" + title + ".jpg")
	defer file.Close()

	file.Write(body)
}
