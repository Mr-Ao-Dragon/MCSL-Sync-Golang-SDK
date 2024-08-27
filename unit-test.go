package main

import (
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
	"sync"
)

func main() {
	dirPath := "./"
	suffix := "_test.go"
	var wg sync.WaitGroup
	results := make([]string, 0)
	errs := make([]error, 100)
	outputBytes := make([]byte, 100)
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			errs = append(errs, err)
		}
		if !info.IsDir() && filepath.Ext(path) == suffix {
			go func() {
				wg.Add(1)
				cmd := exec.Command("go", "test", path)
				outputBytes, err = cmd.CombinedOutput()
				if err != nil {
					errs = append(errs, err)
				}
				results = append(results, string(outputBytes))
				defer wg.Done()
			}()

		}
		return nil
	})
	if err != nil {
		errs = append(errs, err)
	}
	go func() {
		wg.Wait()
	}()
	for _, v := range errs {
		if err != nil {
			log.Printf("%v", v)
		}
	}
	if errs != nil {
		panic(errs)
	}
	if len(errs) != 0 {
		log.Fatalf("Error: %#v", errs)
	}
	for Result := range results {
		log.Println(Result)
	}
	return
}
