# 初始化

```go
package main

import "MCSL-Sync-Golang-SDK/setup"

func main() {
	// 不接受乱序
	client := setup.InitSetupData(
		"sync.mcsl.com.cn", // api address
		false,              // is latest，可选
		"",                 // node address，可选
		"Spigot",           // core name，可选
		"1.18.2",           // minecraft version，可选
		"v42",              // core build version，可选
		"/download/",       // 目标下载路径，推荐填写
		)
}
```
# 获取mc版本

```go
package main

import (
	"MCSL-Sync-Golang-SDK/info"
	"MCSL-Sync-Golang-SDK/setup"
	"log"
)

func get(client setup.Client) {
	mcVersion := new(info.CoreList)
	// client 需要填写MC版本
	errs := mcVersion.ReadCoreList(client)
	// 结果在 mcVersion 内
	if len(errs) != 0 {
		log.Fatalf("%v",errs)
	}
}
```
# 获取核心信息

```go
package main

import (
	"MCSL-Sync-Golang-SDK/info"
	"MCSL-Sync-Golang-SDK/setup"
	"log"
)

func get(client setup.Client) {
	coreData := new(info.CoreInfo)
	// client 需要填写核心名称
	// 不可填写 Arclight，螺端版本号不遵循 x.x.x 的格式
	errs := coreData.getCoreInfo(client)
	// 结果在 coreData 内
	if len(errs) != 0 {
		log.Fatalf("%v", errs)
	}
}
```
# 下载核心

```go
package main

import (
	"MCSL-Sync-Golang-SDK/get"
	"MCSL-Sync-Golang-SDK/info"
	"MCSL-Sync-Golang-SDK/setup"
	"log"
)

func download(client setup.Client) {
	coreData := new(info.CoreInfo)
	coreData.getCoreInfo(client)
	// 请自行对核心构建列表进行排序
	err := get.Download(client, coreData.HistoryVersion[0])
	if err != nil {
		log.Fatalf("%v",err)
	}
}
```