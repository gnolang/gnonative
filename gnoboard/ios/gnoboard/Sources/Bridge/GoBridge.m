#import <React/RCTBridgeModule.h>


@interface RCT_EXTERN_MODULE(GoBridge, NSObject)

RCT_EXTERN_METHOD(initBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(closeBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(listKeyInfo:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(createAccount:(NSString *)nameOrBech32
                  mnemonic:(NSString *)mnemonic
                  bip39Passwd:(NSString *)bip39Passwd
                  password:(NSString *)password
                  account:(nonnull NSNumber *)account
                  index:(nonnull NSNumber *)index
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(selectAccount:(NSString *)nameOrBech32
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(call:(NSString *)packagePath
                  fnc:(NSString *)fnc
                  args:(NSArray *)args
                  gasFee:(NSString *)gasFee
                  gasWanted:(nonnull NSNumber *)gasWanted
                  password:(NSString *)password
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(exportJsonConfig:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

+ (BOOL)requiresMainQueueSetup
{
  return NO;  // only do this if your module initialization relies on calling UIKit!
}

- (dispatch_queue_t)methodQueue
{
  return dispatch_queue_create("com.facebook.React.AsyncLocalStorageQueue", DISPATCH_QUEUE_SERIAL);
}

@end
