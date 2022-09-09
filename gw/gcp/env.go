package gcp

import "os"

func OnGCP() bool {
	return OnAppEngine() || OnCloudRun()
}

func OnAppEngine() bool {
	return os.Getenv("GAE_APPLICATION") != ""
}

func OnCloudRun() bool {
	return os.Getenv("K_SERVICE") != ""
}
