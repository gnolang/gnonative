import Foundation
import GnoCore

@objc(GoBridge)
class GoBridge: NSObject {
    let appRootDir: String

    static func requiresMainQueueSetup() -> Bool {
        return true
    }

    override init() {
        self.appRootDir = try! RootDirGet()
        super.init()
    }

    deinit {
        do {
        } catch let error as NSError {
        }
    }

    @objc func constantsToExport() -> [AnyHashable : Any]! {
        #if DEBUG_LOGS
        let debug = true;
        #else
        let debug = false;
        #endif
        return ["debug": debug];
    }

    // //////// //
    // Protocol //
    // //////// //

    @objc func initBridge(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
        do {
            resolve(true)
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.userInfo.description, error)
        }
    }

    @objc func createDefaultAccount(_ name: String, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
        do {
            var err: NSError?
            resolve(GnoGnomobileCreateDefaultAccount(self.appRootDir, &err))
            if err != nil {
                throw err!
            }
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.userInfo.description, error)
        }
    }

    @objc func createReply(_ message: String, rootdir: String, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
      do {
        resolve(GnoGnomobileCreateReply(message, self.appRootDir))
      } catch let error as NSError {
          reject("\(String(describing: error.code))", error.userInfo.description, error)
      }
  }

    @objc func exportJsonConfig(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
        do {
            resolve(GnoGnomobileExportJsonConfig(self.appRootDir))
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.userInfo.description, error)
        }
    }

    @objc func closeBridge(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
        do {
            resolve(true)
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.userInfo.description, error)
        }
    }
}
