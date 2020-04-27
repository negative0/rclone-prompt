## Rclone Prompt 

An interactive command line application to run and manage rclone [rclone](https://github.com/rclone/rclone)

![Rclone Prompt demo showing how it works](resources/demo.gif)

### Commands supported:

All the commands mentioned on the [rclone website](https://rclone.org/commands/) are supported.

Currently the prompt exits after running one command in the background. Eg ```rcd --rc-web-gui``` will open the web gui in the background and exit the rclone prompt. We are working on a multicommand prompt.

This application is still under development. Use with caution.

### Run 

#### Clone:

```git clone https://github.com/negative0/rclone-prompt```

#### Change rclone Path:

Inside the ```app.go``` file, find the following line:
```		
	...
	path = "/media/negative0/Projects/rclone/rclone-new/rclone/rclone"
	...
```
Here, you have to specify your own path of the rclone executable.

#### Compile:

```go build```

#### Run:

```./rclone-prompt```

### Exit 

use the keyword "exit" to exit the prompt
