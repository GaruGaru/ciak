package daemon

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

type MediaFileSystemChangesTask struct {
	MonitorPath    string
	OnFileCreateFn func()
}

func (t MediaFileSystemChangesTask) Type() string {
	return "MediaFileSystemChangesTask"
}

func (t MediaFileSystemChangesTask) Run() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					t.OnFileCreateFn()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error("error:", err)
			}
		}
	}()

	err = watcher.Add(t.MonitorPath)
	if err != nil {
		return err
	}
	<-done
	return nil
}
