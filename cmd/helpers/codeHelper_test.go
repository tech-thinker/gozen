package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/gozen/mocks"
	"github.com/tech-thinker/gozen/models"
)

func Test_codeHelper_GenerateModel(t *testing.T) {
	type fields struct {
		systemRepo *mocks.MockSystemRepo
	}
	type args struct {
		doc models.Generator
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "success should write all files",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			args: args{
				doc: models.Generator{
					Project: models.Project{
						AppName:     "testapp",
						PackageName: "github.com/test/app",
						Driver:      "postgres",
						WorkingDir:  ".",
					},
					Type:          "model",
					InterfaceName: "IUser",
					StructName:    "User",
					RouteName:     "user",
				},
			},
			wantErr: false,
		},
		{
			name: "write all fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("template error"))
			},
			args: args{
				doc: models.Generator{
					Project: models.Project{
						AppName:     "testapp",
						PackageName: "github.com/test/app",
						Driver:      "postgres",
						WorkingDir:  ".",
					},
					Type:          "model",
					InterfaceName: "IUser",
					StructName:    "User",
					RouteName:     "user",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields = fields{
				systemRepo: mocks.NewMockSystemRepo(t),
			}
			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewCodeHelper(tt.fields.systemRepo)
			err := h.GenerateModel(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("codeHelper.GenerateModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCodeHelper(t *testing.T) {
	systemRepo := mocks.NewMockSystemRepo(t)
	helper := NewCodeHelper(systemRepo)
	if helper == nil {
		t.Error("NewCodeHelper() returned nil")
	}
}