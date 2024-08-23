package get

import (
	"MCSL-Sync-Golang-SDK/info"
	"MCSL-Sync-Golang-SDK/util"
	"github.com/bytedance/sonic"
	"github.com/parnurzeal/gorequest"
	"time"
)

type TaskObj struct {
	TargetCore  string
	VersionInfo info.CoreVersionInfo
	TargetPath  string
	FileName    string
}
type versionMetadata struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []struct {
		Id          string    `json:"id"`
		Type        string    `json:"type"`
		Url         string    `json:"url"`
		Time        time.Time `json:"time"`
		ReleaseTime time.Time `json:"releaseTime"`
	} `json:"versions"`
}

func VersionList(versionType string, latest bool) (versions []string, errs []error) {
	req := gorequest.New()
	_, body, httpErrors := req.Get(util.Https + "bmclapi2.bangbang93.com/mc/game/version_manifest.json").
		End()
	if httpErrors != nil {
		return nil, httpErrors
	}
	data := new(versionMetadata)
	errs = append(errs, sonic.UnmarshalString(body, data))
	if errs != nil {
		return nil, errs
	}
	if latest {
		versions = append(versions, data.Latest.Release)
		errs = nil
		return
	}
	for _, version := range data.Versions {
		if version.Type == versionType {
			versions = append(versions, version.Id)
		}
	}
	errs = nil
	return
}
