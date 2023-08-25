//
//  RootDir.swift
//  Berty
//
//  Created by Antoine Eddi on 09/08/2021.
//

import Foundation

enum RootDirError: Error {
    case path
}
extension RootDirError: LocalizedError {
    public var errorDescription: String? {
        switch self {
        case .path:
            return NSLocalizedString(
                "unable to get app group path url",
                comment: ""
            )
        }
    }
}

func RootDirGet() throws -> String {
  
  let docDir = try! FileManager.default.url(for: .documentDirectory, in: .userDomainMask, appropriateFor: nil, create: true)
  return docDir.appendingPathComponent("gnomobile", isDirectory: true).path
}

@objc(RootDir)
class RootDir: NSObject {
  @objc func get(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    do {
      resolve(try RootDirGet())
    } catch {
      reject("root_dir_failure", error.localizedDescription, error)
    }
  }
  
  @objc static func requiresMainQueueSetup() -> Bool {
      return false
  }
}
