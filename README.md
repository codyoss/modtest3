# modtest3

Project is meant for some Go module testing only. It should never be imported.

## Steps

- Added code
- Tagged v0.1.0
- touch go_mod_tidy_hack.go
- [child] go mod edit -require github.com/codyoss/modtest2@v0.2.0
- [child] go mod edit -replace github.com/codyoss/modtest2@v0.2.0=/Users/codyoss/scratch/modtest2
- [child] go mod tidy
- [child] go mod edit -dropreplace github.com/codyoss/modtest2@v0.2.0
- [parent] go mod edit -require github.com/codyoss/modtest2/cat@v0.1.