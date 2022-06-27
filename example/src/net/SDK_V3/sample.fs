open System.Collections.Generic
open Flagship

let client = FlagshipBuilder.Start("ENV_ID","API_KEY");

let context = new Dictionary<string, obj>();
context.Add("key", "value");

let visitor = client.NewVisitor("visitor_id", context);

visitor.FetchFlags();
let btnColorFlag = visitor.GetFlag("btnColor", 'red');