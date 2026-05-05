package catalog

import "testing"

func TestLoadSharedCatalog(t *testing.T) {
	catalog, err := Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	if catalog.Version != "0.2.0" {
		t.Fatalf("version = %q, want 0.2.0", catalog.Version)
	}
	if len(catalog.Recipes) < 30 {
		t.Fatalf("recipes = %d, want at least 30", len(catalog.Recipes))
	}
}

func TestPlanExpandsStack(t *testing.T) {
	catalog, err := Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}

	plan, err := catalog.Plan("windows", []string{"template.fullstack-web"})
	if err != nil {
		t.Fatal(err)
	}

	ids := map[string]bool{}
	for _, action := range plan.Actions {
		ids[action.ID] = true
		if action.Command == "" {
			t.Fatalf("action %s has empty command", action.ID)
		}
		if action.Risk.Level == "" {
			t.Fatalf("action %s has empty risk level", action.ID)
		}
	}

	for _, id := range []string{"git", "node", "pnpm", "frontend.react", "frontend.vite", "python", "backend.fastapi"} {
		if !ids[id] {
			t.Fatalf("plan missing %s", id)
		}
	}
}

func TestAssessCommandRisk(t *testing.T) {
	tests := []struct {
		name  string
		cmd   string
		level string
	}{
		{
			name:  "package manager",
			cmd:   "winget install --id Git.Git -e",
			level: "low",
		},
		{
			name:  "language package",
			cmd:   "python -m pip install torch",
			level: "medium",
		},
		{
			name:  "remote execution",
			cmd:   "irm https://example.com/install.ps1 | iex",
			level: "high",
		},
		{
			name:  "destructive",
			cmd:   "Remove-Item -Recurse -Force C:\\tmp\\demo",
			level: "high",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			risk := AssessCommandRisk(tt.cmd)
			if risk.Level != tt.level {
				t.Fatalf("level = %q, want %q", risk.Level, tt.level)
			}
			if risk.Summary == "" || len(risk.Reasons) == 0 {
				t.Fatalf("incomplete risk = %#v", risk)
			}
		})
	}
}
