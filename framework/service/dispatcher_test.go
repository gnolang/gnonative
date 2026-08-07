package gnonative

import (
	"encoding/base64"
	"errors"
	"io"
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
	"github.com/gnolang/gnonative/v4/service"
)

func newTestService(t *testing.T) service.GnoNativeService {
	t.Helper()
	// No sockets: disable the UDS listener and don't use TCP.
	svc, err := service.NewGnoNativeService(service.WithDisableUdsListener())
	if err != nil {
		t.Fatalf("NewGnoNativeService: %v", err)
	}
	t.Cleanup(func() { svc.Close() })
	return svc
}

// TestDispatcherCompleteness is the drift guard between the two paths: every proto method must
// appear in exactly one dispatch map with matching streaming-ness.
func TestDispatcherCompleteness(t *testing.T) {
	svc := newTestService(t)
	unary := newUnaryHandlers(svc)
	streams := newStreamHandlers(svc)

	methods := api_gen.File_rpc_proto.Services().Get(0).Methods()
	for i := 0; i < methods.Len(); i++ {
		m := methods.Get(i)
		name := string(m.Name())
		_, inUnary := unary[name]
		_, inStream := streams[name]

		if inUnary && inStream {
			t.Errorf("method %s appears in both unary and stream maps", name)
			continue
		}
		if !inUnary && !inStream {
			t.Errorf("method %s is missing from both dispatch maps", name)
			continue
		}
		if m.IsStreamingServer() && !inStream {
			t.Errorf("streaming method %s is in the unary map", name)
		}
		if !m.IsStreamingServer() && !inUnary {
			t.Errorf("unary method %s is in the stream map", name)
		}
	}

	// Also assert the maps have no extra entries beyond the proto methods.
	total := methods.Len()
	if got := len(unary) + len(streams); got != total {
		t.Errorf("dispatch maps have %d entries, proto has %d methods", got, total)
	}
}

func TestDispatcherHelloUnary(t *testing.T) {
	svc := newTestService(t)
	d := newServiceDispatcher(svc).(*serviceDispatcher)

	out, err := d.invokeMethod("Hello", `{"name":"world"}`)
	if err != nil {
		t.Fatalf("invokeMethod Hello: %v", err)
	}
	raw, err := base64.StdEncoding.DecodeString(out)
	if err != nil {
		t.Fatalf("decode base64: %v", err)
	}
	var res api_gen.HelloResponse
	if err := protojson.Unmarshal(raw, &res); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if res.Greeting != "Hello world" {
		t.Errorf("unexpected greeting: %q", res.Greeting)
	}
	// protojson uses the proto JSON field name (here "Greeting"), not the Go snake_case tag.
	if !strings.Contains(string(raw), `"Greeting"`) {
		t.Errorf("expected protojson field name in output, got %q", string(raw))
	}
}

func TestDispatcherHelloStream(t *testing.T) {
	svc := newTestService(t)
	d := newServiceDispatcher(svc).(*serviceDispatcher)

	id, err := d.createStream("HelloStream", `{"name":"world"}`)
	if err != nil {
		t.Fatalf("createStream HelloStream: %v", err)
	}

	// HelloStream sends 4 messages (2s apart) then returns nil -> EOF.
	got := 0
	for {
		out, err := d.streamReceive(id)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			t.Fatalf("streamReceive: %v", err)
		}
		raw, decErr := base64.StdEncoding.DecodeString(out)
		if decErr != nil {
			t.Fatalf("decode base64: %v", decErr)
		}
		var res api_gen.HelloStreamResponse
		if decErr := protojson.Unmarshal(raw, &res); decErr != nil {
			t.Fatalf("unmarshal: %v", decErr)
		}
		if res.Greeting != "Hello world" {
			t.Errorf("unexpected greeting: %q", res.Greeting)
		}
		got++
	}
	if got != 4 {
		t.Errorf("expected 4 stream messages, got %d", got)
	}
	// After EOF the stream is still registered (matching the legacy path); closing succeeds once.
	if err := d.closeStream(id); err != nil {
		t.Errorf("closeStream after EOF: %v", err)
	}
	// A second close must fail: the stream is already unregistered.
	if err := d.closeStream(id); err == nil {
		t.Errorf("expected error closing already-closed stream")
	}
}

func TestDispatcherErrorCode(t *testing.T) {
	svc := newTestService(t)
	d := newServiceDispatcher(svc).(*serviceDispatcher)

	// GetActivatedAccount with no address returns ErrCode_ErrInvalidAddress.
	_, err := d.invokeMethod("GetActivatedAccount", `{}`)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "ErrInvalidAddress") {
		t.Errorf("expected ErrInvalidAddress in error, got %q", err.Error())
	}
}
