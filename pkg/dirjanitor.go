package dirjanitor

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

type DirectoryCleaner struct {
	Directory string      // Directory path
	Retention uint        // Max age of files to keep (in days)
	Frequency uint        // Frequency of cleaning up (in seconds)
	Logger    *log.Logger // Custom logger object
	exitCh    chan int
}

// PerformCleanup scans the specified directory and removes files older than the retention period.
// It returns an error if it encounters any issues while performing the cleanup.
func (dc *DirectoryCleaner) PerformCleanup() error {
	files, err := os.ReadDir(dc.Directory)
	if err != nil {
		dc.logf("Can not read dir, error: %s\n", err.Error())
		return err
	}

	if len(files) == 0 {
		dc.logf("Skipping cleanup, no files found.\n")
		return err
	}

	dc.logf("Starting cleanup..\n")

	retentionTime := time.Now().AddDate(0, 0, -int(dc.Retention))

	// Loop through files and cleanup files older than specified with Retention setting
	for _, file := range files {
		modTime, errModTime := getFileModificationTime(file)
		if errModTime != nil {
			dc.logf("Can't retrieve last modification time for file: %s\n", file.Name())
			return err
		}

		if modTime.Unix() < retentionTime.Unix() {
			filePath := filepath.Join(dc.Directory, file.Name())
			errRemove := os.Remove(filePath)
			if errRemove != nil {
				dc.logf("Can not remove file, error: %s\n", errRemove.Error())
				return err
			}
			dc.logf("Successfully cleaned up: %s\n", filePath)
		}
	}
	dc.logf("Cleanup finished.\n")
	return nil
}

// StartCleanupInterval initiates the automated cleanup process at the defined frequency.
// It runs the PerformCleanup method at regular intervals.
func (dc *DirectoryCleaner) StartCleanupInterval() {
	go func() {
		dc.exitCh = make(chan int)
		for {
			select {
			case <-dc.exitCh:
				return
			default:
				time.Sleep(time.Duration(dc.Frequency) * time.Second)
				dc.PerformCleanup()

			}
		}
	}()
}

// StopCleanupInterval stops the automated cleanup process and releases associated resources.
func (dc *DirectoryCleaner) StopCleanupInterval() {
	dc.logf("Stopping cleanup interval.\n")
	close(dc.exitCh)
}

// getFileModificationTime retrieves the last modification time of a file.
// It returns the modification time and any error encountered.
func getFileModificationTime(file fs.DirEntry) (time.Time, error) {
	var creationTime time.Time
	fileInfo, err := file.Info()
	if err != nil {
		return creationTime, err
	}

	createTime := fileInfo.ModTime()
	return createTime, nil
}

// logf logs a formatted message using the provided logger, if available.
// It is used for informative output during cleanup operations.
func (dc *DirectoryCleaner) logf(format string, v ...interface{}) {
	if dc.Logger != nil {
		dc.Logger.Printf(format, v...)
	}
}
