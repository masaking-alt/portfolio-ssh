package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/masaking-alt/portfolio-ssh/internal/portfolio"
)

type screen int

const (
	screenHome screen = iota
	screenWorks
	screenWorkDetail
	screenAbout
	screenContact
	screenHelp
)

type Model struct {
	profile       portfolio.Profile
	activeScreen  screen
	history       []screen
	width         int
	height        int
	homeCursor    int
	workCursor    int
	selectedIndex int
	viewport      viewport.Model
	ready         bool
}

type menuItem struct {
	key         string
	label       string
	description string
	target      screen
}

var homeMenu = []menuItem{
	{key: "1", label: "Works", description: "制作物の一覧", target: screenWorks},
	{key: "2", label: "About", description: "自己紹介", target: screenAbout},
	{key: "3", label: "Contact", description: "連絡先とSNS", target: screenContact},
	{key: "4", label: "Help", description: "操作キー", target: screenHelp},
	{key: "q", label: "Quit", description: "接続終了", target: screenHome},
}

const (
	maxWorkRowWidth = 96
	workCategoryGap = 2
)

var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#9FE7D7"))

	asciiNameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9FE7D7"))

	asciiFaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4FB3A5"))

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C6D8D3"))

	sectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F4D35E"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#111827")).
			Background(lipgloss.Color("#9FE7D7")).
			Padding(0, 1)

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8FA7A1"))

	bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F2F7F5"))

	borderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#325D63"))
)

func NewModel(profile portfolio.Profile) Model {
	return Model{
		profile:       profile,
		activeScreen:  screenHome,
		selectedIndex: -1,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Masaking SSH Portfolio")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateViewport(false)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "backspace", "left":
			m.back()
			return m, nil
		case "?", "h":
			m.push(screenHelp)
			return m, nil
		}

		switch m.activeScreen {
		case screenHome:
			return m.updateHome(msg)
		case screenWorks:
			return m.updateWorks(msg)
		case screenWorkDetail, screenAbout, screenContact, screenHelp:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	}

	if m.isScrollable() {
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return "準備中..."
	}

	switch m.activeScreen {
	case screenHome:
		return appStyle.Render(m.viewHome())
	case screenWorks:
		return appStyle.Render(m.viewWorks())
	case screenWorkDetail, screenAbout, screenContact, screenHelp:
		return appStyle.Render(m.viewScrollable())
	default:
		return appStyle.Render(m.viewHome())
	}
}

func (m Model) updateHome(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		m.homeCursor = moveCursor(m.homeCursor, -1, len(homeMenu))
	case "down", "j":
		m.homeCursor = moveCursor(m.homeCursor, 1, len(homeMenu))
	case "enter", "right":
		item := homeMenu[m.homeCursor]
		if item.key == "q" {
			return m, tea.Quit
		}
		m.push(item.target)
	case "1", "2", "3", "4":
		for _, item := range homeMenu {
			if item.key == msg.String() {
				m.push(item.target)
				break
			}
		}
	}
	return m, nil
}

func (m Model) updateWorks(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		m.workCursor = moveCursor(m.workCursor, -1, len(m.profile.Works))
	case "down", "j":
		m.workCursor = moveCursor(m.workCursor, 1, len(m.profile.Works))
	case "enter", "right":
		m.selectedIndex = m.workCursor
		m.push(screenWorkDetail)
	}
	return m, nil
}

func (m *Model) push(next screen) {
	if next == m.activeScreen {
		return
	}

	m.history = append(m.history, m.activeScreen)
	m.activeScreen = next
	m.updateViewport(true)
}

func (m *Model) back() {
	if len(m.history) == 0 {
		return
	}

	last := len(m.history) - 1
	m.activeScreen = m.history[last]
	m.history = m.history[:last]
	m.updateViewport(true)
}

func (m *Model) updateViewport(reset bool) {
	if !m.isScrollable() {
		return
	}

	contentWidth := maxInt(32, m.width-6)
	contentHeight := maxInt(5, m.height-8)

	if !m.ready {
		m.viewport = viewport.New(contentWidth, contentHeight)
		m.ready = true
	} else {
		m.viewport.Width = contentWidth
		m.viewport.Height = contentHeight
	}

	m.viewport.SetContent(m.scrollableContent(contentWidth))
	if reset {
		m.viewport.GotoTop()
	}
}

func (m Model) isScrollable() bool {
	return m.activeScreen == screenWorkDetail ||
		m.activeScreen == screenAbout ||
		m.activeScreen == screenContact ||
		m.activeScreen == screenHelp
}

func (m Model) viewHome() string {
	if m.homeColumnsFit() {
		return m.viewHomeColumns()
	}
	return m.viewHomeStacked()
}

func (m Model) viewHomeStacked() string {
	var builder strings.Builder
	builder.WriteString(m.homeHeader())
	builder.WriteString("\n")
	builder.WriteString(subtitleStyle.Render(m.profile.Title))
	builder.WriteString("\n\n")
	builder.WriteString(bodyStyle.Render(strings.Join(m.profile.HeroLines, "\n")))
	builder.WriteString("\n\n")

	builder.WriteString(m.homeMenu())
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString(m.footer())
	return builder.String()
}

func (m Model) viewHomeColumns() string {
	availableWidth := maxInt(0, m.width-4)
	faceWidth := maxLineWidth(asciiMyFace)
	rightWidth := availableWidth - faceWidth - 4

	face := asciiFaceStyle.Render(strings.Join(asciiMyFace, "\n"))
	right := lipgloss.NewStyle().
		Width(rightWidth).
		MarginLeft(4).
		MarginTop(2).
		Render(m.homeRightColumn(rightWidth))

	layout := lipgloss.JoinHorizontal(lipgloss.Top, face, right)
	return strings.Join([]string{layout, "", m.footer()}, "\n")
}

func (m Model) homeColumnsFit() bool {
	availableWidth := maxInt(0, m.width-4)
	availableHeight := maxInt(0, m.height-2)
	neededWidth := maxLineWidth(asciiMyFace) + 4 + maxLineWidth(asciiMyName)
	neededHeight := len(asciiMyFace) + 2

	return availableWidth >= neededWidth && availableHeight >= neededHeight
}

func (m Model) homeRightColumn(width int) string {
	var builder strings.Builder
	builder.WriteString(asciiNameStyle.Render(strings.Join(asciiMyName, "\n")))
	builder.WriteString("\n\n")
	builder.WriteString(subtitleStyle.Render(m.profile.Title))
	builder.WriteString("\n\n")
	builder.WriteString(bodyStyle.Render(strings.Join(m.profile.HeroLines, "\n")))
	builder.WriteString("\n\n")
	builder.WriteString(mutedStyle.Width(width).Render(m.profile.AboutLead))
	builder.WriteString("\n\n")
	builder.WriteString(m.homeMenu())
	return builder.String()
}

func (m Model) homeHeader() string {
	availableWidth := maxInt(0, m.width-4)
	availableHeight := maxInt(0, m.height-2)
	headerReservedHeight := 11
	parts := []string{}

	if asciiFits(asciiMyName, availableWidth, availableHeight-headerReservedHeight) {
		parts = append(parts, asciiNameStyle.Render(strings.Join(asciiMyName, "\n")))
		availableHeight -= len(asciiMyName) + 1
	} else {
		parts = append(parts, titleStyle.Render("MASAKING PORTFOLIO"))
		availableHeight -= 2
	}

	if asciiFits(asciiMyFace, availableWidth, availableHeight-headerReservedHeight) {
		parts = append(parts, asciiFaceStyle.Render(strings.Join(asciiMyFace, "\n")))
	}

	return strings.Join(parts, "\n")
}

func (m Model) homeMenu() string {
	var builder strings.Builder
	for i, item := range homeMenu {
		line := fmt.Sprintf("[%s] %-7s %s", item.key, item.label, mutedStyle.Render(item.description))
		if i == m.homeCursor {
			line = selectedStyle.Render(line)
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return strings.TrimRight(builder.String(), "\n")
}

func (m Model) viewWorks() string {
	var builder strings.Builder
	total := len(m.profile.Works)
	start, end := visibleWorkRange(m.workCursor, total, m.height)
	rowWidth := workListRowWidth(m.width)

	builder.WriteString(sectionStyle.Render("WORKS"))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(fmt.Sprintf("Enterで詳細を表示します。%d-%d / %d", start+1, end, total)))
	builder.WriteString("\n\n")

	for i := start; i < end; i++ {
		work := m.profile.Works[i]
		titleLine := workTitleLine(work, rowWidth)
		if i == m.workCursor {
			titleLine = selectedStyle.Render(titleLine)
		} else {
			titleLine = bodyStyle.Render(titleLine)
		}

		builder.WriteString(titleLine)
		builder.WriteString("\n")
		builder.WriteString("   ")
		builder.WriteString(mutedStyle.Render(shorten(work.Description, 82)))
		builder.WriteString("\n")
		builder.WriteString("   ")
		builder.WriteString(mutedStyle.Render(strings.Join(firstStrings(work.Technologies, 3), " / ")))
		builder.WriteString("\n\n")
	}

	builder.WriteString(m.footer())
	return builder.String()
}

func (m Model) viewScrollable() string {
	var builder strings.Builder
	builder.WriteString(m.scrollableTitle())
	builder.WriteString("\n")
	builder.WriteString(borderStyle.Render(strings.Repeat("─", maxInt(10, minInt(m.width-4, 72)))))
	builder.WriteString("\n")
	builder.WriteString(m.viewport.View())
	builder.WriteString("\n")
	builder.WriteString(m.footer())
	return builder.String()
}

func (m Model) scrollableTitle() string {
	switch m.activeScreen {
	case screenWorkDetail:
		work := m.selectedWork()
		return sectionStyle.Render(work.Title)
	case screenAbout:
		return sectionStyle.Render("ABOUT")
	case screenContact:
		return sectionStyle.Render("CONTACT")
	case screenHelp:
		return sectionStyle.Render("HELP")
	default:
		return sectionStyle.Render("PORTFOLIO")
	}
}

func (m Model) scrollableContent(width int) string {
	switch m.activeScreen {
	case screenWorkDetail:
		return m.workDetailContent(width)
	case screenAbout:
		return m.aboutContent(width)
	case screenContact:
		return m.contactContent(width)
	case screenHelp:
		return m.helpContent(width)
	default:
		return ""
	}
}

func (m Model) workDetailContent(width int) string {
	work := m.selectedWork()
	lines := []string{
		bodyStyle.Render("種別: " + work.Category),
		bodyStyle.Render("URL: " + work.ExternalURL),
		bodyStyle.Render("画像: " + work.ImagePath),
		"",
		sectionStyle.Render("説明"),
		wrap(work.Description, width),
		"",
		sectionStyle.Render("使用技術"),
		wrap(strings.Join(work.Technologies, " / "), width),
	}
	return strings.Join(lines, "\n")
}

func (m Model) aboutContent(width int) string {
	lines := []string{
		wrap(m.profile.AboutLead, width),
		"",
	}
	for _, paragraph := range m.profile.AboutBody {
		lines = append(lines, wrap(paragraph, width), "")
	}

	lines = append(lines, sectionStyle.Render("Built with"))
	lines = append(lines, wrap(m.profile.BuiltWithIntro, width))
	lines = append(lines, "")
	for _, tech := range m.profile.BuiltWith {
		lines = append(lines, bodyStyle.Render("- "+tech))
	}

	return strings.Join(lines, "\n")
}

func (m Model) contactContent(width int) string {
	lines := []string{
		wrap(m.profile.ContactLead, width),
		"",
		bodyStyle.Render("メール: " + m.profile.Email),
		"",
		sectionStyle.Render("SNS"),
	}

	for _, social := range m.profile.Socials {
		lines = append(lines, bodyStyle.Render("- "+social.Name+": "+social.URL))
	}

	return strings.Join(lines, "\n")
}

func (m Model) helpContent(width int) string {
	lines := []string{
		wrap("このSSHポートフォリオは、通常のシェルではなく専用のTUIアプリです。", width),
		"",
		bodyStyle.Render("↑ / ↓ または k / j: 選択移動、スクロール"),
		bodyStyle.Render("← / Esc / Backspace: 戻る"),
		bodyStyle.Render("→ / Enter: 決定、詳細表示"),
		bodyStyle.Render("? / h: ヘルプ"),
		bodyStyle.Render("q / Ctrl+C: 終了"),
	}
	return strings.Join(lines, "\n")
}

func (m Model) selectedWork() portfolio.Work {
	if m.selectedIndex >= 0 && m.selectedIndex < len(m.profile.Works) {
		return m.profile.Works[m.selectedIndex]
	}
	if len(m.profile.Works) == 0 {
		return portfolio.Work{}
	}
	return m.profile.Works[0]
}

func (m Model) footer() string {
	return mutedStyle.Render("↑/↓ 選択  Enter 決定  Esc 戻る  ? ヘルプ  q 終了")
}

func moveCursor(current int, delta int, total int) int {
	if total <= 0 {
		return 0
	}
	next := current + delta
	if next < 0 {
		return total - 1
	}
	if next >= total {
		return 0
	}
	return next
}

func visibleWorkRange(cursor int, total int, height int) (int, int) {
	if total <= 0 {
		return 0, 0
	}

	visible := maxInt(1, (height-8)/4)
	if visible > total {
		visible = total
	}

	start := cursor - visible/2
	if start < 0 {
		start = 0
	}
	if start+visible > total {
		start = total - visible
	}

	return start, start + visible
}

func firstStrings(values []string, count int) []string {
	if len(values) < count {
		return values
	}
	return values[:count]
}

func workListRowWidth(windowWidth int) int {
	return minInt(maxWorkRowWidth, maxInt(0, windowWidth-4))
}

func workTitleLine(work portfolio.Work, width int) string {
	left := fmt.Sprintf("%2d. %s", work.ID, work.Title)
	category := work.Category
	if width <= 0 {
		return left + " " + category
	}

	categoryWidth := lipgloss.Width(category)
	leftWidth := width - categoryWidth - workCategoryGap
	if leftWidth <= 0 {
		return fitDisplayWidth(left, width)
	}

	left = fitDisplayWidth(left, leftWidth)
	spaces := width - lipgloss.Width(left) - categoryWidth
	if spaces < workCategoryGap {
		spaces = workCategoryGap
	}

	return left + strings.Repeat(" ", spaces) + category
}

func fitDisplayWidth(value string, width int) string {
	if width <= 0 {
		return ""
	}
	if lipgloss.Width(value) <= width {
		return value
	}
	if width <= 3 {
		return trimDisplayWidth(value, width)
	}
	return trimDisplayWidth(value, width-3) + "..."
}

func trimDisplayWidth(value string, width int) string {
	var builder strings.Builder
	currentWidth := 0
	for _, r := range value {
		runeWidth := lipgloss.Width(string(r))
		if currentWidth+runeWidth > width {
			break
		}
		builder.WriteRune(r)
		currentWidth += runeWidth
	}
	return builder.String()
}

func shorten(value string, limit int) string {
	if len([]rune(value)) <= limit {
		return value
	}
	runes := []rune(value)
	return string(runes[:limit]) + "..."
}

func wrap(value string, width int) string {
	if width <= 0 {
		return bodyStyle.Render(value)
	}
	return bodyStyle.Width(width).Render(value)
}

func asciiFits(lines []string, width int, height int) bool {
	if len(lines) == 0 {
		return false
	}
	return width >= maxLineWidth(lines) && height >= len(lines)
}

func maxLineWidth(lines []string) int {
	width := 0
	for _, line := range lines {
		width = maxInt(width, lipgloss.Width(line))
	}
	return width
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
