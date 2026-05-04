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
}

type Plan struct {
	Platform string   `json:"platform"`
	ItemIDs  []string `json:"itemIds"`
	Actions  []Action `json:"actions"`
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
			})
		}
	}

	return &Plan{
		Platform: platform,
		ItemIDs:  append([]string(nil), itemIDs...),
		Actions:  actions,
	}, nil
}
