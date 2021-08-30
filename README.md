# modtest3

Project is meant for some Go module testing only. It should never be imported.

## Steps

- Added code
- Tagged v0.2.0
- touch go_mod_tidy_hack.go
- cd cat
- go mod init github.com/codyoss/modtest3/cat
- go mod edit -require github.com/codyoss/modtest3@v0.3.0
- go mod edit -replace github.com/codyoss/modtest3@v0.3.0=/Users/codyoss/scratch/modtest3
- go mod tidy
- go mod edit -dropreplace github.com/codyoss/modtest3@v0.3.0
- cd ./..
- go mod edit -require github.com/codyoss/modtest3/cat@v0.1.0
