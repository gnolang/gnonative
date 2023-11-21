import { ErrCode, ErrDetails } from '@gno/api/rpc_pb';
import { Code, ConnectError } from '@connectrpc/connect';

class GRPCError extends Error {
  public Details: ErrDetails | undefined;
  public GrpcCode: Code;

  public error: ConnectError;

  constructor(e: ConnectError | null | undefined) {
    if (!e) {
      // this should not happen, but should not break the app either.
      // instead simply create a empty error and warn about this
      console.warn(`GRPCError: (${e}) grpc error provided, empty error returned`);
      e = new ConnectError('');
    }

    super(e.rawMessage);

    this.error = e;
    this.GrpcCode = e.code;
    this.Details = e?.findDetails(ErrDetails).shift();
  }

  public details(): ErrDetails | unknown {
    return this.Details;
  }

  public errCode(): ErrCode {
    if (this.Details === undefined) {
      return ErrCode.Undefined;
    }

    const codes = this.Details.codes;
    if (codes.length == 0) {
      return ErrCode.Undefined;
    }

    return codes[0];
  }

  public grpcErrorCode(): Code {
    return this.GrpcCode;
  }

  public toJSON(): any {
    return {
      message: this.message,
      grpcErrorCode: Code[this.GrpcCode],
      errorCode: ErrCode[this.errCode()],
      details: this.Details,
    };
  }

  public hasErrCode(error: ErrCode): boolean {
    return this.Details?.codes.reduce((ac, v) => ac && v == error, false) || false;
  }
}

export { GRPCError };
