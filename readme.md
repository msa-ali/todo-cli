# TODO CLI Application

A simple command-line tool to manage a list of to-do items to keep track of items left in a project or activity. This tool will save the list of items in a file using the JSON format.

## Command line Arguments

- `-list`: A Boolean flag. When used, the tool will list all to-do items.
- `-add`: A Boolean flag. When used, the tool will include the string argument as a new todo item in the list from stdin or any reader like args.

Example:

```sh
go build
./todo -add including items from args
./todo​​ ​​-list​
​echo​​ ​​"This item comes from STDIN"​​ ​​|​​ ​​./todo​​ ​​-add​
```

- `-complete`: An integer flag. When used, the tool will mark the item number as completed.

## Env vars

SET `export TODO_FILENAME=new-todo.json` to set custom filename.
