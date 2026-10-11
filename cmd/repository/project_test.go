package repository

import (
	"errors"
	"os"
	"testing"

	"github.com/tech-thinker/gozen/mocks"
	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/wrappers"
)

func Test_projectRepo_Create(t *testing.T) {
	type fields struct {
		fileSystemWrapper wrappers.FileSystemWrapper
	}
	type args struct {
		project models.Project
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		args    args
		wantErr bool
		cleanup func()
	}{
		{
			name: "success should create directory and write config",
			fields: fields{},
			prepare: func(f *fields) {
				f.fileSystemWrapper = wrappers.NewFileSystemWrapper()
			},
			args: args{
				project: models.Project{
					AppName:     "testapp",
					PackageName: "github.com/test/app",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: false,
			cleanup: func() {
				os.RemoveAll("./testapp")
			},
		},
		{
			name: "create directory fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				mockWrapper := &mockFileSystemWrapperFail{}
				f.fileSystemWrapper = mockWrapper
			},
			args: args{
				project: models.Project{
					AppName:     "testapp",
					PackageName: "github.com/test/app",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
			cleanup: func() {},
		},
		{
			name: "write json config fail should fail (using mock that fails on WriteFile)",
			fields: fields{},
			prepare: func(f *fields) {
				mockWrapper := &mockFileSystemWrapperWriteFail{}
				f.fileSystemWrapper = mockWrapper
			},
			args: args{
				project: models.Project{
					AppName:     "testapp",
					PackageName: "github.com/test/app",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
			cleanup: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields = fields{}
			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			repo := NewProjectRepo(tt.fields.fileSystemWrapper)
			err := repo.Create(tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestNewProjectRepo(t *testing.T) {
	wrapper := mocks.NewMockFileSystemWrapper(t)
	repo := NewProjectRepo(wrapper)
	if repo == nil {
		t.Error("NewProjectRepo() returned nil")
	}
}

// mockFileSystemWrapperFail is a mock that fails on CreateDirectory
type mockFileSystemWrapperFail struct{}

func (m *mockFileSystemWrapperFail) CreateDirectory(path string) error {
	return errors.New("permission denied")
}

func (m *mockFileSystemWrapperFail) Destroy(path string) error {
	return nil
}

func (m *mockFileSystemWrapperFail) WriteFile(path, data string) error {
	return nil
}

// mockFileSystemWrapperWriteFail is a mock that fails on WriteFile
type mockFileSystemWrapperWriteFail struct{}

func (m *mockFileSystemWrapperWriteFail) CreateDirectory(path string) error {
	return nil
}

func (m *mockFileSystemWrapperWriteFail) Destroy(path string) error {
	return nil
}

func (m *mockFileSystemWrapperWriteFail) WriteFile(path, data string) error {
	return errors.New("write failed")
}