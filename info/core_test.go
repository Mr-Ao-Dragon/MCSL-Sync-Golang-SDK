package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"errors"
	"reflect"
	"testing"
)

func TestCoreInfo_getCoreInfo(t *testing.T) {
	type fields struct {
		Name             string
		SupportMcVersion []string
		HistoryVersion   map[string]CoreVersionInfo
	}
	type args struct {
		client setup.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "Arclight",
			args: args{
				client: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "Arclight",
				},
			},
			want: []error{errors.New("arclight is not supported")},
		},
		{
			name: "Spigot",
			args: args{
				client: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "Spigot",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &CoreInfo{
				Name:             tt.fields.Name,
				SupportMcVersion: tt.fields.SupportMcVersion,
				HistoryVersion:   tt.fields.HistoryVersion,
			}
			if got := receiver.getCoreInfo(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCoreInfo() = %v, want %v", got, tt.want)
			}
			t.Logf("%#v", receiver)
		})
	}
}
