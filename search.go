package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func saucenaoSearch(fileID string) string {
	photo, err := os.Open(fmt.Sprintf("./cache%s.jpg", fileID))
	if err != nil {
		log.Println(err)
	}
	defer photo.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fmt.Sprintf("./cache%s.jpg", fileID))
	if err != nil {
		log.Println(err)
	}
	_, err = io.Copy(part, photo)
	if err != nil {
		log.Println(err)
	}
	writer.Close()

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("https://saucenao.com/search.php?output_type=2&numres=1&hide=0&db=999&api_key=%s", config.SaucenaoAPIKey), body)
	if err != nil {
		log.Println(err)
	}
	httpReq.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	reqDo, err := client.Do(httpReq)
	if err != nil {
		log.Println(err)
	}
	response, err := ioutil.ReadAll(reqDo.Body)
	if err != nil {
		log.Println(err)
	}
	return string(response)
}
