package main

import (
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
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

func (p *Plugin) OnActivate() error {
	p.API.RegisterCommand(getCommand())

	botId, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    "irbot",
		DisplayName: "irbot",
		Description: "Created by the com.irbrad.test-mattermost-plugin plugin.",
	})
	if err != nil {
		return errors.Wrap(err, "failed to ensure bot")
	}
	p.BotUserID = botId

	return nil
}
