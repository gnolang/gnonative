package service

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// Starting the service must not race with the goroutine which serves the listener. Run with -race.
func TestServeAfterStart(t *testing.T) {
	// Not t.TempDir(), since its path can be too long for a Unix socket.
	dir, err := os.MkdirTemp("", "gnonative")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	svc, err := NewGnoNativeService(
		WithRootDir(dir),
		WithTmpDir(dir),
		WithUdsPath(filepath.Join(dir, "gno.sock")),
		WithTcpAddr("127.0.0.1:0"),
		WithUseTcpListener(),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer svc.Close()

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/", svc.GetTcpPort()))
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
}
