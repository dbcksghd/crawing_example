package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://snuarori.snu.ac.kr/renew/admission_guide/typical_guide2.php?board_07=2022"

	list := []string{}
	list2 := []string{}
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tr").Each(func(_ int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			list = append(list, text)
		}
	})
	for i, a := range list {
		if a[0] >= 48 && a[0] <= 57 {
			for j, _ := range a {
				if i < 10 {
					if j > 5 && a[j] <= 57 && a[j] >= 48 {
						list2 = append(list2, a[1:j])
						for k, _ := range a {
							if k >= j && (a[k] >= 49 && a[k] <= 57) {
								list2 = append(list2, a[k+2:])
								break
							}
						}
						break
					}
				} else {
					if j > 5 && a[j] <= 57 && a[j] >= 48 {
						list2 = append(list2, a[2:j])
						for k, _ := range a {
							if k >= j && (a[k] >= 49 && a[k] <= 57) {
								list2 = append(list2, a[k+2:])
								break
							}
						}
						break
					}
				}
			}
		}
	}
	for _, a := range list2 {
		fmt.Println(a)
	}
}
