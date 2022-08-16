package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getWallpaper(ts string) (imgs []Image, err error) {

	url := fmt.Sprintf("https://cn.bing.com/hp/api/v1/imagegallery?format=json&ssd=%s_0700&FORM=BEHPTB", ts)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Header.Add("Cookie", "MUID=2E468C5F8AA568AC05489DAD8BEF6986; SRCHD=AF=BEHPTB; SRCHHPGUSR=SRCHLANG=zh-Hans; SRCHUID=V=2&GUID=6B723D3BD838423EA30ECC069ED6FCAC&dmnchg=1; SRCHUSR=DOB=20220731; SUID=M; _EDGE_S=F=1&SID=1763DC24934764BE3033CDD6920D6582; _EDGE_V=1; _HPVN=CS=eyJQbiI6eyJDbiI6MSwiU3QiOjAsIlFzIjowLCJQcm9kIjoiUCJ9LCJTYyI6eyJDbiI6MSwiU3QiOjAsIlFzIjowLCJQcm9kIjoiSCJ9LCJReiI6eyJDbiI6MSwiU3QiOjAsIlFzIjowLCJQcm9kIjoiVCJ9LCJBcCI6dHJ1ZSwiTXV0ZSI6dHJ1ZSwiTGFkIjoiMjAyMi0wNy0zMVQwMDowMDowMFoiLCJJb3RkIjowLCJHd2IiOjAsIkRmdCI6bnVsbCwiTXZzIjowLCJGbHQiOjAsIkltcCI6MX0=; _SS=SID=1763DC24934764BE3033CDD6920D6582; MUIDB=2E468C5F8AA568AC05489DAD8BEF6986")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	var resp Result
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Printf("Unmarshal err: %v", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s: %d", resp.StatusCode, resp.StatusMessage)
		return
	}

	return resp.Data.Images, nil
}
