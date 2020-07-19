// Copyright 2009 Bart de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"github.com/spf13/cobra"
)

type Command struct {
	Command   *cobra.Command
	GetConfig func(cmd *Command, args []string) interface{}
	ConfigKey string
	parent    *Command
	commands  []*Command
}

func (c *Command) AddCommand(cmds ...*Command) {
	for i, cc := range cmds {
		if cmds[i] == c {
			panic("Command can't be a child of itself")
		}
		cmds[i].parent = c
		// update max lengths
		c.commands = append(c.commands, cc)
		c.Command.AddCommand(cc.Command)
	}
}

func (c *Command) AddCobraCommand(cmds ...*cobra.Command) {
	for _, cc := range cmds {
		c.AddCommand(&Command{
			Command: cc,
		})
	}
}

func (c *Command) Parent() *Command {
	return c.parent
}

func (c *Command) Execute() error {
	return c.Command.Execute()
}
