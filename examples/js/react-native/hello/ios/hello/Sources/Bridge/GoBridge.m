#import <React/RCTBridgeModule.h>


@interface RCT_EXTERN_MODULE(GoBridge, NSObject)

RCT_EXTERN_METHOD(initBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(closeBridge:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject);

RCT_EXTERN_METHOD(getTcpPort:(RCTPromiseResolveBlock)resolve
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
