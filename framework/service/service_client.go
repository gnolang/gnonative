package gnonative

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/gnolang/gnonative/api/gen/go/_goconnect"
)

type PromiseBlock interface {
	CallResolve(reply string)
	CallReject(error error)
}

type ServiceClient interface {
	InvokeGrpcMethodWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string)
	CreateStreamClientWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string)
	StreamClientReceiveWithPromiseBlock(promise PromiseBlock, id string)
	CloseStreamClientWithPromiseBlock(promise PromiseBlock, id string)
}

type serviceClient struct {
	streamIds uint64
	streams   map[string]reflect.Value
	muStreams sync.RWMutex

	client _goconnect.GnoNativeServiceClient
}

func NewServiceClient(client _goconnect.GnoNativeServiceClient) ServiceClient {
	return &serviceClient{
		streamIds: 0,
		streams:   make(map[string]reflect.Value),
		client:    client,
	}
}

func (s *serviceClient) InvokeGrpcMethodWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string) {
	go func() {
		res, err := s.InvokeGrpcMethod(method, jsonMessage)
		// if an internal error occurred generate a new bridge error
		if err != nil {
			promise.CallReject(err)
			return
		}

		promise.CallResolve(res)
	}()
}

func (s *serviceClient) InvokeGrpcMethod(method string, jsonMessage string) (string, error) {
	refClient := reflect.ValueOf(s.client)

	refMethod := refClient.MethodByName(method)
	if !refMethod.IsValid() {
		return "", errors.Errorf("method not found: %s", method)
	}

	// create arguments for the method
	in := make([]reflect.Value, refMethod.Type().NumIn())
	in[0] = reflect.ValueOf(context.Background())

	refReqType := refMethod.Type().In(1)                     // **Request[Req] type
	refReqValue := reflect.New(refReqType.Elem())            // *Request[Req] type
	refMsgReqValue := refReqValue.Elem().FieldByName("Msg")  // *Request[Req].Msg
	refMsgValue := reflect.New(refMsgReqValue.Type().Elem()) // Req type

	err := protojson.Unmarshal([]byte(jsonMessage), refMsgValue.Interface().(proto.Message))
	if err != nil {
		return "", errors.Wrap(err, "unable to unmarshal request")
	}

	refMsgReqValue.Set(refMsgValue)
	in[1] = refReqValue

	outRaw := refMethod.Call(in)
	if len(outRaw) != 2 {
		return "", errors.Errorf("unexpected number of return values: %d", len(outRaw))
	}

	errValue := outRaw[1].Interface()
	if errValue != nil {
		return "", errors.Wrap(errValue.(error), "invoke bridge method error")
	}

	msg := outRaw[0].Elem().FieldByName("Msg").Interface()
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return "", errors.Wrap(err, "unable to marshal response")
	}
	jsonRes := base64.StdEncoding.EncodeToString(jsonMsg)

	return jsonRes, nil
}

func (s *serviceClient) CreateStreamClientWithPromiseBlock(promise PromiseBlock, method string, jsonMessage string) {
	go func() {
		streamId, err := s.CreateStreamClient(method, jsonMessage)
		// if an internal error occurred generate a new bridge error
		if err != nil {
			promise.CallReject(err)
			return
		}
		promise.CallResolve(streamId)
	}()
}

// CreateStreamClient create a stream and returns the stream ID if there are no errors.
func (s *serviceClient) CreateStreamClient(method string, jsonMessage string) (string, error) {
	refClient := reflect.ValueOf(s.client)

	refMethod := refClient.MethodByName(method)
	if !refMethod.IsValid() {
		return "", errors.Errorf("method not found: %s", method)
	}

	// create arguments for the method
	in := make([]reflect.Value, refMethod.Type().NumIn())
	in[0] = reflect.ValueOf(context.Background())

	refReqType := refMethod.Type().In(1)                     // **Request[Req] type
	refReqValue := reflect.New(refReqType.Elem())            // *Request[Req] type
	refMsgReqValue := refReqValue.Elem().FieldByName("Msg")  // *Request[Req].Msg
	refMsgValue := reflect.New(refMsgReqValue.Type().Elem()) // Req type

	err := protojson.Unmarshal([]byte(jsonMessage), refMsgValue.Interface().(proto.Message))
	if err != nil {
		return "", errors.Wrap(err, "unable to unmarshal request")
	}

	refMsgReqValue.Set(refMsgValue)
	in[1] = refReqValue

	outRaw := refMethod.Call(in)
	if len(outRaw) != 2 {
		return "", errors.Errorf("unexpected number of return values: %d", len(outRaw))
	}

	errValue := outRaw[1].Interface()
	if errValue != nil {
		return "", errors.Wrap(errValue.(error), "stream bridge method error")
	}

	streamId := strconv.FormatUint(atomic.AddUint64(&s.streamIds, 1), 16)
	stream := outRaw[0]
	s.registerStream(streamId, stream)

	return streamId, nil
}

func (s *serviceClient) StreamClientReceiveWithPromiseBlock(promise PromiseBlock, id string) {
	go func() {
		jsonRes, err := s.StreamClientReceive(id)
		// if an internal error occurred generate a new bridge error
		if err != nil {
			promise.CallReject(err)
			return
		}
		promise.CallResolve(jsonRes)
	}()
}

func (s *serviceClient) StreamClientReceive(id string) (string, error) {
	s.muStreams.Lock()
	stream := s.streams[id]
	s.muStreams.Unlock()

	refVals := stream.MethodByName("Receive").Call([]reflect.Value{})
	toContinue := refVals[0].Bool()
	if !toContinue {
		refErr := stream.MethodByName("Err").Call([]reflect.Value{})
		if refErr[0].Interface() != nil {
			err := refErr[0].Interface().(error)
			return "", errors.Wrap(err, "stream's reveived method error")
		}
		return "", io.EOF
	}

	refResMsg := stream.MethodByName("Msg").Call([]reflect.Value{})

	refMsg := refResMsg[0].Interface()

	jsonMsg, err := json.Marshal(refMsg)
	if err != nil {
		return "", errors.Wrap(err, "unable to marshal response")
	}

	jsonMessage := base64.StdEncoding.EncodeToString(jsonMsg)

	return jsonMessage, nil
}

// Close the given stream
func (s *serviceClient) CloseStreamClientWithPromiseBlock(promise PromiseBlock, id string) {
	go func() {
		err := s.closeStreamClient(id)

		if err != nil {
			err = errors.Wrap(err, "unable to close bridge stream")
			promise.CallReject(err)
			return
		}
		promise.CallResolve("")
	}()
}

func (s *serviceClient) closeStreamClient(id string) error {
	stream, err := s.getSream(id)
	if err != nil {
		return err
	}
	refErr := stream.MethodByName("Close").Call([]reflect.Value{})
	if refErr[0].Interface() != nil {
		err := refErr[0].Interface().(error)
		return err
	}
	return s.unregisterStream(id)
}

func (s *serviceClient) registerStream(id string, cstream reflect.Value) {
	s.muStreams.Lock()
	s.streams[id] = cstream
	s.muStreams.Unlock()
}

func (s *serviceClient) unregisterStream(id string) error {
	s.muStreams.Lock()
	defer s.muStreams.Unlock()
	if _, ok := s.streams[id]; !ok {
		return fmt.Errorf("invalid stream id")
	}

	delete(s.streams, id)

	return nil
}

func (s *serviceClient) getSream(id string) (reflect.Value, error) {
	s.muStreams.RLock()
	defer s.muStreams.RUnlock()

	if cstream, ok := s.streams[id]; ok {
		return cstream, nil
	}
	return reflect.Value{}, fmt.Errorf("invalid stream id")
}
