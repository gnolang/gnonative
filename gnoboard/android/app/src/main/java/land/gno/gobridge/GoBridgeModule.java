package land.gno.gobridge;

import android.util.Log;

import com.facebook.react.bridge.Arguments;
import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableArray;
import com.facebook.react.bridge.WritableArray;
import com.facebook.react.bridge.WritableMap;
import com.google.protobuf.ByteString;

import java.io.File;
import java.util.ArrayList;
import android.util.Base64 ;
import java.util.List;

import gnolang.gno.gnomobile.Gnomobile;
import gnolang.gno.gnomobile.Bridge;
import gnolang.gno.gnomobile.BridgeConfig;

import io.grpc.Channel;
import android.net.LocalSocketAddress.Namespace;

import io.grpc.StatusRuntimeException;
import land.gno.gnomobile.v1.GnomobileServiceGrpc;
import land.gno.gnomobile.v1.Rpc;
import land.gno.gnomobile.v1.Gnomobiletypes;
import land.gno.gnomobile.v1.Rpc.KeyInfo;
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
    public void setPassword(String password, Promise promise) {
        Gnomobiletypes.SetPassword_Request request = Gnomobiletypes.SetPassword_Request.newBuilder()
            .setPassword(password)
            .build();

        Gnomobiletypes.SetPassword_Reply reply;
        try {
            reply = blockingStub.setPassword(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC setPassword failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }
        promise.resolve(true);
    }

    @ReactMethod
    public void generateRecoveryPhrase(Promise promise) {
        Gnomobiletypes.GenerateRecoveryPhrase_Request request = Gnomobiletypes.GenerateRecoveryPhrase_Request.newBuilder()
            .build();
        Gnomobiletypes.GenerateRecoveryPhrase_Reply reply;
        try {
            reply = blockingStub.generateRecoveryPhrase(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC generateRecoveryPhrase failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }
        promise.resolve(reply.getPhrase().toString());
    }

    @ReactMethod
    public void listKeyInfo(Promise promise) {
        Rpc.ListKeyInfo.Request request = Rpc.ListKeyInfo.Request.newBuilder()
            .build();
        Rpc.ListKeyInfo.Reply reply;
        try {
            reply = blockingStub.listKeyInfo(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC listKeyInfo failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }

        List<Rpc.KeyInfo> listKey = reply.getKeysList();
        WritableArray promiseArray= Arguments.createArray();
        for(int i=0;i<listKey.size();i++){
            promiseArray.pushMap(convertKeyInfo(listKey.get(i)));
        }

        promise.resolve(promiseArray);
    }

    @ReactMethod
    public void createAccount(String nameOrBech32, String mnemonic, String bip39Passwd, String password, int account, int index, Promise promise) {
        Rpc.CreateAccount.Request request = Rpc.CreateAccount.Request.newBuilder()
            .setNameOrBech32(nameOrBech32)
            .setMnemonic(mnemonic)
            .setBip39Passwd(bip39Passwd)
            .setPassword(password)
            .setAccount(account)
            .setIndex(index)
            .build();

        Rpc.CreateAccount.Reply reply;
        try {
            reply = blockingStub.createAccount(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC createAccount failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }

        promise.resolve(convertKeyInfo(reply.getKey()));
    }

    @ReactMethod
    public void selectAccount(String nameOrBech32, Promise promise) {
        Rpc.SelectAccount.Request request = Rpc.SelectAccount.Request.newBuilder()
            .setNameOrBech32(nameOrBech32)
            .build();

        Rpc.SelectAccount.Reply reply;
        try {
            reply = blockingStub.selectAccount(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC selectAccount failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }
        promise.resolve(convertKeyInfo(reply.getKey()));
    }

    @ReactMethod
    public void getActiveAccount(Promise promise) {
        Rpc.GetActiveAccount.Request request = Rpc.GetActiveAccount.Request.newBuilder()
            .build();
        Rpc.GetActiveAccount.Reply reply;
        try {
            reply = blockingStub.getActiveAccount(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC getActiveAccount failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }
        promise.resolve(convertKeyInfo(reply.getKey()));
    }

    @ReactMethod
    public void query(String path, String data_b64, Promise promise) {
        Gnomobiletypes.Query_Request request = Gnomobiletypes.Query_Request.newBuilder()
            .setPath(path)
            .setData(ByteString.copyFrom(Base64.decode(data_b64, Base64.DEFAULT)))
            .build();

        Gnomobiletypes.Query_Reply reply;
        try {
            reply = blockingStub.query(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC call failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }
        promise.resolve(Base64.encodeToString(reply.getResult().toByteArray(), Base64.NO_WRAP));
    }

    @ReactMethod
    public void call(String packagePath, String fnc, ReadableArray args, String gasFee, int gasWanted, Promise promise) {
        List<String> argList = new ArrayList<>();
        for (int i = 0; i < args.size(); i++) {
            argList.add(args.getString(i));
        }

        Gnomobiletypes.Call_Request request = Gnomobiletypes.Call_Request.newBuilder()
            .setPackagePath(packagePath)
            .setFnc(fnc)
            .addAllArgs(argList)
            .setGasFee(gasFee)
            .setGasWanted(gasWanted)
            .build();

        Gnomobiletypes.Call_Reply reply;
        try {
            reply = blockingStub.call(request);
        } catch (StatusRuntimeException e) {
            Log.d(TAG, String.format("RPC call failed: {%s}", e.getStatus()));
            promise.reject(e);
            return;
        }
        promise.resolve(Base64.encodeToString(reply.getResult().toByteArray(), Base64.NO_WRAP));
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

    /**
     * Convert the gRPC type KeyInfo into a WritableMap for passing to ReactNative.
     */
    public static WritableMap convertKeyInfo(KeyInfo keyInfo) {
        WritableMap map = Arguments.createMap();
        map.putInt("type", keyInfo.getType().getNumber());
        map.putString("name", keyInfo.getName());
        map.putString("address_b64", Base64.encodeToString(keyInfo.getAddress().toByteArray(), Base64.NO_WRAP));
        map.putString("pubKey_b64", Base64.encodeToString(keyInfo.getPubKey().toByteArray(), Base64.NO_WRAP));
        return map;
    }

    @Override
    public void finalize() {
        try {
        } catch (Exception e) {
        }
    }
}
