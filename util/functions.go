package util

import (
	"github.com/bytedance/sonic"
	"github.com/hashicorp/go-version"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
)

func JavaVersionCatch(mcVersion string) (javaVersion int) {
	java7, _ := version.NewVersion("1.7.2")
	java8, _ := version.NewVersion("1.12.2")
	java11, _ := version.NewVersion("1.14.4")
	java15, _ := version.NewVersion("1.16.5")
	java17, _ := version.NewVersion("1.20.6")
	java21, _ := version.NewVersion("1.21.1")
	unknownVersion, _ := version.NewVersion(mcVersion)
	switch {
	case unknownVersion.LessThanOrEqual(java7):
		return 7
	case unknownVersion.LessThanOrEqual(java8):
		return 8
	case unknownVersion.LessThanOrEqual(java11):
		return 11
	case unknownVersion.LessThanOrEqual(java15):
		return 15
	case unknownVersion.LessThanOrEqual(java17):
		return 17
	case unknownVersion.LessThanOrEqual(java21):
		return 21
	default:
		return 21
	}
}

func VersionList(versionType string, latest bool) (versions []string, errs []error) {
	req := gorequest.New()
	reqUrl := new(url.URL)
	reqUrl.Scheme = "Https"
	reqUrl.Host = "bmclapi2.bangbang93.com"
	reqUrl.Path = "/mc/game/version_manifest.json"
	_, body, httpErrors := req.Get(reqUrl.String()).
		End()
	if httpErrors != nil {
		return nil, httpErrors
	}
	data := new(VersionMetadata)
	errs = append(errs, sonic.UnmarshalString(body, data))
	if errs != nil {
		return nil, errs
	}
	if latest {
		versions = append(versions, (*data).Latest.Release)
		errs = nil
		return
	}
	for _, MCVersion := range data.Versions {
		if MCVersion.Type == versionType {
			versions = append(versions, MCVersion.Id)
		}
	}
	errs = nil
	return
}

type VersionMetadata struct {
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
