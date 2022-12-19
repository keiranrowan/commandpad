# CommandPad

This project allows you to run arbitrary commands from a Novation LaunchPad. 

## Usage

Adjust config.toml to match the settings required by your Novation LaunchPad. In the "notes" section, add commands to be executed when each pad is pressed. The configuration can be reloaded while the program is running by sending a Ctrl-C interrupt. 

## Installation

Download the source code to your Go workspace and run the following commands in the working directory.

```bash
$ go get
$ go build
$ go install
```
