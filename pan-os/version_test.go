package version_test

import (
	"reflect"
	"testing"

	version "github.com/MaineK00n/go-paloalto-version/pan-os"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    version.Version
		wantErr bool
	}{
		{
			name: "10.0",
			args: args{
				ver: "10.0",
			},
			wantErr: true,
		},
		{
			name: "10.0.1",
			args: args{
				ver: "10.0.1",
			},
			want: version.Version{
				Major:       10,
				Minor:       0,
				Maintenance: 1,
				Hotfix:      nil,
			},
		},
		{
			name: "10.0.a",
			args: args{
				ver: "10.0.a",
			},
			wantErr: true,
		},
		{
			name: "10.0.0-h1",
			args: args{
				ver: "10.0.0-h1",
			},
			want: version.Version{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      toPtr(1),
			},
		},
		{
			name: "10.0.0-f1",
			args: args{
				ver: "10.0.0-f1",
			},
			wantErr: true,
		},
		{
			name: "10.0.0-1",
			args: args{
				ver: "10.0.0-1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := version.NewVersion(tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	type fields struct {
		Major       int
		Minor       int
		Maintenance int
		Hotfix      *int
	}
	type args struct {
		v2 version.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "10.0.0 = 10.0.0",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      nil,
			},
			args: args{
				v2: version.Version{
					Major:       10,
					Minor:       0,
					Maintenance: 0,
					Hotfix:      nil,
				},
			},
			want: 0,
		},
		{
			name: "10.0.0-h1 = 10.0.0-h1",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      toPtr(1),
			},
			args: args{
				v2: version.Version{
					Major:       10,
					Minor:       0,
					Maintenance: 0,
					Hotfix:      toPtr(1),
				},
			},
			want: 0,
		},
		{
			name: "10.0.0 < 10.0.1",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      nil,
			},
			args: args{
				v2: version.Version{
					Major:       10,
					Minor:       0,
					Maintenance: 1,
					Hotfix:      nil,
				},
			},
			want: -1,
		},
		{
			name: "10.0.0 < 10.0.0-h1",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      nil,
			},
			args: args{
				v2: version.Version{
					Major:       10,
					Minor:       0,
					Maintenance: 0,
					Hotfix:      toPtr(1),
				},
			},
			want: -1,
		},
		{
			name: "10.0.0-h2 > 10.0.0-h1",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      toPtr(2),
			},
			args: args{
				v2: version.Version{
					Major:       10,
					Minor:       0,
					Maintenance: 0,
					Hotfix:      toPtr(1),
				},
			},
			want: +1,
		},
		{
			name: "10.0.1 > 10.0.0-h1",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 1,
				Hotfix:      nil,
			},
			args: args{
				v2: version.Version{
					Major:       10,
					Minor:       0,
					Maintenance: 0,
					Hotfix:      toPtr(1),
				},
			},
			want: +1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := version.Version{
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Maintenance: tt.fields.Maintenance,
				Hotfix:      tt.fields.Hotfix,
			}
			if got := v1.Compare(tt.args.v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		Major       int
		Minor       int
		Maintenance int
		Hotfix      *int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "10.0.0",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      nil,
			},
			want: "10.0.0",
		},
		{
			name: "10.0.0-h1",
			fields: fields{
				Major:       10,
				Minor:       0,
				Maintenance: 0,
				Hotfix:      toPtr(1),
			},
			want: "10.0.0-h1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := version.Version{
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Maintenance: tt.fields.Maintenance,
				Hotfix:      tt.fields.Hotfix,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toPtr[T any](v T) *T {
	return &v
}
