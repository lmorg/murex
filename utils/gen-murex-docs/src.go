package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func scanSrc(src string) {
	err := filepath.Walk(src, walkCallback)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if f.IsDir() {
		return nil
	}

	if len(f.Name()) < 5 {
		return nil
	}

	name := f.Name()[:len(f.Name())-4]
	ext := f.Name()[len(f.Name())-4:]
	switch ext {
	case ".def":
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		Define[name] = strings.TrimSpace(string(b))

	case ".dig":
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		Digest[name] = strings.TrimSpace(string(b))

	case ".rel":
		s, err := readLines(path)
		if err != nil {
			return err
		}
		Related[name] = s
	}

	return nil
}

func readLines(filename string) ([]string, error) {
	var s []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return s, nil
}
