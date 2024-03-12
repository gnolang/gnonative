import Foundation
import os
import GnoCore

@available(iOS 14.0, *)
@objc(GoBridge)
class GoBridge: NSObject {
  let logger: Logger
  let appRootDir: String
  let tmpDir: String
  var bridge: GnoGnonativeBridge?
  var socketPort: Int = 0
  
  static func requiresMainQueueSetup() -> Bool {
    return false
  }
  
  override init() {
    self.logger = Logger(
      subsystem: Bundle.main.bundleIdentifier!,
      category: String(describing: "gobridge")
    )
    self.appRootDir = try! RootDirGet()
    self.tmpDir = try! TempDirGet()
    
    super.init()
  }
  
  deinit {
    do {
      if self.bridge != nil {
        try self.bridge?.close()
        self.bridge = nil
      }
    } catch let error as NSError {
      self.logger.error("\(String(describing: error.code))")
    }
  }
  
  @objc func constantsToExport() -> [AnyHashable : Any]! {
#if DEBUG
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
    var err: NSError?
    
    do {
      if self.bridge != nil {
        throw NSError(domain: "land.gno.gnonative", code: 1, userInfo: [NSLocalizedDescriptionKey : "already started"])
      }
      
      // init the bridge service
      
      guard let config = GnoGnonativeBridgeConfig() else {
        throw NSError(domain: "land.gno.gnonative", code: 2, userInfo: [NSLocalizedDescriptionKey : "unable to create config"])
      }
      config.rootDir = self.appRootDir
      config.tmpDir = self.tmpDir

      // On simulator we can't create an UDS, see comment below
#if targetEnvironment(simulator)
      config.useTcpListener = true
      config.disableUdsListener = true
#endif
      
      let bridge = GnoGnonativeNewBridge(config, &err);
      if err != nil {
        throw err!
      }
      self.bridge = bridge

      self.socketPort = bridge!.getTcpPort()
      self.logger.info("gRPC server port: \(self.socketPort)")
      
      resolve(true)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
      return
    }
  }

  @objc func getTcpPort(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge == nil {
        throw NSError(domain: "land.gno.gnonative", code: 2, userInfo: [NSLocalizedDescriptionKey : "bridge not init"])
      }
      resolve(self.socketPort)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
    }
  }
  
  @objc func closeBridge(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge != nil {
        try self.bridge?.close()
        self.bridge = nil
      }
    } catch let error as NSError {
      self.logger.error("\(String(describing: error.code))")
    }
  }
  
  @objc func invokeGrpcMethod(_ method: NSString, jsonMessage: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge == nil {
        throw NSError(domain: "land.gno.gnonative", code: 2, userInfo: [NSLocalizedDescriptionKey : "bridge not init"])
      }
      let promise = PromiseBlock(resolve, reject)
      self.bridge?.invokeGrpcMethod(with: promise, method: method as String, jsonMessage: jsonMessage as String)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
    }
  }
  
  @objc func createStreamClient(_ method: NSString, jsonMessage: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge == nil {
        throw NSError(domain: "land.gno.gnonative", code: 2, userInfo: [NSLocalizedDescriptionKey : "bridge not init"])
      }
      let promise = PromiseBlock(resolve, reject)
      self.bridge?.createStreamClient(with: promise, method: method as String, jsonMessage: jsonMessage as String)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
    }
  }
  
  @objc func streamClientReceive(_ id: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge == nil {
        throw NSError(domain: "land.gno.gnonative", code: 2, userInfo: [NSLocalizedDescriptionKey : "bridge not init"])
      }
      let promise = PromiseBlock(resolve, reject)
      self.bridge?.streamClientReceive(with: promise, id_: id as String)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
    }
  }
  
  @objc func closeStreamClient(_ id: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge == nil {
        throw NSError(domain: "land.gno.gnonative", code: 2, userInfo: [NSLocalizedDescriptionKey : "bridge not init"])
      }
      let promise = PromiseBlock(resolve, reject)
      self.bridge?.closeStreamClient(with: promise, id_: id as String)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
    }
  }
}
