package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/masaking-alt/portfolio-ssh/internal/portfolio"
)

func TestViewContainsHomeNavigation(t *testing.T) {
	model := NewModel(portfolio.DefaultProfile())
	updated, _ := model.Update(tea.WindowSizeMsg{Width: 40, Height: 24})
	view := updated.(Model).View()

	for _, want := range []string{"MASAKING PORTFOLIO", "Works", "About", "Contact"} {
		if !strings.Contains(view, want) {
			t.Fatalf("表示に %q が含まれていません", want)
		}
	}
}

func TestHomeHeaderUsesAsciiArtWhenItFits(t *testing.T) {
	model := NewModel(portfolio.DefaultProfile())
	model.width = maxLineWidth(asciiMyName) + 4
	model.height = len(asciiMyName) + 13

	view := model.View()
	if !strings.Contains(view, strings.TrimSpace(asciiMyName[0])) {
		t.Fatal("十分な表示領域がある場合に名前のASCIIアートが表示されていません")
	}
}

func TestHomeUsesTwoColumnLayoutWhenItFits(t *testing.T) {
	model := NewModel(portfolio.DefaultProfile())
	model.width = maxLineWidth(asciiMyFace) + 4 + maxLineWidth(asciiMyName) + 4
	model.height = len(asciiMyFace) + 4

	view := model.View()
	for _, want := range []string{asciiMyFace[0], strings.TrimSpace(asciiMyName[0]), "Works"} {
		if !strings.Contains(view, want) {
			t.Fatalf("2カラム表示に %q が含まれていません", want)
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
