package watcher

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/state"
)

// FileWatcher polls a local targetstate file for updates (e.g. K8s ConfigMap volume mount).
type FileWatcher struct {
	filePath string
	aesKey   string
	interval time.Duration
	callback func(*state.TargetState)
}

// NewFileWatcher constructs a new file polling watcher.
func NewFileWatcher(filePath, aesKey string, interval time.Duration, callback func(*state.TargetState)) *FileWatcher {
	if interval <= 0 {
		interval = 10 * time.Second
	}
	return &FileWatcher{
		filePath: filePath,
		aesKey:   aesKey,
		interval: interval,
		callback: callback,
	}
}

// ReadState reads and parses the target state once from disk.
func (w *FileWatcher) ReadState() (*state.TargetState, error) {
	data, err := os.ReadFile(w.filePath)
	if err != nil {
		return nil, err
	}
	return state.ParseTargetState(data, w.aesKey)
}

// Watch runs a background polling loop checking for mod-time changes on the target state file.
func (w *FileWatcher) Watch(ctx context.Context) {
	const category = "FileWatcher"
	logrus.WithField("category", category).Infof("Starting watcher for %s (polling every %v)", w.filePath, w.interval)

	var lastModTime time.Time
	if info, err := os.Stat(w.filePath); err == nil {
		lastModTime = info.ModTime()
	}

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			info, err := os.Stat(w.filePath)
			if err != nil {
				continue
			}

			if info.ModTime().After(lastModTime) {
				logrus.WithFields(logrus.Fields{
					"category": category,
					"path":     w.filePath,
				}).Info("Detected modification to initial state file. Re-applying...")

				lastModTime = info.ModTime()

				st, err := w.ReadState()
				if err != nil {
					logrus.WithField("category", category).Errorf("Failed to read modified state file: %v", err)
					continue
				}

				if w.callback != nil {
					w.callback(st)
				}
			}
		}
	}
}

