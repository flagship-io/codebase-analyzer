/// Create visitor
FSVisitor * visitor1 = [[[[Flagship sharedInstance] newVisitor:@"visitor_1" instanceType:InstanceSHARED_INSTANCE] withContextWithContext:@{@"age":@18} ] build];

/// Fetch flags
[visitor1 fetchFlagsOnFetchCompleted:^{
  
  // Ex: get flag for vip feature
  FSFlag * flag = [visitor1 getFlagWithKey:@"btnColor" defaultValue:@"red"];
  FSFlag * flag = [visitor1 getFlagWithKey:@"showBtn" defaultValue:TRUE];
  FSFlag * flag = [visitor1 getFlagWithKey:@"btnSize" defaultValue:13];
  FSFlag * flag = [visitor1 getFlagWithKey:@"btnSizeFloat" defaultValue:14.5];
  // Use this flag to enable displaying the vip feature
}];
