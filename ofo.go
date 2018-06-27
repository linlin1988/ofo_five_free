package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ofoUrl = "https://activity.api.ofo.com/ofo/Api/aff"
)

func ofo(inviteKey *string, phone int) {

	var r http.Request
	r.ParseForm()

	//电话转成string
	stringPhone := strconv.Itoa(phone)
	//添加post参数
	r.Form.Add("inviteKey", *inviteKey)
	r.Form.Add("tel", stringPhone)
	bodystr := strings.TrimSpace(r.Form.Encode())

	//fmt.Println("req.body:",bodystr)
	//发起一个post请求
	request, _ := http.NewRequest("POST", ofoUrl, strings.NewReader(bodystr))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	//读取返回
	respContent, _ := ioutil.ReadAll(resp.Body)
	ret := string(respContent)

	if ok := strings.Contains(ret, "50004"); ok {
		fmt.Println("操作太频繁了,等等再搞...")
		fmt.Println("resp.body:", ret)
		time.Sleep(time.Second * 3600)
	}

	if ok := strings.Contains(ret, "isNewuser\":1"); ok {
		fmt.Printf("[%s]邀请的号码成功\n", stringPhone)
	}

	//

}

func main() {

	//定义2个命令行参数
	inviteKey := flag.String("k", "xxxx", "inviteKey")
	tel := flag.Int("t", 13510630203, "mobile phone number")
	flag.Parse()

	counter := *tel
	for {
		ofo(inviteKey, counter)
		counter++
		time.Sleep(time.Second * 8)
	}

}
