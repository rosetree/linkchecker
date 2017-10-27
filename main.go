package main

import (
	"fmt"
	"io/ioutil"
	"mvdan.cc/xurls"
	"os"
	"path/filepath"
	"regexp"
)

var collection linkCollection

func main() {
	wd, _ := os.Getwd()

	err := filepath.Walk(wd, extractLinks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, link := range collection.links {
		err = link.FetchStatus()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if link.StatusSuccess() {
			continue
		}

		fmt.Println(link)
	}
}

func extractLinks(path string, fi os.FileInfo, err error) error {
	if !fi.Mode().IsRegular() {
		return nil
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	links := xurls.Strict().FindAllString(string(f), -1)
	if links == nil {
		return nil
	}

	for _, link := range links {
		isHTTP, err := regexp.MatchString("^http.*", link)
		if err != nil || !isHTTP {
			continue
		}

		collection.Add(link, path)
	}

	return nil
}
