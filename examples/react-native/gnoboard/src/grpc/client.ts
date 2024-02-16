import { createPromiseClient } from '@connectrpc/connect';
import { createNativeGrpcTransport } from './transport_native';
import { GnoNativeService } from '@gno/api/rpc_connect';

// Create a GnoNativeService client
export function createClient(_port: number) {
  return createPromiseClient(
    GnoNativeService,
    createNativeGrpcTransport({
      baseUrl: '',
    }),
  );
}
