package land.gno.gobridge;

import com.facebook.react.bridge.Promise;
import gnolang.gno.gnonative.PromiseBlock;
import java.util.ArrayList;
import java.util.List;

class JavaPromiseBlock implements PromiseBlock {
    public static List<JavaPromiseBlock> promises = new ArrayList<>();
    private final Promise promise;

    public JavaPromiseBlock(Promise promise) {
        this.promise = promise;
        this.store();
    }
    public void callResolve(String reply) {
        this.promise.resolve(reply);
        this.remove();
    }

    public void callReject(java.lang.Exception error) {
        this.promise.reject(error);
        this.remove();
    }

    private void remove() {
        promises.remove(this);
    }

    private void store() {
        promises.add(this);
    }
}
