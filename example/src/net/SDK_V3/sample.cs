using Flagship.Main;

var visitor = Fs.NewVisitor("<VISITOR_ID>")
  .IsAuthenticated(true)
  .HasConsented(true)
  .WithContext(new Dictionary<string, object> {
    ["isVIP"] = true,
    ["country"] = "NL",
    ["loginProvider"] = "Google"
    })
  .Build();

  await visitor.FetchFlags();

var flag = visitor.GetFlag("btnColor", 'red');
var flag = visitor.GetFlag("btnSize", 13);
var flag = visitor.GetFlag("showBtn", true);