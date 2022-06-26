# fileSystemUtilities
Go (Lang) based file system utilities for cleanup and organization

# To Run
``` go run main.go [directory path] ```
example:
``` go run main.go c:/temp ```

This will then go through that folder and do the following actions:
- Delete folder that have a New Folder in them that are empty.
- Remove images that are the same as each other within a certain confidence level.