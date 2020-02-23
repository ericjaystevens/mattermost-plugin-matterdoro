package main

import (
	"strings"
	"testing"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/mock"
)

type spyTimer struct {
	Calls int
}

func (s *spyTimer) AfterFunc(t time.Duration, f func()) *time.Timer {
	s.Calls++
	return time.NewTimer(1 * time.Millisecond)
}

func TestGetCommand(t *testing.T) {
	got := getCommand()
	want := &model.Command{
		Trigger: "doro",
	}

	if got.Trigger != want.Trigger {
		t.Errorf("got %s but wanted %s", got.Trigger, want.Trigger)
	}
}

func TestGetHelp(t *testing.T) {
	got := getHelp()
	want := &model.CommandResponse{
		Text: "###### Matterdoro - Slash Command Help",
	}

	if !strings.Contains(got.Text, want.Text) {
		t.Errorf("got %s but wanted it to start with %s", got.Text, want.Text)
	}
}

func TestStartPomodoro(t *testing.T) {
	timer := &spyTimer{}
	p := &Plugin{}
	api := &plugintest.API{}
	sessionID := "1234567"
	p.SetAPI(api)

	mockedSession := &model.Session{Id: "1234567", UserId: "erico"}
	api.On("GetSession", mock.Anything).Return(mockedSession)

	got := p.startPomodoro(timer, sessionID)
	want := &model.CommandResponse{
		Text: "start working!",
	}

	if got.Text != want.Text {
		t.Errorf("got %s but wanted %s", got.Text, want.Text)
	}

}
