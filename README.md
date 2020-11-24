<p align="center">
	<a href=""><img src="https://project.golanger.org/img/golang-project/task/screenshot/taskGo-screenshot-screely.com.png" width="666"></a>
	<p align="center">
		<img src="https://github.com/YUbuntu0109/task/workflows/Go/badge.svg?branch=master"></img>
		<img src="https://goreportcard.com/badge/github.com/YUbuntu0109/task"></img>
		<img src="https://img.shields.io/github/commit-activity/m/google-golang/task?color=ff69b4"></img>
		<img src="https://img.shields.io/github/repo-size/google-golang/task"></img>
		<img src="https://img.shields.io/github/license/google-golang/task.svg"></img>
	</p>	
</p>


# Task.go
> It's not only a todo list but also a notebook and file mangage system, which base on vue, gin framework and sqlite3 database. so no complicated configurations and run it by one command !
* `go` v1.14.5
* `vue` v2.0
* `gin` v1.6.3
* ~`docker` v2.5.0.1~
* `gorm` v1.20.1
* `gopkg` v1.62.0
* `sqlite3` v3.33.0

## How To Run
### build
```shell script
go build
```

### run
```shell script
# Mac / Unix
./task conf/config.ini

# Windows
task.exe conf/config.ini
```

finally, open your browser and input the url : *http://127.0.0.1:9000/*, please modify the content of config.ini file if you want to custom the configuration information bro.
