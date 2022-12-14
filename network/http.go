package _network

import (
	"fmt"
	"net/url"
)

func Http() {
	//resp, _ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// プログラム上でurlが正しいか判定するときなどに使用する
	// http://example.com\fasfyuのような間違ったurlを渡してもベースのhttp://example.comを取得できる
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

}
