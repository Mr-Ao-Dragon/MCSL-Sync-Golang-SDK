package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"reflect"
	"testing"
)

func TestCoreInfo_GetTargetBuildInfo(t *testing.T) {
	type fields struct {
		Name             string
		SupportMcVersion []string
		HistoryVersion   map[string]CoreVersionInfo
	}
	type args struct {
		setupData   setup.Client
		targetBuild string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "Vanilla",
			args: args{
				setupData: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "vanilla",
					IsLatest:  true,
					MCVersion: "1.20.4",
				},
				targetBuild: "Latest",
			},
		},
		{
			name: "Forge",
			args: args{
				setupData: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "Forge",
					MCVersion: "1.21",
				},
				targetBuild: "51.0.8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &CoreInfo{
				Name:             tt.fields.Name,
				SupportMcVersion: tt.fields.SupportMcVersion,
				HistoryVersion:   tt.fields.HistoryVersion,
			}
			if got := receiver.GetTargetBuildInfo(tt.args.setupData, tt.args.targetBuild); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTargetBuildInfo() = %v, want %v", got, tt.want)
			}
			t.Logf("result: %#v", receiver)
		})
	}
}
