import { createPromiseClient } from '@connectrpc/connect';

import { createNativeGrpcTransport } from './transport_native';
import { GnoNativeService } from '../api/rpc_connect';

// Create a GnoNativeService client
export function createClient(_port: number) {
  const transport = createNativeGrpcTransport({
    baseUrl: '',
    useBinaryFormat: false,
  });

  return createPromiseClient(GnoNativeService, transport);
}
