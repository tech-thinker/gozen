package wrappers

import (
	"testing"
)

func Test_fileSystemWrapper_CreateDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create directory should success",
			args: args{
				path: "./_test_2A60D77E-EE78-46E6-B218-CB6877135016",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			repo := NewFileSystemWrapper()
			if err := repo.CreateDirectory(tt.args.path); (err != nil) != tt.wantErr {
				repo.Destroy(tt.args.path)
				t.Errorf("fileSystemWrapper.CreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
			repo.Destroy(tt.args.path)
		})
	}
}

func Test_fileSystemWrapper_Destroy(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		repo    *fileSystemWrapper
		args    args
		wantErr bool
	}{
		{
			name: "delete directory should success",
			args: args{
				path: "./_test_E6E4D318-D4F2-413E-B1EE-2C88138BB4D1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewFileSystemWrapper()
			repo.CreateDirectory(tt.args.path)
			if err := repo.Destroy(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("fileSystemWrapper.Destroy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystemWrapper_WriteFile(t *testing.T) {
	type args struct {
		path string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create file should success",
			args: args{
				path: "./_test_file_5C3D878D-3072-42D1-89F4-EDE74FF21A53.txt",
				data: "sample data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewFileSystemWrapper()
			if err := repo.WriteFile(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				repo.Destroy(tt.args.path)
				t.Errorf("fileSystemWrapper.WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			repo.Destroy(tt.args.path)
		})
	}
}
