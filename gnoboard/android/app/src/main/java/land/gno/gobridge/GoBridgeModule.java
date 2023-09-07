package land.gno.gobridge;

import android.util.Log;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import java.io.File;
import java.util.logging.Level;

import gnoland.gno.gnomobile.Gnomobile;
import gnoland.gno.gnomobile.Bridge;
import gnoland.gno.gnomobile.BridgeConfig;

import io.grpc.Channel;
import android.net.LocalSocketAddress.Namespace;

import io.grpc.StatusRuntimeException;
import land.gno.gnomobile.GnomobileServiceGrpc;
import land.gno.udschannel.UdsChannelBuilder;

public class GoBridgeModule extends ReactContextBaseJavaModule {
    private final static String TAG = "GoBridge";
    private final ReactApplicationContext reactContext;
    private final File rootDir;
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

            bridgeGnomobile = Gnomobile.newBridge(config);

            String path = rootDir.getAbsolutePath() + "/gnomobile.sock";

//            Channel channel = UdsChannelBuilder.forPath(path, Namespace.FILESYSTEM).build();
//            GnomobileServiceGrpc.GnomobileServiceBlockingStub blockingStub = GnomobileServiceGrpc.newBlockingStub(channel);
//            Hello.Request request = Hello.Request.newBuilder().setName("d4ryl00").build();
//            Hello.Reply greeting;
//            try {
//                greeting = blockingStub.hello(request);
//            } catch (StatusRuntimeException e) {
//                Log.d("Gnomobile", String.format("RPC failed: {%s}", e.getStatus()));
//                return;
//            }

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
    public void clientExec(String command, Promise promise) {
        promise.resolve(Gnomobile.clientExec(command, rootDir.getAbsolutePath()));
    }

    @ReactMethod
    public void exportJsonConfig(Promise promise) {
        try {
            promise.resolve(Gnomobile.exportJsonConfig(rootDir.getAbsolutePath()));
        } catch (Exception err) {
            promise.reject(err);
        }
    }

    @Override
    public void finalize() {
        try {
        } catch (Exception e) {
        }
    }
}
