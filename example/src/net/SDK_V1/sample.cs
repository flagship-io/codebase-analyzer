using Flagship;
var client = FlagshipBuilder.Start(
  "ENV_ID",
  "API_KEY"
);

var context = new Dictionary<string, object>();
context.Add("key", "value");
var visitor = client.NewVisitor("visitor_id", context);
await visitor.SynchronizeModifications();

var btnColorFlag = visitor.GetModification("btnColor", "red", true);
