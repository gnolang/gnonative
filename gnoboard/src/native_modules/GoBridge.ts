import { GoBridgeInterface } from "./GoBridgeInterface";

class NoopGoBridge implements GoBridgeInterface {
  initBridge() {
    return Promise.reject();
  }

  closeBridge() {
    return Promise.reject();
  }

<<<<<<< HEAD
  createDefaultAccount(_: string) {
    return Promise.reject();
  }

  createReply(_: string, __: string) {
=======
  clientExec(_: string) {
>>>>>>> 0abf827 (Update GoBridge to use clientExec.)
    return Promise.reject();
  }

  exportJsonConfig() {
    return Promise.reject();
  }
}

export const GoBridge: GoBridgeInterface = new NoopGoBridge();
