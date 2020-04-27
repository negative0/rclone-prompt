package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"log"
	"os/exec"
	"strings"
)

var commands = []prompt.Suggest{
	{Text: "about", Description: "Get quota information from the remote"},
	{Text: "authorize", Description: "Remote authorization"},
	{Text: "cachestats", Description: "Print cache stats for a remote"},
	{Text: "cat", Description: "Concatenates any files and sends them to stdout"},
	{Text: "check", Description: "Checks the files in the source and destination match"},
	{Text: "cleanup", Description: "Clean up the remote if possible"},
	{Text: "config", Description: "Enter an interactive configuration session"},
	{Text: "copy", Description: "Copy files from source to dest, skipping already copied"},
	{Text: "copyto", Description: "Copy files from source to dest, skipping already copied"},
	{Text: "copyurl", Description: "Copy url content to dest"},
	{Text: "cryptcheck", Description: "Cryptcheck checks the integrity of a crypted remote"},
	{Text: "cryptdecode", Description: "Cryptdecode returns unencrypted file names"},
	{Text: "dbhashsum", Description: "Produces a Dropbox hash file for all the objects in the path"},
	{Text: "dedupe", Description: "Interactively find duplicate files and delete/rename them"},
	{Text: "delete", Description: "Remove the contents of path"},
	{Text: "deletefile", Description: "Remove a single file from remote"},
	{Text: "genautocomplete", Description: "Output completion script for a given shell"},
	{Text: "gendocs", Description: "Output markdown docs for rclone to the directory supplied"},
	{Text: "hashsum", Description: "Produces an hashsum file for all the objects in the path"},
	{Text: "help", Description: "Show help for rclone commands, flags ande backends"},
	{Text: "link", Description: "Generate public link to file/folder"},
	{Text: "listremotes", Description: "List all the remotes in the config file"},
	{Text: "ls", Description: "List the objects in the path with size and path"},
	{Text: "lsd", Description: "List all directories/containers/buckets in the path"},
	{Text: "lsf", Description: "List directories and objects in remote"},
	{Text: "lsjson", Description: "List directories and objects in the path in JSON format"},
	{Text: "lsl", Description: "List the objects in path with modification time, size and path"},
	{Text: "md5sum", Description: "Produces an md5sum file for all the objects in the path"},
	{Text: "mkdir", Description: "Make the path if it doesn't already exist"},
	{Text: "move", Description: "Move files from source to dest"},
	{Text: "moveto", Description: "Move file or directory from source to dest"},
	{Text: "ncdu", Description: "Explore a remote with a text based user interface"},
	{Text: "obscure", Description: "Obscure password for use in the rclone"},
	{Text: "purge", Description: "Remove the path and all of its contents"},
	{Text: "rc", Description: "Run a command against a running rclone"},
	{Text: "rcat", Description: "Copies standard input to file on remote"},
	{Text: "rcd", Description: "Run rclone listening to remote control commands only"},
	{Text: "rmdir", Description: "Remove the path if empty"},
	{Text: "rmdirs", Description: "Remove empty directories under the path"},
	{Text: "serve", Description: "Serve a remote over a protocol"},
	{Text: "settier", Description: "Changes storage class/tier of objects in remote"},
	{Text: "sha1sum", Description: "Produces an sha1sum file for all the objects in the path"},
	{Text: "size", Description: "Prints the total size and number of objects in remote"},
	{Text: "sync", Description: "Make source and dest identical, modifying destination only"},
	{Text: "touch", Description: "Create new file or change file modification time"},
	{Text: "tree", Description: "List the contents of the remote in a tree like fashion"},
	{Text: "version", Description: "Show the version number"},
}

var configCommands = []prompt.Suggest{
	{Text: "create", Description: "Create a new remote with name, type and options"},
	{Text: "delete", Description: "Delete an existing remote <name>"},
	{Text: "disconnect", Description: "Disconnects user from remote"},
	{Text: "dump", Description: "Dump the config file as JSON"},
	{Text: "edit", Description: "Enter an interactive configuration session"},
	{Text: "file", Description: "Show path of configuration file in use"},
	{Text: "password", Description: "Update password in an existing remote"},
	{Text: "providers", Description: "List in JSON format all the providers and options"},
	{Text: "reconnect", Description: "Re-authenticates user with remote"},
	{Text: "show", Description: "Print (decrypted) config file, or the config for a single remote"},
	{Text: "update", Description: "Update options in an existing remote"},
	{Text: "userinfo", Description: "Prints info about logged in user of remote"},
}

var genAutoComplete = []prompt.Suggest{
	{Text: "bash", Description: "Output bash completion script for rclone"},
	{Text: "zsh", Description: "Output zsh completion script for rclone"},
}

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

func main() {
	/*Rclone prompt*/

	/*Look for rclone installation in the path*/
	path, err := exec.LookPath("rclone")
	if err != nil {
		path = "/media/negative0/Projects/rclone/rclone-new/rclone/rclone"
	}

	fmt.Println("Rclone Prompt")
	t := prompt.Input("rclone> ", completer)

	// Execute the command
	cmd := exec.Command(path, strings.Split(t, " ")...)

	//stdin, err := cmd.StdinPipe()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(cmd)
	}

	fmt.Println("You selected " + t)
}
