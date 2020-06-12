/**
 * User: Trunks(Gao)
 * Date: 2020/06/09
 * 关闭HTTP的响应
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Close = true
	//or do this:
	//req.Header.Add("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		// resp.Body.Close()的原始实现也会读取并丢弃剩余的响应主体数据
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}
