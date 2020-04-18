package core

import "time"

const (
	DefaultLimit              = 20
	DefaultHTTPTimeOut        = 10 * time.Second
	DefaultSyncTimeout        = 1 * time.Hour
	DefaultGoRequestRetry     = 3
	DefaultGoRequestRetryTime = 5 * time.Second
	DefaultDockerRepo         = "docker.io"
	DefaultK8sRepo            = "k8s.gcr.io"
	ManifestDir               = "manifests"
	ChangeLog                 = "CHANGELOG-%s.md"
	DockerHubImage            = "https://hub.docker.com/v2/repositories/%s/?page_size=100"
	DockerHubTags             = "https://hub.docker.com/v2/repositories/%s/%s/tags/?page_size=100"

	defaultSyncRetry     = 3
	defaultSyncRetryTime = 10 * time.Second
)

func retry(count int, interval time.Duration, f func() error) error {
	var err error
redo:
	if err = f(); err != nil {
		if count > 0 {
			count--
			if interval > 0 {
				<-time.After(interval)
			}
			goto redo
		}
	}
	return err
}