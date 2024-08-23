package info

import (
	"net/url"
	"time"
)

type CoreList []string
type CoreInfo struct {
	Name             string
	SupportMcVersion []string
	HistoryVersion   map[string]CoreVersionInfo
}
type CoreVersionInfo struct {
	TargetMcVersion    string
	SyncTime           time.Time
	RequestJavaVersion int
	DownloadUrl        url.URL
	SHA256             string
	Size               int
}

func NewCoreInfo() *CoreInfo {
	info := new(CoreInfo)
	return info
}
