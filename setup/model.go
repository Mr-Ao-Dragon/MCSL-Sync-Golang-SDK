package setup

import (
	"MCSL-Sync-Golang-SDK/util"
	"log"
)

type Client struct {
	ApiDomain   string
	Node        string
	CoreName    string
	MCVersion   string
	IsLatest    bool
	CoreVersion string
	TargetPath  string
}

func InitSetupData(ApiDomain string, isLatest bool, moreArgs ...string) *Client {
	clientData := new(Client)
	clientData.ApiDomain = ApiDomain
	if isLatest {
		versions, err := util.VersionList("release", true)
		if err != nil {
			log.Fatalf("%#v", err)
		}
		clientData.CoreVersion = versions[0]
	}
	for argNum := 0; len(moreArgs) != argNum; argNum++ {
		switch argNum {
		case 0:
			continue
		case 1:
			if moreArgs[1] != "" {
				clientData.Node = moreArgs[argNum]
			}
			continue
		case 2:
			clientData.CoreName = moreArgs[argNum]
		case 3:
			clientData.MCVersion = moreArgs[argNum]
		case 4:
			clientData.CoreVersion = moreArgs[argNum]
		case 5:
			clientData.TargetPath = moreArgs[argNum]
		default:
			log.Fatalf("client setting have Too many arguments,all arguments are %#v", moreArgs)
		}
	}
	return clientData
}
