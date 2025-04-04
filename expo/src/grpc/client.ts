import { createPromiseClient, PromiseClient } from '@connectrpc/connect';

import { createNativeGrpcTransport } from './transport_native';
import { GnoNativeService } from '../api/vendor/rpc_connect';

// Create a GnoNativeService client
export function createClient(_port: number): PromiseClient<typeof GnoNativeService> {
  const transport = createNativeGrpcTransport({
    baseUrl: '',
    useBinaryFormat: false,
  });

  return createPromiseClient(GnoNativeService, transport);
}
