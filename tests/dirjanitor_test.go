package dirjanitor

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	dirjanitor "github.com/zizekuros/dir-janitor/pkg"
)

func TestDirectoryCleaner_CleanupInterval(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test-dirjanitor-1")
	if err != nil {
		t.Fatalf("Failed to create a temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create some test files with different modification times
	createTestFiles(tmpDir)

	// Create a DirectoryCleaner instance
	cleaner := &dirjanitor.DirectoryCleaner{
		Directory: tmpDir,
		Retention: 2, // Files younger than 2 days will be retained
		Frequency: 1, // Cleanup every second
		Logger:    log.New(os.Stdout, "DirJanitor: ", log.Ldate|log.Ltime),
	}

	// Start the cleanup interval
	go cleaner.StartCleanupInterval()

	// Let the cleaner run for a few seconds
	time.Sleep(4 * time.Second)

	// Stop the cleanup routine
	cleaner.StopCleanupInterval()

	// Verify that only the test files older than 2 days were cleaned up
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			t.Errorf("Failed to get file info for %s: %v", file.Name(), err)
			continue
		}
		if fileInfo.ModTime().Before(time.Now().AddDate(0, 0, -2)) {
			t.Errorf("File %s was not cleaned up as expected.", file.Name())
		}
	}
}

func TestDirectoryCleaner_PerformCleanup(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test-dirjanitor-2")
	if err != nil {
		t.Fatalf("Failed to create a temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create some test files with different modification times
	createTestFiles(tmpDir)

	// Create a DirectoryCleaner instance
	cleaner := &dirjanitor.DirectoryCleaner{
		Directory: tmpDir,
		Retention: 2, // Files younger than 2 days will be retained
		Frequency: 1,
		Logger:    log.New(os.Stdout, "DirJanitor: ", log.Ldate|log.Ltime),
	}

	// Perform manual cleanup
	errCleanup := cleaner.PerformCleanup()
	if errCleanup != nil {
		t.Errorf("PerformCleanup failed with error: %v", err)
	}

	// Verify that only the test files older than 2 days were cleaned up
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			t.Errorf("Failed to get file info for %s: %v", file.Name(), err)
			continue
		}
		if fileInfo.ModTime().Before(time.Now().AddDate(0, 0, -2)) {
			t.Errorf("File %s was not cleaned up as expected.", file.Name())
		}
	}
}

func createTestFiles(dir string) {
	// Create test files with different modification times (some older than 2 days and some younger)
	for i := 0; i < 5; i++ {
		filename := filepath.Join(dir, fmt.Sprintf("file%d.txt", i))
		if i%2 == 0 {
			// Set the modification time to more than 2 days ago
			modTime := time.Now().AddDate(0, 0, -3)
			os.WriteFile(filename, []byte("Test file content"), os.ModePerm)
			os.Chtimes(filename, modTime, modTime)
		} else {
			// Set the modification time to less than 2 days ago
			modTime := time.Now().AddDate(0, 0, -1)
			os.WriteFile(filename, []byte("Test file content"), os.ModePerm)
			os.Chtimes(filename, modTime, modTime)
		}
	}
}
