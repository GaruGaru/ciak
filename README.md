# Ciak is a lightweight media server written in go

Ciak allows you to show and stream your personal media tv series, movies, etc with a simple and clean web ui.

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
