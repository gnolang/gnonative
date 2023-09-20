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
  var socketPath: String?
  
  var eventLoopGroup: EventLoopGroup?
  var channel: GRPCChannel?
  var client: Gnomobile_V1_GnomobileServiceAsyncClient?
  
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
      
      let bridge = GnoGnomobileNewBridge(config, &err);
      if err != nil {
        throw err!
      }
      self.bridge = bridge
      
      self.socketPath = bridge?.getSocketPath()
      self.logger.info("socket path: \(self.socketPath!)")
      
      // init the gPRC client
      
      self.eventLoopGroup = MultiThreadedEventLoopGroup(numberOfThreads: 1)
      
      // Configure the channel, we're not using TLS so the connection is `insecure`.
      self.channel = try GRPCChannelPool.with(
        target: .unixDomainSocket(self.socketPath!),
        transportSecurity: .plaintext,
        eventLoopGroup: self.eventLoopGroup!
      )
      
      self.client = Gnomobile_V1_GnomobileServiceAsyncClient(channel: self.channel!)
      
      //            resolve(true)
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.localizedDescription, error)
      return
    }
    
    // TODO: to remove
    // Simulate account creation
    // To be handled by the frontend
    
    Task {
      // Get the list of account and determine if we have to create a account
      let listKeyReq = Gnomobile_V1_ListKeyInfo.Request()
      var listKeyRep: Gnomobile_V1_ListKeyInfo.Reply?;
      
      do {
        listKeyRep = try await self.client?.listKeyInfo(listKeyReq)
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
        return
      }
      
      self.logger.info("list account size: \(listKeyRep!.keys.count)")
      
      // if no account, create a new one
      let createAccountReq = Gnomobile_V1_CreateAccount.Request.with {
        $0.nameOrBech32 = "jefft0"
        $0.mnemonic = "enable until hover project know foam join table educate room better scrub clever powder virus army pitch ranch fix try cupboard scatter dune fee"
        $0.bip39Passwd = ""
        $0.password = "password"
        $0.account = 0
        $0.index = 0
      }
      
      do {
        try await self.client?.createAccount(createAccountReq)
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
        return
      }
      
      // select the account
      let selectAccountReq = Gnomobile_V1_SelectAccount.Request.with {
        $0.nameOrBech32 = "jefft0"
      }
      
      do {
        try await self.client?.selectAccount(selectAccountReq)
      } catch let error as NSError {
        reject("\(String(describing: error.code))", error.localizedDescription, error)
        return
      }
      
      resolve(true)
    }
  }
  
  @objc func call(_ packagePath: NSString, fnc: NSString, args: NSArray, gasFee: NSString, gasWanted: NSNumber, password: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      guard self.client != nil else {
        throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
      }
    } catch let error as NSError {
      reject("\(String(describing: error.code))", error.userInfo.description, error)
    }
    
    var argArray: [String] = []
    
    for arg in args {
      argArray.append(arg as? String ?? "")
    }
    
    let request = Gnomobile_V1_Call.Request.with {
      $0.packagePath = packagePath as String
      $0.fnc = fnc as String
      $0.args = argArray
      $0.gasFee = gasFee as String
      $0.gasWanted = gasWanted as? Int64 ?? 0
      $0.password = password as String
    }
    
    Task {
      do {
        let resp = try await client!.call(request)
        resolve(resp.result)
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
}
