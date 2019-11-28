package client

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
func stringOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

// ReadImage create image from the path
func ReadImage(path string) *string {
	img, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	imgBase64 := base64.StdEncoding.EncodeToString(img)
	return stringOrNil(imgBase64)
}
