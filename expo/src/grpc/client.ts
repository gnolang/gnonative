import { createClient as cc, Client } from '@connectrpc/connect';

import { createNativeGrpcTransport } from './transport_native';
import { GnoNativeService } from '../api/vendor/rpc_pb';

// Create a GnoNativeService client
export function createClient(_port: number): Client<typeof GnoNativeService> {
  const transport = createNativeGrpcTransport({
    baseUrl: '',
    useBinaryFormat: false,
  });

  return cc(GnoNativeService, transport);
}
