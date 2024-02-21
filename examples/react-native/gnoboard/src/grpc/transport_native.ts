import type { AnyMessage, MethodInfo, PartialMessage, ServiceType } from '@bufbuild/protobuf';

import type { StreamResponse, Transport, UnaryRequest, UnaryResponse } from '@connectrpc/connect';
import { createClientMethodSerializers, createMethodUrl, runStreamingCall, runUnaryCall } from '@connectrpc/connect/protocol';
import { requestHeader } from '@connectrpc/connect/protocol-connect';
import { requestHeader as webRequestHeader } from '@connectrpc/connect/protocol-grpc-web';
import { GrpcWebTransportOptions } from '@connectrpc/connect-web';
import { Message, MethodKind } from '@bufbuild/protobuf';
import { GoBridge } from '@gno/native_modules';

function base64ToBytes(base64: string): Uint8Array {
  const binString = atob(base64);
  return Uint8Array.from(binString, (m) => m.codePointAt(0));
}

export function createNativeGrpcTransport(options: GrpcWebTransportOptions): Transport {
  const useBinaryFormat = options.useBinaryFormat ?? true;
  return {
    async unary<I extends Message<I> = AnyMessage, O extends Message<O> = AnyMessage>(
      service: ServiceType,
      method: MethodInfo<I, O>,
      signal: AbortSignal | undefined,
      timeoutMs: number | undefined,
      header: Headers,
      message: PartialMessage<I>,
    ): Promise<UnaryResponse<I, O>> {
      const { parse } = createClientMethodSerializers(method, false, options.jsonOptions, options.binaryOptions);

      return await runUnaryCall<I, O>({
        signal,
        interceptors: options.interceptors,
        req: {
          stream: false,
          service,
          method,
          url: createMethodUrl(options.baseUrl, service, method),
          init: {
            method: 'POST',
            mode: 'cors',
          },
          header: webRequestHeader(useBinaryFormat, timeoutMs, header),
          message,
        },
        next: async (req: UnaryRequest<I, O>): Promise<UnaryResponse<I, O>> => {
          try {
            const res = await GoBridge.invokeGrpcMethod(req.method.name, req.message.toJsonString());

            const header: Headers | undefined = new Headers();
            const trailer: Headers | undefined = new Headers();

            const data = base64ToBytes(res);
            const message = parse(data);

            return <UnaryResponse<I, O>>{
              stream: false,
              header,
              message,
              trailer,
            };
          } catch (e) {
            console.log('next: unary call error:', e);
            throw e;
          }
        },
      });
    },

    async stream<I extends Message<I> = AnyMessage, O extends Message<O> = AnyMessage>(
      service: ServiceType,
      method: MethodInfo<I, O>,
      signal: AbortSignal | undefined,
      timeoutMs: number | undefined,
      header: HeadersInit | undefined,
      input: AsyncIterable<PartialMessage<I>>,
    ): Promise<StreamResponse<I, O>> {
      const { parse } = createClientMethodSerializers(method, false, options.jsonOptions, options.binaryOptions);

      async function createRequestBody(input: AsyncIterable<I>): Promise<string> {
        if (method.kind != MethodKind.ServerStreaming) {
          throw 'The fetch API does not support streaming request bodies';
        }

        const r = await input[Symbol.asyncIterator]().next();
        if (r.done == true) {
          throw 'missing request message';
        }

        return r.value.toJsonString();
      }

      return await runStreamingCall<I, O>({
        interceptors: options.interceptors,
        timeoutMs,
        signal,
        req: {
          stream: true,
          service,
          method,
          url: createMethodUrl(options.baseUrl, service, method),
          init: {
            method: 'POST',
            credentials: options.credentials ?? 'same-origin',
            mode: 'cors',
          },
          header: requestHeader(method.kind, useBinaryFormat, timeoutMs, header),
          message: input,
        },
        next: async (req) => {
          const header: Headers | undefined = new Headers();
          const trailer: Headers | undefined = new Headers();

          const body = await createRequestBody(req.message);

          let streamId: string;
          try {
            streamId = await GoBridge.createStreamClient(req.method.name, body);
          } catch (e) {
            console.log('createStreamClient error:', e);
            throw e;
          }

          const generator = {
            async *[Symbol.asyncIterator]() {
              for (;;) {
                try {
                  const res = await GoBridge.streamClientReceive(streamId);
                  const data = base64ToBytes(res);
                  const message = parse(data);
                  yield message;
                } catch (e) {
                  // close the stream
                  try {
                    await GoBridge.closeStreamClient(streamId);
                  } catch (e) {
                    console.log('closeStreamClient error:', e);
                  }

                  if (!(e instanceof Error && e.message === 'EOF')) {
                    console.log('streamClientReceive error:', e);
                    throw e;
                  }
                  break;
                }
              }
            },
          };

          const res: StreamResponse<I, O> = {
            ...req,
            header: header,
            trailer,
            message: generator,
          };
          return res;
        },
      });
    },
  };
}
