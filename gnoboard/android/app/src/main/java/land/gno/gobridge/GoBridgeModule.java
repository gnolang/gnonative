package land.gno.gobridge;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import java.io.File;

import gnolang.gno.gnomobile.Gnomobile;
import gnolang.gno.gnomobile.Bridge;
import gnolang.gno.gnomobile.BridgeConfig;

public class GoBridgeModule extends ReactContextBaseJavaModule {
    private final static String TAG = "GoBridge";
    private final ReactApplicationContext reactContext;
    private final File rootDir;
    private int socketPort;
    private static Bridge bridgeGnomobile = null;

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
            final BridgeConfig config = Gnomobile.newBridgeConfig();
            if (config == null) {
                throw new Exception("");
            }

            config.setRootDir(rootDir.getAbsolutePath());
            config.setUseTcpListener(true);
            config.setUseUdsListener(true);

            bridgeGnomobile = Gnomobile.newBridge(config);
            socketPort = (int)bridgeGnomobile.getTcpPort();

            promise.resolve(true);
        } catch (Exception err) {
            promise.reject(err);
        }
    }

    @ReactMethod
    public void closeBridge(Promise promise) {
        try {
            if (bridgeGnomobile != null) {
                bridgeGnomobile.close();
                bridgeGnomobile = null;
            }
            promise.resolve(true);
        } catch (Exception err) {
            promise.reject(err);
        }
    }

    @ReactMethod
    public void getTcpPort(Promise promise) {
        if (bridgeGnomobile == null) {
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
