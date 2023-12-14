package land.gno.gobridge;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import java.io.File;

import gnolang.gno.gnonative.Gnonative;
import gnolang.gno.gnonative.Bridge;
import gnolang.gno.gnonative.BridgeConfig;

public class GoBridgeModule extends ReactContextBaseJavaModule {
    private final static String TAG = "GoBridge";
    private final ReactApplicationContext reactContext;
    private final File rootDir;
    private int socketPort;
    private static Bridge bridgeGnoNative = null;

    public GoBridgeModule(ReactApplicationContext reactContext) {
        super(reactContext);
        this.reactContext = reactContext;

        rootDir = new File(new land.gno.rootdir.RootDirModule(reactContext).getRootDir());
    }

    @Override
    public String getName() {
        return "GoBridge";
    }

    //////////////
    // Protocol //
    //////////////

    @ReactMethod
    public void initBridge(Promise promise) {
        try {
            final BridgeConfig config = Gnonative.newBridgeConfig();
            if (config == null) {
                throw new Exception("");
            }

            config.setRootDir(rootDir.getAbsolutePath());
            config.setUseTcpListener(true);
            config.setDisableUdsListener(true);

            bridgeGnoNative = Gnonative.newBridge(config);
            socketPort = (int)bridgeGnoNative.getTcpPort();

            promise.resolve(true);
        } catch (Exception err) {
            promise.reject(err);
        }
    }

    @ReactMethod
    public void closeBridge(Promise promise) {
        try {
            if (bridgeGnoNative != null) {
                bridgeGnoNative.close();
                bridgeGnoNative = null;
            }
            promise.resolve(true);
        } catch (Exception err) {
            promise.reject(err);
        }
    }

    @ReactMethod
    public void getTcpPort(Promise promise) {
        if (bridgeGnoNative == null) {
            promise.reject(new Exception("bridge not init"));
            return ;
        }
        promise.resolve(socketPort);
    }

    @Override
    public void finalize() {
        try {
        } catch (Exception e) {
        }
    }
}
