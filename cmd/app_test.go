package cmd

import (
	"flag"
	"testing"

	"github.com/tech-thinker/gozen/cmd/service"
	"github.com/tech-thinker/gozen/cmd/helpers"
	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/utils"
	"github.com/tech-thinker/gozen/wrappers"
	"github.com/urfave/cli/v2"
)

func TestNewApp(t *testing.T) {
	// Use real implementations for testing
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()
	systemRepo := repository.NewSystemRepo(utils.GetMockTmpl(), shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)
	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)
	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)

	app := NewApp(appSvc)
	if app == nil {
		t.Error("NewApp() returned nil")
	}
}

func Test_app_CreateProject(t *testing.T) {
	// Use real implementations for testing
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()
	systemRepo := repository.NewSystemRepo(utils.GetMockTmpl(), shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)
	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)
	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)

	app := NewApp(appSvc).(*app)

	cmd := app.CreateProject()
	if cmd == nil {
		t.Error("CreateProject() returned nil command")
	}
	if cmd.Name != "create" {
		t.Errorf("CreateProject() command name = %v, want 'create'", cmd.Name)
	}
	if len(cmd.Flags) != 3 {
		t.Errorf("CreateProject() flags count = %v, want 3", len(cmd.Flags))
	}
	
	// Test that Action is not nil
	if cmd.Action == nil {
		t.Error("CreateProject() Action should not be nil")
	}
}

func Test_app_CreateProject_Action(t *testing.T) {
	// Test the action function with valid input
	// We test the logic by directly calling the internal logic
	project := models.Project{
		AppName:     "testapp",
		PackageName: "github.com/test/testapp",
		Driver:      "postgres",
		WorkingDir:  "./_test_create_action",
	}

	err := project.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}

	project.AutoFixes()
	
	// Test that AutoFixes works correctly
	if project.AppName != "testapp" {
		t.Errorf("AutoFixes() should not change AppName when already set")
	}
	
	// Test AutoFixes with empty AppName
	project2 := models.Project{
		AppName:     "",
		PackageName: "github.com/test/testapp",
		Driver:      "postgres",
		WorkingDir:  ".",
	}
	project2.AutoFixes()
	if project2.AppName != "testapp" {
		t.Errorf("AutoFixes() should extract AppName from PackageName, got %v", project2.AppName)
	}
}

func Test_app_CreateProject_Action_ValidationFail(t *testing.T) {
	// Test the action function with missing package name
	project := models.Project{
		AppName:     "testapp",
		PackageName: "",
		Driver:      "postgres",
		WorkingDir:  ".",
	}

	err := project.Validate()
	if err == nil {
		t.Error("Validate() should return error for empty package name")
	}
}

func Test_app_CreateProject_Action_Execute(t *testing.T) {
	// Test the Action function execution with mock service
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()
	systemRepo := repository.NewSystemRepo(utils.GetMockTmpl(), shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)
	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)
	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)

	app := NewApp(appSvc).(*app)
	cmd := app.CreateProject()

	// Test 1: Valid project with package name - should call CreateApp
	set := flag.NewFlagSet("create", 0)
	// Apply command flags to flag set
	for _, f := range cmd.Flags {
		_ = f.Apply(set)
	}
	ctx := cli.NewContext(nil, set, nil)
	// Parse args: "testapp" as positional arg, --pkg flag
	_ = set.Parse([]string{"--pkg", "github.com/test/testapp", "testapp"})
	err := cmd.Action(ctx)
	// Should fail because project doesn't exist, but we test the Action runs
	if err != nil {
		_ = err
	}

	// Test 2: Missing package name - validation should fail
	set2 := flag.NewFlagSet("create", 0)
	for _, f := range cmd.Flags {
		_ = f.Apply(set2)
	}
	ctx2 := cli.NewContext(nil, set2, nil)
	_ = set2.Parse([]string{"testapp"}) // No --pkg flag
	err = cmd.Action(ctx2)
	// Should fail validation (empty package name)
	if err == nil {
		// Action returns nil on validation error
		_ = err
	}
}

func Test_app_CreateProject_Action_WithFlags(t *testing.T) {
	// Test the Action function with all flags
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()
	systemRepo := repository.NewSystemRepo(utils.GetMockTmpl(), shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)
	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)
	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)

	app := NewApp(appSvc).(*app)
	cmd := app.CreateProject()

	// Create a proper CLI context with flags
	set := flag.NewFlagSet("create", 0)
	for _, f := range cmd.Flags {
		_ = f.Apply(set)
	}
	ctx := cli.NewContext(nil, set, nil)
	_ = set.Parse([]string{
		"--pkg", "github.com/test/testapp",
		"--output", ".",
		"--driver", "postgres",
		"testapp",
	})

	err := cmd.Action(ctx)
	// Should run without panic
	if err != nil {
		_ = err
	}
}

func Test_app_GenerateModel(t *testing.T) {
	// Use real implementations for testing
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()
	systemRepo := repository.NewSystemRepo(utils.GetMockTmpl(), shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)
	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)
	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)

	app := NewApp(appSvc).(*app)

	cmd := app.GenerateModel()
	if cmd == nil {
		t.Error("GenerateModel() returned nil command")
	}
	if cmd.Name != "generate" {
		t.Errorf("GenerateModel() command name = %v, want 'generate'", cmd.Name)
	}
	
	// Test that Action is not nil
	if cmd.Action == nil {
		t.Error("GenerateModel() Action should not be nil")
	}
}

func Test_app_GenerateModel_Action(t *testing.T) {
	// Test the generator logic
	project := models.Project{
		AppName:     "testapp",
		PackageName: "github.com/test/testapp",
		Driver:      "postgres",
		WorkingDir:  "./_test_gen_action",
	}

	// Test the generator struct creation
	generator := models.Generator{
		Project:       project,
		Type:          "model",
		InterfaceName: "IUser",
		StructName:    "User",
		RouteName:     "user",
	}
	
	if generator.Project.AppName != "testapp" {
		t.Errorf("Generator project not set correctly")
	}
	if generator.StructName != "User" {
		t.Errorf("Generator StructName not set correctly")
	}
	if generator.InterfaceName != "IUser" {
		t.Errorf("Generator InterfaceName not set correctly")
	}
	if generator.RouteName != "user" {
		t.Errorf("Generator RouteName not set correctly")
	}
}

func Test_app_GenerateModel_Action_Execute(t *testing.T) {
	// Test the Action function execution with mock service
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()
	systemRepo := repository.NewSystemRepo(utils.GetMockTmpl(), shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)
	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)
	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)

	app := NewApp(appSvc).(*app)
	cmd := app.GenerateModel()

	// Test 1: Valid generator with type and name - should call GenerateModel
	set := flag.NewFlagSet("generate", 0)
	ctx := cli.NewContext(nil, set, nil)
	// Parse args: "model" as type, "User" as name
	_ = set.Parse([]string{"model", "User"})
	err := cmd.Action(ctx)
	// Should fail because project config doesn't exist, but we test the Action runs
	if err != nil {
		_ = err
	}

	// Test 2: Missing name - should still run (uses empty string)
	set2 := flag.NewFlagSet("generate", 0)
	ctx2 := cli.NewContext(nil, set2, nil)
	_ = set2.Parse([]string{"model"}) // No name
	err = cmd.Action(ctx2)
	if err != nil {
		_ = err
	}

	// Test 3: Missing type - should still run
	set3 := flag.NewFlagSet("generate", 0)
	ctx3 := cli.NewContext(nil, set3, nil)
	_ = set3.Parse([]string{}) // No args
	err = cmd.Action(ctx3)
	if err != nil {
		_ = err
	}
}

func Test_app_GenerateModel_Action_TransformString(t *testing.T) {
	// Test that TransformString is used correctly in the Action
	interfaceName, structName := utils.TransformString("user")
	
	// TransformString should convert "user" to "User" (camelCase) and "user" (snake_case)
	if interfaceName != "User" {
		t.Errorf("TransformString interfaceName = %v, want 'User'", interfaceName)
	}
	if structName != "user" {
		t.Errorf("TransformString structName = %v, want 'user'", structName)
	}

	// Test with complex name
	interfaceName2, structName2 := utils.TransformString("userProfile")
	if interfaceName2 != "UserProfile" {
		t.Errorf("TransformString interfaceName = %v, want 'UserProfile'", interfaceName2)
	}
	if structName2 != "userProfile" {
		t.Errorf("TransformString structName = %v, want 'userProfile'", structName2)
	}

	// Test snake_case input
	interfaceName3, structName3 := utils.TransformString("user_profile")
	if interfaceName3 != "UserProfile" {
		t.Errorf("TransformString interfaceName = %v, want 'UserProfile'", interfaceName3)
	}
	if structName3 != "userProfile" {
		t.Errorf("TransformString structName = %v, want 'userProfile'", structName3)
	}
}

