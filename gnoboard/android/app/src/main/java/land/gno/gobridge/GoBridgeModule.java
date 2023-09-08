package land.gno.gobridge;

import android.util.Log;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableArray;

import java.io.File;
import java.util.ArrayList;
import java.util.List;

import gnoland.gno.gnomobile.Gnomobile;
import gnoland.gno.gnomobile.Bridge;
import gnoland.gno.gnomobile.BridgeConfig;

import io.grpc.Channel;
import android.net.LocalSocketAddress.Namespace;

import io.grpc.StatusRuntimeException;
import land.gno.gnomobile.GnomobileServiceGrpc;
import land.gno.gnomobile.Gnomobiletypes;
import land.gno.udschannel.UdsChannelBuilder;

public class GoBridgeModule extends ReactContextBaseJavaModule {
    private final static String TAG = "GoBridge";
    private final ReactApplicationContext reactContext;
    private final File rootDir;
    private String socketPath;
    private GnomobileServiceGrpc.GnomobileServiceBlockingStub blockingStub;
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

            socketPath = bridgeGnomobile.getSocketPath();

            Channel channel = UdsChannelBuilder.forPath(socketPath, Namespace.FILESYSTEM).build();
            blockingStub = GnomobileServiceGrpc.newBlockingStub(channel);

            // TODO: to remove
            // Simulate account creation
            // To be handled by the frontend

            // Get the list of account and determine if we have to create a account
            Gnomobiletypes.ListKeyInfo.Request listKeyReq = Gnomobiletypes.ListKeyInfo.Request.newBuilder().build();
            Gnomobiletypes.ListKeyInfo.Reply listKeyRep;
            try {
                listKeyRep = blockingStub.listKeyInfo(listKeyReq);
            } catch (StatusRuntimeException e) {
                Log.d(TAG, String.format("RPC failed: {%s}", e.getStatus()));
                return;
            }

            Log.i(TAG, String.format("list accounts size: %d", listKeyRep.getKeysCount()));

            // if no account, create a new one
            if (listKeyRep.getKeysCount() == 0) {
                Gnomobiletypes.CreateAccount.Request createAccReq = Gnomobiletypes.CreateAccount.Request.newBuilder()
                    .setNameOrBech32("jefft0")
                    .setMnemonic("enable until hover project know foam join table educate room better scrub clever powder virus army pitch ranch fix try cupboard scatter dune fee")
                    .setBip39Passwd("")
                    .setPassword("password")
                    .setAccount(0)
                    .setIndex(0)
                    .build();
                Gnomobiletypes.CreateAccount.Reply createAccRep;
                try {
                    createAccRep = blockingStub.createAccount(createAccReq);
                } catch (StatusRuntimeException e) {
                    Log.d(TAG, String.format("RPC failed: {%s}", e.getStatus()));
                    return;
                }
            }

            // select the account
            Gnomobiletypes.SelectAccount.Request selectAccReq = Gnomobiletypes.SelectAccount.Request.newBuilder()
                .setNameOrBech32("jefft0")
                .build();
            Gnomobiletypes.SelectAccount.Reply selectAccRep;
            try {
                selectAccRep = blockingStub.selectAccount(selectAccReq);
            } catch (StatusRuntimeException e) {
                Log.d(TAG, String.format("RPC failed: {%s}", e.getStatus()));
                return;
            }

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
    public void call(String packagePath, String fnc, ReadableArray args, String gasFee, int gasWanted, String password, Promise promise) {
        List<String> argList = new ArrayList<>();
        for (int i = 0; i < args.size(); i++) {
            argList.add(args.getString(i));
        }

        Gnomobiletypes.Call.Request request = Gnomobiletypes.Call.Request.newBuilder()
            .setPackagePath(packagePath)
            .setFnc(fnc)
            .addAllArgs(argList)
            .setGasFee(gasFee)
            .setGasWanted(gasWanted)
            .setPassword(password)
            .build();

        Gnomobiletypes.Call.Reply reply;
        try {
            reply = blockingStub.call(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC failed: {%s}", e.getStatus()));
            return;
        }
        promise.resolve(reply.getResult().toString());
    }

    @ReactMethod
    public void exportJsonConfig(Promise promise) {
        try {
//            promise.resolve(Gnomobile.exportJsonConfig(rootDir.getAbsolutePath()));
            promise.resolve("{}");
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
