package containertest

import (
	"testing"
	"time"

	"github.com/ory/dockertest"
)

/*
RunTestDockerContainer (presently) runs a Docker container to be used in tests.
It returns the dockertest.Pool & dockertest.Container resources, and a
deferrable function used to stop the started container. The container address
can be retrieved from the container.GetHostPort() method, and if desired, a
readiness probe func can be passed using the pool.Retry() method.

For examples of how to use this function, refer to its *own* tests, as they
serve as the most basic & straightforward usage.

In the future, this package is expected to support more than just Docker as a
container runtime.
*/
func RunTestDockerContainer(t *testing.T, runOpts *dockertest.RunOptions) (*dockertest.Pool, *dockertest.Resource, func()) {
	t.Helper()

	// rely on "sensible default" that dockertest takes per platform
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("could not construct a new dockertest pool: %v", err)
	}

	// Lower the default wait time for pool.Retry (callers can still override though)
	pool.MaxWait = 5 * time.Second

	t.Logf("starting container with opts: %+v", runOpts)
	container, err := pool.RunWithOptions(runOpts)
	if err != nil {
		t.Fatalf("could not run test container: %v", err)
	}

	stopContainer := func() {
		err := pool.Purge(container)
		if err != nil {
			t.Fatalf("could not purge test container: %v", err)
		}
	}

	return pool, container, stopContainer
}
