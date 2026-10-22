package main

const (
	cmdEcho      = "ECHO"
	cmdTime      = "TIME"
	cmdClose     = "CLOSE"
	cmdQuit      = "QUIT"
	cmdBroadcast = "BROADCAST"
	cmdMsg       = "MSG"
)

const (
	welcomeMsg        = "[SERVER] welcome to the dispatcher gateway"
	farewellMsg       = "[SERVER] cya later"
	echoFormat        = "[ECHO]: %s"
	timeFormat        = "[TIME]: %s"
	unknownFormat     = "[ERROR] unknown command %q, available commands: %s"
	missingTextFormat = "[ERROR] %s needs some text after it"
	broadcastFormat   = "[MESSAGE from %s]: %s"
	joinedFormat      = "[SYSTEM] new client hopped in: %s"
	leftFormat        = "[SYSTEM] client dipped: %s"
)

const (
	bannerFormat       = "== dispatcher gateway (tcp: %s) =="
	listeningFormat    = "[SERVER] tcp listener is up on port %s"
	listenFailedFormat = "[SERVER] cant open the port: %v"
	acceptFailedFormat = "[SERVER] accept failed: %v"
	connectedFormat    = "[CONNECTION] new client connected: %s"
	disconnectedFormat = "[CONNECTION] client left: %s"
	sendFailedFormat   = "[SERVER] failed to write to %s: %v"
	broadcastLogFormat = "[BROADCAST] from %s -> %q (delivered to %d)"
)
