package helper

import (
	"embed"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/gozen/utils"
	"github.com/tech-thinker/gozen/wrappers"
)

func Test_commonHelper_Write(t *testing.T) {
	type fields struct {
		templatesFS       embed.FS
		shellWrapper      *wrappers.ShellWrapperMock
		fileSystemWrapper *wrappers.FileSystemWrapperMock
	}
	type args struct {
		templatePath string
		outputPath   string
		data         interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name:   "with valid data should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.fileSystemWrapper.On("CreateDirectory", mock.Anything).Return(nil)
				f.fileSystemWrapper.On("WriteFile", mock.Anything, mock.Anything).Return(nil)
			},
			args: args{
				templatePath: "test.tmpl",
				outputPath:   "./test.go",
				data:         struct{ Name string }{Name: "Test"},
			},
			wantErr: false,
		},
		{
			name:   "with directory creation fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.fileSystemWrapper.On("CreateDirectory", mock.Anything).Return(errors.New("error"))
			},
			args: args{
				templatePath: "test.tmpl",
				outputPath:   "./test.go",
				data:         struct{ Name string }{Name: "Test"},
			},
			wantErr: true,
		},
		{
			name:   "with file write fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.fileSystemWrapper.On("CreateDirectory", mock.Anything).Return(nil)
				f.fileSystemWrapper.On("WriteFile", mock.Anything, mock.Anything).Return(errors.New("error"))
			},
			args: args{
				templatePath: "test.tmpl",
				outputPath:   "./test.go",
				data:         struct{ Name string }{Name: "Test"},
			},
			wantErr: true,
		},
		{
			name:   "with code generation fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.fileSystemWrapper.On("CreateDirectory", mock.Anything).Return(nil)
			},
			args: args{
				templatePath: "test.tmpl",
				outputPath:   "./test.go",
				data:         struct{ name string }{name: "Test"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparing mocks
			tt.fields = fields{
				templatesFS:       utils.GetMockTmpl(),
				shellWrapper:      wrappers.NewShellWrapperMock(t),
				fileSystemWrapper: wrappers.NewFileSystemWrapperMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			helper := NewCommonHelper(
				tt.fields.templatesFS,
				tt.fields.shellWrapper,
				tt.fields.fileSystemWrapper,
			)
			if err := helper.Write(tt.args.templatePath, tt.args.outputPath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("commonHelper.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commonHelper_ExecShell(t *testing.T) {
	type fields struct {
		templatesFS       embed.FS
		shellWrapper      *wrappers.ShellWrapperMock
		fileSystemWrapper *wrappers.FileSystemWrapperMock
	}
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:   "with valid command should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.shellWrapper.On("Exec", mock.Anything).Return([]byte("test.txt"), nil)
			},
			args: args{
				command: "ls",
			},
			want:    []string{"test.txt"},
			wantErr: false,
		},
		{
			name:   "with invalid command should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.shellWrapper.On("Exec", mock.Anything).Return([]byte(""), errors.New("ls no such command found"))
			},
			args: args{
				command: "ls",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Preparing mocks
			tt.fields = fields{
				templatesFS:       utils.GetMockTmpl(),
				shellWrapper:      wrappers.NewShellWrapperMock(t),
				fileSystemWrapper: wrappers.NewFileSystemWrapperMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			helper := NewCommonHelper(
				tt.fields.templatesFS,
				tt.fields.shellWrapper,
				tt.fields.fileSystemWrapper,
			)
			got, err := helper.ExecShell(tt.args.command, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("commonHelper.ExecShell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonHelper.ExecShell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commonHelper_ExecShellRaw(t *testing.T) {
	type fields struct {
		templatesFS       embed.FS
		shellWrapper      *wrappers.ShellWrapperMock
		fileSystemWrapper *wrappers.FileSystemWrapperMock
	}
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "with valid command should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.shellWrapper.On("Exec", mock.Anything).Return([]byte("test.txt"), nil)
			},
			args: args{
				command: "ls",
			},
			want:    []byte("test.txt"),
			wantErr: false,
		},
		{
			name:   "with invalid command should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.shellWrapper.On("Exec", mock.Anything).Return([]byte(""), errors.New("ls no such command found"))
			},
			args: args{
				command: "ls",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparing mocks
			tt.fields = fields{
				templatesFS:       utils.GetMockTmpl(),
				shellWrapper:      wrappers.NewShellWrapperMock(t),
				fileSystemWrapper: wrappers.NewFileSystemWrapperMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			helper := NewCommonHelper(
				tt.fields.templatesFS,
				tt.fields.shellWrapper,
				tt.fields.fileSystemWrapper,
			)
			got, err := helper.ExecShellRaw(tt.args.command, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("commonHelper.ExecShellRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonHelper.ExecShellRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}
