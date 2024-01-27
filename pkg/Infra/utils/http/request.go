package http

// HTTP 流程為
// 建立 TCP 連線後
// 發送 Requst 到伺服器
// 伺服器回傳 Response 給 Client
// 適合無狀態連線機制

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ToJSON convert interface{} -> json
func ToJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		// return "", err
		return ""
	}
	return string(js)
}

// Request 為 HTTP 單一請求結構
type Request struct {
	*http.Request
	// bindModel interface{}
	// OnSuccess func(resp *TaskResponse)
	// CallerFileLine string
}

// Send 發送 HTTP 請求
func (r *Request) Send() (response *Response, err error) {

	var client = &http.Client{}
	resp, err := client.Do(r.Request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			raw:    resp,
			Status: http.StatusText(http.StatusBadRequest),
		}, err
	}
	return &Response{
		raw:    resp,
		Status: resp.Status,
		Body:   string(body),
	}, nil
}

// String ...
func (r *Request) String() string {
	result := fmt.Sprintf("%s %s", r.Request.Method, r.Request.URL)

	if len(r.Request.Header) > 0 {
		result += fmt.Sprintf("\nheader:%v", r.Request.Header)
	}

	if r.Request.ContentLength <= 0 {
		return result
	}

	body, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return result
	}
	result += fmt.Sprintf("\nbody:%v", string(body))
	return result
}

// Response 為 HTTP 單一請求之後的回覆結構
type Response struct {
	raw    *http.Response
	Status string
	Body   string
}

// String ...
func (r *Response) String() string {
	return fmt.Sprintf("%s \nbody:%s", r.Status, string(r.Body))
}

// BindModel ...
func (r *Response) BindModel(obj interface{}) (err error) {
	var decoder = json.NewDecoder(strings.NewReader(r.Body))
	return decoder.Decode(obj)
}

// GroupTask 為管理 HTTP 多次請求的結構
type GroupTask struct {
	requests  []*Request
	responses []*Response
	costtime  time.Duration
}

// AddGetTask 新增 GET 任務
func (g *GroupTask) AddGetTask(url string) {
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	g.requests = append(g.requests, &Request{req})
}

// AddPostTask 新增 POST 任務
func (g *GroupTask) AddPostTask(url string, postParams map[string]interface{}) {
	req, err := http.NewRequest(POST, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	g.requests = append(g.requests, &Request{req})
}

// Run 執行 HTTP 群組任務
func (g *GroupTask) Run() []*Response {
	var timebegin = time.Now()

	var wg sync.WaitGroup

	retChans := make([]chan *Response, len(g.requests))
	for i := range retChans {
		retChans[i] = make(chan *Response, 1)
	}

	for i, request := range g.requests {
		wg.Add(1)
		go func(index int, req *Request) {
			defer wg.Done()
			response, _ := req.Send()
			retChans[index] <- response
		}(i, request)
	}

	wg.Wait()
	for _, retCh := range retChans {
		select {
		case val := <-retCh:
			close(retCh)
			g.responses = append(g.responses, val)
		case <-time.After(RequestTimeout):
			panic(fmt.Errorf("request timeout over %v", RequestTimeout))
		}
	}

	g.costtime = time.Since(timebegin)
	return g.responses
}

// Output ...
func (g *GroupTask) Output() {
	var requestCount = len(g.requests)
	fmt.Println("count:", requestCount)
	fmt.Println("[request]")
	for _, v := range g.requests {
		fmt.Println(v)
	}
	fmt.Println("")
	fmt.Println("[response]")
	for _, v := range g.responses {
		fmt.Println(v.Status)
	}
	fmt.Println("")
	fmt.Println("cost", g.costtime)
}

// OutputWithBody ...
func (g *GroupTask) OutputWithBody() {
	var requestCount = len(g.requests)
	fmt.Println("count:", requestCount)
	fmt.Println("[request]")
	for _, v := range g.requests {
		fmt.Println(v)
	}
	fmt.Println("")
	fmt.Println("[response]")
	for _, v := range g.responses {
		fmt.Println(v.String())
	}
	fmt.Println("")
	fmt.Println("cost", g.costtime)
}
