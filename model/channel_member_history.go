// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

type ChannelMemberHistory struct {
	ChannelId int
	UserId    int
	JoinTime  int64
	LeaveTime *int64
}
