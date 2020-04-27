package main

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	args := strings.Split(d.TextBeforeCursor(), " ")

	var returnCommands []prompt.Suggest

	if len(args) <= 1 {
		returnCommands = commands
	} else if len(args) >= 2 {
		switch args[0] {
		case "config":
			returnCommands = configCommands
		case "genautocomplete":
			returnCommands = genAutoComplete
		}
		if strings.HasPrefix(args[len(args)-1], "--") {
			var tempArgs []prompt.Suggest
			switch args[0] {
			case "about":
				tempArgs = Args_rclone_about

			case "authorize":
				tempArgs = Args_rclone_authorize

			case "cachestats":
				tempArgs = Args_rclone_cachestats

			case "cat":
				tempArgs = Args_rclone_cat

			case "check":
				tempArgs = Args_rclone_check

			case "cleanup":
				tempArgs = Args_rclone_cleanup

			case "config":
				tempArgs = Args_rclone_config

			case "copy":
				tempArgs = Args_rclone_copy

			case "copyto":
				tempArgs = Args_rclone_copyto

			case "copyurl":
				tempArgs = Args_rclone_copyurl

			case "cryptcheck":
				tempArgs = Args_rclone_cryptcheck

			case "cryptdecode":
				tempArgs = Args_rclone_cryptdecode

			case "dbhashsum":
				tempArgs = Args_rclone_dbhashsum

			case "dedupe":
				tempArgs = Args_rclone_dedupe

			case "delete":
				tempArgs = Args_rclone_delete

			case "deletefile":
				tempArgs = Args_rclone_deletefile

			case "genautocomplete":
				tempArgs = Args_rclone_genautocomplete

			case "gendocs":
				tempArgs = Args_rclone_gendocs

			case "hashsum":
				tempArgs = Args_rclone_hashsum

			case "help":
				tempArgs = Args_rclone_help

			case "link":
				tempArgs = Args_rclone_link

			case "listremotes":
				tempArgs = Args_rclone_listremotes

			case "ls":
				tempArgs = Args_rclone_ls

			case "lsd":
				tempArgs = Args_rclone_lsd

			case "lsf":
				tempArgs = Args_rclone_lsf

			case "lsjson":
				tempArgs = Args_rclone_lsjson

			case "lsl":
				tempArgs = Args_rclone_lsl

			case "md5sum":
				tempArgs = Args_rclone_md5sum

			case "mkdir":
				tempArgs = Args_rclone_mkdir

			case "move":
				tempArgs = Args_rclone_move

			case "moveto":
				tempArgs = Args_rclone_moveto

			case "ncdu":
				tempArgs = Args_rclone_ncdu

			case "obscure":
				tempArgs = Args_rclone_obscure

			case "purge":
				tempArgs = Args_rclone_purge

			case "rc":
				tempArgs = Args_rclone_rc

			case "rcat":
				tempArgs = Args_rclone_rcat

			case "rcd":
				tempArgs = Args_rclone_rcd

			case "rmdir":
				tempArgs = Args_rclone_rmdir

			case "rmdirs":
				tempArgs = Args_rclone_rmdirs

			case "serve":
				tempArgs = Args_rclone_serve

			case "settier":
				tempArgs = Args_rclone_settier

			case "sha1sum":
				tempArgs = Args_rclone_sha1sum

			case "size":
				tempArgs = Args_rclone_size

			case "sync":
				tempArgs = Args_rclone_sync

			case "touch":
				tempArgs = Args_rclone_touch

			case "tree":
				tempArgs = Args_rclone_tree

			case "version":
				tempArgs = Args_rclone_version

			}

			return prompt.FilterHasPrefix(append(CommonArgs, tempArgs...), d.GetWordBeforeCursor(), true)

		}

	}

	return prompt.FilterHasPrefix(returnCommands, d.GetWordBeforeCursor(), true)
}
