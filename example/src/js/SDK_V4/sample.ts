import { Flagship } from "@flagship.io/js-sdk";

Flagship.start("your_env_id", "your_api_key");

const visitor = Flagship.newVisitor({
  visitorId: "your_visitor_id",
  context: { isVip: true },
});

visitor.fetchFlags();

const variableKey = visitor.getFlag("flagKey");
const boxSizeFlagDefaultValue = variableKey.getValue("flagDefaultValue");

const variableKey1: any = visitor.getFlag("flagKey1");

const variableKey3: any = visitor.getFlag("flagKey3");
const boxSizeFlagDefaultValue1 = variableKey3.getValue(16);

const variableKey4: any = visitor.getFlag("flagKey4");

const variableKey5: any = visitor.getFlag("flagKey5").getValue(false);

// fe:flag: flagKey6, true
const variable6: any = visitor.getFlag("flagKey6");

visitor.getFlag("flagKey5").getValue(false);

visitor.getFlagValue("FlagKey5", false);
