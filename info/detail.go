package info

import (
	"errors"
	"github.com/Mr-Ao-Dragon/MCSL-Sync-Golang-SDK/setup"
	"github.com/Mr-Ao-Dragon/MCSL-Sync-Golang-SDK/util"
	"github.com/bytedance/sonic"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
)

func (receiver *CoreInfo) GetTargetBuildInfo(setupData setup.Client, targetBuild string) []error {
	request := gorequest.New()
	reqUrl := new(url.URL)
	reqUrl.Scheme = "https"
	reqUrl.Host = setupData.ApiDomain

	reqUrl.Path = "/api/core/" + setupData.CoreName + "/" + setupData.MCVersion + "/" + targetBuild
	_, body, httpErrs := request.Get(reqUrl.String()).
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
	dlUrl, err := url.Parse(data.Data.Build.DownloadUrl)
	if err != nil {
		return []error{err}
	}
	buildData := CoreVersionInfo{
		DownloadUrl:        *dlUrl,
		TargetMcVersion:    (*data).Data.Build.McVersion,
		RequestJavaVersion: util.JavaVersionCatch(setupData.MCVersion),
		SyncTime:           (*data).Data.Build.SyncTime,
	}
	if receiver.HistoryVersion == nil {
		receiver.HistoryVersion = make(map[string]CoreVersionInfo)
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
