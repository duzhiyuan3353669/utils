package http_utils

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	url2 "net/url"
	"os"
)

//Get 请求
func Get(url string, header http.Header) (*http.Response, error) {
	var req http.Request
	TargetUrl, err := url2.Parse(url)
	if err != nil {
		return nil, errors.New("url地址异常")
	}
	req.URL = TargetUrl
	req.Method = "GET"
	req.Header = header
	var client http.Client
	resp, err := client.Do(&req)
	if err != nil {
		return nil, errors.New("调用GET异常")
	}
	return resp, nil
}

//Post 请求
func Post(url string, body []byte, header http.Header) (*http.Response, error) {
	var req http.Request
	TargetUrl, err := url2.Parse(url)
	if err != nil {
		return nil, errors.New("url地址异常")
	}
	req.URL = TargetUrl
	req.Method = "POST"
	bd := io.NopCloser(bytes.NewReader(body))
	req.Body = bd
	req.ContentLength = int64(len(body))
	req.Header = header
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	var client http.Client
	resp, err := client.Do(&req)
	if err != nil {
		return nil, errors.New("调用GET异常")
	}
	return resp, nil
}

//UploadFile 请求
func UploadFile(filename string, url string, header http.Header) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	defer bodyWriter.Close()
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, url, bodyBuf)
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("调用上传文件异常")
	}

	return resp, nil
}
