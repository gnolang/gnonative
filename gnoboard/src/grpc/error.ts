import { ErrCode } from '@gno/api/rpc_pb';
import { Code, ConnectError } from '@connectrpc/connect';

class GRPCError extends Error {
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
  }

  private extractErrCode(match: RegExpMatchArray | null) {
    if (match === null || match.length < 2) {
      return ErrCode.Undefined;
    }

    const code = parseInt(match[1]);
    if (typeof code !== 'number') {
      return ErrCode.Undefined;
    }

    return code;
  }

  // errCodes parses the error message formatted like `ErrType(#ErrNumber): WrappedErrType(#WrappedErrNumber)`
  // and returns the corresponding ErrCodes or [ErrCode.Undefined] if some errors occur.
  private errCodes(): ErrCode[] {
    const errCodes: ErrCode[] = [];

    if (this.message === '') {
      return [ErrCode.Undefined];
    }

    const matches = this.message.matchAll(/\w+\(#(\d+)\)/g);

    for (const match of matches) {
      const code = this.extractErrCode(match);
      errCodes.push(code);
    }

    return errCodes;
  }

  // errCode parses the error message formatted like `ErrType(#ErrNumber): WrappedErrType(#WrappedErrNumber)`
  // and returns the corresponding parent ErrCode or ErrCode.Undefined if some errors occur.
  public errCode(): ErrCode {
    if (this.message === '') {
      return ErrCode.Undefined;
    }

    const match = this.message.match(/\w+\(#(\d+)\)/);

    return this.extractErrCode(match);
  }

  public grpcErrorCode(): Code {
    return this.GrpcCode;
  }

  public toJSON(): any {
    return {
      message: this.message,
      grpcErrorCode: Code[this.GrpcCode],
      errorCode: ErrCode[this.errCode()],
    };
  }

  public hasErrCode(error: ErrCode): boolean {
    for (const err of this.errCodes()) {
      if (err === error) {
        return true;
      }
    }
    return false;
  }
}

export { GRPCError };
