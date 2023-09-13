#import <React/RCTBridgeModule.h>


@interface RCT_EXTERN_MODULE(GoBridge, NSObject)

RCT_EXTERN_METHOD(initBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(closeBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(call:(NSString *)packagePath
                  fnc:(NSString *)fnc
                  args:(NSArray *)args
                  gasFee:(NSString *)gasFee
                  gasWanted:(nonnull NSNumber *)gasWanted
                  password:(NSString *)password
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(createDefaultAccount:(NSString)name
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(exportJsonConfig:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

+ (BOOL)requiresMainQueueSetup
{
 return YES;  // only do this if your module initialization relies on calling UIKit!
}

@end
