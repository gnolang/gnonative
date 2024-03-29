//
//  PromiseBlock.swift
//  GnoExpo
//
//  Created by Guilhem Fanton on 10/07/2023.
//

import ExpoModulesCore
import GnoCore

var promises = Set<PromiseBlock>()

// PromiseBlock aim to keep reference over promise object so go can play with
// until the promise is resolved
class PromiseBlock: NSObject, GnoGnonativePromiseBlockProtocol {
    static func == (lhs: PromiseBlock, rhs: PromiseBlock) -> Bool {
        return lhs === rhs
    }
    
    var promise: Promise
    
    init(promise: Promise) {
        self.promise = promise
        super.init()
        self.store()
    }
    
    func callResolve(_ reply: String?) {
        self.promise.resolve(reply ?? "")
        self.remove() // cleanup the promise
    }
    
    func callReject(_ error: Error?) {
        self.promise.reject(error ?? Exception(name: "Unknown Error", description: "unknown reject error"))
        self.remove() // cleanup the promise
    }
    
    func store() {
        promises.insert(self)
    }
    
    func remove() {
        promises.remove(self)
    }
}
