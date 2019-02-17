# Ciak is a lightweight media server written in go

[![Build Status](https://travis-ci.org/GaruGaru/ciak.svg?branch=master)](https://travis-ci.org/GaruGaru/ciak)
[![Go Report Card](https://goreportcard.com/badge/github.com/GaruGaru/ciak)](https://goreportcard.com/report/github.com/GaruGaru/ciak)
![license](https://img.shields.io/github/license/GaruGaru/ciak.svg)

Ciak allows you to show and stream your personal media tv series, movies, etc with a simple and clean web ui.
The server also provide on the fly video encoding in order to stream non standard formats such as avi, mkv...

<img src="https://github.com/garugaru/ciak/raw/master/res/ciak-media-list.png" width="1000">


## Run ciak

### Using go

Install ciak


    go get -u github.com/garugaru/ciak


Launch the media server (on 0.0.0.0:8082)


    ciak --data=<your/media/directory>



### Using docker


    docker run -v <your/media/directory>:/data -p 8082:8082 garugaru/ciak



### Configuration

You can configure Ciak using the command line flags


* **--bind** binding for the webserver interface:port (default 0.0.0.0:8082)

* **--media** media files directory (default /data)

* **--auth** enable web server authentication (default false) the authentication is configured by the env variables **CIAK_USERNAME** and **CIAK_PASSWORD**

* **--auto-convert-media** automatically converts the medias to a streamable format (eg: mkv->mp4) (default false)

* **--delete-original-media** delete original media after conversion to save up space (default false)

* **--transfer-path** enable transfer feature that allows media copy from a device to another using the web ui 

* **--omdb-api-key** omdbapi.com api key used for movie metadata retrieving 

* **--db** database file path (default /tmp/ciak_daemon.db)