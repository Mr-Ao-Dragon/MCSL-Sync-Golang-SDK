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
	ch := make(chan string, 100)
	errs := make(chan error, 100)
	outputBytes := make([]byte, 100)
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			errs <- err
		}
		if !info.IsDir() && filepath.Ext(path) == suffix {
			go func() {
				wg.Add(1)
				cmd := exec.Command("go", "test", path)
				outputBytes, err = cmd.CombinedOutput()
				if err != nil {
					errs <- err
				}
				ch <- string(outputBytes)
				defer wg.Done()
			}()

		}
		return nil
	})
	if err != nil {
		errs <- err
	}
	go func() {
		wg.Wait()
	}()
	if len(errs) != 0 {
		log.Fatalf("Error: %#v", <-errs)
	}
	for Result := range ch {
		log.Println(Result)
	}
	defer func() {
		close(ch)
		close(errs)
	}()
	return
}
