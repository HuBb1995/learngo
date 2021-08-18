package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

var listRe = regexp.MustCompile(`{"uuid":([0-9]+),"cN":"([^"]+)","uT":"([^"]+)","cnt":([0-9]+),"cU":"([^"]+)","id":([0-9]+),"sS":([0-9]+)}`)

var contentRe = regexp.MustCompile(`<div class="read-content j_readContent" id="">([\s\S]+?)</div>`)
var pRe = regexp.MustCompile(`p>([^<]+)<`)

//目录
//https://book.qidian.com/ajax/book/category?_csrfToken=NgHXj4yaL4qeJ3HKcFI1yz2jLxv4Ljm37egnvH19&bookId=1004608738

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://book.qidian.com/ajax/book/category?_csrfToken=NgHXj4yaL4qeJ3HKcFI1yz2jLxv4Ljm37egnvH19&bookId=1010868264", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("cookie", "_csrfToken=NgHXj4yaL4qeJ3HKcFI1yz2jLxv4Ljm37egnvH19; newstatisticUUID=1623512778_788461596; qdrs=0%7C3%7C0%7C0%7C2; showSectionCommentGuide=1; qdgd=1; _yep_uuid=514e48f7-c813-cbf6-c561-a4691ee7a3a4; _gid=GA1.2.384447619.1629073851; e1=%7B%22pid%22%3A%22qd_p_qidian%22%2C%22eid%22%3A%22qd_A17%22%2C%22l1%22%3A3%7D; e2=%7B%22pid%22%3A%22qd_p_qidian%22%2C%22eid%22%3A%22qd_A06%22%2C%22l1%22%3A1%7D; rcr=1004608738%2C1019664125%2C1024617405%2C1021617576%2C2643379%2C1016150754; bc=1004608738; lrbc=1004608738%7C350158384%7C1%2C1019664125%7C535760902%7C1; _gat_gtag_UA_199934072_2=1; ywguid=800217964283; ywkey=ywo5SOTR39Xa; ywopenid=CDE3E788A117B5CFFA28127192B5CD81; _ga_FZMMH98S83=GS1.1.1629073849.3.1.1629079441.0; _ga_PFYW0QLV3P=GS1.1.1629073849.9.1.1629079441.0; _ga=GA1.2.944493772.1625113226")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("qidiancrawler/诡秘之主.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body, err := io.ReadAll(resp.Body)
	matches := listRe.FindAllSubmatch(body, -1)

	bodyChan := make(chan []byte)

	for _, m := range matches {
		log.Printf("uuid: %s, cN: %s, uT: %s, cnt: %s, cU: %s, id: %s, sS: %s\n", m[1], m[2], m[3], m[4], m[5], m[6], m[7])
		go fetch("https://read.qidian.com/chapter/1010868264/"+string(m[6]), bodyChan)
		body = <-bodyChan
		contentMatch := contentRe.FindSubmatch(body)
		pMatches := pRe.FindAllSubmatch(contentMatch[1], -1)
		file.Write(m[2])
		file.WriteString("\n")
		for _, p := range pMatches {
			fmt.Printf("%s\n", p[1])
			file.Write(p[1])
			file.WriteString("\n")
		}
	}
}

func fetch(url string, bodyChan chan []byte) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("cookie", "__yep_uuid=7e30a5ea-579f-8cc6-de60-17718e1ea6c3; _csrfToken=NgHXj4yaL4qeJ3HKcFI1yz2jLxv4Ljm37egnvH19; newstatisticUUID=1623512778_788461596; qdrs=0%7C3%7C0%7C0%7C2; showSectionCommentGuide=1; qdgd=1; _gid=GA1.2.384447619.1629073851; ywguid=800217964283; ywkey=ywo5SOTR39Xa; ywopenid=CDE3E788A117B5CFFA28127192B5CD81; e1=%7B%22pid%22%3A%22qd_P_Searchresult%22%2C%22eid%22%3A%22qd_S05%22%2C%22l1%22%3A3%7D; e2=%7B%22pid%22%3A%22qd_p_qidian%22%2C%22eid%22%3A%22qd_H_Search%22%2C%22l1%22%3A2%7D; rcr=1010868264%2C1004608738%2C1019664125%2C1024617405%2C1021617576%2C2643379%2C1016150754; bc=1010868264; lrbc=1010868264%7C402760766%7C0; pageOps=1; _ga_FZMMH98S83=GS1.1.1629109829.4.1.1629110587.0; _ga_PFYW0QLV3P=GS1.1.1629109829.10.1.1629110587.0; _ga=GA1.2.944493772.1625113226; _gat_gtag_UA_199934072_2=1")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	bodyChan <- body
}
