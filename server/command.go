package main

import (
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const commandHelp = `###### Matterdoro - Slash Command Help
* |/doro start| - Starts a timer.
* |/doro help| - Dispays Matterdoro help 
`

const pomodoroSeconds = 1500

// TimerAble is an interface for the Timer method
type TimerAble interface {
	AfterFunc(time.Duration, func()) *time.Timer
}

type defaultTimer struct{}

// AfterFunc calls the default time AfterFunc method
func (d *defaultTimer) AfterFunc(t time.Duration, f func()) *time.Timer {
	return time.AfterFunc(t, f)
}

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "doro",
		DisplayName:      "Matterdoro",
		Description:      "Pomodor Plugin for focus and task tracking.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: start",
		AutoCompleteHint: "[command]",
	}
}

//ExecuteCommand runs parses /doro commands
func (p *Plugin) ExecuteCommand(cont *plugin.Context, args *model.CommandArgs) (response *model.CommandResponse, modelErr *model.AppError) {

	split := strings.Fields(args.Command)

	if len(split) == 1 {
		response = getHelp()
	} else {
		action := split[1]
		if action == "start" {
			timer := &defaultTimer{}
			response = p.startPomodoro(timer, cont.SessionId)
		}
	}

	return response, modelErr
}

func getHelp() (response *model.CommandResponse) {
	response = &model.CommandResponse{
		Text: strings.Replace(commandHelp, "|", "`", -1),
	}
	return response
}

// startPomodoro starts the timer
func (p *Plugin) startPomodoro(timer TimerAble, sessionID string) (response *model.CommandResponse) {
	response = &model.CommandResponse{
		Text: "start working!",
	}

	// create a delayed completion message from bot
	timer.AfterFunc(time.Second*pomodoroSeconds, func() {
		ses, _ := p.API.GetSession(sessionID)
		userID := ses.UserId
		_ = p.CreateBotDMPost(string(userID), "Time is up, take a break!")
	})

	return response
}
