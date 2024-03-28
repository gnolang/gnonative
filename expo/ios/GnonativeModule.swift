import ExpoModulesCore
import GnoCore

public class GnonativeModule: Module {
  // Each module class must implement the definition function. The definition consists of components
  // that describes the module's functionality and behavior.
  // See https://docs.expo.dev/modules/module-api for more details about available components.
  var logger: Logger
  var appRootDir: String
  var tmpDir: String
  var bridge: GnoGnonativeBridge?
  var socketPort: Int = 0
  public func definition() -> ModuleDefinition {
    // Sets the name of the module that JavaScript code will use to refer to the module. Takes a string as an argument.
    // Can be inferred from module's class name, but it's recommended to set it explicitly for clarity.
    // The module will be accessible from `requireNativeModule('Gnonative')` in JavaScript.
    Name("Gnonative")
    
    // Sets constant properties on the module. Can take a dictionary or a closure that returns a dictionary.
    Constants([
      "PI": Double.pi
    ])
    
    OnCreate {
      self.logger = Logger(
        category: String(describing: "GnoNative")
      )
      self.appRootDir = try! FileManager.default.url(for: .documentDirectory, in: .userDomainMask, appropriateFor: nil, create: true).path
      self.tmpDir = try! FileManager.default.compatTemporaryDirectory.path
    }
    
    OnDestroy {
      do {
        if self.bridge != nil {
          try self.bridge?.close()
          self.bridge = nil
        }
      } catch let error as NSError {
        self.logger.error("\(String(describing: error.code))")
      }
    }
    
    // Defines event names that the module can send to JavaScript.
    Events("onChange")
    
    AsyncFunction("initBridge") { (promise: Promise) in
      var err: NSError?
      
      do {
        if self.bridge != nil {
          throw GnoError(.alreadyStarted)
        }
        
        // init the bridge service
        
        guard let config = GnoGnonativeBridgeConfig() else {
          throw GnoError(.createConfig)
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
        
        resolve(true)
      } catch let error {
        promise.reject(err)
      }
    }
    
    AsyncFunction("getTcpPort") { (promise: Promise) in
      do {
        guard let service = self.bridge else {
          throw GnoError(.notStarted)
        }
        self.socketPort = service.getTcpPort()
        self.logger.info("gRPC server port: \(self.socketPort)")
        resolve(self.socketPort)
      } catch let error {
        promise.reject(err)
      }
    }
    
    AsyncFunction("closeBridge") { (promise: Promise) in
      do {
        guard let service = self.bridge else {
          throw GnoError(.notStarted)
        }
        
        service.close()
        self.bridge = nil
      } catch let error {
        promise.reject(err)
      }
    }
    
    AsyncFunction("invokeGrpcMethod") { (method: String, b64message: String, promise: Promise) in
      do {
        guard let service = self.bridge else {
          throw GnoError(.notStarted)
        }
        
        let block = PromiseBlock(promise: promise)
        service.invokeGrpcMethod(with: block, method: method as String, jsonMessage: jsonMessage as String)
      } catch let err {
        promise.reject(err)
      }
    }
    
    AsyncFunction("createStreamClient") { (method: String, b64message: String, promise: Promise) in
      do {
        guard let service = self.bridge else {
          throw GnoError(.notStarted)
        }
        
        let block = PromiseBlock(promise: promise)
        service.createStreamClient(with: block, method: method as String, jsonMessage: jsonMessage as String)
      } catch let err {
        promise.reject(err)
      }
    }
    
    AsyncFunction("streamClientReceive") { (method: String, b64message: String, promise: Promise) in
      do {
        guard let service = self.bridge else {
          throw GnoError(.notStarted)
        }
        
        let block = PromiseBlock(promise: promise)
        service.streamClientReceive(with: block, id_: id as String)
      } catch let err {
        promise.reject(err)
      }
    }
    
    AsyncFunction("closeStreamClient") { (id: String, promise: Promise) in
      do {
        guard let service = self.bridge else {
          throw GnoError(.notStarted)
        }
        
        let block = PromiseBlock(promise: promise)
        service.closeStreamClient(with: block, id_: id as String)
      } catch let err {
        promise.reject(err)
      }
    }

    // Enables the module to be used as a native view. Definition components that are accepted as part of the
    // view definition: Prop, Events.
    View(GnonativeView.self) {
      // Defines a setter for the `name` prop.
      Prop("name") { (view: GnonativeView, prop: String) in
        print(prop)
      }
    }
  }
}

extension FileManager {
  public var compatTemporaryDirectory: URL {
    if #available(iOS 10.0, *) {
      return temporaryDirectory
    } else {
      return (try? url(
        for: .itemReplacementDirectory,
        in: .userDomainMask,
        appropriateFor: nil,
        create: true)
      ) ?? URL(fileURLWithPath: NSTemporaryDirectory())
    }
  }
}
