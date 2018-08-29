// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

// Registers a given function to be called when the cluster leader may have changed. Returns a unique ID for the
// listener which can later be used to remove it. If clustering is not enabled in this build, the callback will never
// be called.
func (a *App) AddClusterLeaderChangedListener(listener func()) int {
	// note: fix this!
	id := 0
	a.clusterLeaderListeners.Store(id, listener)
	return id
}

// Removes a listener function by the unique ID returned when AddConfigListener was called
func (a *App) RemoveClusterLeaderChangedListener(id int) {
	a.clusterLeaderListeners.Delete(id)
}

func (a *App) InvokeClusterLeaderChangedListeners() {
	a.Go(func() {
		a.clusterLeaderListeners.Range(func(_, listener interface{}) bool {
			listener.(func())()
			return true
		})
	})
}
