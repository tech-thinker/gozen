package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/gozen/cmd/helpers"
	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/models"
)

func Test_appService_CreateApp(t *testing.T) {
	type fields struct {
		projectRepo   *repository.MockProjectRepo
		projectHelper *helpers.MockProjectHelper
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
				projectRepo:   repository.NewMockProjectRepo(t),
				projectHelper: helpers.NewMockProjectHelper(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			cmd := NewAppService(
				tt.fields.projectRepo,
				tt.fields.projectHelper,
			)
			if err := cmd.CreateApp(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("appService.CreateApp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
