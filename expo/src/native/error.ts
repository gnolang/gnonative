// Error wrapper for the GnoNative client. The native module rejects with a plain Error whose message
// is the Go error string (e.g. "ErrInvalidAddress(#205)"); this parses the Name(#code) format into
// an ErrCode.
import { ErrCode } from './apitypes';

export class GnoNativeError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'GnoNativeError';
  }

  private extractErrCode(match: RegExpMatchArray | null): ErrCode {
    if (match === null || match.length < 2) {
      return ErrCode.Undefined;
    }

    const code = parseInt(match[1], 10);
    if (Number.isNaN(code)) {
      return ErrCode.Undefined;
    }

    return code;
  }

  // errCodes parses the error message formatted like `ErrType(#ErrNumber): WrappedErrType(#WrappedErrNumber)`
  // and returns the corresponding ErrCodes, or [ErrCode.Undefined] if none can be parsed.
  private errCodes(): ErrCode[] {
    const errCodes: ErrCode[] = [];

    if (this.message === '') {
      return [ErrCode.Undefined];
    }

    const matches = this.message.matchAll(/\w+\(#(\d+)\)/g);

    for (const match of matches) {
      errCodes.push(this.extractErrCode(match));
    }

    if (errCodes.length === 0) {
      return [ErrCode.Undefined];
    }

    return errCodes;
  }

  // errCode parses the error message and returns the parent ErrCode, or ErrCode.Undefined.
  public errCode(): ErrCode {
    if (this.message === '') {
      return ErrCode.Undefined;
    }

    const match = this.message.match(/\w+\(#(\d+)\)/);
    return this.extractErrCode(match);
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
