// Copyright 2009 Bart de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"github.com/bartdeboer/cobra"
)

type Command struct {
	Command   *cobra.Command
	GetConfig func(cmd *Command, args []string) interface{}
	parent    *Command
	commands  []*Command
}

func (c *Command) GetCommand() *cobra.Command {
	return c.Command
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
