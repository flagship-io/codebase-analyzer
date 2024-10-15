using Flagship.Main;

//Step 2: Create a visitor
var visitor = Fs.NewVisitor("<VISITOR_ID>", true)
  .SetContext(new Dictionary<string, object>(){["isVip"] = true})
  .Build();


//Step 3: Fetch flag from the Flagship platform 
await visitor.FetchFlags();

/* Step 4: Retrieves a flag named "displayVipFeature", 
 */
var flag = visitor.GetFlag("showBtn");

//Step 5: get the flag value and if the flag does not exist, it returns the default value "false"
var flagValue = flag.GetValue(false);

var flag_ = visitor.GetFlag("btnSize").GetValue(15);

var flag1 = visitor.GetFlag("btnColor");
var flagValue = flag1.GetValue("red");

Console.WriteLine($"Flag {flagValue}");