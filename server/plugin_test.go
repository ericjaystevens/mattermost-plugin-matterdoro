package main

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/mock"
)

func TestCreateBotDMPost(t *testing.T) {

	p := new(Plugin)
	api := &plugintest.API{}

	testChannel := &model.Channel{Id: "12345"}
	api.On("GetDirectChannel", mock.Anything, mock.Anything).Return(testChannel, nil)
	api.On("CreatePost", mock.Anything).Return(nil, nil)

	p.SetAPI(api)

	message := "Done"
	userID := "12345"

	got := p.CreateBotDMPost(userID, message)

	if got != nil {
		t.Errorf("got: %v\n wanted: %v\n", got, nil)
	}

}
