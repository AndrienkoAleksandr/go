// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "intern/poll"

var (
	// Placeholders for saving original socket system calls.
	origSocket      = socketFunc
	origWSASocket   = wsaSocketFunc
	origClosesocket = poll.CloseFunc
	origConnect     = connectFunc
	origConnectEx   = poll.ConnectExFunc
	origListen      = listenFunc
	origAccept      = poll.AcceptFunc
)

func installTestHooks() {
	socketFunc = sw.Socket
	wsaSocketFunc = sw.WSASocket
	poll.CloseFunc = sw.Closesocket
	connectFunc = sw.Connect
	poll.ConnectExFunc = sw.ConnectEx
	listenFunc = sw.Listen
	poll.AcceptFunc = sw.AcceptEx
}

func uninstallTestHooks() {
	socketFunc = origSocket
	wsaSocketFunc = origWSASocket
	poll.CloseFunc = origClosesocket
	connectFunc = origConnect
	poll.ConnectExFunc = origConnectEx
	listenFunc = origListen
	poll.AcceptFunc = origAccept
}

// forceCloseSockets must be called only from TestMain.
func forceCloseSockets() {
	for s := range sw.Sockets() {
		poll.CloseFunc(s)
	}
}
