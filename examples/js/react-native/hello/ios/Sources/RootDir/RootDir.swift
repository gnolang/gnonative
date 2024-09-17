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

func RootDirGet() throws -> String {
  return try! FileManager.default.url(for: .documentDirectory, in: .userDomainMask, appropriateFor: nil, create: true).path
}

func TempDirGet() throws -> String {
  return FileManager.default.compatTemporaryDirectory.path
}

extension FileManager {
  public var compatTemporaryDirectory: URL {
    if #available(iOS 10.0, *) {
      return temporaryDirectory
    } else {
      return (try? url(
        for: .itemReplacementDirectory,
        in: .userDomainMask,
        appropriateFor: nil,
        create: true)
      ) ?? URL(fileURLWithPath: NSTemporaryDirectory())
    }
  }
}
