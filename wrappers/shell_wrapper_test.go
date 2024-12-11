package wrappers

import (
	"testing"
)

func Test_shellWrapper_Exec(t *testing.T) {
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "with correct command should success",
			args: args{
				name: "pwd",
			},
			wantErr: false,
		},
		{
			name: "with correct command should success",
			args: args{
				name: "sample-command-for-test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShellWrapper()
			got, err := s.Exec(tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("shellWrapper.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 && !tt.wantErr {
				t.Errorf("shellWrapper.Exec() = %v, want non empty", got)
			}
		})
	}
}
