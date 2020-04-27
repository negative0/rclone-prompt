package main

import "github.com/c-bata/go-prompt"

var Args_rclone = []prompt.Suggest{}
var Args_rclone_about = []prompt.Suggest{
	{Text: "--full", Description: "Full numbers instead of SI units"},
	{Text: "--json", Description: "Format output as JSON"},
}
var Args_rclone_authorize = []prompt.Suggest{
	{Text: "--auth-no-open-browser", Description: "Do not automatically open auth link in default browser"},
}
var Args_rclone_cachestats = []prompt.Suggest{}
var Args_rclone_cat = []prompt.Suggest{
	{Text: "--count", Description: "Only print N characters."},
	{Text: "--discard", Description: "Discard the output instead of printing."},
	{Text: "--head", Description: "Only print the first N characters."},
	{Text: "--offset", Description: "Start printing at offset N (or from end if -ve)."},
	{Text: "--tail", Description: "Only print the last N characters."},
}
var Args_rclone_check = []prompt.Suggest{
	{Text: "--download", Description: "Check by downloading rather than with hash."},
	{Text: "--one-way", Description: "Check one way only, source files must exist on remote"},
}
var Args_rclone_cleanup = []prompt.Suggest{}
var Args_rclone_config = []prompt.Suggest{}
var Args_rclone_config_create = []prompt.Suggest{}
var Args_rclone_config_delete = []prompt.Suggest{}
var Args_rclone_config_disconnect = []prompt.Suggest{}
var Args_rclone_config_dump = []prompt.Suggest{}
var Args_rclone_config_edit = []prompt.Suggest{}
var Args_rclone_config_file = []prompt.Suggest{}
var Args_rclone_config_password = []prompt.Suggest{}
var Args_rclone_config_providers = []prompt.Suggest{}
var Args_rclone_config_reconnect = []prompt.Suggest{}
var Args_rclone_config_show = []prompt.Suggest{}
var Args_rclone_config_update = []prompt.Suggest{}
var Args_rclone_config_userinfo = []prompt.Suggest{
	{Text: "--json", Description: "Format output as JSON"},
}
var Args_rclone_copy = []prompt.Suggest{
	{Text: "--create-empty-src-dirs", Description: "Create empty source dirs on destination after copy"},
}
var Args_rclone_copyto = []prompt.Suggest{}
var Args_rclone_copyurl = []prompt.Suggest{
	{Text: "(-a --auto-filename)'{-a,--auto-filename}'", Description: "Get the file name from the URL and use it for destination file path"},
	{Text: "--stdout", Description: "Write the output to stdout rather than a file"},
}
var Args_rclone_cryptcheck = []prompt.Suggest{
	{Text: "--one-way", Description: "Check one way only, source files must exist on destination"},
}
var Args_rclone_cryptdecode = []prompt.Suggest{
	{Text: "--reverse", Description: "Reverse cryptdecode, encrypts filenames"},
}
var Args_rclone_dbhashsum = []prompt.Suggest{}
var Args_rclone_dedupe = []prompt.Suggest{
	{Text: "--dedupe-mode", Description: "Dedupe mode interactive|skip|first|newest|oldest|largest|smallest|rename."},
}
var Args_rclone_delete = []prompt.Suggest{}
var Args_rclone_deletefile = []prompt.Suggest{}
var Args_rclone_genautocomplete = []prompt.Suggest{}
var Args_rclone_genautocomplete_bash = []prompt.Suggest{}
var Args_rclone_genautocomplete_zsh = []prompt.Suggest{
	{Text: "(-h --help)'{-h,--help}'", Description: "help for zsh"},
}
var Args_rclone_gendocs = []prompt.Suggest{}
var Args_rclone_hashsum = []prompt.Suggest{
	{Text: "--base64", Description: "Output base64 encoded hashsum"},
}
var Args_rclone_help = []prompt.Suggest{}
var Args_rclone_help_backend = []prompt.Suggest{}
var Args_rclone_help_backends = []prompt.Suggest{}
var Args_rclone_help_flags = []prompt.Suggest{}
var Args_rclone_link = []prompt.Suggest{}
var Args_rclone_listremotes = []prompt.Suggest{
	{Text: "--long", Description: "Show the type as well as names."},
}
var Args_rclone_ls = []prompt.Suggest{}
var Args_rclone_lsd = []prompt.Suggest{
	{Text: "(-R --recursive)'{-R,--recursive}'", Description: "Recurse into the listing."},
}
var Args_rclone_lsf = []prompt.Suggest{
	{Text: "--absolute", Description: "Put a leading / in front of path names."},
	{Text: "--csv", Description: "Output in CSV format."},
	{Text: "(-d --dir-slash)'{-d,--dir-slash}'", Description: "Append a slash to directory names."},
	{Text: "--dirs-only", Description: "Only list directories."},
	{Text: "--files-only", Description: "Only list files."},
	{Text: "(-F --format)'{-F,--format}'", Description: "Output format - see  help for details"},
	{Text: "--hash", Description: "Use this hash when `h` is used in the format MD5|SHA-1|DropboxHash"},
	{Text: "(-R --recursive)'{-R,--recursive}'", Description: "Recurse into the listing."},
	{Text: "(-s --separator)'{-s,--separator}'", Description: "Separator for the items in the format."},
}
var Args_rclone_lsjson = []prompt.Suggest{
	{Text: "--dirs-only", Description: "Show only directories in the listing."},
	{Text: "(-M --encrypted)'{-M,--encrypted}'", Description: "Show the encrypted names."},
	{Text: "--files-only", Description: "Show only files in the listing."},
	{Text: "--hash", Description: "Include hashes in the output (may take longer)."},
	{Text: "--no-mimetype", Description: "Don'''t read the mime type (can speed things up)."},
	{Text: "--no-modtime", Description: "Don'''t read the modification time (can speed things up)."},
	{Text: "--original", Description: "Show the ID of the underlying Object."},
	{Text: "(-R --recursive)'{-R,--recursive}'", Description: "Recurse into the listing."},
}
var Args_rclone_lsl = []prompt.Suggest{}
var Args_rclone_md5sum = []prompt.Suggest{}
var Args_rclone_mkdir = []prompt.Suggest{}
var Args_rclone_move = []prompt.Suggest{
	{Text: "--create-empty-src-dirs", Description: "Create empty source dirs on destination after move"},
	{Text: "--delete-empty-src-dirs", Description: "Delete empty source dirs after move"},
}
var Args_rclone_moveto = []prompt.Suggest{}
var Args_rclone_ncdu = []prompt.Suggest{}
var Args_rclone_obscure = []prompt.Suggest{}
var Args_rclone_purge = []prompt.Suggest{}
var Args_rclone_rc = []prompt.Suggest{
	{Text: "--json", Description: "Input JSON - use instead of key=value args."},
	{Text: "--loopback", Description: "If set connect to this rclone instance not via HTTP."},
	{Text: "--no-output", Description: "If set don'''t output the JSON result."},
	{Text: "--pass", Description: "Password to use to connect to rclone remote control."},
	{Text: "--url", Description: "URL to connect to rclone remote control."},
	{Text: "--user", Description: "Username to use to rclone remote control."},
}
var Args_rclone_rcat = []prompt.Suggest{}
var Args_rclone_rcd = []prompt.Suggest{}
var Args_rclone_rmdir = []prompt.Suggest{}
var Args_rclone_rmdirs = []prompt.Suggest{
	{Text: "--leave-root", Description: "Do not remove root directory if empty"},
}
var Args_rclone_serve = []prompt.Suggest{}
var Args_rclone_serve_dlna = []prompt.Suggest{
	{Text: "--addr", Description: "ip:port or :port to bind the DLNA http server to."},
	{Text: "--dir-cache-time", Description: "Time to cache directory entries for."},
	{Text: "--dir-perms", Description: "Directory permissions"},
	{Text: "--file-perms", Description: "File permissions"},
	{Text: "--gid", Description: "Override the gid field set by the filesystem."},
	{Text: "--log-trace", Description: "enable trace logging of SOAP traffic"},
	{Text: "--name", Description: "name of DLNA server"},
	{Text: "--no-checksum", Description: "Don'''t compare checksums on up/download."},
	{Text: "--no-modtime", Description: "Don'''t read/write the modification time (can speed things up)."},
	{Text: "--no-seek", Description: "Don'''t allow seeking in files."},
	{Text: "--poll-interval", Description: "Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable."},
	{Text: "--read-only", Description: "Mount read-only."},
	{Text: "--uid", Description: "Override the uid field set by the filesystem."},
	{Text: "--umask", Description: "Override the permission bits set by the filesystem."},
	{Text: "--vfs-cache-max-age", Description: "Max age of objects in the cache."},
	{Text: "--vfs-cache-max-size", Description: "Max total size of objects in the cache."},
	{Text: "--vfs-cache-mode", Description: "Cache mode off|minimal|writes|full"},
	{Text: "--vfs-cache-poll-interval", Description: "Interval to poll the cache for stale objects."},
	{Text: "--vfs-case-insensitive", Description: "If a file name not found, find a case insensitive match."},
	{Text: "--vfs-read-chunk-size", Description: "Read the source objects in chunks."},
	{Text: "--vfs-read-chunk-size-limit", Description: "If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached. '''off''' is unlimited."},
	{Text: "--vfs-read-wait", Description: "Time to wait for in-sequence read before seeking."},
	{Text: "--vfs-write-wait", Description: "Time to wait for in-sequence write before giving error."},
}
var Args_rclone_serve_ftp = []prompt.Suggest{
	{Text: "--addr", Description: "IPaddress:Port or :Port to bind server to."},
	{Text: "--auth-proxy", Description: "A program to use to create the backend from the auth."},
	{Text: "--dir-cache-time", Description: "Time to cache directory entries for."},
	{Text: "--dir-perms", Description: "Directory permissions"},
	{Text: "--file-perms", Description: "File permissions"},
	{Text: "--gid", Description: "Override the gid field set by the filesystem."},
	{Text: "--no-checksum", Description: "Don'''t compare checksums on up/download."},
	{Text: "--no-modtime", Description: "Don'''t read/write the modification time (can speed things up)."},
	{Text: "--no-seek", Description: "Don'''t allow seeking in files."},
	{Text: "--pass", Description: "Password for authentication. (empty value allow every password)"},
	{Text: "--passive-port", Description: "Passive port range to use."},
	{Text: "--poll-interval", Description: "Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable."},
	{Text: "--public-ip", Description: "Public IP address to advertise for passive connections."},
	{Text: "--read-only", Description: "Mount read-only."},
	{Text: "--uid", Description: "Override the uid field set by the filesystem."},
	{Text: "--umask", Description: "Override the permission bits set by the filesystem."},
	{Text: "--user", Description: "User name for authentication."},
	{Text: "--vfs-cache-max-age", Description: "Max age of objects in the cache."},
	{Text: "--vfs-cache-max-size", Description: "Max total size of objects in the cache."},
	{Text: "--vfs-cache-mode", Description: "Cache mode off|minimal|writes|full"},
	{Text: "--vfs-cache-poll-interval", Description: "Interval to poll the cache for stale objects."},
	{Text: "--vfs-case-insensitive", Description: "If a file name not found, find a case insensitive match."},
	{Text: "--vfs-read-chunk-size", Description: "Read the source objects in chunks."},
	{Text: "--vfs-read-chunk-size-limit", Description: "If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached. '''off''' is unlimited."},
	{Text: "--vfs-read-wait", Description: "Time to wait for in-sequence read before seeking."},
	{Text: "--vfs-write-wait", Description: "Time to wait for in-sequence write before giving error."},
}
var Args_rclone_serve_http = []prompt.Suggest{
	{Text: "--addr", Description: "IPaddress:Port or :Port to bind server to."},
	{Text: "--baseurl", Description: "Prefix for URLs - leave blank for root."},
	{Text: "--cert", Description: "SSL PEM key (concatenation of certificate and CA certificate)"},
	{Text: "--client-ca", Description: "Client certificate authority to verify clients with"},
	{Text: "--dir-cache-time", Description: "Time to cache directory entries for."},
	{Text: "--dir-perms", Description: "Directory permissions"},
	{Text: "--file-perms", Description: "File permissions"},
	{Text: "--gid", Description: "Override the gid field set by the filesystem."},
	{Text: "--htpasswd", Description: "htpasswd file - if not provided no authentication is done"},
	{Text: "--key", Description: "SSL PEM Private key"},
	{Text: "--max-header-bytes", Description: "Maximum size of request header"},
	{Text: "--no-checksum", Description: "Don'''t compare checksums on up/download."},
	{Text: "--no-modtime", Description: "Don'''t read/write the modification time (can speed things up)."},
	{Text: "--no-seek", Description: "Don'''t allow seeking in files."},
	{Text: "--pass", Description: "Password for authentication."},
	{Text: "--poll-interval", Description: "Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable."},
	{Text: "--read-only", Description: "Mount read-only."},
	{Text: "--realm", Description: "realm for authentication"},
	{Text: "--server-read-timeout", Description: "Timeout for server reading data"},
	{Text: "--server-write-timeout", Description: "Timeout for server writing data"},
	{Text: "--uid", Description: "Override the uid field set by the filesystem."},
	{Text: "--umask", Description: "Override the permission bits set by the filesystem."},
	{Text: "--user", Description: "User name for authentication."},
	{Text: "--vfs-cache-max-age", Description: "Max age of objects in the cache."},
	{Text: "--vfs-cache-max-size", Description: "Max total size of objects in the cache."},
	{Text: "--vfs-cache-mode", Description: "Cache mode off|minimal|writes|full"},
	{Text: "--vfs-cache-poll-interval", Description: "Interval to poll the cache for stale objects."},
	{Text: "--vfs-case-insensitive", Description: "If a file name not found, find a case insensitive match."},
	{Text: "--vfs-read-chunk-size", Description: "Read the source objects in chunks."},
	{Text: "--vfs-read-chunk-size-limit", Description: "If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached. '''off''' is unlimited."},
	{Text: "--vfs-read-wait", Description: "Time to wait for in-sequence read before seeking."},
	{Text: "--vfs-write-wait", Description: "Time to wait for in-sequence write before giving error."},
}
var Args_rclone_serve_restic = []prompt.Suggest{
	{Text: "--addr", Description: "IPaddress:Port or :Port to bind server to."},
	{Text: "--append-only", Description: "disallow deletion of repository data"},
	{Text: "--baseurl", Description: "Prefix for URLs - leave blank for root."},
	{Text: "--cert", Description: "SSL PEM key (concatenation of certificate and CA certificate)"},
	{Text: "--client-ca", Description: "Client certificate authority to verify clients with"},
	{Text: "--htpasswd", Description: "htpasswd file - if not provided no authentication is done"},
	{Text: "--key", Description: "SSL PEM Private key"},
	{Text: "--max-header-bytes", Description: "Maximum size of request header"},
	{Text: "--pass", Description: "Password for authentication."},
	{Text: "--private-repos", Description: "users can only access their private repo"},
	{Text: "--realm", Description: "realm for authentication"},
	{Text: "--server-read-timeout", Description: "Timeout for server reading data"},
	{Text: "--server-write-timeout", Description: "Timeout for server writing data"},
	{Text: "--stdio", Description: "run an HTTP2 server on stdin/stdout"},
	{Text: "--user", Description: "User name for authentication."},
}
var Args_rclone_serve_sftp = []prompt.Suggest{
	{Text: "--addr", Description: "IPaddress:Port or :Port to bind server to."},
	{Text: "--auth-proxy", Description: "A program to use to create the backend from the auth."},
	{Text: "--authorized-keys", Description: "Authorized keys file"},
	{Text: "--dir-cache-time", Description: "Time to cache directory entries for."},
	{Text: "--dir-perms", Description: "Directory permissions"},
	{Text: "--file-perms", Description: "File permissions"},
	{Text: "--gid", Description: "Override the gid field set by the filesystem."},
	{Text: "--key", Description: "SSH private host key file (leave blank to auto generate)"},
	{Text: "--no-auth", Description: "Allow connections with no authentication if set."},
	{Text: "--no-checksum", Description: "Don'''t compare checksums on up/download."},
	{Text: "--no-modtime", Description: "Don'''t read/write the modification time (can speed things up)."},
	{Text: "--no-seek", Description: "Don'''t allow seeking in files."},
	{Text: "--pass", Description: "Password for authentication."},
	{Text: "--poll-interval", Description: "Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable."},
	{Text: "--read-only", Description: "Mount read-only."},
	{Text: "--uid", Description: "Override the uid field set by the filesystem."},
	{Text: "--umask", Description: "Override the permission bits set by the filesystem."},
	{Text: "--user", Description: "User name for authentication."},
	{Text: "--vfs-cache-max-age", Description: "Max age of objects in the cache."},
	{Text: "--vfs-cache-max-size", Description: "Max total size of objects in the cache."},
	{Text: "--vfs-cache-mode", Description: "Cache mode off|minimal|writes|full"},
	{Text: "--vfs-cache-poll-interval", Description: "Interval to poll the cache for stale objects."},
	{Text: "--vfs-case-insensitive", Description: "If a file name not found, find a case insensitive match."},
	{Text: "--vfs-read-chunk-size", Description: "Read the source objects in chunks."},
	{Text: "--vfs-read-chunk-size-limit", Description: "If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached. '''off''' is unlimited."},
	{Text: "--vfs-read-wait", Description: "Time to wait for in-sequence read before seeking."},
	{Text: "--vfs-write-wait", Description: "Time to wait for in-sequence write before giving error."},
}
var Args_rclone_serve_webdav = []prompt.Suggest{
	{Text: "--addr", Description: "IPaddress:Port or :Port to bind server to."},
	{Text: "--auth-proxy", Description: "A program to use to create the backend from the auth."},
	{Text: "--baseurl", Description: "Prefix for URLs - leave blank for root."},
	{Text: "--cert", Description: "SSL PEM key (concatenation of certificate and CA certificate)"},
	{Text: "--client-ca", Description: "Client certificate authority to verify clients with"},
	{Text: "--dir-cache-time", Description: "Time to cache directory entries for."},
	{Text: "--dir-perms", Description: "Directory permissions"},
	{Text: "--disable-dir-list", Description: "Disable HTML directory list on GET request for a directory"},
	{Text: "--etag-hash", Description: "Which hash to use for the ETag, or auto or blank for off"},
	{Text: "--file-perms", Description: "File permissions"},
	{Text: "--gid", Description: "Override the gid field set by the filesystem."},
	{Text: "--htpasswd", Description: "htpasswd file - if not provided no authentication is done"},
	{Text: "--key", Description: "SSL PEM Private key"},
	{Text: "--max-header-bytes", Description: "Maximum size of request header"},
	{Text: "--no-checksum", Description: "Don'''t compare checksums on up/download."},
	{Text: "--no-modtime", Description: "Don'''t read/write the modification time (can speed things up)."},
	{Text: "--no-seek", Description: "Don'''t allow seeking in files."},
	{Text: "--pass", Description: "Password for authentication."},
	{Text: "--poll-interval", Description: "Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable."},
	{Text: "--read-only", Description: "Mount read-only."},
	{Text: "--realm", Description: "realm for authentication"},
	{Text: "--server-read-timeout", Description: "Timeout for server reading data"},
	{Text: "--server-write-timeout", Description: "Timeout for server writing data"},
	{Text: "--uid", Description: "Override the uid field set by the filesystem."},
	{Text: "--umask", Description: "Override the permission bits set by the filesystem."},
	{Text: "--user", Description: "User name for authentication."},
	{Text: "--vfs-cache-max-age", Description: "Max age of objects in the cache."},
	{Text: "--vfs-cache-max-size", Description: "Max total size of objects in the cache."},
	{Text: "--vfs-cache-mode", Description: "Cache mode off|minimal|writes|full"},
	{Text: "--vfs-cache-poll-interval", Description: "Interval to poll the cache for stale objects."},
	{Text: "--vfs-case-insensitive", Description: "If a file name not found, find a case insensitive match."},
	{Text: "--vfs-read-chunk-size", Description: "Read the source objects in chunks."},
	{Text: "--vfs-read-chunk-size-limit", Description: "If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached. '''off''' is unlimited."},
	{Text: "--vfs-read-wait", Description: "Time to wait for in-sequence read before seeking."},
	{Text: "--vfs-write-wait", Description: "Time to wait for in-sequence write before giving error."},
}
var Args_rclone_settier = []prompt.Suggest{}
var Args_rclone_sha1sum = []prompt.Suggest{}
var Args_rclone_size = []prompt.Suggest{
	{Text: "--json", Description: "format output as JSON"},
}
var Args_rclone_sync = []prompt.Suggest{
	{Text: "--create-empty-src-dirs", Description: "Create empty source dirs on destination after sync"},
}
var Args_rclone_touch = []prompt.Suggest{
	{Text: "--localtime", Description: "Use localtime for timestamp, not UTC."},
	{Text: "(-C --no-create)'{-C,--no-create}'", Description: "Do not create the file if it does not exist."},
	{Text: "(-t --timestamp)'{-t,--timestamp}'", Description: "Use specified time instead of the current time of day."},
}
var Args_rclone_tree = []prompt.Suggest{
	{Text: "(-a --all)'{-a,--all}'", Description: "All files are listed (list . files too)."},
	{Text: "(-C --color)'{-C,--color}'", Description: "Turn colorization on always."},
	{Text: "(-d --dirs-only)'{-d,--dirs-only}'", Description: "List directories only."},
	{Text: "--dirsfirst", Description: "List directories before files (-U disables)."},
	{Text: "--full-path", Description: "Print the full path prefix for each file."},
	{Text: "--human", Description: "Print the size in a more human readable way."},
	{Text: "--level", Description: "Descend only level directories deep."},
	{Text: "(-D --modtime)'{-D,--modtime}'", Description: "Print the date of last modification."},
	{Text: "(-i --noindent)'{-i,--noindent}'", Description: "Don'''t print indentation lines."},
	{Text: "--noreport", Description: "Turn off file/directory count at end of tree listing."},
	{Text: "(-o --output)'{-o,--output}'", Description: "Output to file instead of stdout."},
	{Text: "(-p --protections)'{-p,--protections}'", Description: "Print the protections for each file."},
	{Text: "(-Q --quote)'{-Q,--quote}'", Description: "Quote filenames with double quotes."},
	{Text: "(-s --size)'{-s,--size}'", Description: "Print the size in bytes of each file."},
	{Text: "--sort", Description: "Select sort: name,version,size,mtime,ctime."},
	{Text: "--sort-ctime", Description: "Sort files by last status change time."},
	{Text: "(-t --sort-modtime)'{-t,--sort-modtime}'", Description: "Sort files by last modification time."},
	{Text: "(-r --sort-reverse)'{-r,--sort-reverse}'", Description: "Reverse the order of the sort."},
	{Text: "(-U --unsorted)'{-U,--unsorted}'", Description: "Leave files unsorted."},
	{Text: "--version", Description: "Sort files alphanumerically by version."},
}
var Args_rclone_version = []prompt.Suggest{
	{Text: "--check", Description: "Check for new version."},
}
