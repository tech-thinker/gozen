package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/models"
)

func Test_projectHelper_InitProject(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(
				tt.fields.systemRepo,
			)

			if err := h.InitProject(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.InitProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupEnv(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupEnv(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupDocker(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupDocker(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupDocker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupUtils(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupUtils(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupUtils() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupRepository(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupRepository(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupService(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupService(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupModel(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupModel(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupLogger(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupLogger(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupConfig(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupConfig(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_SetupConstants(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("Write", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.SetupConstants(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.SetupConstants() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_CreateRestAPI(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.CreateRestAPI(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.CreateRestAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_projectHelper_CreategRPCAPI(t *testing.T) {
	type fields struct {
		systemRepo *repository.MockSystemRepo
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
			name:   "if WriteAll success should success",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			name:   "if WriteAll fail should fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.systemRepo.On("WriteAll", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(`error`))
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
				systemRepo: repository.NewMockSystemRepo(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			h := NewProjectHelper(tt.fields.systemRepo)

			if err := h.CreategRPCAPI(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("projectHelper.CreategRPCAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
