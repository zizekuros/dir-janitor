package dirjanitor

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

type DirectoryCleaner struct {
	Directory string
	Retention int // max age of files to keep (in days)
	Frequency int // frequency of cleaning up (in seconds)
	exitCh    chan int
}

func (dc *DirectoryCleaner) StartCleanupInterval() {

	go func() {

		dc.exitCh = make(chan int)

		for {

			select {

			case <-dc.exitCh:
				log.Println("Cleanup interval stopped")
				return

			default:
				// Good night for some time
				time.Sleep(time.Duration(dc.Frequency) * time.Second)

				// Read dir
				files, err := os.ReadDir(dc.Directory)
				if err != nil {
					log.Println(err.Error())
					continue
				}

				if len(files) == 0 {
					log.Println("Skipping cleanup, no files found.")
					continue
				}

				log.Println("Starting cleanup.")

				// Loop through files and cleanup files older than specified with Retention setting
				retentionTime := time.Now().AddDate(0, 0, -dc.Retention)

				for _, file := range files {
					modTime, errModTime := getFileModificationTime(file)
					if errModTime != nil {
						log.Printf("Can't retrieve last modification time for file: %s\n", file.Name())
						continue
					}

					if modTime.Unix() < retentionTime.Unix() {
						filePath := filepath.Join(dc.Directory, file.Name())
						errRemove := os.Remove(filePath)
						if errRemove != nil {
							log.Printf("Can not remove file, error: %s", errRemove.Error())
							continue
						}
						log.Printf("Sucessfuly cleaned up: %s\n", filePath)
					}
				}
				log.Println("Cleanup finished.")
			}

		}

	}()
}

func (fc *DirectoryCleaner) StopCleanupInterval() {
	close(fc.exitCh)
}

func getFileModificationTime(file fs.DirEntry) (time.Time, error) {
	var creationTime time.Time
	fileInfo, err := file.Info()
	if err != nil {
		return creationTime, err
	}

	createTime := fileInfo.ModTime()
	return createTime, nil
}
