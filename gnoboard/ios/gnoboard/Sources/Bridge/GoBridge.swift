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
