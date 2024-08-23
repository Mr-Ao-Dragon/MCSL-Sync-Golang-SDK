package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"sync"
)

var wg sync.WaitGroup

func (receiver *CoreInfo) getCoreInfo(client setup.Client) []error {
	errors := receiver.GetCoreSupportMcList(client)
	if errors != nil {
		return errors
	}
	receiver.GetCoreSupportMcList(client)
	versionOrigin := client.MCVersion
	for _, mcVersion := range receiver.SupportMcVersion {
		client.MCVersion = mcVersion
		wg.Add(1)
		go receiver.GetCoreBuildListSingleMCVersion(client)
	}
	wg.Wait()
	client.MCVersion = versionOrigin

	return nil
}
