package info

import (
	"errors"
	"github.com/Mr-Ao-Dragon/MCSL-Sync-Golang-SDK/setup"
	"github.com/bytedance/sonic"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

type coreListApiRepose struct {
	Data []string `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}
type coreInfoApiRepose struct {
	Data struct {
		Type     string   `json:"type"`
		Versions []string `json:"versions"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type CoreVersionInfoRepose struct {
	Data struct {
		Type   string   `json:"type"`
		Builds []string `json:"builds"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (receiver *CoreList) coreListAdd(name []string) {
	*receiver = append(*receiver, name...)
}
func (receiver *CoreList) ReadCoreList(setupData setup.Client) (errs []error) {
	request := gorequest.New()
	errs = make([]error, 0)
	reqUrl := new(url.URL)
	reqUrl.Scheme = "https"
	reqUrl.Host = setupData.ApiDomain
	reqUrl.Path = "/api/core"
	_, body, httpErrs := request.Get(reqUrl.String()).End()
	if httpErrs != nil {
		errs = append(errs, httpErrs...)
	}
	data := new(coreListApiRepose)
	err := sonic.UnmarshalString(body, &data)
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	if (*data).Code != 200 {
		errs = append(errs, errors.New(data.Msg))
		return
	}
	receiver.coreListAdd(data.Data)
	return nil
}
func (receiver *CoreInfo) GetCoreSupportMcList(setupData setup.Client) []error {
	request := gorequest.New()
	reqUrl := new(url.URL)
	reqUrl.Scheme = "https"
	reqUrl.Host = setupData.ApiDomain
	reqUrl.Path = "/api/core" + "/" + setupData.CoreName
	if setupData.CoreName == "Arclight" {
		return []error{errors.New("arclight is not supported")}
	}
	_, body, httpErrs := request.Get(reqUrl.String()).
		End()
	if httpErrs != nil {
		return httpErrs
	}
	data := new(coreInfoApiRepose)
	if err := sonic.UnmarshalString(body, &data); err != nil {
		return []error{err}
	}
	if (*data).Code != 200 {
		return []error{errors.New(data.Msg)}
	}
	receiver.SupportMcVersion = (*data).Data.Versions
	receiver.Name = setupData.CoreName
	return nil
}
func (receiver *CoreInfo) GetCoreBuildListSingleMCVersion(setupData setup.Client) []error {
	request := gorequest.New()
	reqUrl := new(url.URL)
	reqUrl.Scheme = "https"
	reqUrl.Host = setupData.ApiDomain
	reqUrl.Path = "/api/core/" + setupData.CoreName + "/" + setupData.MCVersion
	_, body, httpErrs := request.Get(reqUrl.String()).
		End()
	if httpErrs != nil {
		return httpErrs
	}
	data := new(CoreVersionInfoRepose)
	if err := sonic.UnmarshalString(body, &data); err != nil {
		return []error{err}
	}
	if (*data).Code != 200 {
		return []error{errors.New((*data).Msg)}
	}
	receiver.SupportMcVersion = append(receiver.SupportMcVersion, setupData.MCVersion)
	receiver.HistoryVersion = make(map[string]CoreVersionInfo)
	for _, buildVerNum := range (*data).Data.Builds {
		buildData := CoreVersionInfo{
			TargetMcVersion: setupData.MCVersion,
		}
		receiver.HistoryVersion[buildVerNum] = buildData
	}
	return nil

}
