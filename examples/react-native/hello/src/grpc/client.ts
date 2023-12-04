import { createPromiseClient } from '@connectrpc/connect';
import { createXHRGrpcWebTransport } from './transport';
import { GnomobileService } from '@gno/api/rpc_connect';

// Create a GnomobileService client
export function createClient(port: number) {
  return createPromiseClient(
    GnomobileService,
    createXHRGrpcWebTransport({
      baseUrl: 'http://127.0.0.1:' + port,
    }),
  );
}
