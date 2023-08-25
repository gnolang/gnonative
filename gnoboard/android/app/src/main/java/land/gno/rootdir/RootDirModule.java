package land.gno.rootdir;

import android.content.Context;

import androidx.annotation.NonNull;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import java.io.File;

public class RootDirModule extends ReactContextBaseJavaModule {
    private final static String nameFolder = "gnomobile";

    public RootDirModule(ReactApplicationContext reactContext) {
        super(reactContext);
    }

    public String getRootDir() {
        String rootDir = getReactApplicationContext().getFilesDir().getAbsolutePath();
        return new File(rootDir + "/" + nameFolder).getAbsolutePath();
    }

    @NonNull
    @Override
    public String getName() {
        return "RootDir";
    }

    @ReactMethod
    public void get(Promise promise) {
        promise.resolve(getRootDir());
    }
}

