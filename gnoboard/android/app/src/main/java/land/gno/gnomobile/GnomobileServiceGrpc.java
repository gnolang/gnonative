package land.gno.gnomobile;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * GnomobileService is the service to interact with the Gno blockchain
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.57.2)",
    comments = "Source: gnomobiletypes.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class GnomobileServiceGrpc {

  private GnomobileServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "gomobile.v1.GnomobileService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.SetRemote.Request,
      land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply> getSetRemoteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetRemote",
      requestType = land.gno.gnomobile.Gnomobiletypes.SetRemote.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.SetRemote.Request,
      land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply> getSetRemoteMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.SetRemote.Request, land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply> getSetRemoteMethod;
    if ((getSetRemoteMethod = GnomobileServiceGrpc.getSetRemoteMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSetRemoteMethod = GnomobileServiceGrpc.getSetRemoteMethod) == null) {
          GnomobileServiceGrpc.getSetRemoteMethod = getSetRemoteMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.SetRemote.Request, land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetRemote"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.SetRemote.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSetRemoteMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request,
      land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply> getInitKeyBaseFromDirMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "InitKeyBaseFromDir",
      requestType = land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request,
      land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply> getInitKeyBaseFromDirMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request, land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply> getInitKeyBaseFromDirMethod;
    if ((getInitKeyBaseFromDirMethod = GnomobileServiceGrpc.getInitKeyBaseFromDirMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getInitKeyBaseFromDirMethod = GnomobileServiceGrpc.getInitKeyBaseFromDirMethod) == null) {
          GnomobileServiceGrpc.getInitKeyBaseFromDirMethod = getInitKeyBaseFromDirMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request, land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "InitKeyBaseFromDir"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getInitKeyBaseFromDirMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.SetAccount.Request,
      land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply> getSetAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetAccount",
      requestType = land.gno.gnomobile.Gnomobiletypes.SetAccount.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.SetAccount.Request,
      land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply> getSetAccountMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.SetAccount.Request, land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply> getSetAccountMethod;
    if ((getSetAccountMethod = GnomobileServiceGrpc.getSetAccountMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSetAccountMethod = GnomobileServiceGrpc.getSetAccountMethod) == null) {
          GnomobileServiceGrpc.getSetAccountMethod = getSetAccountMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.SetAccount.Request, land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.SetAccount.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSetAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request,
      land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply> getGetKeyCountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetKeyCount",
      requestType = land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request,
      land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply> getGetKeyCountMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request, land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply> getGetKeyCountMethod;
    if ((getGetKeyCountMethod = GnomobileServiceGrpc.getGetKeyCountMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getGetKeyCountMethod = GnomobileServiceGrpc.getGetKeyCountMethod) == null) {
          GnomobileServiceGrpc.getGetKeyCountMethod = getGetKeyCountMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request, land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetKeyCount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getGetKeyCountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request,
      land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply> getCreateAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateAccount",
      requestType = land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request,
      land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply> getCreateAccountMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request, land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply> getCreateAccountMethod;
    if ((getCreateAccountMethod = GnomobileServiceGrpc.getCreateAccountMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getCreateAccountMethod = GnomobileServiceGrpc.getCreateAccountMethod) == null) {
          GnomobileServiceGrpc.getCreateAccountMethod = getCreateAccountMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request, land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getCreateAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.Query.Request,
      land.gno.gnomobile.Gnomobiletypes.Query.Reply> getQueryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Query",
      requestType = land.gno.gnomobile.Gnomobiletypes.Query.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.Query.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.Query.Request,
      land.gno.gnomobile.Gnomobiletypes.Query.Reply> getQueryMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.Query.Request, land.gno.gnomobile.Gnomobiletypes.Query.Reply> getQueryMethod;
    if ((getQueryMethod = GnomobileServiceGrpc.getQueryMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getQueryMethod = GnomobileServiceGrpc.getQueryMethod) == null) {
          GnomobileServiceGrpc.getQueryMethod = getQueryMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.Query.Request, land.gno.gnomobile.Gnomobiletypes.Query.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Query"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.Query.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.Query.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getQueryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.Call.Request,
      land.gno.gnomobile.Gnomobiletypes.Call.Reply> getCallMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Call",
      requestType = land.gno.gnomobile.Gnomobiletypes.Call.Request.class,
      responseType = land.gno.gnomobile.Gnomobiletypes.Call.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.Call.Request,
      land.gno.gnomobile.Gnomobiletypes.Call.Reply> getCallMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.Gnomobiletypes.Call.Request, land.gno.gnomobile.Gnomobiletypes.Call.Reply> getCallMethod;
    if ((getCallMethod = GnomobileServiceGrpc.getCallMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getCallMethod = GnomobileServiceGrpc.getCallMethod) == null) {
          GnomobileServiceGrpc.getCallMethod = getCallMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.Gnomobiletypes.Call.Request, land.gno.gnomobile.Gnomobiletypes.Call.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Call"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.Call.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.Gnomobiletypes.Call.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getCallMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static GnomobileServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GnomobileServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GnomobileServiceStub>() {
        @java.lang.Override
        public GnomobileServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GnomobileServiceStub(channel, callOptions);
        }
      };
    return GnomobileServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static GnomobileServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GnomobileServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GnomobileServiceBlockingStub>() {
        @java.lang.Override
        public GnomobileServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GnomobileServiceBlockingStub(channel, callOptions);
        }
      };
    return GnomobileServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static GnomobileServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GnomobileServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GnomobileServiceFutureStub>() {
        @java.lang.Override
        public GnomobileServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GnomobileServiceFutureStub(channel, callOptions);
        }
      };
    return GnomobileServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * GnomobileService is the service to interact with the Gno blockchain
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Set the connection info for the remote node. If you don't call this, the default is
     * remote = "127.0.0.1:26657" and chainID = "dev"
     * </pre>
     */
    default void setRemote(land.gno.gnomobile.Gnomobiletypes.SetRemote.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetRemoteMethod(), responseObserver);
    }

    /**
     * <pre>
     * InitKeyBaseFromDir initializes a keybase in the given subdirectory of the app's root directory.
     * If the keybase already exists then this opens it, otherwise this creates a new empty keybase.
     * </pre>
     */
    default void initKeyBaseFromDir(land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInitKeyBaseFromDirMethod(), responseObserver);
    }

    /**
     * <pre>
     * Set the name and password for the account in the keybase, used for later operations
     * </pre>
     */
    default void setAccount(land.gno.gnomobile.Gnomobiletypes.SetAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetAccountMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get the count of keys in the keybase that was specified by InitKeyBaseFromDir
     * </pre>
     */
    default void getKeyCount(land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetKeyCountMethod(), responseObserver);
    }

    /**
     * <pre>
     * Create a new account the keybase that was specified by InitKeyBaseFromDir, using
     * the name an password specified by SetAccount
     * </pre>
     */
    default void createAccount(land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAccountMethod(), responseObserver);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    default void query(land.gno.gnomobile.Gnomobiletypes.Query.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.Query.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryMethod(), responseObserver);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    default void call(land.gno.gnomobile.Gnomobiletypes.Call.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.Call.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCallMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service GnomobileService.
   * <pre>
   * GnomobileService is the service to interact with the Gno blockchain
   * </pre>
   */
  public static abstract class GnomobileServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return GnomobileServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service GnomobileService.
   * <pre>
   * GnomobileService is the service to interact with the Gno blockchain
   * </pre>
   */
  public static final class GnomobileServiceStub
      extends io.grpc.stub.AbstractAsyncStub<GnomobileServiceStub> {
    private GnomobileServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GnomobileServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GnomobileServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Set the connection info for the remote node. If you don't call this, the default is
     * remote = "127.0.0.1:26657" and chainID = "dev"
     * </pre>
     */
    public void setRemote(land.gno.gnomobile.Gnomobiletypes.SetRemote.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetRemoteMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * InitKeyBaseFromDir initializes a keybase in the given subdirectory of the app's root directory.
     * If the keybase already exists then this opens it, otherwise this creates a new empty keybase.
     * </pre>
     */
    public void initKeyBaseFromDir(land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInitKeyBaseFromDirMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Set the name and password for the account in the keybase, used for later operations
     * </pre>
     */
    public void setAccount(land.gno.gnomobile.Gnomobiletypes.SetAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get the count of keys in the keybase that was specified by InitKeyBaseFromDir
     * </pre>
     */
    public void getKeyCount(land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetKeyCountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Create a new account the keybase that was specified by InitKeyBaseFromDir, using
     * the name an password specified by SetAccount
     * </pre>
     */
    public void createAccount(land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    public void query(land.gno.gnomobile.Gnomobiletypes.Query.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.Query.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    public void call(land.gno.gnomobile.Gnomobiletypes.Call.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.Call.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service GnomobileService.
   * <pre>
   * GnomobileService is the service to interact with the Gno blockchain
   * </pre>
   */
  public static final class GnomobileServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<GnomobileServiceBlockingStub> {
    private GnomobileServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GnomobileServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GnomobileServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Set the connection info for the remote node. If you don't call this, the default is
     * remote = "127.0.0.1:26657" and chainID = "dev"
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply setRemote(land.gno.gnomobile.Gnomobiletypes.SetRemote.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetRemoteMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * InitKeyBaseFromDir initializes a keybase in the given subdirectory of the app's root directory.
     * If the keybase already exists then this opens it, otherwise this creates a new empty keybase.
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply initKeyBaseFromDir(land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInitKeyBaseFromDirMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Set the name and password for the account in the keybase, used for later operations
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply setAccount(land.gno.gnomobile.Gnomobiletypes.SetAccount.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get the count of keys in the keybase that was specified by InitKeyBaseFromDir
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply getKeyCount(land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetKeyCountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Create a new account the keybase that was specified by InitKeyBaseFromDir, using
     * the name an password specified by SetAccount
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply createAccount(land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.Query.Reply query(land.gno.gnomobile.Gnomobiletypes.Query.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    public land.gno.gnomobile.Gnomobiletypes.Call.Reply call(land.gno.gnomobile.Gnomobiletypes.Call.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCallMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service GnomobileService.
   * <pre>
   * GnomobileService is the service to interact with the Gno blockchain
   * </pre>
   */
  public static final class GnomobileServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<GnomobileServiceFutureStub> {
    private GnomobileServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GnomobileServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GnomobileServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Set the connection info for the remote node. If you don't call this, the default is
     * remote = "127.0.0.1:26657" and chainID = "dev"
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply> setRemote(
        land.gno.gnomobile.Gnomobiletypes.SetRemote.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetRemoteMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * InitKeyBaseFromDir initializes a keybase in the given subdirectory of the app's root directory.
     * If the keybase already exists then this opens it, otherwise this creates a new empty keybase.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply> initKeyBaseFromDir(
        land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInitKeyBaseFromDirMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Set the name and password for the account in the keybase, used for later operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply> setAccount(
        land.gno.gnomobile.Gnomobiletypes.SetAccount.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetAccountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get the count of keys in the keybase that was specified by InitKeyBaseFromDir
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply> getKeyCount(
        land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetKeyCountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Create a new account the keybase that was specified by InitKeyBaseFromDir, using
     * the name an password specified by SetAccount
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply> createAccount(
        land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAccountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.Query.Reply> query(
        land.gno.gnomobile.Gnomobiletypes.Query.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.Gnomobiletypes.Call.Reply> call(
        land.gno.gnomobile.Gnomobiletypes.Call.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SET_REMOTE = 0;
  private static final int METHODID_INIT_KEY_BASE_FROM_DIR = 1;
  private static final int METHODID_SET_ACCOUNT = 2;
  private static final int METHODID_GET_KEY_COUNT = 3;
  private static final int METHODID_CREATE_ACCOUNT = 4;
  private static final int METHODID_QUERY = 5;
  private static final int METHODID_CALL = 6;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SET_REMOTE:
          serviceImpl.setRemote((land.gno.gnomobile.Gnomobiletypes.SetRemote.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply>) responseObserver);
          break;
        case METHODID_INIT_KEY_BASE_FROM_DIR:
          serviceImpl.initKeyBaseFromDir((land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply>) responseObserver);
          break;
        case METHODID_SET_ACCOUNT:
          serviceImpl.setAccount((land.gno.gnomobile.Gnomobiletypes.SetAccount.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply>) responseObserver);
          break;
        case METHODID_GET_KEY_COUNT:
          serviceImpl.getKeyCount((land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply>) responseObserver);
          break;
        case METHODID_CREATE_ACCOUNT:
          serviceImpl.createAccount((land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply>) responseObserver);
          break;
        case METHODID_QUERY:
          serviceImpl.query((land.gno.gnomobile.Gnomobiletypes.Query.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.Query.Reply>) responseObserver);
          break;
        case METHODID_CALL:
          serviceImpl.call((land.gno.gnomobile.Gnomobiletypes.Call.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.Gnomobiletypes.Call.Reply>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getSetRemoteMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.SetRemote.Request,
              land.gno.gnomobile.Gnomobiletypes.SetRemote.Reply>(
                service, METHODID_SET_REMOTE)))
        .addMethod(
          getInitKeyBaseFromDirMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Request,
              land.gno.gnomobile.Gnomobiletypes.InitKeyBaseFromDir.Reply>(
                service, METHODID_INIT_KEY_BASE_FROM_DIR)))
        .addMethod(
          getSetAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.SetAccount.Request,
              land.gno.gnomobile.Gnomobiletypes.SetAccount.Reply>(
                service, METHODID_SET_ACCOUNT)))
        .addMethod(
          getGetKeyCountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Request,
              land.gno.gnomobile.Gnomobiletypes.GetKeyCount.Reply>(
                service, METHODID_GET_KEY_COUNT)))
        .addMethod(
          getCreateAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.CreateAccount.Request,
              land.gno.gnomobile.Gnomobiletypes.CreateAccount.Reply>(
                service, METHODID_CREATE_ACCOUNT)))
        .addMethod(
          getQueryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.Query.Request,
              land.gno.gnomobile.Gnomobiletypes.Query.Reply>(
                service, METHODID_QUERY)))
        .addMethod(
          getCallMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.Gnomobiletypes.Call.Request,
              land.gno.gnomobile.Gnomobiletypes.Call.Reply>(
                service, METHODID_CALL)))
        .build();
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (GnomobileServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getSetRemoteMethod())
              .addMethod(getInitKeyBaseFromDirMethod())
              .addMethod(getSetAccountMethod())
              .addMethod(getGetKeyCountMethod())
              .addMethod(getCreateAccountMethod())
              .addMethod(getQueryMethod())
              .addMethod(getCallMethod())
              .build();
        }
      }
    }
    return result;
  }
}
