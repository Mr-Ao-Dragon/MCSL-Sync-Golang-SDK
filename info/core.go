package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"sync"
)

var wg sync.WaitGroup

// getCoreInfo 获取核心信息。该方法用于通过指定的客户端信息来获取系统的核心支持、构建列表等信息。
// 参数:
//
//	client: setup.Client 类型的客户端信息，用于连接和交互。
//
// 返回值:
//
//	[]error: 返回一个错误切片，表示在获取核心信息过程中出现的任何错误。
func (receiver *CoreInfo) getCoreInfo(client setup.Client) []error {
	// 首先，调用GetCoreSupportMcList方法来获取支持的MC版本列表。
	// 如果出现错误，则直接返回该错误列表。
	GCSMErrors := receiver.GetCoreSupportMcList(client)
	if GCSMErrors != nil {
		return GCSMErrors
	}
	// versionOrigin 用于备份原始的MC版本，以便在所有操作完成后恢复。
	versionOrigin := client.MCVersion
	// 遍历支持的MC版本列表，对每个版本执行操作。
	errCh := make(chan error)
	for _, mcVersion := range receiver.SupportMcVersion {
		// 设置当前的MC版本为正在处理的版本。
		client.MCVersion = mcVersion
		// 使用WaitGroup来确保所有并发操作完成。
		wg.Add(1)
		// 启动一个goroutine并传入当前的客户端信息，用于并发地处理每个MC版本。
		go func(client setup.Client, errCh chan<- error) {
			if err := receiver.GetCoreBuildListSingleMCVersion(client); err != nil {
				for _, errSig := range err {
					errCh <- errSig
				}
			}
			defer wg.Done()
		}(client, errCh)
	}
	for buildNum := range receiver.HistoryVersion {
		wg.Add(1)
		go func(buildNum string, client setup.Client, errch chan<- error) {
			if err := receiver.GetTargetBuildInfo(client, buildNum); err != nil {
				for _, errSig := range err {
					errCh <- errSig
				}
			}
			defer wg.Done()
		}(buildNum, client, errCh)
	}
	close(errCh)
	errs := make([]error, 0)
	for err := range errCh {
		errs = append(errs, err)
	}
	// 等待所有goroutine完成，确保所有数据都已处理完毕。
	wg.Wait()
	// 恢复原始的MC版本。
	client.MCVersion = versionOrigin
	// 如果整个过程中没有出现错误，返回nil。
	if len(errs) != 0 {
		return errs
	}
	return nil
}
