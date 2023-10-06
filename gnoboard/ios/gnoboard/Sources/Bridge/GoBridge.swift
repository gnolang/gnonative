import Foundation
import os
import GnoCore
import GRPC
import NIO

@available(iOS 14.0, *)
@objc(GoBridge)
class GoBridge: NSObject {
  let logger: Logger
  let appRootDir: String
  let tmpDir: String
  var bridge: GnoGnomobileBridge?
  var socketPort: Int = 0

  var eventLoopGroup: EventLoopGroup?
  var channel: GRPCChannel?
  var client: Land_Gno_Gnomobile_V1_GnomobileServiceAsyncClient?
  
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
        
        // Close the gRPC connection
        try! self.channel?.close().wait()
        
        try! self.eventLoopGroup?.syncShutdownGracefully()
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
        throw NSError(domain: "land.gno.gnomobile", code: 1, userInfo: [NSLocalizedDescriptionKey : "already started"])
      }
      
      // init the bridge service
      
      guard let config = GnoGnomobileBridgeConfig() else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "unable to create config"])
      }
      config.rootDir = self.appRootDir
      config.tmpDir = self.tmpDir

      // On simulator we can't create an UDS, see comment below
      #if !targetEnvironment(simulator)
      config.useUdsListener = true
      #endif

      config.useTcpListener = true

      let bridge = GnoGnomobileNewBridge(config, &err);
      if err != nil {
        throw err!
      }
      self.bridge = bridge
      
      // init the gRPC client

      /*
      ** On iOS simulator, temporary directory's absolute path exceeds
      ** the length limit for Unix Domain Socket, since simulator is
      ** only used for debug, we can safely fallback over TCP
      */
      #if !targetEnvironment(simulator)
      let socketPath = bridge!.getSocketPath()
      self.logger.info("gRPC server socket path: \(socketPath)")
      let target: ConnectionTarget = .unixDomainSocket(socketPath)
      #else
      self.socketPort = bridge!.getTcpPort()
      self.logger.info("gRPC server port: \(self.socketPort)")
      let target: ConnectionTarget = .host("localhost", port: self.socketPort)
      #endif
      
      // init the gPRC client
      
      // With React Native, we have to use MultiThreadedEventLoopGroup
      // instead of PlatformSupport.makeEventLoopGroup
      self.eventLoopGroup = MultiThreadedEventLoopGroup(numberOfThreads: 1)
      
      // Configure the channel, we're not using TLS so the connection is `insecure`.
      self.channel = try GRPCChannelPool.with(
        target: target,
        transportSecurity: .plaintext,
        eventLoopGroup: self.eventLoopGroup!
      )
      
      self.client = Land_Gno_Gnomobile_V1_GnomobileServiceAsyncClient(channel: self.channel!)
      
      resolve(true)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
      return
    }
  }

  @objc func getTcpPort(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge == nil {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "bridge not init"])
      }
      resolve(self.socketPort)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
    }
  }

  @objc func setPassword(_ password: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }

    let req = Land_Gno_Gnomobile_V1_SetPasswordRequest.with {
      $0.password = password as String
    }

    Task {
      do {
        try await client!.setPassword(req)
        resolve(true)
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }

  @objc func generateRecoveryPhrase(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }

    Task {
      do {
        let resp = try await client!.generateRecoveryPhrase(Land_Gno_Gnomobile_V1_GenerateRecoveryPhraseRequest())
        resolve(resp.phrase)
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }

  @objc func listKeyInfo(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }

    Task {
      do {
        let resp = try await client!.listKeyInfo(Land_Gno_Gnomobile_V1_ListKeyInfoRequest())
        var list = [NSDictionary]()
        for keyInfo in resp.keys {
          list.append(GoBridge.convertKeyInfo(keyInfo))
        }
        resolve(list as NSArray)
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }

  @objc func createAccount(_ nameOrBech32: NSString, mnemonic: NSString, bip39Passwd: NSString, password: NSString, account: NSNumber, index: NSNumber, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }

    let req = Land_Gno_Gnomobile_V1_CreateAccountRequest.with {
      $0.nameOrBech32 = nameOrBech32 as String
      $0.mnemonic = mnemonic as String
      $0.bip39Passwd = bip39Passwd as String
      $0.password = password as String
      $0.account = account as? UInt32 ?? 0
      $0.index = index as? UInt32 ?? 0
    }

    Task {
      do {
        let resp = try await client!.createAccount(req)
        resolve(GoBridge.convertKeyInfo(resp.key))
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }

  @objc func selectAccount(_ nameOrBech32: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }

    let req = Land_Gno_Gnomobile_V1_SelectAccountRequest.with {
      $0.nameOrBech32 = nameOrBech32 as String
    }

    Task {
      do {
        let resp = try await client!.selectAccount(req)
        resolve(GoBridge.convertKeyInfo(resp.key))
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }

  @objc func getActiveAccount(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }
    
    Task {
      do {
        let resp = try await client!.getActiveAccount(Land_Gno_Gnomobile_V1_GetActiveAccountRequest())
        resolve(GoBridge.convertKeyInfo(resp.key))
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }
  
  @objc func query(_ path: NSString, data_b64: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    var data: Data?
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }

      if let asData = (data_b64 as String).data(using: .utf8) {
        data = Data(base64Encoded: asData)
      } else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "query: can't get data_b64 as Data"])
      }
      if data == nil {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "query: data_b64 must be valid base64"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }

    let request = Land_Gno_Gnomobile_V1_QueryRequest.with {
      $0.path = path as String
      $0.data = data!
    }
    
    Task {
      do {
        let resp = try await client!.query(request)
        resolve(resp.result.base64EncodedString())
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }

  @objc func call(_ packagePath: NSString, fnc: NSString, args: NSArray, gasFee: NSString, gasWanted: NSNumber, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
      return
    }
    
    var argArray: [String] = []
    
    for arg in args {
      argArray.append(arg as? String ?? "")
    }
    
    let request = Land_Gno_Gnomobile_V1_CallRequest.with {
      $0.packagePath = packagePath as String
      $0.fnc = fnc as String
      $0.args = argArray
      $0.gasFee = gasFee as String
      $0.gasWanted = gasWanted as? Int64 ?? 0
    }
    
    Task {
      do {
        let resp = try await client!.call(request)
        resolve(resp.result.base64EncodedString())
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
      }
    }
  }
  
  @objc func exportJsonConfig(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      //            resolve(GnoGnomobileExportJsonConfig(self.appRootDir))
      resolve("{}")
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
    }
  }
  
  @objc func closeBridge(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      if self.bridge != nil {
        try self.bridge?.close()
        self.bridge = nil
        
        // Close the gRPC connection
        try! self.channel?.close().wait()
        
        try! self.eventLoopGroup?.syncShutdownGracefully()
      }
    } catch let error as NSError {
      self.logger.error("\(String(describing: error.code))")
    }
  }
  
  /**
   * Convert the gRPC type KeyInfo into an NSDictionary for passing to ReactNative.
   */
  public static func convertKeyInfo(_ keyInfo: Land_Gno_Gnomobile_V1_KeyInfo) -> NSDictionary {
    let dict: NSMutableDictionary = [:]
    dict["type"] = keyInfo.type.rawValue;
    dict["name"] = keyInfo.name;
    dict["address_b64"] = keyInfo.address.base64EncodedString()
    dict["pubKey_b64"] = keyInfo.pubKey.base64EncodedString();
    return dict;
  }
}
