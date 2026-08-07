package gnonative

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"

	"github.com/gnolang/gnonative/v4/service"
)

// ServiceDispatcher is the gomobile-facing promise API for the connect/gRPC-free path. It mirrors
// ServiceClient but calls the plain service.GnoNativeApi directly (no in-process HTTP round-trip).
// Resolved payloads are base64(protojson) exactly like ServiceClient, so the native modules and the
// PromiseBlock plumbing are reused unchanged. On stream end, StreamReceive rejects with the literal
// "EOF" message (from io.EOF), which the JS side matches to terminate the AsyncIterable.
type ServiceDispatcher interface {
	InvokeMethodWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string)
	CreateStreamWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string)
	StreamReceiveWithPromiseBlock(promise PromiseBlock, id string)
	CloseStreamWithPromiseBlock(promise PromiseBlock, id string)
}

// dispatchStream adapts a push-based streaming method to the pull-based JS boundary. The method
// goroutine pushes protojson bytes into ch (cap 1, providing backpressure); Receive pulls them.
// When the method returns, its error (nil on clean end) is stored and finished is closed.
type dispatchStream struct {
	ch       chan []byte
	finished chan struct{}
	cancel   context.CancelFunc

	mu  sync.Mutex
	err error
}

type serviceDispatcher struct {
	svc     service.GnoNativeApi
	unary   map[string]unaryHandler
	streams map[string]streamHandler

	streamIds uint64
	active    map[string]*dispatchStream
	muActive  sync.RWMutex
}

func newServiceDispatcher(svc service.GnoNativeApi) ServiceDispatcher {
	return &serviceDispatcher{
		svc:     svc,
		unary:   newUnaryHandlers(svc),
		streams: newStreamHandlers(svc),
		active:  make(map[string]*dispatchStream),
	}
}

func (s *serviceDispatcher) InvokeMethodWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string) {
	go func() {
		res, err := s.invokeMethod(method, jsonMessage)
		if err != nil {
			promise.CallReject(err)
			return
		}
		promise.CallResolve(res)
	}()
}

func (s *serviceDispatcher) invokeMethod(method string, jsonMessage string) (string, error) {
	handler, ok := s.unary[method]
	if !ok {
		return "", errors.Errorf("method not found: %s", method)
	}

	out, err := handler(context.Background(), jsonMessage)
	if err != nil {
		return "", errors.Wrap(err, "invoke method error")
	}

	return base64.StdEncoding.EncodeToString(out), nil
}

func (s *serviceDispatcher) CreateStreamWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string) {
	go func() {
		id, err := s.createStream(method, jsonMessage)
		if err != nil {
			promise.CallReject(err)
			return
		}
		promise.CallResolve(id)
	}()
}

func (s *serviceDispatcher) createStream(method string, jsonMessage string) (string, error) {
	handler, ok := s.streams[method]
	if !ok {
		return "", errors.Errorf("method not found: %s", method)
	}

	ctx, cancel := context.WithCancel(context.Background())
	st := &dispatchStream{
		ch:       make(chan []byte, 1),
		finished: make(chan struct{}),
		cancel:   cancel,
	}

	// send blocks on the cap-1 channel (backpressure) and unblocks if the stream is closed.
	send := func(b []byte) error {
		select {
		case st.ch <- b:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	go func() {
		err := handler(ctx, jsonMessage, send)
		st.mu.Lock()
		st.err = err
		st.mu.Unlock()
		close(st.finished)
	}()

	id := strconv.FormatUint(atomic.AddUint64(&s.streamIds, 1), 16)
	s.registerStream(id, st)
	return id, nil
}

func (s *serviceDispatcher) StreamReceiveWithPromiseBlock(promise PromiseBlock, id string) {
	go func() {
		res, err := s.streamReceive(id)
		if err != nil {
			promise.CallReject(err)
			return
		}
		promise.CallResolve(res)
	}()
}

func (s *serviceDispatcher) streamReceive(id string) (string, error) {
	st, err := s.getStream(id)
	if err != nil {
		return "", err
	}

	select {
	case msg := <-st.ch:
		return base64.StdEncoding.EncodeToString(msg), nil
	case <-st.finished:
		// Deliver any message buffered before completion before reporting the end.
		select {
		case msg := <-st.ch:
			return base64.StdEncoding.EncodeToString(msg), nil
		default:
		}
		st.mu.Lock()
		methodErr := st.err
		st.mu.Unlock()
		if methodErr != nil {
			return "", errors.Wrap(methodErr, "stream method error")
		}
		// Return the raw io.EOF (message "EOF") so the JS side can match it literally.
		return "", io.EOF
	}
}

func (s *serviceDispatcher) CloseStreamWithPromiseBlock(promise PromiseBlock, id string) {
	go func() {
		if err := s.closeStream(id); err != nil {
			promise.CallReject(errors.Wrap(err, "unable to close stream"))
			return
		}
		promise.CallResolve("")
	}()
}

func (s *serviceDispatcher) closeStream(id string) error {
	st, err := s.getStream(id)
	if err != nil {
		return err
	}
	st.cancel()
	return s.unregisterStream(id)
}

func (s *serviceDispatcher) registerStream(id string, st *dispatchStream) {
	s.muActive.Lock()
	s.active[id] = st
	s.muActive.Unlock()
}

func (s *serviceDispatcher) unregisterStream(id string) error {
	s.muActive.Lock()
	defer s.muActive.Unlock()
	if _, ok := s.active[id]; !ok {
		return fmt.Errorf("invalid stream id")
	}
	delete(s.active, id)
	return nil
}

func (s *serviceDispatcher) getStream(id string) (*dispatchStream, error) {
	s.muActive.RLock()
	defer s.muActive.RUnlock()
	if st, ok := s.active[id]; ok {
		return st, nil
	}
	return nil, fmt.Errorf("invalid stream id")
}

// disabledServiceClient is the ServiceClient used when the gRPC servers are disabled. All its
// methods reject: callers must use the ServiceDispatcher (InvokeMethod/CreateStream/...) instead.
type disabledServiceClient struct{}

func newDisabledServiceClient() ServiceClient { return &disabledServiceClient{} }

var errGrpcDisabled = errors.New("grpc servers disabled; use InvokeMethod")

func (*disabledServiceClient) InvokeGrpcMethodWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string) {
	promise.CallReject(errGrpcDisabled)
}

func (*disabledServiceClient) CreateStreamClientWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string) {
	promise.CallReject(errGrpcDisabled)
}

func (*disabledServiceClient) StreamClientReceiveWithPromiseBlock(promise PromiseBlock, id string) {
	promise.CallReject(errGrpcDisabled)
}

func (*disabledServiceClient) CloseStreamClientWithPromiseBlock(promise PromiseBlock, id string) {
	promise.CallReject(errGrpcDisabled)
}
