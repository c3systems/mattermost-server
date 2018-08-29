// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type UserSearch struct {
	Term           string `json:"term"`
	TeamId         int    `json:"team_id"`
	NotInTeamId    int    `json:"not_in_team_id"`
	InChannelId    int    `json:"in_channel_id"`
	NotInChannelId int    `json:"not_in_channel_id"`
	AllowInactive  bool   `json:"allow_inactive"`
	WithoutTeam    bool   `json:"without_team"`
}

// ToJson convert a User to a json string
func (u *UserSearch) ToJson() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// UserSearchFromJson will decode the input and return a User
func UserSearchFromJson(data io.Reader) *UserSearch {
	var us *UserSearch
	json.NewDecoder(data).Decode(&us)
	return us
}
