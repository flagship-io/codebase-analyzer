open System.Collections.Generic
open Flagship

let client = FlagshipBuilder.Start("ENV_ID","API_KEY");

let context = new Dictionary<string, obj>();
context.Add("key", "value");

let visitor = client.NewVisitor("visitor_id", context);

visitor.FetchFlags();

let btnColorFlag = visitor.GetFlag("btnColor");
let btnColorFlagValue = btnColorFlag.GetValue('red');

let flag = visitor.GetFlag("btnSize");
let flagValue = flag.GetValue(13);

let showBtnFlag = visitor.GetFlag("showBtn");
let showBtnFlagValue = showBtnFlag.GetValue(true);