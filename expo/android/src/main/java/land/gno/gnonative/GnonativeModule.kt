package land.gno.gnonative

import android.content.Context
import android.util.Log
import expo.modules.kotlin.Promise
import expo.modules.kotlin.exception.CodedException
import expo.modules.kotlin.modules.Module
import expo.modules.kotlin.modules.ModuleDefinition
import gnolang.gno.gnonative.Bridge
import gnolang.gno.gnonative.BridgeConfig
import gnolang.gno.gnonative.Gnonative
import gnolang.gno.gnonative.PromiseBlock
import java.io.File

class GnonativeModule : Module() {
  private var context: Context? = null
  private var rootDir: File? = null
  private var socketPort = 0
  private var bridgeGnoNative: Bridge? = null

  // Each module class must implement the definition function. The definition consists of components
  // that describes the module's functionality and behavior.
  // See https://docs.expo.dev/modules/module-api for more details about available components.
  override fun definition() = ModuleDefinition {
    // Sets the name of the module that JavaScript code will use to refer to the module. Takes a string as an argument.
    // Can be inferred from module's class name, but it's recommended to set it explicitly for clarity.
    // The module will be accessible from `requireNativeModule('Gnonative')` in JavaScript.
    Name("Gnonative")

    // Sets constant properties on the module. Can take a dictionary or a closure that returns a dictionary.
    Constants()

    OnCreate {
      context = appContext.reactContext
      rootDir = context!!.filesDir
    }

    // Defines event names that the module can send to JavaScript.
    Events("onChange")

    AsyncFunction("initBridge") { promise: Promise ->
      try {
        val config: BridgeConfig = Gnonative.newBridgeConfig() ?: throw Exception("")
        config.rootDir = rootDir!!.absolutePath
        bridgeGnoNative = Gnonative.newBridge(config)

        promise.resolve(true)
      } catch (err: CodedException) {
        promise.reject(err)
      }
    }

    AsyncFunction("closeBridge") { promise: Promise ->
      try {
        bridgeGnoNative?.let {
          bridgeGnoNative!!.close()
          bridgeGnoNative = null
          promise.resolve(true)
        } ?: run {
          throw GoBridgeNotStartedError()
        }
      } catch (err: CodedException) {
        promise.reject(err)
      }
    }

    AsyncFunction("getTcpPort") { promise: Promise ->
    try {
      bridgeGnoNative?.let {
        socketPort = bridgeGnoNative!!.tcpPort.toInt()
        promise.resolve(socketPort)
      } ?: run {
        throw GoBridgeNotStartedError()
      }
    } catch (err: CodedException) {
      promise.reject(err)
      }
    }

    AsyncFunction("invokeGrpcMethod") { method: String, jsonMessage: String, promise: Promise ->
      try {
        bridgeGnoNative?.let {
          val promiseBlock: PromiseBlock = PromiseBlock(promise)
          bridgeGnoNative!!.invokeGrpcMethodWithPromiseBlock(promiseBlock, method, jsonMessage)
        } ?: run {
          throw GoBridgeNotStartedError()
        }
      } catch (err: CodedException) {
        promise.reject(err)
      }
    }

    AsyncFunction("createStreamClient") { method: String, jsonMessage: String, promise: Promise ->
      try {
        bridgeGnoNative?.let {
          val promiseBlock: PromiseBlock = PromiseBlock(promise)
          bridgeGnoNative!!.createStreamClientWithPromiseBlock(promiseBlock, method, jsonMessage)
        } ?: run {
          throw GoBridgeNotStartedError()
        }
      } catch (err: CodedException) {
        promise.reject(err)
      }
    }

    AsyncFunction("streamClientReceive") { id: String, promise: Promise ->
      try {
        bridgeGnoNative?.let {
          val promiseBlock: PromiseBlock = PromiseBlock(promise)
          bridgeGnoNative!!.streamClientReceiveWithPromiseBlock(promiseBlock, id)
        } ?: run {
          throw GoBridgeNotStartedError()
        }
      } catch (err: CodedException) {
        promise.reject(err)
      }
    }

    AsyncFunction("closeStreamClient") { id: String, promise: Promise ->
      try {
        bridgeGnoNative?.let {
          val promiseBlock: PromiseBlock = PromiseBlock(promise)
          bridgeGnoNative!!.closeStreamClientWithPromiseBlock(promiseBlock, id)
        } ?: run {
          throw GoBridgeNotStartedError()
        }
      } catch (err: CodedException) {
        promise.reject(err)
      }
    }

    // Defines a JavaScript function that always returns a Promise and whose native code
    // is by default dispatched on the different thread than the JavaScript runtime runs on.
    AsyncFunction("setValueAsync") { value: String ->
      // Send an event to JavaScript.
      sendEvent("onChange", mapOf(
        "value" to value
      ))
    }

    // Enables the module to be used as a native view. Definition components that are accepted as part of
    // the view definition: Prop, Events.
    View(GnonativeView::class) {
      // Defines a setter for the `name` prop.
      Prop("name") { view: GnonativeView, prop: String ->
        println(prop)
      }
    }
  }
}
