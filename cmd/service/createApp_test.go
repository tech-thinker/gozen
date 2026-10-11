package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/gozen/mocks"
	"github.com/tech-thinker/gozen/models"
)

func Test_appService_CreateApp(t *testing.T) {
	type fields struct {
		projectRepo   *mocks.MockProjectRepo
		projectHelper *mocks.MockProjectHelper
		codeHelper    *mocks.MockCodeHelper
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
	}{
		{
			name:   "if create project and helper success then should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(nil)
				f.projectHelper.On("SetupProfiler", mock.Anything).Return(nil)
				f.projectHelper.On("SetupRepository", mock.Anything).Return(nil)
				f.projectHelper.On("SetupService", mock.Anything).Return(nil)
				f.projectHelper.On("SetupUtils", mock.Anything).Return(nil)
				f.projectHelper.On("CreateRestAPI", mock.Anything).Return(nil)
				f.projectHelper.On("CreategRPCAPI", mock.Anything).Return(nil)
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: false,
		},
		{
			name:   "if create fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if InitProject fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupEnv fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupDocker fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupConfig fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupConstants fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupLogger fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupModel fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupRepository fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(nil)
				f.projectHelper.On("SetupProfiler", mock.Anything).Return(nil)
				f.projectHelper.On("SetupRepository", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupService fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(nil)
				f.projectHelper.On("SetupProfiler", mock.Anything).Return(nil)
				f.projectHelper.On("SetupRepository", mock.Anything).Return(nil)
				f.projectHelper.On("SetupService", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if SetupUtils fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(nil)
				f.projectHelper.On("SetupProfiler", mock.Anything).Return(nil)
				f.projectHelper.On("SetupRepository", mock.Anything).Return(nil)
				f.projectHelper.On("SetupService", mock.Anything).Return(nil)
				f.projectHelper.On("SetupUtils", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if CreateRestAPI fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(nil)
				f.projectHelper.On("SetupProfiler", mock.Anything).Return(nil)
				f.projectHelper.On("SetupRepository", mock.Anything).Return(nil)
				f.projectHelper.On("SetupService", mock.Anything).Return(nil)
				f.projectHelper.On("SetupUtils", mock.Anything).Return(nil)
				f.projectHelper.On("CreateRestAPI", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
		{
			name:   "if CreategRPCAPI fail then should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.projectRepo.On("Create", mock.Anything).Return(nil)
				f.projectHelper.On("InitProject", mock.Anything).Return(nil)
				f.projectHelper.On("SetupEnv", mock.Anything).Return(nil)
				f.projectHelper.On("SetupDocker", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConfig", mock.Anything).Return(nil)
				f.projectHelper.On("SetupConstants", mock.Anything).Return(nil)
				f.projectHelper.On("SetupLogger", mock.Anything).Return(nil)
				f.projectHelper.On("SetupModel", mock.Anything).Return(nil)
				f.projectHelper.On("SetupProfiler", mock.Anything).Return(nil)
				f.projectHelper.On("SetupRepository", mock.Anything).Return(nil)
				f.projectHelper.On("SetupService", mock.Anything).Return(nil)
				f.projectHelper.On("SetupUtils", mock.Anything).Return(nil)
				f.projectHelper.On("CreateRestAPI", mock.Anything).Return(nil)
				f.projectHelper.On("CreategRPCAPI", mock.Anything).Return(errors.New(`error`))
			},
			args: args{
				project: models.Project{
					AppName:     "test",
					PackageName: "test",
					Driver:      "postgres",
					WorkingDir:  ".",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields = fields{
				projectRepo:   mocks.NewMockProjectRepo(t),
				projectHelper: mocks.NewMockProjectHelper(t),
				codeHelper:    mocks.NewMockCodeHelper(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			cmd := NewAppService(
				tt.fields.projectRepo,
				tt.fields.projectHelper,
				tt.fields.codeHelper,
			)
			if err := cmd.CreateApp(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("appService.CreateApp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_appService_GenerateModel(t *testing.T) {
	type fields struct {
		projectRepo   *mocks.MockProjectRepo
		projectHelper *mocks.MockProjectHelper
		codeHelper    *mocks.MockCodeHelper
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
			name: "success should call codeHelper.GenerateModel",
			fields: fields{},
			prepare: func(f *fields) {
				f.codeHelper.On("GenerateModel", mock.Anything).Return(nil)
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
			name: "codeHelper.GenerateModel fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.codeHelper.On("GenerateModel", mock.Anything).Return(errors.New("generate error"))
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
				projectRepo:   mocks.NewMockProjectRepo(t),
				projectHelper: mocks.NewMockProjectHelper(t),
				codeHelper:    mocks.NewMockCodeHelper(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			cmd := NewAppService(
				tt.fields.projectRepo,
				tt.fields.projectHelper,
				tt.fields.codeHelper,
			)
			if err := cmd.GenerateModel(tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("appService.GenerateModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}