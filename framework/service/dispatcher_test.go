package gnonative

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	api_gen "github.com/gnolang/gnonative/v5/api"
	"github.com/gnolang/gnonative/v5/service"
)

func newTestService(t *testing.T) service.GnoNativeService {
	t.Helper()
	// There is no server; the service methods are called directly.
	svc, err := service.NewGnoNativeService()
	if err != nil {
		t.Fatalf("NewGnoNativeService: %v", err)
	}
	t.Cleanup(func() { svc.Close() })
	return svc
}

// TestDispatcherCompleteness is the drift guard for the dispatch maps: every method of
// service.GnoNativeApi must appear in exactly one map, matching its signature shape. Unary methods
// are func(ctx, *Req) (*Res, error) (2 in, 2 out); streaming methods are
// func(ctx, *Req, func(*Res) error) error (3 in, 1 out).
func TestDispatcherCompleteness(t *testing.T) {
	svc := newTestService(t)
	unary := newUnaryHandlers(svc)
	streams := newStreamHandlers(svc)

	apiType := reflect.TypeOf((*service.GnoNativeApi)(nil)).Elem()
	for i := 0; i < apiType.NumMethod(); i++ {
		m := apiType.Method(i)
		name := m.Name
		ft := m.Type
		_, inUnary := unary[name]
		_, inStream := streams[name]

		switch {
		case ft.NumIn() == 2 && ft.NumOut() == 2:
			if !inUnary {
				t.Errorf("unary method %s is missing from the unary map", name)
			}
			if inStream {
				t.Errorf("unary method %s also appears in the stream map", name)
			}
		case ft.NumIn() == 3 && ft.NumOut() == 1:
			if !inStream {
				t.Errorf("streaming method %s is missing from the stream map", name)
			}
			if inUnary {
				t.Errorf("streaming method %s also appears in the unary map", name)
			}
		default:
			t.Errorf("method %s has an unexpected signature (in=%d out=%d)", name, ft.NumIn(), ft.NumOut())
		}
	}

	if got, want := len(unary)+len(streams), apiType.NumMethod(); got != want {
		t.Errorf("dispatch maps have %d entries, GnoNativeApi has %d methods", got, want)
	}
}

func TestDispatcherHelloUnary(t *testing.T) {
	svc := newTestService(t)
	d := newServiceDispatcher(svc).(*serviceDispatcher)

	out, err := d.invokeMethod("Hello", `{"Name":"world"}`)
	if err != nil {
		t.Fatalf("invokeMethod Hello: %v", err)
	}
	raw, err := base64.StdEncoding.DecodeString(out)
	if err != nil {
		t.Fatalf("decode base64: %v", err)
	}
	var res api_gen.HelloResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if res.Greeting != "Hello world" {
		t.Errorf("unexpected greeting: %q", res.Greeting)
	}
	// The wire field uses the proto json_name "Greeting", not a lowerCamel default.
	if !strings.Contains(string(raw), `"Greeting"`) {
		t.Errorf("expected \"Greeting\" field name in output, got %q", string(raw))
	}
}

func TestDispatcherHelloStream(t *testing.T) {
	svc := newTestService(t)
	d := newServiceDispatcher(svc).(*serviceDispatcher)

	id, err := d.createStream("HelloStream", `{"Name":"world"}`)
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
		if decErr := json.Unmarshal(raw, &res); decErr != nil {
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
	// After EOF the stream is still registered; closing succeeds once.
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
