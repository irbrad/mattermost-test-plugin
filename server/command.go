package main

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "irbot",
		DisplayName:      "irbot",
		Description:      "Test plugin.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: test1, test2",
		AutoCompleteHint: "[command]",
	}
}

func (p *Plugin) postCommandResponse(args *model.CommandArgs, text string) {
	post := &model.Post{
		UserId:    p.BotUserID,
		ChannelId: args.ChannelId,
		Message:   text,
	}
	_ = p.API.SendEphemeralPost(args.UserId, post)
}

func (p *Plugin) ExecuteBaseCommand(action string, parameters []string, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	if action == "" {
		p.postCommandResponse(args, "irbot_base_command, no actions")
		return &model.CommandResponse{}, nil
	}

	switch action {
	case "test1":
		p.postCommandResponse(args, "irbot_base_command, action: test1")
		break
	case "test2":
		p.postCommandResponse(args, "irbot_base_command, action: test2")
		break
	default:
		p.postCommandResponse(args, "irbot_base_command, action: default")
		break
	}

	return &model.CommandResponse{}, nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	p.Helpers.RegisterSlashCommand(args, "irbot", p.ExecuteBaseCommand)
	return &model.CommandResponse{}, nil
}
