/// Create visitor
FSVisitor * visitor1 = [[[[Flagship sharedInstance] newVisitor:@"visitor_1" instanceType:InstanceSHARED_INSTANCE] withContextWithContext:@{@"age":@18} ] build];

/// Fetch flags
[visitor1 fetchFlagsOnFetchCompleted:^{
  // Ex: get flag for vip feature
  FSFlag * flag = [visitor1 getFlagWithKey:@"btnColor" defaultValue:FALSE];
  // Use this flag to enable displaying the vip feature
}];
