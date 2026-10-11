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

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single word lowercase", "user", "User"},
		{"single word uppercase", "USER", "USER"},
		{"snake_case", "user_name", "UserName"},
		{"multiple underscores", "my_user_name", "MyUserName"},
		{"with spaces", "my user name", "MyUserName"},
		{"with hyphens", "my-user-name", "MyUserName"},
		{"mixed separators", "my_user-name test", "MyUserNameTest"},
		{"already camelCase", "myUserName", "MyUserName"},
		{"with numbers", "user2_name", "User2Name"},
		{"leading underscore", "_user_name", "UserName"},
		{"trailing underscore", "user_name_", "UserName"},
		{"only underscores", "___", ""},
		{"single char", "a", "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCamelCase(tt.in)
			if got != tt.want {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single word lowercase", "user", "user"},
		{"single word uppercase", "USER", "uSER"},
		{"camelCase", "UserName", "userName"},
		{"PascalCase", "UserName", "userName"},
		{"already snake_case", "user_name", "user_name"},
		{"with numbers", "User2Name", "user2Name"},
		{"single char", "A", "a"},
		{"single char lowercase", "a", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSnakeCase(tt.in)
			if got != tt.want {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestTransformString(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		wantCamel  string
		wantSnake  string
	}{
		{"empty string", "", "", ""},
		{"camelCase input", "userName", "UserName", "userName"},
		{"PascalCase input", "UserName", "UserName", "userName"},
		{"snake_case input (no uppercase)", "user_name", "UserName", "userName"},
		{"multiple words snake_case (no uppercase)", "my_user_name", "MyUserName", "myUserName"},
		{"single word (no uppercase)", "user", "User", "user"},
		{"UPPERCASE (has uppercase)", "USER", "USER", "uSER"},
		{"with numbers (has uppercase)", "user2Name", "User2Name", "user2Name"},
		{"with hyphens (no uppercase)", "user-name", "UserName", "userName"},
		{"with spaces (no uppercase)", "user name", "UserName", "userName"},
		{"mixed case (has uppercase)", "myUserName", "MyUserName", "myUserName"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCamel, gotSnake := TransformString(tt.in)
			if gotCamel != tt.wantCamel {
				t.Errorf("TransformString(%q) camel = %q, want %q", tt.in, gotCamel, tt.wantCamel)
			}
			if gotSnake != tt.wantSnake {
				t.Errorf("TransformString(%q) snake = %q, want %q", tt.in, gotSnake, tt.wantSnake)
			}
		})
	}
}
