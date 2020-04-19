package com_utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strings"
)

func Md5(str string) string {

	hash := md5.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}

func IsFileExists(filename string) bool {

	_, err := os.Stat(filename)

	if err != nil {
		if err == os.ErrExist {
			return true
		}
		return false
	}
	return true
}

func GetExt(filename string) string {
	return filename[strings.LastIndex(filename, ".")+1:]
}
