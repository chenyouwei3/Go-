package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	HUNTER_API_URL = "https://hunter.qianxin.com/api/search"
)

// Payload 结构体用于存储请求的参数
type Payload struct {
	Search          string `json:"search"`
	PageSize        int    `json:"page_size"`
	Page            int    `json:"page"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	IsWeb           string `json:"is_web"`
	StatusCode      string `json:"status_code"`
	SyntaxCondition []int  `json:"syntax_condition"`
	AssetTag        []int  `json:"asset_tag"`
	IPTag           []int  `json:"ip_tag"`
	PortFilter      bool   `json:"port_filter"`
}

// Response 结构体用于解析API响应
type Response struct {
	Data struct {
		Total int `json:"total"`
	} `json:"data"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file_name> <output_file_name>")
		return
	}
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	inputFile, err := os.Open(inputFileName) // 打开输入文件
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputFileName) // 创建输出文件
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	headers := map[string]string{
		"Cookie":           "wzws_sessionid=gmQ0ZjAxMoEzNjEwNmSAMTExLjQ2LjU3Ljk3oGYnbsg=; __8qcehdE7ZaRq2q6M__=b9623d06d01035df623d547185bf5a12; Hm_lvt_64787111d439a06146c3a4be00dda632=1713860300; next=https%3A//hunter.qianxin.com/api/uLogin; User-Center=d0436176-6c81-48b0-864e-bb6a4ba483fb; token=kYxAB%3Ae105926f-ca12-476e-8e36-2965c452295a; session_id=2c81989123b8135cceaac693d771711ae75b1a2351928c7e9e10df091f09faa2; Hm_lpvt_64787111d439a06146c3a4be00dda632=1713860440; csrf_token=1713864042##1d8aa021909bfad8d655b7b4e41e0528f0067244",
		"Accept":           "application/json, text/plain, */*",
		"Content-Type":     "application/json",
		"Sec-Ch-Ua-Mobile": "?0",
		"Authorization":    "Bearer kYxAB:e105926f-ca12-476e-8e36-2965c452295a",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.127 Safari/537.36",
		"Origin":           "https://hunter.qianxin.com",
		"Sec-Fetch-Site":   "same-origin",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-Dest":   "empty",
		"Accept-Language":  "zh-CN,zh;q=0.9",
		"Connection":       "close",
	}
	client := &http.Client{}
	scanner := bufio.NewScanner(inputFile) // 创建文件扫描器
	for scanner.Scan() {
		line := scanner.Text()
		lineWithoutCR := strings.TrimSpace(line)
		fmt.Println(lineWithoutCR)

		searchString := fmt.Sprintf(`icp.name="%s" OR web.body="%s" OR web.title="%s"`, lineWithoutCR, lineWithoutCR, lineWithoutCR)
		fmt.Println(searchString)
		encodedSearch := base64.URLEncoding.EncodeToString([]byte(searchString))
		fmt.Println(encodedSearch)
		payload := Payload{
			Search:     encodedSearch,
			PageSize:   10,
			Page:       1,
			StartTime:  "2024-03-25",
			EndTime:    "2024-04-23",
			IsWeb:      "0",
			StatusCode: "",
			PortFilter: false,
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshalling payload: %v\n", err)
			continue
		}

		req, err := http.NewRequest("POST", HUNTER_API_URL, bytes.NewReader(payloadBytes))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}

		for key, value := range headers {
			req.Header.Set(key, value)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			continue
		}

		var response Response
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			fmt.Printf("Error unmarshalling response: %v\n", err)
			continue
		}

		total := response.Data.Total
		fmt.Println(resp.StatusCode)
		fmt.Println(total)
		outputFile.WriteString(fmt.Sprintf("%d\n", total))

		time.Sleep(10 * time.Second)
	}
}
