package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	var avilable int64 = 10000000
	var min = avilable
	var max = avilable + 10
	succeed := make(chan int64, 0)
	lost := func() {
		for {
			uid := rand.Int63n(max-min) + min
			err := handle(uid)
			if err == nil {
				succeed <- uid
			}

		}
	}
	for i := 0; i < 10; i++ {
		go lost()
	}
	go func() {
		for {
			err := check(min - 1)
			if err != nil {
				return
			} else {
				min = min - 1
			}
		}
	}()
	go func() {
		for {
			err := check(max + 1)
			if err != nil {
				return
			} else {
				max = max + 1
			}
		}
	}()
	var counter uint64 = 0
	var circle uint64 = 0
	var circle2 uint64 = 0
	var CIRCLE uint64 = 1 << 62

	for {
		select {
		case <-succeed:
			counter++
			if counter > CIRCLE {
				counter = counter % CIRCLE
				circle += 1
				if circle > CIRCLE {
					circle = circle % CIRCLE
					circle2 += 1
					if circle2 > CIRCLE {
						fmt.Printf("%v :%d -> %d -> %d\n", time.Now(), circle2, circle, counter)
						circle2 = 1
					}
				}
			}
		}
	}
}

func check(uid int64) error {
	// 获取TOKEN
	response, err := http.Get(fmt.Sprintf("http://gg.liujiaweixiaoman.cn/chabai/v1/?key=123456&uid=%d", uid))
	if err != nil || response == nil || response.StatusCode != 200 {
		return fmt.Errorf("err:%v", err)
	}
	token := TokenResponse{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&token)
	if err != nil {
		return err
	}
	if token.Code != 200 {
		return fmt.Errorf("err:%v", token.Msg)
	}
	return nil
}

func handle(uid int64) error {
	// 获取TOKEN
	response, err := http.Get(fmt.Sprintf("http://gg.liujiaweixiaoman.cn/chabai/v1/?key=123456&uid=%d", uid))
	if err != nil || response == nil || response.StatusCode != 200 {
		return fmt.Errorf("err:%v", err)
	}
	token := TokenResponse{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&token)
	if err != nil {
		return err
	}
	if token.Code != 200 {
		return fmt.Errorf("err:%v", token.Msg)
	}
	// 请求
	request := TokenRequest{
		Type:  0,
		Time:  3,
		Token: token.Msg,
	}
	url := "http://gg.liujiaweixiaoman.cn/chabai/v1/"
	contentType := "application/json"
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(b)
	resp, err := http.Post(url, contentType, body)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return fmt.Errorf("err:%v", err)
	}
	return nil
}
