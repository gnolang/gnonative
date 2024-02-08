import { createPromiseClient } from '@connectrpc/connect';
import { createNativeGrpcTransport } from './transport';
import { GnoNativeService } from '@gno/api/rpc_connect';

// Create a GnoNativeService client
export function createClient(port: number) {
  return createPromiseClient(
    GnoNativeService,
    createNativeGrpcTransport({
      baseUrl: 'http://127.0.0.1:' + port,
    }),
  );
}
