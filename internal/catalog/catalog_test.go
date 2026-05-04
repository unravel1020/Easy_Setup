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
	}

	for _, id := range []string{"git", "node", "pnpm", "frontend.react", "frontend.vite", "python", "backend.fastapi"} {
		if !ids[id] {
			t.Fatalf("plan missing %s", id)
		}
	}
}
