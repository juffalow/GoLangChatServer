# [GoLangChatServer](https://github.com/juffalow/GoLangChatServer)

Chat server in Go lang.

## Tech

* [iris](http://iris-go.com/)

## Install

### Download dependency

[go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) command documentation

[Iris install](https://kataras.gitbooks.io/iris/content/install.html) page

```
go get -v github.com/kataras/iris/iris
```

### Build project

```
go install GoLangChatServer
```

## Versions

### v1.0.*

* v1.0.0 - basic functionality
* v1.0.1 - unnecessary code removed
* v1.0.2 - chat with message time

### v2.0.*

* v2.0.0 - messages in json format, list of connected users, caching last 10 messages
