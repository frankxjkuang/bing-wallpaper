package main

const domain = "https://cn.bing.com"

type ImageUrl struct {
	//HighDef      string `json:"highDef"`
	UltraHighDef string `json:"ultraHighDef"` // 使用该url
	//Wallpaper    string `json:"wallpaper"`
}

type Image struct {
	Caption     string `json:"caption"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Copyright   string `json:"copyright"`
	//date        string `json:"date"`
	ImageUrls struct {
		// PC端
		Landscape ImageUrl `json:"landscape"`
		// 移动端
		Portrait ImageUrl `json:"portrait"`
	} `json:"imageUrls"`
	DescriptionPara2 string `json:"descriptionPara2"`
	DescriptionPara3 string `json:"descriptionPara3"`
	IsoDate          string `json:"isoDate"`
}

type data struct {
	Images []Image `json:"images"`
}

type Result struct {
	Title         string `json:"title"`
	Data          data   `json:"data"`
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}
