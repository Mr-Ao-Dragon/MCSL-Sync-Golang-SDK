package get

import (
	"MCSL-Sync-Golang-SDK/info"
	"MCSL-Sync-Golang-SDK/setup"
	"errors"
	"github.com/hashicorp/go-getter"
	"log"
)

func Download(client setup.Client, info info.CoreVersionInfo, filename ...string) error {
	if len(filename) == 0 {
		filename = append(filename, client.CoreName+"-"+info.TargetMcVersion+".jar")
		log.Printf("filename is empty, use default filename: %s\n", filename[0])
	} else if filename[0] == "" {
		filename = append(filename, client.CoreName+"-"+info.TargetMcVersion+".jar")
		log.Printf("filename is empty, use default filename: %s\n", filename[0])
	}
	if len(filename) > 1 {
		return errors.New("too many filename")
	}
	err := getter.GetFile(client.TargetPath+"/"+filename[0], info.DownloadUrl.String())
	if err != nil {
		return err
	}
	log.Printf("download %s success\n", filename[0])
	log.Printf("file path is %s", client.TargetPath+"/"+filename[0])
	return nil
}
