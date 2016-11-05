package portfolio

import (
	"testing"
)

func newProjectOrFatal(t *testing.T, title string) *Project {
	project, err := NewProject(title)

	if err != nil {
		t.Fatalf("new project: %v", err)
	}
	return project
}

func TestNewProject(t *testing.T) {
	title := "project1"
	project := newProjectOrFatal(t, title)
	if project.Title != title {
		t.Errorf("expected title %q, got %q", title, project.Title)
	}
}

func TestAll(t *testing.T) {
	project1 := newProjectOrFatal(t, "project1")
	project2 := newProjectOrFatal(t, "project1")

	m := NewProjectManager()
	m.Save(project1)
	m.Save(project2)

	all := m.All()
	if len(all) != 2 {
		t.Errorf("expected 2 projects, got %v", len(all))
	}
	if *all[0] != *project1 && *all[1] != *project1 {
		t.Errorf("missing project %v", project1)
	}
	if *all[0] != *project2 && *all[1] != *project2 {
		t.Errorf("missing project %v", project2)
	}
}
