package land.gno.gnonative
import expo.modules.kotlin.exception.CodedException

class GoBridgeNotStartedError : CodedException("NotStarted", "Service hasn't started yet", null)
class GoBridgeCoreEOF : CodedException("EOF", "EOF", null)
