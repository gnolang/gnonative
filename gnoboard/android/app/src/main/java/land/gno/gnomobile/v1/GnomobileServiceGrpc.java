package land.gno.gnomobile.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * GnomobileService is the service to interact with the Gno blockchain
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.57.2)",
    comments = "Source: rpc.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class GnomobileServiceGrpc {

  private GnomobileServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "land.gno.gnomobile.v1.GnomobileService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply> getSetRemoteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetRemote",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply> getSetRemoteMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply> getSetRemoteMethod;
    if ((getSetRemoteMethod = GnomobileServiceGrpc.getSetRemoteMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSetRemoteMethod = GnomobileServiceGrpc.getSetRemoteMethod) == null) {
          GnomobileServiceGrpc.getSetRemoteMethod = getSetRemoteMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetRemote"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSetRemoteMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply> getSetChainIDMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetChainID",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply> getSetChainIDMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply> getSetChainIDMethod;
    if ((getSetChainIDMethod = GnomobileServiceGrpc.getSetChainIDMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSetChainIDMethod = GnomobileServiceGrpc.getSetChainIDMethod) == null) {
          GnomobileServiceGrpc.getSetChainIDMethod = getSetChainIDMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetChainID"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSetChainIDMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply> getSetNameOrBech32Method;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetNameOrBech32",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply> getSetNameOrBech32Method() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply> getSetNameOrBech32Method;
    if ((getSetNameOrBech32Method = GnomobileServiceGrpc.getSetNameOrBech32Method) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSetNameOrBech32Method = GnomobileServiceGrpc.getSetNameOrBech32Method) == null) {
          GnomobileServiceGrpc.getSetNameOrBech32Method = getSetNameOrBech32Method =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetNameOrBech32"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSetNameOrBech32Method;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply> getSetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetPassword",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply> getSetPasswordMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply> getSetPasswordMethod;
    if ((getSetPasswordMethod = GnomobileServiceGrpc.getSetPasswordMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSetPasswordMethod = GnomobileServiceGrpc.getSetPasswordMethod) == null) {
          GnomobileServiceGrpc.getSetPasswordMethod = getSetPasswordMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request, land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSetPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply> getGenerateRecoveryPhraseMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GenerateRecoveryPhrase",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply> getGenerateRecoveryPhraseMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request, land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply> getGenerateRecoveryPhraseMethod;
    if ((getGenerateRecoveryPhraseMethod = GnomobileServiceGrpc.getGenerateRecoveryPhraseMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getGenerateRecoveryPhraseMethod = GnomobileServiceGrpc.getGenerateRecoveryPhraseMethod) == null) {
          GnomobileServiceGrpc.getGenerateRecoveryPhraseMethod = getGenerateRecoveryPhraseMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request, land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GenerateRecoveryPhrase"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getGenerateRecoveryPhraseMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request,
      land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply> getListKeyInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListKeyInfo",
      requestType = land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request.class,
      responseType = land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request,
      land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply> getListKeyInfoMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request, land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply> getListKeyInfoMethod;
    if ((getListKeyInfoMethod = GnomobileServiceGrpc.getListKeyInfoMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getListKeyInfoMethod = GnomobileServiceGrpc.getListKeyInfoMethod) == null) {
          GnomobileServiceGrpc.getListKeyInfoMethod = getListKeyInfoMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request, land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListKeyInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getListKeyInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.CreateAccount.Request,
      land.gno.gnomobile.v1.Rpc.CreateAccount.Reply> getCreateAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateAccount",
      requestType = land.gno.gnomobile.v1.Rpc.CreateAccount.Request.class,
      responseType = land.gno.gnomobile.v1.Rpc.CreateAccount.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.CreateAccount.Request,
      land.gno.gnomobile.v1.Rpc.CreateAccount.Reply> getCreateAccountMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.CreateAccount.Request, land.gno.gnomobile.v1.Rpc.CreateAccount.Reply> getCreateAccountMethod;
    if ((getCreateAccountMethod = GnomobileServiceGrpc.getCreateAccountMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getCreateAccountMethod = GnomobileServiceGrpc.getCreateAccountMethod) == null) {
          GnomobileServiceGrpc.getCreateAccountMethod = getCreateAccountMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Rpc.CreateAccount.Request, land.gno.gnomobile.v1.Rpc.CreateAccount.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.CreateAccount.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.CreateAccount.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getCreateAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.SelectAccount.Request,
      land.gno.gnomobile.v1.Rpc.SelectAccount.Reply> getSelectAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SelectAccount",
      requestType = land.gno.gnomobile.v1.Rpc.SelectAccount.Request.class,
      responseType = land.gno.gnomobile.v1.Rpc.SelectAccount.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.SelectAccount.Request,
      land.gno.gnomobile.v1.Rpc.SelectAccount.Reply> getSelectAccountMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.SelectAccount.Request, land.gno.gnomobile.v1.Rpc.SelectAccount.Reply> getSelectAccountMethod;
    if ((getSelectAccountMethod = GnomobileServiceGrpc.getSelectAccountMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getSelectAccountMethod = GnomobileServiceGrpc.getSelectAccountMethod) == null) {
          GnomobileServiceGrpc.getSelectAccountMethod = getSelectAccountMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Rpc.SelectAccount.Request, land.gno.gnomobile.v1.Rpc.SelectAccount.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SelectAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.SelectAccount.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.SelectAccount.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSelectAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request,
      land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply> getGetActiveAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetActiveAccount",
      requestType = land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request.class,
      responseType = land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request,
      land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply> getGetActiveAccountMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request, land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply> getGetActiveAccountMethod;
    if ((getGetActiveAccountMethod = GnomobileServiceGrpc.getGetActiveAccountMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getGetActiveAccountMethod = GnomobileServiceGrpc.getGetActiveAccountMethod) == null) {
          GnomobileServiceGrpc.getGetActiveAccountMethod = getGetActiveAccountMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request, land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActiveAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getGetActiveAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.Query_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply> getQueryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Query",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.Query_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.Query_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply> getQueryMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.Query_Request, land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply> getQueryMethod;
    if ((getQueryMethod = GnomobileServiceGrpc.getQueryMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getQueryMethod = GnomobileServiceGrpc.getQueryMethod) == null) {
          GnomobileServiceGrpc.getQueryMethod = getQueryMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.Query_Request, land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Query"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.Query_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply.getDefaultInstance()))
              .build();
        }
      }
    }
    return getQueryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.Call_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply> getCallMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Call",
      requestType = land.gno.gnomobile.v1.Gnomobiletypes.Call_Request.class,
      responseType = land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.Call_Request,
      land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply> getCallMethod() {
    io.grpc.MethodDescriptor<land.gno.gnomobile.v1.Gnomobiletypes.Call_Request, land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply> getCallMethod;
    if ((getCallMethod = GnomobileServiceGrpc.getCallMethod) == null) {
      synchronized (GnomobileServiceGrpc.class) {
        if ((getCallMethod = GnomobileServiceGrpc.getCallMethod) == null) {
          GnomobileServiceGrpc.getCallMethod = getCallMethod =
              io.grpc.MethodDescriptor.<land.gno.gnomobile.v1.Gnomobiletypes.Call_Request, land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Call"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.Call_Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply.getDefaultInstance()))
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
     * Set the connection addresse for the remote node. If you don't call this,
     * the default is "127.0.0.1:26657"
     * </pre>
     */
    default void setRemote(land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetRemoteMethod(), responseObserver);
    }

    /**
     * <pre>
     * Set the chain ID for the remote node. If you don't call this, the default
     * is "dev"
     * </pre>
     */
    default void setChainID(land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetChainIDMethod(), responseObserver);
    }

    /**
     * <pre>
     * Set the nameOrBech32 for the account in the keybase, used for later
     * operations
     * </pre>
     */
    default void setNameOrBech32(land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetNameOrBech32Method(), responseObserver);
    }

    /**
     * <pre>
     * Set the password for the account in the keybase, used for later operations
     * </pre>
     */
    default void setPassword(land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetPasswordMethod(), responseObserver);
    }

    /**
     * <pre>
     * Generate a recovery phrase of BIP39 mnemonic words using entropy from the crypto library
     * random number generator. This can be used as the mnemonic in CreateAccount.
     * </pre>
     */
    default void generateRecoveryPhrase(land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGenerateRecoveryPhraseMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get the keys informations in the keybase
     * </pre>
     */
    default void listKeyInfo(land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListKeyInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * Create a new account the keybase using the name an password specified by
     * SetAccount
     * </pre>
     */
    default void createAccount(land.gno.gnomobile.v1.Rpc.CreateAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.CreateAccount.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAccountMethod(), responseObserver);
    }

    /**
     * <pre>
     * SelectAccount selects the active account to use for later operations
     * </pre>
     */
    default void selectAccount(land.gno.gnomobile.v1.Rpc.SelectAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.SelectAccount.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSelectAccountMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetActiveAccount gets the active account which was set by SelectAccount.
     * If there is no active account, then return ErrNoActiveAccount.
     * (To check if there is an active account, use ListKeyInfo and check the length of the result.)
     * </pre>
     */
    default void getActiveAccount(land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActiveAccountMethod(), responseObserver);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    default void query(land.gno.gnomobile.v1.Gnomobiletypes.Query_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryMethod(), responseObserver);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    default void call(land.gno.gnomobile.v1.Gnomobiletypes.Call_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply> responseObserver) {
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
     * Set the connection addresse for the remote node. If you don't call this,
     * the default is "127.0.0.1:26657"
     * </pre>
     */
    public void setRemote(land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetRemoteMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Set the chain ID for the remote node. If you don't call this, the default
     * is "dev"
     * </pre>
     */
    public void setChainID(land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetChainIDMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Set the nameOrBech32 for the account in the keybase, used for later
     * operations
     * </pre>
     */
    public void setNameOrBech32(land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetNameOrBech32Method(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Set the password for the account in the keybase, used for later operations
     * </pre>
     */
    public void setPassword(land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Generate a recovery phrase of BIP39 mnemonic words using entropy from the crypto library
     * random number generator. This can be used as the mnemonic in CreateAccount.
     * </pre>
     */
    public void generateRecoveryPhrase(land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGenerateRecoveryPhraseMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get the keys informations in the keybase
     * </pre>
     */
    public void listKeyInfo(land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListKeyInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Create a new account the keybase using the name an password specified by
     * SetAccount
     * </pre>
     */
    public void createAccount(land.gno.gnomobile.v1.Rpc.CreateAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.CreateAccount.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * SelectAccount selects the active account to use for later operations
     * </pre>
     */
    public void selectAccount(land.gno.gnomobile.v1.Rpc.SelectAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.SelectAccount.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSelectAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetActiveAccount gets the active account which was set by SelectAccount.
     * If there is no active account, then return ErrNoActiveAccount.
     * (To check if there is an active account, use ListKeyInfo and check the length of the result.)
     * </pre>
     */
    public void getActiveAccount(land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActiveAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    public void query(land.gno.gnomobile.v1.Gnomobiletypes.Query_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    public void call(land.gno.gnomobile.v1.Gnomobiletypes.Call_Request request,
        io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply> responseObserver) {
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
     * Set the connection addresse for the remote node. If you don't call this,
     * the default is "127.0.0.1:26657"
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply setRemote(land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetRemoteMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Set the chain ID for the remote node. If you don't call this, the default
     * is "dev"
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply setChainID(land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetChainIDMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Set the nameOrBech32 for the account in the keybase, used for later
     * operations
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply setNameOrBech32(land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetNameOrBech32Method(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Set the password for the account in the keybase, used for later operations
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply setPassword(land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Generate a recovery phrase of BIP39 mnemonic words using entropy from the crypto library
     * random number generator. This can be used as the mnemonic in CreateAccount.
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply generateRecoveryPhrase(land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGenerateRecoveryPhraseMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get the keys informations in the keybase
     * </pre>
     */
    public land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply listKeyInfo(land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListKeyInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Create a new account the keybase using the name an password specified by
     * SetAccount
     * </pre>
     */
    public land.gno.gnomobile.v1.Rpc.CreateAccount.Reply createAccount(land.gno.gnomobile.v1.Rpc.CreateAccount.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * SelectAccount selects the active account to use for later operations
     * </pre>
     */
    public land.gno.gnomobile.v1.Rpc.SelectAccount.Reply selectAccount(land.gno.gnomobile.v1.Rpc.SelectAccount.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSelectAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetActiveAccount gets the active account which was set by SelectAccount.
     * If there is no active account, then return ErrNoActiveAccount.
     * (To check if there is an active account, use ListKeyInfo and check the length of the result.)
     * </pre>
     */
    public land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply getActiveAccount(land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply query(land.gno.gnomobile.v1.Gnomobiletypes.Query_Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    public land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply call(land.gno.gnomobile.v1.Gnomobiletypes.Call_Request request) {
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
     * Set the connection addresse for the remote node. If you don't call this,
     * the default is "127.0.0.1:26657"
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply> setRemote(
        land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetRemoteMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Set the chain ID for the remote node. If you don't call this, the default
     * is "dev"
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply> setChainID(
        land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetChainIDMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Set the nameOrBech32 for the account in the keybase, used for later
     * operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply> setNameOrBech32(
        land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetNameOrBech32Method(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Set the password for the account in the keybase, used for later operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply> setPassword(
        land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetPasswordMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Generate a recovery phrase of BIP39 mnemonic words using entropy from the crypto library
     * random number generator. This can be used as the mnemonic in CreateAccount.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply> generateRecoveryPhrase(
        land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGenerateRecoveryPhraseMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get the keys informations in the keybase
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply> listKeyInfo(
        land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListKeyInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Create a new account the keybase using the name an password specified by
     * SetAccount
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Rpc.CreateAccount.Reply> createAccount(
        land.gno.gnomobile.v1.Rpc.CreateAccount.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAccountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * SelectAccount selects the active account to use for later operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Rpc.SelectAccount.Reply> selectAccount(
        land.gno.gnomobile.v1.Rpc.SelectAccount.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSelectAccountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetActiveAccount gets the active account which was set by SelectAccount.
     * If there is no active account, then return ErrNoActiveAccount.
     * (To check if there is an active account, use ListKeyInfo and check the length of the result.)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply> getActiveAccount(
        land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActiveAccountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Make an ABCI query to the remote node.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply> query(
        land.gno.gnomobile.v1.Gnomobiletypes.Query_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Call a specific realm function.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply> call(
        land.gno.gnomobile.v1.Gnomobiletypes.Call_Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SET_REMOTE = 0;
  private static final int METHODID_SET_CHAIN_ID = 1;
  private static final int METHODID_SET_NAME_OR_BECH32 = 2;
  private static final int METHODID_SET_PASSWORD = 3;
  private static final int METHODID_GENERATE_RECOVERY_PHRASE = 4;
  private static final int METHODID_LIST_KEY_INFO = 5;
  private static final int METHODID_CREATE_ACCOUNT = 6;
  private static final int METHODID_SELECT_ACCOUNT = 7;
  private static final int METHODID_GET_ACTIVE_ACCOUNT = 8;
  private static final int METHODID_QUERY = 9;
  private static final int METHODID_CALL = 10;

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
          serviceImpl.setRemote((land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply>) responseObserver);
          break;
        case METHODID_SET_CHAIN_ID:
          serviceImpl.setChainID((land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply>) responseObserver);
          break;
        case METHODID_SET_NAME_OR_BECH32:
          serviceImpl.setNameOrBech32((land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply>) responseObserver);
          break;
        case METHODID_SET_PASSWORD:
          serviceImpl.setPassword((land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply>) responseObserver);
          break;
        case METHODID_GENERATE_RECOVERY_PHRASE:
          serviceImpl.generateRecoveryPhrase((land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply>) responseObserver);
          break;
        case METHODID_LIST_KEY_INFO:
          serviceImpl.listKeyInfo((land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply>) responseObserver);
          break;
        case METHODID_CREATE_ACCOUNT:
          serviceImpl.createAccount((land.gno.gnomobile.v1.Rpc.CreateAccount.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.CreateAccount.Reply>) responseObserver);
          break;
        case METHODID_SELECT_ACCOUNT:
          serviceImpl.selectAccount((land.gno.gnomobile.v1.Rpc.SelectAccount.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.SelectAccount.Reply>) responseObserver);
          break;
        case METHODID_GET_ACTIVE_ACCOUNT:
          serviceImpl.getActiveAccount((land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply>) responseObserver);
          break;
        case METHODID_QUERY:
          serviceImpl.query((land.gno.gnomobile.v1.Gnomobiletypes.Query_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply>) responseObserver);
          break;
        case METHODID_CALL:
          serviceImpl.call((land.gno.gnomobile.v1.Gnomobiletypes.Call_Request) request,
              (io.grpc.stub.StreamObserver<land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply>) responseObserver);
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
              land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.SetRemote_Reply>(
                service, METHODID_SET_REMOTE)))
        .addMethod(
          getSetChainIDMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.SetChainID_Reply>(
                service, METHODID_SET_CHAIN_ID)))
        .addMethod(
          getSetNameOrBech32Method(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.SetNameOrBech32_Reply>(
                service, METHODID_SET_NAME_OR_BECH32)))
        .addMethod(
          getSetPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.SetPassword_Reply>(
                service, METHODID_SET_PASSWORD)))
        .addMethod(
          getGenerateRecoveryPhraseMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.GenerateRecoveryPhrase_Reply>(
                service, METHODID_GENERATE_RECOVERY_PHRASE)))
        .addMethod(
          getListKeyInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Rpc.ListKeyInfo.Request,
              land.gno.gnomobile.v1.Rpc.ListKeyInfo.Reply>(
                service, METHODID_LIST_KEY_INFO)))
        .addMethod(
          getCreateAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Rpc.CreateAccount.Request,
              land.gno.gnomobile.v1.Rpc.CreateAccount.Reply>(
                service, METHODID_CREATE_ACCOUNT)))
        .addMethod(
          getSelectAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Rpc.SelectAccount.Request,
              land.gno.gnomobile.v1.Rpc.SelectAccount.Reply>(
                service, METHODID_SELECT_ACCOUNT)))
        .addMethod(
          getGetActiveAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Rpc.GetActiveAccount.Request,
              land.gno.gnomobile.v1.Rpc.GetActiveAccount.Reply>(
                service, METHODID_GET_ACTIVE_ACCOUNT)))
        .addMethod(
          getQueryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Gnomobiletypes.Query_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.Query_Reply>(
                service, METHODID_QUERY)))
        .addMethod(
          getCallMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              land.gno.gnomobile.v1.Gnomobiletypes.Call_Request,
              land.gno.gnomobile.v1.Gnomobiletypes.Call_Reply>(
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
              .addMethod(getSetChainIDMethod())
              .addMethod(getSetNameOrBech32Method())
              .addMethod(getSetPasswordMethod())
              .addMethod(getGenerateRecoveryPhraseMethod())
              .addMethod(getListKeyInfoMethod())
              .addMethod(getCreateAccountMethod())
              .addMethod(getSelectAccountMethod())
              .addMethod(getGetActiveAccountMethod())
              .addMethod(getQueryMethod())
              .addMethod(getCallMethod())
              .build();
        }
      }
    }
    return result;
  }
}
