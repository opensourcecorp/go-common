package containertest

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest"
)

func TestRunTestDockerContainer(t *testing.T) {
	t.Run("basic nginx http", func(t *testing.T) {
		runOpts := &dockertest.RunOptions{
			Name:       "nginx",
			Repository: "docker.io/library/nginx",
			Tag:        "alpine",
			// I don't need it here, but users can specify a Docker network ID, such
			// as you might need/find on a CI/CD system
			NetworkID:    "",
			ExposedPorts: []string{"80/tcp"},
		}
		pool, container, stopContainer := RunTestDockerContainer(t, runOpts)
		defer stopContainer()

		addr := container.GetHostPort(runOpts.ExposedPorts[0])
		t.Logf("addr: %s", addr)

		err := pool.Retry(func() error {
			_, err := http.Get("http://" + addr)
			if err != nil {
				return fmt.Errorf("could not hit NGINX container: %v", err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("could not hit NGINX container in readiness deadline: %v", err)
		}

		t.Log("pass")
	})

	t.Run("postgres with a DDL", func(t *testing.T) {
		runOpts := &dockertest.RunOptions{
			Name:         "postgres",
			Repository:   "docker.io/library/postgres",
			Tag:          "alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env: []string{
				"POSTGRES_USER=test",
				"POSTGRES_PASSWORD=test",
			},
		}
		pool, container, stopContainer := RunTestDockerContainer(t, runOpts)
		defer stopContainer()

		addr := container.GetHostPort(runOpts.ExposedPorts[0])
		t.Logf("addr: %s", addr)

		ctx := context.Background()

		err := pool.Retry(func() error {
			db, err := pgx.Connect(ctx, fmt.Sprintf("postgres://test:test@%s/postgres", addr))
			if err != nil {
				return fmt.Errorf("could not hit Postgres container: %v", err)
			}
			defer db.Close(ctx)

			_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS test (id INTEGER);`)
			if err != nil {
				return fmt.Errorf("could not hit Postgres container: %v", err)
			}

			return nil
		})
		if err != nil {
			t.Fatalf("could not hit Postgres container in readiness deadline: %v", err)
		}

		t.Log("pass")
	})
}
