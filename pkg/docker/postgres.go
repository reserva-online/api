package docker

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// Database connection
const (
	PostgresDBName       = "postgres"
	PostgresUsername     = "postgres"
	PostgresPassword     = "postgres"
	PostgresDatabaseHost = "localhost"
)

type DockerPostgres struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource
	port     string
}

// NewPostgres creates a new instance of postgres container
func NewPostgres() *DockerPostgres {
	return &DockerPostgres{}
}

func (d DockerPostgres) WithPort(port string) DockerPostgres {
	d.port = port
	return d
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (kc DockerPostgres) WithTestPort(t *testing.T) DockerPostgres {
	seed := hash(t.Name())
	port := seed % 3000
	basePort := 5432
	kc.port = strconv.Itoa(int(port) + basePort)
	return kc
}

func (dp *DockerPostgres) Wait() error {
	var errReturn error
	dataSource := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		PostgresUsername,
		PostgresPassword,
		PostgresDBName,
		PostgresDatabaseHost,
		dp.port,
	)
	for i := 0; i < 5; i++ {
		_, err := sqlx.Connect("postgres", dataSource)

		if err != nil {
			errReturn = err
		}
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return errReturn
}

func (dp *DockerPostgres) GetHost() string {
	return fmt.Sprintf("%s:%s", PostgresDatabaseHost, dp.port)
}

func (dp *DockerPostgres) CreateImage() (pool *dockertest.Pool, resource *dockertest.Resource, err error) {
	pool, err = dockertest.NewPool("")
	if err != nil {
		return nil, nil, err
	}

	opts := dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "11.9-alpine",
		ExposedPorts: []string{dp.port},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("5432/tcp"): {
				{HostIP: "0.0.0.0", HostPort: dp.port},
			},
		},
		Env: []string{
			fmt.Sprintf("POSTGRES_DB=%s", PostgresDBName),
			fmt.Sprintf("POSTGRES_USER=%s", PostgresUsername),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", PostgresPassword),
		},
	}

	resource, err = pool.RunWithOptions(&opts)

	if err != nil {
		return nil, nil, err
	}
	return
}

func (d *DockerPostgres) Stop() error {
	return d.pool.Purge(d.resource)
}

func (d *DockerPostgres) Start(t *testing.T) {
	if d.port == "" {
		t.Fatalf("please inform a port to dockerPostgres, ex: dockerPostgres{}.WithTestPort(t)")
	}

	pool, resource, err := d.CreateImage()
	if err != nil {
		t.Fatalf("failed to create docker test container %s", err)
	}

	d.pool = pool
	d.resource = resource

	if err := d.Wait(); err != nil {
		d.Stop()
		t.Fatalf("postgres container didnt start on time %s", err)
	}

}

func (d *DockerPostgres) GetPort() string {
	return d.port
}
