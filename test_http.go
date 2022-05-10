package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func requestUrl(url string) {
	result, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Body.Close()
	all, _ := ioutil.ReadAll(result.Body)
	fmt.Printf("all结果:%v\n", string(all))
}

func parseUrl() {

	params := url.Values{}
	parse, err := url.Parse("https://apis.juhe.cn/simpleWeather/query")
	if err != nil {
		log.Panic("错误处理")
	}
	params.Set("key", "087d7d10f700d20e27bb753cd806e40b")
	params.Set("city", "重庆")
	parse.RawQuery = params.Encode()
	urlPath := parse.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("读取错误")
	}
	//var weather struct{}

	//json.Unmarshal(all, &weather)
	//fmt.Printf("读取结果:%v\n", weather)
	fmt.Printf("读取结果:%v\n", string(all))

}

func testParamsToJson() {
	type result struct {
		Args    map[string]string `json:"args"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
	}
	resp, err := http.Get("https://httpbin.org/get?testParams=test")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("结果:%v\n", string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("json结果:%v\n", res)
}

func testAddHeader() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("test-user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	request.Header.Add("test-accept-encoding", "gzip, deflate, br")
	resp, err1 := client.Do(request)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(resp)
	readBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("result:%v\n", string(readBody))
}

func testHttpPost() {
	params := url.Values{}
	//pares, err := url.Parse("https://apis.juhe.cn/simpleWeather/query")
	//if err != nil {
	//	log.Fatal(err)
	//}
	params.Set("key", "087d7d10f700d20e27bb753cd806e40b")
	params.Set("city", "重庆")
	//pares.RawQuery = params.Encode()
	//urlPath := pares.String()
	form, err := http.PostForm("https://apis.juhe.cn/simpleWeather/query", params)
	if err != nil {
		log.Fatal(err)
	}
	all, err1 := ioutil.ReadAll(form.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer form.Body.Close()
	fmt.Printf("result:%v\v", string(all))

}

func testHttpPost2() {
	params := url.Values{
		"name": {"banana"},
		"sex":  {"1"},
		"test": {"sbx"},
		"tags": {"隆小镇"},
	}
	requestBody := params.Encode()
	response, _ := http.Post("https://httpbin.org/post", "text/html; charset=UTF-8", strings.NewReader(requestBody))
	//form, _ := http.PostForm("https://httpbin.org/post", params)/
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result:%v\n", string(all))

}

func testHttpJson() {
	params := make(map[string]interface{})
	params["name"] = "隆小镇"
	params["desc"] = []string{"名称", "性别"}
	params["age"] = 20
	marshal, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}
	post, err1 := http.Post("https://httpbin.org/post", "application/json", bytes.NewReader(marshal))
	if err1 != nil {
		log.Fatal(err1)
	}
	all, _ := ioutil.ReadAll(post.Body)
	fmt.Printf("result:%v\n", string(all))
}

func testHttpServer() {
	f := func(resp http.ResponseWriter, r *http.Request) {
		io.WriteString(resp, "test server")
	}
	http.HandleFunc("/test", f)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type countHandler struct {
	mutex sync.Mutex
	count int
}

func (ch *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()
	ch.count++
	_, err := fmt.Fprintf(w, "count num is :%v\n", ch.count)
	if err != nil {
		log.Fatal(err)
	}
}

func testHttpServer2() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//func getSum(n int) (sum int) {
//	if n < 1 {
//		return 0
//	}
//	for i := 1; i <= n; i++ {
//		sum += i
//	}
//	return sum
//}
//
//func GetSumRecursive(n int) (sum int) {
//	if n <= 1 {
//		return n
//	}
//	return n + GetSumRecursive(n-1)
//}
//
//func FibonacciNumbersRecursive(n int) (sum int) {
//	if n == 1 || n == 0 {
//		return 1
//	}
//	return FibonacciNumbersRecursive(n-1) + FibonacciNumbersRecursive(n-2)
//}

func main() {
	//requestUrl("https://www.baidu.com")
	//parseUrl()
	//testParamsToJson()
	//testAddHeader()
	//testHttpPost()
	//testHttpPost2()
	//testHttpJson()
	//testHttpServer()
	//testHttpServer2()
	//sum := getSum(10)
	//recursive := GetSumRecursive(10)
	recursive := FibonacciNumbersRecursive(10)
	fmt.Printf("结果%v\n", recursive)
}
