package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	month := now.Format("200601")
	date := now.Format("20060102")

	imgs, err := getWallpaper(date)
	if err != nil {
		fmt.Printf("get imgs failed: %v", err)
		return
	}
	content := `## bing wallpaper`
	for i, v := range imgs {
		// PC
		landscape := domain + v.ImageUrls.Landscape.UltraHighDef
		// Mobile
		portrait := domain + v.ImageUrls.Portrait.UltraHighDef

		item := fmt.Sprintf(
			tmpl,
			v.IsoDate,
			v.Title,
			v.Copyright,
			v.Description,
			v.DescriptionPara2,
			v.DescriptionPara3,
			v.Caption,
			landscape,
			v.Caption,
			portrait,
			landscape,
			portrait,
		)
		content += item
		// archive the latest documents
		if i == 0 {
			archivePath := fmt.Sprintf("archive/%s_README.md", month)
			if isExist(archivePath) {
				writeAppend(archivePath, item)
			} else {
				write(archivePath, item)
			}
		}
	}

	write("./README.md", content)
}
