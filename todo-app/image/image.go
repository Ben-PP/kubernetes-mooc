package image

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"todo-app/directory"
)

func getLastUpdated(dir string) int64 {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	if len(entries) == 0 {
		return int64(0)
	}
	fmt.Println(entries)
	lastImage := entries[len(entries)-1].Name()
	lastUpdated := lastImage[:len(lastImage)-4]
	value, err := strconv.ParseInt(lastUpdated, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func makeDataUri(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	enc := base64.StdEncoding.EncodeToString(b)
	uri := fmt.Sprintf("data:image/jpeg;base64,%s", enc)
	return uri, nil
}

func GetImage(dataDir string) (string, error) {
	imageDir := fmt.Sprintf("%s/images", dataDir)
	err := directory.MustExist(imageDir)
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	var lastUpdated int64 = getLastUpdated(imageDir)
	var fileName string = fmt.Sprintf("%s/%d.jpg", imageDir, lastUpdated)
	timeDif := timestamp - lastUpdated
	if timeDif > 600{

		fmt.Println(lastUpdated)
		resp, err := http.Get("https://picsum.photos/1200")
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return "", err
		}
		file, err := os.Create(fmt.Sprintf("%s/%d.jpg",imageDir, timestamp))
		if err != nil {
			return "", err
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return "", err
		}
		fileName = fmt.Sprintf("%s/%d.jpg", imageDir, timestamp)
		if lastUpdated != 0 {
			err := os.Remove(fmt.Sprintf("%s/%d.jpg", imageDir, lastUpdated))
			if err != nil {
				return "", err
			}
		}
	}
	imageURI, err := makeDataUri(fileName)
	if err != nil {
		return "", err
	}
	return imageURI, nil
}
