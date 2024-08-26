package get

import (
	"MCSL-Sync-Golang-SDK/info"
	"MCSL-Sync-Golang-SDK/setup"
	"net/url"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	dlUrl, _ := url.Parse("https://hunan.node.sync.mcsl.com.cn:4524/core/Vanilla/1.21/Latest/download")
	userProfile := os.Getenv("USERPROFILE")

	type args struct {
		client   setup.Client
		info     info.CoreVersionInfo
		filename []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "download test for win",
			args: args{
				client: setup.Client{
					CoreName:   "Vanilla",
					TargetPath: userProfile + "\\AppData\\Local\\Temp\\",
					MCVersion:  "1.21.1",
				},
				info: info.CoreVersionInfo{
					DownloadUrl: *dlUrl,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Download(tt.args.client, tt.args.info, tt.args.filename...); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	t.Logf("cleanup")
	_ = os.RemoveAll(userProfile + "\\AppData\\Local\\Temp\\*.jar")
}
