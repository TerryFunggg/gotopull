# GoToPull

GoToPull is a simple executable command script for help you pull all the list git repositories on your local machine.

## Platform
Package only test on Arch linux and MacOS currently.

## Installation

Make sure your $GOPATH is set up correctly.
And $GOBIN or `~/go/bin` should included in your $PATH

For install this package, just simply type below:
```sh
go install github.com/terryfunggg/gotopull@latest
```
## Usage
If first time install this packge, run command once:
```
gotopull
```
Script will create an empty config file call `.target-git-pull` in your home directory if the file not exist.

open the config file, and type some directory path(Absolute path), for example:
```
/home/peter/.config/nvim
/home/peter/repo/some/repo
```

Save it. And run the command again:
```
gotopull
```
Script will run git pull inside each list directory. If some reason git fail to pull, it will show the directory path at the end.

## TODO
- [ ] support custom config file
- [ ] support multi config file
- [ ] Display more detail about fail git pull message, like conflict or other reason.
