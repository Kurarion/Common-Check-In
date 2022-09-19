package core

import (
	types "CommonCheckIn/types"
	util "CommonCheckIn/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	cron "github.com/robfig/cron"
)

//json file name
const jsonFile = "assets/list.json"

//time format
const timeFormat = "2006-01-02 15:04:05"

func Run() {
	//read JSON file from disk
	list := make(types.CheckInDatas, 0)
	buf, err := util.GetLocalJSONBytes(jsonFile)
	if err != nil {
		fail(err, nil)
		return
	}
	json.Unmarshal(buf.Bytes(), &list)
	//create Cron
	c := cron.New()
	//loop access
	for _, v := range list {
		//check enable
		if v.Disable {
			continue
		}
		//http
		fn := getCheckInFunc(v)
		for _, vv := range v.Specs {
			err := c.AddFunc(vv, fn)
			if err != nil {
				fail(err, &v)
				return
			}
		}
	}
	c.Start()
	defer c.Stop()
	select {}
}

func getCheckInFunc(v types.CheckInData) func() {
	return func() {
		client := new(http.Client)
		req, _ := http.NewRequest(strings.ToUpper(v.Method), v.Url, strings.NewReader(v.Payload))
		for _, vv := range v.Headers {
			req.Header.Add(vv.Key, vv.Value)
		}
		switch 1 {
		case 1:
			resp, err := client.Do(req)
			if err != nil {
				fail(err, &v)
				break
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fail(err, &v)
				break
			}
			info(string(body), &v)
		}
	}
}

func printData(data *types.CheckInData) {
	if data != nil {
		currentTime := time.Now()
		fmt.Println("======[" + currentTime.Format(timeFormat) + "]======")
		fmt.Println("URL: ", data.Url)
		fmt.Println("Method: ", data.Method)
		fmt.Println("Payload: ", data.Payload)
		fmt.Println()
	}
}

func fail(err error, data *types.CheckInData) {
	printData(data)
	fmt.Println("[Fail]: ", err)
	fmt.Println("======[XXXXXXXXXXXXXXXXXXXXXX]======")
}

func info(s string, data *types.CheckInData) {
	printData(data)
	fmt.Println("[Info]: ", s)
	fmt.Println("====================================")
}
