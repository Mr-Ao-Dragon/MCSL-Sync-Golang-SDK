package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"MCSL-Sync-Golang-SDK/util"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
)

func (receiver *CoreInfo) GetTargetBuildInfo(setupData setup.Client, targetBuild string) []error {
	request := gorequest.New()
	_, body, httpErrs := request.Get(util.Https + setupData.ApiDomain + "/core/" + setupData.CoreName + "/" + setupData.MCVersion + "/" + targetBuild).
		End()
	if httpErrs != nil {
		return httpErrs
	}
	data := new(TargetBuildInfoRepose)
	if err := sonic.Unmarshal([]byte(body), &data); err != nil {
		return []error{err}
	}
	if data.Code != 200 {
		return []error{errors.New(data.Msg)}
	}
	dlUrl, _ := url.Parse(data.Data.Build.DownloadUrl)
	buildData := CoreVersionInfo{
		DownloadUrl:        *dlUrl,
		TargetMcVersion:    (*data).Data.Build.McVersion,
		RequestJavaVersion: util.JavaVersionCatch(setupData.MCVersion),
		SyncTime:           (*data).Data.Build.SyncTime,
	}
	receiver.HistoryVersion[targetBuild] = buildData
	return nil
}

type TargetBuildInfoRepose struct {
	Data struct {
		Type  string `json:"type"`
		Build struct {
			SyncTime    time.Time `json:"sync_time"`
			DownloadUrl string    `json:"download_url"`
			CoreType    string    `json:"core_type"`
			McVersion   string    `json:"mc_version"`
			CoreVersion string    `json:"core_version"`
		} `json:"build"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
