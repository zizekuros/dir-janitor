# DirJanitor - A Go Package for Directory Cleanup

![Go Version](https://img.shields.io/github/go-mod/go-version/zizekuros/dir-janitor)
[![Latest Release](https://img.shields.io/badge/latest%20release-v0.1.0-blue)](https://github.com/zizekuros/dir-janitor/releases/tag/v0.1.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/zizekuros/dir-janitor)](https://goreportcard.com/report/github.com/zizekuros/dir-janitor)

DirJanitor is a Go package that provides an easy and automated way to clean up a directory based on file ages. You can specify the maximum age of files to keep and the frequency of the cleanup operation.

## Getting Started

### Installation

To use the DirJanitor package in your Go project, you can add it as a dependency to your project's `go.mod` file:

```shell
go get github.com/zizekuros/dir-janitor@v0.1.0
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

    // Start the cleanup interval
    cleaner.StartCleanupInterval()

    // Your application logic here

    // Stop the cleanup routine when you're done
    // Make sure to call this to release resources
    cleaner.StopCleanupInterval()

    log.Println("Cleanup stopped")
}
```
In the above example, the DirectoryCleaner is configured to clean up files older than 30 days at a frequency of 1 hour (3600 seconds).

## License
This package is open source and is available under the MIT License.

## Contributing
Contributions and feedback are welcome! Feel free to fork the repository, open issues, or submit pull requests.

If you encounter any issues or have questions, please don't hesitate to create an issue.

Happy cleaning!
