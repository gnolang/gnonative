package land.gno.gnonative
import expo.modules.kotlin.exception.CodedException

class GoBridgeNotStartedError : CodedException("NotStarted", "Service hasn't started yet", null)
class GoBridgeCoreError(err: Exception?) : CodedException("CoreError", err)
