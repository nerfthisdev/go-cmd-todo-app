# Todolite

A simple CLI todo app that I created to learn Cobra CLI for go.


Availible functionality at this stage of the project
```
Usage:
  todolite [command]

Available Commands:
  add         Add a task to your todo list!
  complete    Mark task as completed!
  help        Help about any command
  list        Shows a list of current tasks
```

Also there are some flags you can use with commands

For example

```
Usage:
  todolite list [flags]

Flags:
  -a, --all    Include all items
  -h, --help   help for list
  -i, --id     Show id all items
```
```
Mark task as completed!

Usage:
  todolite complete [flags]

Flags:
  -h, --help   help for complete
  -i, --id     Use Id to complete a task
```

