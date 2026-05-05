package catalog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Catalog struct {
	Version    string     `json:"version"`
	Categories []Category `json:"categories"`
	Recipes    []Recipe   `json:"recipes"`
	Stacks     []Stack    `json:"stacks"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Recipe struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Category string            `json:"category"`
	Detect   string            `json:"detect"`
	Verify   []string          `json:"verify"`
	Install  map[string]string `json:"install"`
	Trust    Trust             `json:"trust,omitempty"`
}

type Stack struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	RecipeIDs   []string `json:"recipeIds"`
}

type Action struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Command  string   `json:"command"`
	Verify   []string `json:"verify"`
	Risk     Risk     `json:"risk"`
	Trust    Trust    `json:"trust"`
}

type Plan struct {
	Platform string   `json:"platform"`
	ItemIDs  []string `json:"itemIds"`
	Actions  []Action `json:"actions"`
}

type Risk struct {
	Level   string   `json:"level"`
	Summary string   `json:"summary"`
	Reasons []string `json:"reasons"`
}

type Trust struct {
	Trusted bool   `json:"trusted"`
	Source  string `json:"source"`
	Reason  string `json:"reason"`
}

func Load(path string) (*Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	var catalog Catalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return nil, err
	}
	if catalog.Version == "" {
		return nil, errors.New("catalog version is required")
	}
	return &catalog, nil
}

func (c *Catalog) Plan(platform string, itemIDs []string) (*Plan, error) {
	if platform == "" {
		platform = "windows"
	}

	recipesByID := map[string]Recipe{}
	for _, recipe := range c.Recipes {
		recipesByID[recipe.ID] = recipe
	}

	stacksByID := map[string]Stack{}
	for _, stack := range c.Stacks {
		stacksByID[stack.ID] = stack
	}

	seen := map[string]bool{}
	actions := []Action{}
	for _, itemID := range itemIDs {
		recipeIDs := []string{itemID}
		if stack, ok := stacksByID[itemID]; ok {
			recipeIDs = stack.RecipeIDs
		}

		for _, recipeID := range recipeIDs {
			if seen[recipeID] {
				continue
			}
			recipe, ok := recipesByID[recipeID]
			if !ok {
				return nil, fmt.Errorf("unknown recipe %q from item %q", recipeID, itemID)
			}
			command := recipe.Install[platform]
			if command == "" {
				return nil, fmt.Errorf("recipe %q does not support platform %q", recipeID, platform)
			}
			seen[recipeID] = true
			actions = append(actions, Action{
				ID:       recipe.ID,
				Name:     recipe.Name,
				Category: recipe.Category,
				Command:  command,
				Verify:   recipe.Verify,
				Risk:     AssessCommandRisk(command),
				Trust:    normalizeTrust(recipe.Trust),
			})
		}
	}

	return &Plan{
		Platform: platform,
		ItemIDs:  append([]string(nil), itemIDs...),
		Actions:  actions,
	}, nil
}

func normalizeTrust(trust Trust) Trust {
	if trust.Source == "" && trust.Reason == "" {
		return Trust{
			Trusted: true,
			Source:  "builtin catalog",
			Reason:  "Bundled with Easy_Setup",
		}
	}
	if trust.Source == "" {
		trust.Source = "custom catalog"
	}
	if trust.Reason == "" {
		if trust.Trusted {
			trust.Reason = "Trusted by catalog metadata"
		} else {
			trust.Reason = "Not in trusted allowlist"
		}
	}
	return trust
}

func AssessCommandRisk(command string) Risk {
	normalized := string(bytes.ToLower([]byte(command)))
	reasons := []string{}
	level := "low"

	if containsAny(normalized, []string{"irm ", "iwr ", "invoke-webrequest", "invoke-restmethod", "curl ", "wget "}) &&
		containsAny(normalized, []string{"| iex", "invoke-expression", "bash", "sh"}) {
		level = "high"
		reasons = append(reasons, "downloads and executes remote script")
	}
	if containsAny(normalized, []string{"set-executionpolicy", "new-itemproperty", "set-itemproperty", "[environment]::setenvironmentvariable"}) {
		level = maxRisk(level, "medium")
		reasons = append(reasons, "changes system or user configuration")
	}
	if containsAny(normalized, []string{"winget install", "brew install", "apt install", "dnf install", "pacman -s", "choco install", "scoop install"}) {
		reasons = append(reasons, "installs packages through a package manager")
	}
	if containsAny(normalized, []string{"pip install", "npm install", "pnpm add", "cargo install", "go install"}) {
		level = maxRisk(level, "medium")
		reasons = append(reasons, "installs language ecosystem packages")
	}
	if containsAny(normalized, []string{"sudo ", "runas", "start-process powershell -verb runas"}) {
		level = maxRisk(level, "high")
		reasons = append(reasons, "may require elevated privileges")
	}
	if containsAny(normalized, []string{"rm -rf", "remove-item", "del /f", "format "}) {
		level = maxRisk(level, "high")
		reasons = append(reasons, "contains destructive file operation")
	}
	if len(reasons) == 0 {
		reasons = append(reasons, "runs a local command")
	}

	return Risk{
		Level:   level,
		Summary: riskSummary(level),
		Reasons: reasons,
	}
}

func containsAny(value string, needles []string) bool {
	for _, needle := range needles {
		if bytes.Contains([]byte(value), []byte(needle)) {
			return true
		}
	}
	return false
}

func maxRisk(current string, candidate string) string {
	if riskRank(candidate) > riskRank(current) {
		return candidate
	}
	return current
}

func riskRank(level string) int {
	switch level {
	case "high":
		return 3
	case "medium":
		return 2
	default:
		return 1
	}
}

func riskSummary(level string) string {
	switch level {
	case "high":
		return "Review carefully before running"
	case "medium":
		return "Review package and configuration changes"
	default:
		return "Low risk command"
	}
}
