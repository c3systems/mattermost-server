// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

type MessageExport struct {
	TeamId          *int
	TeamName        *string
	TeamDisplayName *string

	ChannelId          *int
	ChannelName        *string
	ChannelDisplayName *string
	ChannelType        *string

	UserId    *int
	UserEmail *string
	Username  *string

	PostId         *int
	PostCreateAt   *int64
	PostMessage    *string
	PostType       *string
	PostRootId     *int
	PostOriginalId *int
	PostFileIds    StringArray
}
