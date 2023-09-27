import { Code, ConnectError, createPromiseClient } from "@connectrpc/connect";
import { createXHRGrpcWebTransport } from "./transport";
import { GnomobileService } from "@gno/api/rpc_connect";

// Create a GnomobileService client
export function createClient() {
  return createPromiseClient(
    GnomobileService,
    createXHRGrpcWebTransport({
      baseUrl: "https://demo.connectrpc.com",
    }),
  );
}
