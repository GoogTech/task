<p align="center">
	<a href=""><img src="https://user-images.githubusercontent.com/43493852/147871269-c0bbe5ea-e844-45ba-9bac-58214e2e1ab4.png" width="666"></a>
	<p align="center">
		<img src="https://github.com/GoogTech/task/workflows/Go/badge.svg?branch=master"></img>
		<img src="https://goreportcard.com/badge/github.com/GoogTech/task"></img>
		<img src="https://img.shields.io/github/commit-activity/m/google-golang/task?color=ff69b4"></img>
		<img src="https://img.shields.io/github/repo-size/google-golang/task"></img>
		<img src="https://img.shields.io/github/license/google-golang/task.svg"></img>
	</p>	
</p>


# Task.go
> It's a mini todo list which is based on vue, gin framework, and sqlite3 database. no complicated configurations and run it by one command !
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
