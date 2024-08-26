package info

import (
	"MCSL-Sync-Golang-SDK/setup"
	"errors"
	"reflect"
	"testing"
)

func TestCoreList_ReadCoreList(t *testing.T) {
	type args struct {
		setupData setup.Client
	}
	tests := []struct {
		name     string
		receiver CoreList
		args     args
		wantErrs []error
	}{
		// TODO: Add test cases.
		{
			name:     "ReadCoreList",
			receiver: CoreList{},
			args: args{
				setupData: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
				},
			},
			wantErrs: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrs := tt.receiver.ReadCoreList(tt.args.setupData); !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("ReadCoreList() = %v, want %v", gotErrs, tt.wantErrs)
			}
			t.Logf("result: %#v", tt.receiver)
		})
	}
}

func TestCoreInfo_GetCoreSupportMcList(t *testing.T) {
	type fields struct {
		Name             string
		SupportMcVersion []string
		HistoryVersion   map[string]CoreVersionInfo
	}
	type args struct {
		setupData setup.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "GetCoreSupportMcList",
			args: args{
				setupData: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "Mohist",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &CoreInfo{}
			if got := receiver.GetCoreSupportMcList(tt.args.setupData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoreSupportMcList() = %v, want %v", got, tt.want)
			} else if len(receiver.SupportMcVersion) == 0 {
				t.Errorf("enpty result")
			}
			t.Logf("result: %#v", receiver.SupportMcVersion)
		})
	}
}

func TestCoreInfo_GetCoreBuildListSingleMCVersion(t *testing.T) {
	type fields struct {
		Name             string
		SupportMcVersion []string
		HistoryVersion   map[string]CoreVersionInfo
	}
	type args struct {
		setupData setup.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "ok input",
			args: args{
				setupData: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "Mohist",
					MCVersion: "1.20.1",
				},
			},
			want: nil,
		},
		{
			name: "not support mc version",
			args: args{
				setupData: setup.Client{
					ApiDomain: "sync.mcsl.com.cn",
					CoreName:  "Mohist",
					MCVersion: "1.30.2",
				},
			},
			want: []error{
				errors.New("Error: No data were found."),
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
			if got := receiver.GetCoreBuildListSingleMCVersion(tt.args.setupData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoreBuildListSingleMCVersion() = %v, want %v", got, tt.want)
			}
			t.Logf("result: %#v", receiver.HistoryVersion)

		})
	}
}
