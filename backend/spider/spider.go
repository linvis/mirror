package spider

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Spider struct {
	userAgent string
	url       string
	cookies   string
}

func NewSpider(url string, cookies string) *Spider {
	return &Spider{
		userAgent: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
		url:       url,
		cookies:   cookies,
	}
}

func (spider *Spider) DoSpider() (string, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", spider.url, nil)

	req.Header.Set("User-Agent", spider.userAgent)

	req.Header.Set("Cookie", spider.cookies)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return "", err
	}
	// fmt.Println(string(body))

	return string(body), nil
}

func LoadCookies(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	input := bufio.NewScanner(f)
	var cookies string

	defer f.Close()

	for input.Scan() {
		cookies += input.Text()
	}

	return cookies, nil
}
