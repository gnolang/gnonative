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
    var bridge: GnoGnomobileBridge?
    var socketPath: String?
    var channel: GRPCChannel?
    var client: Gnomobile_V1_GnomobileServiceNIOClient?

    static func requiresMainQueueSetup() -> Bool {
        return true
    }

    override init() {
        self.logger = Logger(
            subsystem: Bundle.main.bundleIdentifier!,
            category: String(describing: "gobridge")
        )
        self.appRootDir = try! RootDirGet()
        super.init()
    }

    deinit {
        do {
            if self.bridge != nil {
                try self.bridge?.close()
                self.bridge = nil

                // Close the gRPC connection
                try! self.channel?.close().wait()
            }
        } catch let error as NSError {
            self.logger.error("\(String(describing: error.code))")
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

            let bridge = GnoGnomobileNewBridge(config, &err);
            if err != nil {
                throw err!
            }
            self.bridge = bridge

            guard let socketPath = bridge?.getSocketPath() else {
                throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "unable to get the socket path from the bridge"])
            }
            self.socketPath = socketPath

            // init the gPRC client

            let group = PlatformSupport.makeEventLoopGroup(loopCount: 1)
            defer {
                try? group.syncShutdownGracefully()
            }

            // Configure the channel, we're not using TLS so the connection is `insecure`.
            let channel = try GRPCChannelPool.with(
                target: .unixDomainSocket(self.socketPath!),
//                target: .socketAddress(SocketAddress(unixDomainSocketPath: self.socketPath!)),
                transportSecurity: .plaintext,
                eventLoopGroup: group
            )
            self.channel = channel

            self.client = Gnomobile_V1_GnomobileServiceNIOClient(channel: self.channel!)

            resolve(true)
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.localizedDescription, error)
        }
    }

    @objc func call(_ packagePath: NSString, fnc: NSString, args: NSArray, gasFee: NSString, gasWanted: NSNumber, password: NSString, resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
        self.logger.info("remi: calling call")
        do {
            guard self.client != nil else {
                throw NSError(domain: "land.gno.gnomobile", code: 2, userInfo: [NSLocalizedDescriptionKey : "gRPC client not init"])
            }
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.userInfo.description, error)
        }
        self.logger.info("remi: call 2")
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

        let call = client!.call(request)
        self.logger.info("remi: call 3, args: \(argArray)")
        do {
            let response = try call.response.wait()
            self.logger.info("remi: call 4")
            resolve(response)
        } catch let error as NSError {
            reject("\(String(describing: error.code))", error.localizedDescription, error)
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
            }
        } catch let error as NSError {
            self.logger.error("\(String(describing: error.code))")
        }
    }
}
