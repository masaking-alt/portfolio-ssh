package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/masaking-alt/portfolio-ssh/internal/portfolio"
)

func TestViewContainsHomeNavigation(t *testing.T) {
	model := NewModel(portfolio.DefaultProfile())
	updated, _ := model.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	view := updated.(Model).View()

	for _, want := range []string{"MASAKING PORTFOLIO", "Works", "About", "Contact"} {
		if !strings.Contains(view, want) {
			t.Fatalf("表示に %q が含まれていません", want)
		}
	}
}

func TestEnterFromWorksShowsDetail(t *testing.T) {
	model := NewModel(portfolio.DefaultProfile())
	updated, _ := model.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	model = updated.(Model)

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	model = updated.(Model)
	if model.activeScreen != screenWorks {
		t.Fatalf("画面 = %v, want screenWorks", model.activeScreen)
	}

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	model = updated.(Model)
	if model.activeScreen != screenWorkDetail {
		t.Fatalf("画面 = %v, want screenWorkDetail", model.activeScreen)
	}

	if !strings.Contains(model.View(), "Marple") {
		t.Fatal("詳細画面に作品名が含まれていません")
	}
}
