import type {
  DescMessage,
  DescMethodUnary,
  DescMethodStreaming,
  MessageShape,
  MessageInitShape,
} from '@bufbuild/protobuf';
import { toJsonString } from '@bufbuild/protobuf';
import type {
  ContextValues,
  StreamResponse,
  Transport,
  UnaryRequest,
  UnaryResponse,
} from '@connectrpc/connect';
import { createContextValues } from '@connectrpc/connect';
import {
  createClientMethodSerializers,
  createMethodUrl,
  runStreamingCall,
  runUnaryCall,
} from '@connectrpc/connect/protocol';
import { requestHeader } from '@connectrpc/connect/protocol-connect';
import { GrpcWebTransportOptions } from '@connectrpc/connect-web';
import { toByteArray } from 'base64-js';
import { CodedError } from 'expo-modules-core';

import { GoBridge } from '../GoBridge';

export function createNativeGrpcTransport(options: GrpcWebTransportOptions): Transport {
  const useBinaryFormat = options.useBinaryFormat ?? true;
  return {
    async unary<I extends DescMessage, O extends DescMessage>(
      method: DescMethodUnary<I, O>,
      signal: AbortSignal | undefined,
      timeoutMs: number | undefined,
      header: HeadersInit | undefined,
      message: MessageInitShape<I>,
      contextValues?: ContextValues,
    ): Promise<UnaryResponse<I, O>> {
      const { parse } = createClientMethodSerializers(
        method,
        useBinaryFormat,
        options.jsonOptions,
        options.binaryOptions,
      );

      return await runUnaryCall<I, O>({
        signal,
        interceptors: options.interceptors,
        req: {
          stream: false,
          service: method.parent,
          method,
          requestMethod: 'POST',
          url: createMethodUrl(options.baseUrl, method),
          header: requestHeader(method.methodKind, useBinaryFormat, timeoutMs, header, false),
          contextValues: contextValues ?? createContextValues(),
          message,
        },
        next: async (req: UnaryRequest<I, O>): Promise<UnaryResponse<I, O>> => {
          try {
            const messageJson = toJsonString(req.method.input, req.message);
            const res = await GoBridge.invokeGrpcMethod(req.method.name, messageJson);

            const header: Headers | undefined = new Headers();
            const trailer: Headers | undefined = new Headers();

            const data = toByteArray(res);
            const message = parse(data);

            return {
              stream: false,
              service: method.parent,
              method,
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

    async stream<I extends DescMessage, O extends DescMessage>(
      method: DescMethodStreaming<I, O>,
      signal: AbortSignal | undefined,
      timeoutMs: number | undefined,
      header: HeadersInit | undefined,
      input: AsyncIterable<MessageInitShape<I>>,
      contextValues?: ContextValues,
    ): Promise<StreamResponse<I, O>> {
      const { parse } = createClientMethodSerializers(
        method,
        useBinaryFormat,
        options.jsonOptions,
        options.binaryOptions,
      );

      async function createRequestBody(
        schema: DescMessage,
        input: AsyncIterable<MessageShape<I>>,
      ): Promise<string> {
        if (method.methodKind !== 'server_streaming') {
          throw 'The fetch API does not support streaming request bodies';
        }

        const r = await input[Symbol.asyncIterator]().next();
        if (r.done === true) {
          throw 'missing request message';
        }

        const messageJson = toJsonString(schema, r.value);
        return messageJson;
      }

      return await runStreamingCall<I, O>({
        interceptors: options.interceptors,
        timeoutMs,
        signal,
        req: {
          stream: true,
          service: method.parent,
          method,
          requestMethod: 'POST',
          url: createMethodUrl(options.baseUrl, method),
          header: requestHeader(method.methodKind, useBinaryFormat, timeoutMs, header, false),
          message: input,
          contextValues: contextValues ?? createContextValues(),
        },
        next: async req => {
          const header: Headers | undefined = new Headers();
          const trailer: Headers | undefined = new Headers();

          const body = await createRequestBody(req.method.input, req.message);

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
                  const data = toByteArray(res);
                  const message = parse(data);
                  yield message;
                } catch (e) {
                  // close the stream
                  try {
                    await GoBridge.closeStreamClient(streamId);
                  } catch (e) {
                    console.log('closeStreamClient error:', e);
                  }

                  if (!(e instanceof CodedError && e.message === 'EOF')) {
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
            header,
            trailer,
            message: generator,
          };
          return res;
        },
      });
    },
  };
}
