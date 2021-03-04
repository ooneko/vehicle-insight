package version

import (
	"fmt"
	"runtime"
)

var (
	gitVersion = "v0.0.0"
	gitCommit  = "unknown"
	buildDate  = "unknown"
)

type Info struct {
	GitVersion string `json:"gitVersion"`
	GitCommit  string `json:"gitCommmit`
	BuildDate  string `json:"buildDate"`
	Platform   string `json:platform`
}

func Get() Info {
	return Info{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
