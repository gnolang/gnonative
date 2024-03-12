import { GoBridgeInterface } from "./GoBridgeInterface";

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
}

export const GoBridge: GoBridgeInterface = new NoopGoBridge();
