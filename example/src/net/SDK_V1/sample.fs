open System.Collections.Generic
open Flagship

let client = FlagshipBuilder.Start("ENV_ID","API_KEY");

let context = new Dictionary<string, obj>();
context.Add("key", "value");

let visitor = client.NewVisitor("visitor_id", context);

let btnColorFlag = visitor.GetModification("btnColor", "red", true);
let btnSizeFlag = visitor.GetModification("btnSize", 13, true);
let showBtnFlag = visitor.GetModification("showBtn", false, true);
