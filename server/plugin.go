package main

import (
	"fmt"
	"sync"

	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	BotUserID string
}

//OnActivate is called when the matterdoro plugin is activated
func (p *Plugin) OnActivate() (err error) {
	if err = p.API.RegisterCommand(getCommand()); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to register command: %v", getCommand()))
	}

	//configure doro bot
	botID, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    "Doro",
		DisplayName: "Doro Yon Time",
		Description: "Created by the Matterdoro plugin.",
	})
	if err != nil {
		return errors.Wrap(err, "failed to ensure doro bot")
	}
	p.BotUserID = botID

	return
}

// CreateBotDMPost takes a destination userID and message and sends a direct
// message from doro.
func (p *Plugin) CreateBotDMPost(userID, message string) *model.AppError {
	channel, err := p.API.GetDirectChannel(userID, p.BotUserID)
	if err != nil {
		mlog.Error("Couldn't get bot's DM channel", mlog.String("user_id", userID))
		return err
	}

	post := &model.Post{
		UserId:    p.BotUserID,
		ChannelId: channel.Id,
		Message:   message,
	}

	if _, err := p.API.CreatePost(post); err != nil {
		mlog.Error(err.Error())
		return err
	}

	return nil
}
