package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/parnurzeal/gorequest"
)

type ApiRepose struct {
	Data []string `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}

func (receiver *CoreList) add(name []string) {
	*receiver = append(*receiver, name...)
}
func (receiver *CoreList) ReadCoreList(setupData setup.Client) (errs []error) {
	request := gorequest.New()
	_, body, httpErrs := request.Get(setupData.ApiDomain + "/core").End()
	if httpErrs != nil {
		errs = append(errs, httpErrs...)
	}
	var data ApiRepose
	err := sonic.Unmarshal([]byte(body), &data)
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	if data.Code != 200 {
		errs = append(errs, errors.New(data.Msg))
		return
	}
	receiver.add(data.Data)
	return nil
}
