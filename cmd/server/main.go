package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	wishtea "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"

	"github.com/masaking-alt/portfolio-ssh/internal/portfolio"
	"github.com/masaking-alt/portfolio-ssh/internal/ui"
)

func main() {
	configureTerminalRenderer()

	cfg := configFromEnv()

	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(cfg.host, cfg.port)),
		cfg.hostKeyOption(),
		wish.WithMiddleware(
			wishtea.Middleware(func(session ssh.Session) (tea.Model, []tea.ProgramOption) {
				return ui.NewModel(portfolio.DefaultProfile()), []tea.ProgramOption{
					tea.WithAltScreen(),
				}
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("サーバの初期化に失敗しました: %v", err)
	}

	go func() {
		log.Printf("SSHポートフォリオを %s で待ち受けています", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Fatalf("サーバが異常終了しました: %v", err)
		}
	}()

	waitForShutdown(server)
}

func configureTerminalRenderer() {
	lipgloss.SetColorProfile(termenv.TrueColor)
}

type config struct {
	host        string
	port        string
	hostKeyPath string
	hostKeyPEM  string
}

func configFromEnv() config {
	return config{
		host:        envOrDefault("HOST", "0.0.0.0"),
		port:        envOrDefault("PORT", "2222"),
		hostKeyPath: envOrDefault("HOST_KEY_PATH", ".ssh/ssh_host_ed25519_key"),
		hostKeyPEM:  os.Getenv("SSH_HOST_KEY"),
	}
}

func (c config) hostKeyOption() ssh.Option {
	if c.hostKeyPEM != "" {
		return wish.WithHostKeyPEM([]byte(c.hostKeyPEM))
	}
	return wish.WithHostKeyPath(c.hostKeyPath)
}

func envOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func waitForShutdown(server *ssh.Server) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("終了シグナルを受け取りました。SSHサーバを停止します。")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("SSHサーバの停止中にエラーが発生しました: %v", err)
	}
}
