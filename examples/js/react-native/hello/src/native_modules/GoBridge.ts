import { GoBridgeInterface } from './GoBridgeInterface';

class NoopGoBridge implements GoBridgeInterface {
  initBridge() {
    return Promise.reject();
  }

  closeBridge() {
    return Promise.reject();
  }

  getTcpPort() {
    return Promise.reject();
  }

  invokeGrpcMethod() {
    return Promise.reject();
  }

  createStreamClient() {
    return Promise.reject();
  }

  streamClientReceive() {
    return Promise.reject();
  }

  closeStreamClient() {
    return Promise.reject();
  }
}

export const GoBridge: GoBridgeInterface = new NoopGoBridge();
