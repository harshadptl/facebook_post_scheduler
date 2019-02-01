package main

import (
	"fmt"
	fb "github.com/huandu/facebook"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const FileExtJPEG = ".jpeg"
const FileExtJPG = ".jpg"
const FileExtPNG = ".png"
const FileExtGIF = ".gif"

const PageID = "282170032653457"


const AccessToken = "EAANtBO5jw70BAPdLSMfFuanmlsoPrrhN5RrZAVjZBWqZA33kT9B31L0ooHRgTS6cVJxPZCbD9wxQULSiErjsL66NVZCfrn0wUlBs53RtUGUp0HSKZCUARRX9g61SEsKkmUEkUw5Or0EXPYbkMEDuMnHUp6gB1wGTesqBNLtOQePF7QlEFuBVfBkDM6lN2j2ZCvACtkhIrLO5wZDZD"

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	scheduled_publish_time := time.Now().UTC()
	count := 0

	for _, f := range files {
		if isImageFileName(f.Name()) {
			fmt.Println(f.Name(), scheduled_publish_time.Unix()+720)
			t := scheduled_publish_time.Unix()+720
			if count %2 == 1 {
				t = t + 12*60*60
			}
			res, _ := fb.Post("/v3.2/" + PageID + "/photos", fb.Params{
				"source":          fb.File(f.Name()),
				"access_token": AccessToken,
				"scheduled_publish_time": t,
				"published": false,
			})
			fmt.Println("data: ", res["id"], res["post_id"])
			fmt.Printf("%+v \n", res)

			if count %2 == 1{
				scheduled_publish_time = scheduled_publish_time.AddDate(0,0,1)
			}
			count = count + 1
		}
	}
}

func isImageFileName(fName string) bool {
	if strings.HasSuffix(fName, FileExtJPEG) {
		return true
	}

	if strings.HasSuffix(fName, FileExtJPG) {
		return true
	}

	if strings.HasSuffix(fName, FileExtPNG) {
		return true
	}

	if strings.HasSuffix(fName, FileExtGIF) {
		return true
	}

	return false
}
