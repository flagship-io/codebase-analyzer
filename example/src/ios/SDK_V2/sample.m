/// init config object and set the running mode and timeout
FSConfig * config = [[FSConfig alloc] init:FlagshipModeDECISION_API timeout:0.2];

/// Start the sdk
[[Flagship sharedInstance] startWithEnvId:@"you envId" apiKey:@"your apiKey" visitorId:@"visitorId" config:config onStartDone:^(enum FlagshipResult result) {

  if (result == FlagshipResultReady){

    dispatch_async(dispatch_get_main_queue(), ^{

              /// update UI
    });
  }else{

            /// An error occurs or the SDK is disabled
  }

}];

// Retrieve modification and activate
 NSString * title = [[Flagship sharedInstance] getModification:@"btnColor" defaultString:@"More Infos" activate:YES];
 NSString * title = [[Flagship sharedInstance] getModification:@"showBtn" defaultBool:TRUE activate:YES];
 NSString * title = [[Flagship sharedInstance] getModification:@"btnSize" defaultInt:13 activate:YES];
 NSString * title = [[Flagship sharedInstance] getModification:@"btnSizeFloat" defaultFloat:14.5 activate:YES];




// END
