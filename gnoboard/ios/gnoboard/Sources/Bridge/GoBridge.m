#import <React/RCTBridgeModule.h>

@interface RCT_EXTERN_MODULE(GoBridge, NSObject)

RCT_EXTERN_METHOD(initBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(closeBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(hello:(NSString)name
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(createDefaulAccount:(NSString)name
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(exportJsonConfig:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(createReply:(NSString)message
                  rootdir:(NSString)rootdir
                  resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

+ (BOOL)requiresMainQueueSetup
{
 return YES;  // only do this if your module initialization relies on calling UIKit!
}

@end
