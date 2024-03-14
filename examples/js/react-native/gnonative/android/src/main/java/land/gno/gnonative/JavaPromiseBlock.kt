package land.gno.gnonative

//  Created by Guilhem Fanton on 10/07/2023.

import expo.modules.kotlin.Promise
import gnolang.gno.gnonative.PromiseBlock as IPromiseBlock

class PromiseBlock(val promise: Promise): IPromiseBlock {
  // gnolang.gno.gnonative.PromiseBlock aims to keep reference over promise object so go can play with
  // until the promise is resolved
  companion object {
    private var promises = mutableSetOf<PromiseBlock>()
  }

  init {
    store()
  }

  override fun callResolve(reply: String?) {
    this.promise.resolve(reply ?: "")
    this.remove() // cleanup the promise
  }

  override fun callReject(err: Exception?) {
    this.promise.reject(GoBridgeCoreError(err))
    this.remove() // cleanup the promise
  }

  private fun store() {
    promises.add(this)
  }

  private fun remove() {
    promises.remove(this)
  }
}
