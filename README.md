# DirJanitor - A Go Package for Directory Cleanup

[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/zizekuros/dir-janitor/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/go%20version-1.20-blue)](https://tip.golang.org/doc/go1.20)
[![Latest Release](https://img.shields.io/badge/latest%20release-v0.3.0-blue)](https://github.com/zizekuros/dir-janitor/releases/tag/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/zizekuros/dir-janitor)](https://goreportcard.com/report/github.com/zizekuros/dir-janitor)
![Test Status](https://img.shields.io/badge/tests-passed-brightgreen)

DirJanitor is tiny [Go package](https://pkg.go.dev/github.com/zizekuros/dir-janitor/pkg) that provides an easy and automated way to clean up a directory based on file ages. You can specify the maximum age of files to keep and the frequency of the cleanup operation.

I created this package for one of my private projects and decided to share it in case it's helpful to others. 👌

### Alternative

If you're running on Linux, this is the shell command that you can use to achieve a similar job, so you don't need to use this package.

```shell
find /path/to/directory -type f -mtime +1 -delete
```

This command will delete files older than one day. Adjust the value after -mtime to match the desired retention period in days.

## Getting Started

### Installation

To use the DirJanitor package in your Go project, you can add it as a dependency to your project's `go.mod` file:

```shell
go get github.com/zizekuros/dir-janitor
```
This command will download and add the package to your project.

### Importing the Package
In your Go source code, import the dir-janitor package like this:

```go
import "github.com/zizekuros/dir-janitor/pkg"
```

You can also use an alias for the package to make it more concise:
```go
import dirjanitor "github.com/zizekuros/dir-janitor/pkg"
```

### Using the Package
Here's an example of how to use the DirJanitor package to automate directory cleanup:
```go
package main

import (
    "log"
    "time"
    dirjanitor "github.com/zizekuros/dir-janitor/pkg"
)

func main() {
    // Create a DirectoryCleaner instance
    cleaner := &dirjanitor.DirectoryCleaner{
        Directory: "/path/to/directory",
        Retention: 30,      // Max age of files to keep (in days)
        Frequency: 3600,    // Frequency of cleanup (in seconds)
    }

    // Set a custom logger for the cleaner (optional)
    // cleaner.Logger = log.New(os.Stdout, "DirJanitor: ", log.Ldate|log.Ltime)

    // Start the cleanup interval
    cleaner.StartCleanupInterval()

    // You can also perform cleanup job manually if you want like that:
    // cleaner.PerformCleanup()

    // Your application logic here

    // Stop the cleanup routine when you're done
    // Make sure to call this to release resources
    cleaner.StopCleanupInterval()
}
```
In the above example, the DirectoryCleaner is configured to clean up files older than 30 days at a frequency of 1 hour (3600 seconds).

### Running Tests
To run tests for the DirJanitor package, follow these steps:

1. Open a terminal or command prompt.
2. Navigate to the root directory of your Go project where the package is located.
3. Run the following command to execute the tests:

```shell
go test -count=1 ./tests/...
```

## Release Log

**v0.3.0** (Most Recent)
  - Added support for manually performing cleanup actions, addressing [Issue #2](https://github.com/zizekuros/dir-janitor/issues/2).
  - Additional unit test provided (separate tests for "PerformCleanup" and "CleanupInterval").
  - Improved documentation.

**v0.2.1**
  - Fixed spelling errors and improved documentation.

**v0.2.0**
  - Added support for custom Logger, addressing [Issue #1](https://github.com/zizekuros/dir-janitor/issues/1).
  - Added unit tests to ensure package stability.

**v0.1.0** 
- Initial Release

## License
This package is open source and is available under the MIT License.

## Contributing
Feel free to fork the repository, open issues, or submit pull requests. If you encounter any issues or have questions, please don't hesitate to create an issue.

Happy cleaning! 🧹
