go-mastery-labs
===============

A template project for 'Get Going with Go' Mastery Labs

##Download Golang
[http://golang.org/dl/](http://golang.org/dl/)
Version 1.3.1

##Install Golang
[http://golang.org/doc/install](http://golang.org/doc/install)


```
sudo tar -C /usr/local -xf go1.3.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

##Cloning this Repo

```
cd $WHEREVER_YOU_WANNA_STICKIT
git clone https://github.com/sheenathejunglegirl/go-mastery-labs.git
```

##Set GOPATH
```
cd go-mastery-labs
export GOPATH=`pwd`
```

##Hello World
###the code
view `$GOPATH/src/infusionsoft.com/user/hello/hello.go` in your favorite editor.

* package
* import
* main
* why is Printf capitalized?

###compiling
```
cd $GOPATH
go install infusionsoft.com/user/hello
```

###running

```
cd bin
./hello
```

##The actual project name here
###code walkthrough?
###installing
Describe go get?


```
cd $GOPATH
go get github.com/go-martini/martini
go get github.com/codegangsta/martini-contrib/render
go get github.com/mattn/go-sqlite3
```
###compiling
```
cd $GOPATH
go install github.com/sheenathejunglegirl/get-going-with-go
```
###running
```
cd github.com/sheenathejunglegirl/get-going-with-go/
$GOPATH/bin/get-going-with-go
```
> [martini] listening on :3000 (development)

open a browser and head to `http://localhost:3000/questions/unanswered/random`
