package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type commandHandler func(c *Client, args string) bool

type command struct {
	name      string
	needsText bool
	handler   commandHandler
}

func (s *Server) registerCommands() {
	table := []command{
		{name: cmdEcho, needsText: true, handler: s.handleEcho},
		{name: cmdTime, handler: s.handleTime},
		{name: cmdClose, handler: s.handleClose},
		{name: cmdQuit, handler: s.handleClose},
	}

	s.commands = make(map[string]command, len(table))
	names := make([]string, 0, len(table))
	for _, cmd := range table {
		s.commands[cmd.name] = cmd
		names = append(names, cmd.name)
	}
	s.commandHint = strings.Join(names, ", ")
}

func (s *Server) dispatch(c *Client, line string) bool {
	if line == "" {
		return true
	}

	name, args := splitCommand(line)
	cmd, ok := s.commands[strings.ToUpper(name)]
	if !ok {
		return c.Send(fmt.Sprintf(unknownFormat, name, s.commandHint)) == nil
	}
	if cmd.needsText && args == "" {
		return c.Send(fmt.Sprintf(missingTextFormat, cmd.name)) == nil
	}
	return cmd.handler(c, args)
}

func splitCommand(line string) (string, string) {
	idx := strings.IndexFunc(line, unicode.IsSpace)
	if idx < 0 {
		return line, ""
	}
	return line[:idx], strings.TrimSpace(line[idx:])
}

func (s *Server) handleEcho(c *Client, args string) bool {
	return c.Send(fmt.Sprintf(echoFormat, args)) == nil
}

func (s *Server) handleTime(c *Client, _ string) bool {
	return c.Send(fmt.Sprintf(timeFormat, time.Now().Format(time.RFC3339))) == nil
}

func (s *Server) handleClose(c *Client, _ string) bool {
	_ = c.Send(farewellMsg)
	return false
}
