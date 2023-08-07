package main

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

func GbkToUtf8(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
