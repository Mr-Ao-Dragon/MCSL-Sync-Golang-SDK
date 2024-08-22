package get

import "MCSL-Sync-Golang-SDK/info"

type TaskObj struct {
	targetCore  string
	versionInfo info.CoreVersionInfo
	targetPath  string
	fileName    string
}
