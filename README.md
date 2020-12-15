# Ciak is a lightweight media server written in go

[![Go Report Card](https://goreportcard.com/badge/github.com/GaruGaru/ciak)](https://goreportcard.com/report/github.com/GaruGaru/ciak)
![go-build-test](https://github.com/GaruGaru/ciak/workflows/go-build-test/badge.svg)
![docker-build](https://github.com/GaruGaru/ciak/workflows/docker-build/badge.svg)
![release](https://github.com/GaruGaru/ciak/workflows/release/badge.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/garugaru/ciak)
![license](https://img.shields.io/github/license/GaruGaru/ciak.svg)


Ciak allows you to show and stream your personal media tv series, movies, etc with a simple and clean web ui.
The server also provide on the fly video encoding in order to stream non standard formats such as avi, mkv...

<img src="https://github.com/garugaru/ciak/raw/master/res/ciak-media-list.png" width="1000">


## Run ciak

### Using go

Install ciak


    go get -u github.com/garugaru/ciak


Launch the media server (on 0.0.0.0:8082)


    ciak --media=<your/media/directory>



### Using docker


    docker run -v <your/media/directory>:/data -p 8082:8082 garugaru/ciak



### Configuration

You can configure Ciak using the command line flags


* **--bind** binding for the webserver interface:port (default 0.0.0.0:8082)

* **--media** media files directory (default /data)

* **--auth** enable web server authentication (default false) the authentication is configured by the env variables **CIAK_USERNAME** and **CIAK_PASSWORD**

* **--omdb-api-key** omdbapi.com api key used for movie metadata retrieving 

* **--db** database file path (default /ciak_daemon.db)