import { GnoNativeService } from '@buf/gnolang_gnonative.connectrpc_es/rpc_connect';
import { createPromiseClient, PromiseClient } from '@connectrpc/connect';

import { createNativeGrpcTransport } from './transport_native';

// Create a GnoNativeService client
export function createClient(_port: number): PromiseClient<typeof GnoNativeService> {
  const transport = createNativeGrpcTransport({
    baseUrl: '',
    useBinaryFormat: false,
  });

  return createPromiseClient(GnoNativeService, transport);
}
