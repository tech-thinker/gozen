package models

import (
	"os"
	"testing"
)

func TestProject_Validate(t *testing.T) {
	tests := []struct {
		name    string
		project Project
		wantErr bool
	}{
		{
			name:    "with package name should succeed",
			project: Project{PackageName: "github.com/test/app"},
			wantErr: false,
		},
		{
			name:    "without package name should fail",
			project: Project{PackageName: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.project.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil {
				if err.Error() != "Please provide package name." {
					t.Errorf("Project.Validate() error message = %v, want 'Please provide package name.'", err.Error())
				}
			}
		})
	}
}

func TestProject_AutoFixes(t *testing.T) {
	tests := []struct {
		name     string
		project  Project
		wantName string
	}{
		{
			name:     "with app name should keep it",
			project:  Project{AppName: "myapp", PackageName: "github.com/test/other"},
			wantName: "myapp",
		},
		{
			name:     "without app name should extract from package",
			project:  Project{AppName: "", PackageName: "github.com/test/myapp"},
			wantName: "myapp",
		},
		{
			name:     "without app name with complex package",
			project:  Project{AppName: "", PackageName: "github.com/user/project-name"},
			wantName: "project-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.project
			p.AutoFixes()
			if p.AppName != tt.wantName {
				t.Errorf("Project.AutoFixes() AppName = %v, want %v", p.AppName, tt.wantName)
			}
		})
	}
}

func TestProject_GetCWD(t *testing.T) {
	tests := []struct {
		name     string
		project  Project
		wantCWD  string
	}{
		{
			name:     "with empty working dir should return current directory",
			project:  Project{WorkingDir: ""},
			wantCWD:  ".",
		},
		{
			name:     "with single char working dir should return current directory",
			project:  Project{WorkingDir: "."},
			wantCWD:  ".",
		},
		{
			name:     "with custom working dir should return it",
			project:  Project{WorkingDir: "/custom/path"},
			wantCWD:  "/custom/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cwd := tt.project.GetCWD()
			if cwd != tt.wantCWD {
				t.Errorf("Project.GetCWD() = %v, want %v", cwd, tt.wantCWD)
			}
		})
	}
}

func TestProject_GetAppDir(t *testing.T) {
	tests := []struct {
		name     string
		project  Project
		wantDir  string
	}{
		{
			name:     "with default working dir",
			project:  Project{AppName: "myapp", WorkingDir: ""},
			wantDir:  "./myapp",
		},
		{
			name:     "with custom working dir",
			project:  Project{AppName: "myapp", WorkingDir: "/custom/path"},
			wantDir:  "/custom/path/myapp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.project.GetAppDir()
			if dir != tt.wantDir {
				t.Errorf("Project.GetAppDir() = %v, want %v", dir, tt.wantDir)
			}
		})
	}
}

func TestProject_ToJSON(t *testing.T) {
	tests := []struct {
		name    string
		project Project
		want    string
	}{
		{
			name:    "valid project should return json",
			project: Project{AppName: "test", PackageName: "github.com/test/app", Driver: "postgres"},
			want:    `{"app_name":"test","package_name":"github.com/test/app","driver":"postgres"}`,
		},
		{
			name:    "empty project should return empty json object",
			project: Project{},
			want:    `{"app_name":"","package_name":"","driver":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonStr := tt.project.ToJSON()
			if jsonStr != tt.want {
				t.Errorf("Project.ToJSON() = %v, want %v", jsonStr, tt.want)
			}
		})
	}
}

func TestProject_ToJSON_Error(t *testing.T) {
	// Test error path - this is hard to trigger but we can verify the function handles it
	// by ensuring it returns "{}" on error (though json.Marshal rarely fails for structs)
	project := Project{AppName: "test", PackageName: "github.com/test/app", Driver: "postgres"}
	jsonStr := project.ToJSON()
	if jsonStr == "{}" {
		t.Error("ToJSON should not return empty object for valid project")
	}
}

func TestProject_FromJSON(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    Project
		wantErr bool
	}{
		{
			name:    "valid json should parse",
			jsonStr: `{"app_name":"test","package_name":"github.com/test/app","driver":"postgres"}`,
			want:    Project{AppName: "test", PackageName: "github.com/test/app", Driver: "postgres"},
			wantErr: false,
		},
		{
			name:    "invalid json should fail",
			jsonStr: `{invalid}`,
			want:    Project{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Project
			err := p.FromJSON(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.FromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if p.AppName != tt.want.AppName || p.PackageName != tt.want.PackageName || p.Driver != tt.want.Driver {
					t.Errorf("Project.FromJSON() = %+v, want %+v", p, tt.want)
				}
			}
		})
	}
}

func TestProject_LoadFromJsonFile(t *testing.T) {
	// Create a temporary file for testing
	tmpFile := "./test_gozen.json"
	defer os.Remove(tmpFile)

	// Write test data
	testJSON := `{"app_name":"test","package_name":"github.com/test/app","driver":"postgres"}`
	err := os.WriteFile(tmpFile, []byte(testJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Save original constant and modify for test
	// We can't modify constants, so we'll test with the current directory
	// The LoadFromJsonFile uses constants.CURRENT_WORKING_DIRECTORY which is "."
	// So we need to write to ./gozen.json

	testFile := "gozen.json"
	defer os.Remove(testFile)
	err = os.WriteFile(testFile, []byte(testJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	tests := []struct {
		name    string
		project *Project
		want    Project
		wantErr bool
	}{
		{
			name:    "load existing file should succeed",
			project: &Project{},
			want:    Project{AppName: "test", PackageName: "github.com/test/app", Driver: "postgres"},
			wantErr: false,
		},
		{
			name:    "load non-existent file should fail",
			project: &Project{},
			want:    Project{},
			wantErr: true,
		},
		{
			name:    "load invalid json should fail",
			project: &Project{},
			want:    Project{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "load non-existent file should fail" {
				os.Remove(testFile)
			} else if tt.name == "load invalid json should fail" {
				err = os.WriteFile(testFile, []byte(`{invalid json`), 0644)
				if err != nil {
					t.Fatalf("Failed to write test file: %v", err)
				}
			} else {
				err = os.WriteFile(testFile, []byte(testJSON), 0644)
				if err != nil {
					t.Fatalf("Failed to write test file: %v", err)
				}
			}
			err := tt.project.LoadFromJsonFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.LoadFromJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.project.AppName != tt.want.AppName || tt.project.PackageName != tt.want.PackageName || tt.project.Driver != tt.want.Driver {
					t.Errorf("Project.LoadFromJsonFile() = %+v, want %+v", tt.project, tt.want)
				}
			}
		})
	}
}

func TestProject_WriteToJsonFile(t *testing.T) {
	tests := []struct {
		name     string
		project  Project
		setup    func()
		wantErr  bool
		cleanup  func()
	}{
		{
			name:    "write to valid path should succeed",
			project: Project{AppName: "testapp", PackageName: "github.com/test/app", Driver: "postgres", WorkingDir: "."},
			setup:   func() { os.MkdirAll("./testapp", os.ModePerm) },
			wantErr: false,
			cleanup: func() { os.RemoveAll("./testapp") },
		},
		{
			name:     "write with empty app name creates file in working dir",
			project:  Project{AppName: "", PackageName: "github.com/test/app", Driver: "postgres", WorkingDir: "."},
			setup:    func() {},
			wantErr:  false, // Implementation doesn't validate empty app name
			cleanup:  func() { os.Remove("./gozen.json") },
		},
		{
			name:     "write to invalid path should fail",
			project:  Project{AppName: "testapp", PackageName: "github.com/test/app", Driver: "postgres", WorkingDir: "/invalid/path/that/does/not/exist"},
			setup:    func() {},
			wantErr:  true,
			cleanup:  func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			err := tt.project.WriteToJsonFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.WriteToJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}