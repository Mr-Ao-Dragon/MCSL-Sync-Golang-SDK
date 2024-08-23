package util

import "testing"

func TestJavaVersionCatch(t *testing.T) {
	type args struct {
		mcVersion string
	}
	tests := []struct {
		name            string
		args            args
		wantJavaVersion int
	}{
		{
			name: "1.20",
			args: args{
				mcVersion: "1.20",
			},
			wantJavaVersion: 17,
		},
		{
			name: "1.19",
			args: args{
				mcVersion: "1.19.2",
			},
			wantJavaVersion: 17,
		},
		{
			name: "1.18",
			args: args{
				mcVersion: "1.18.1",
			},
			wantJavaVersion: 17,
		},
		{
			name: "1.17",
			args: args{
				mcVersion: "1.17.2",
			},
			wantJavaVersion: 17,
		},
		{
			name: "1.16",
			args: args{
				mcVersion: "1.16.5",
			},
			wantJavaVersion: 15,
		},
		{
			name: "1.15",
			args: args{
				mcVersion: "1.15.1",
			},
			wantJavaVersion: 15,
		},
		{
			name: "1.14",
			args: args{
				mcVersion: "1.14.1",
			},
			wantJavaVersion: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJavaVersion := JavaVersionCatch(tt.args.mcVersion); gotJavaVersion != tt.wantJavaVersion {
				t.Errorf("JavaVersionCatch() = %v, want %v", gotJavaVersion, tt.wantJavaVersion)
			}
		})
	}
}
