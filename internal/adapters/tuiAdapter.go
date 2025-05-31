package adapters

// Backend API adapter for Bubbletea, the Update() function for a model should re-subscribe to the ListenEvents Cmd to get new updates from the backend asyncronously
// Backend will handle message parsing from user

import (
	"log"
	"quiteline/internal/backend"

	tea "github.com/charmbracelet/bubbletea"
)

type TUIBackendAdapter struct {
	backend *backend.Backend
}

func NewTUIBackendAdapter() *TUIBackendAdapter {
	return &TUIBackendAdapter{
		backend: backend.NewBackend(),
	}
}

func (a *TUIBackendAdapter) ListenEvents() tea.Cmd {
	return func() tea.Msg {
		evt := <-a.backend.Events
		return evt // convert Event to tea.Msg as needed
	}
}

func (a *TUIBackendAdapter) SendMessageCmd(msg string) tea.Cmd {
	log.Println("TUIADAPTER: SendMessageCmd sent")
	return func() tea.Msg {
		a.backend.SendMessage(msg)
		return nil
	}
}
