package utils

import (
	"embed"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	type objData struct {
		Name string `json:"Name"`
	}

	type args struct {
		tplFS   embed.FS
		tplFile string
		data    interface{}
	}

	tests := []struct {
		name    string
		args    args
		prepare func(f args)
		want    string
		wantErr bool
	}{
		{
			name: "with valid template should pass",
			args: args{
				tplFS:   GetMockTmpl(),
				tplFile: "test.tmpl",
				data:    objData{Name: "Test"},
			},
			want:    "Hello Test!",
			wantErr: false,
		},
		{
			name: "with invalid template should fail",
			args: args{
				tplFS:   GetMockTmpl(),
				tplFile: "invalid.tmpl",
				data:    objData{Name: "Test"},
			},
			want:    "{}",
			wantErr: true,
		},
		{
			name: "with template generate issue should fail",
			args: args{
				tplFS:   GetMockTmpl(),
				tplFile: "test.tmpl",
				data:    []byte{},
			},
			want:    "{}",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.prepare != nil {
				tt.prepare(tt.args)
			}

			got, err := GenerateCode(tt.args.tplFS, tt.args.tplFile, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyEscapeChar(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "text with escape char entities",
			args: args{
				data: "&lt;Hello World &amp; Golang is best!&gt;",
			},
			want: "<Hello World & Golang is best!>",
		},
		{
			name: "text with out escape char entities",
			args: args{
				data: "Hello World & Golang is best!",
			},
			want: "Hello World & Golang is best!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyEscapeChar(tt.args.data); got != tt.want {
				t.Errorf("ApplyEscapeChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
