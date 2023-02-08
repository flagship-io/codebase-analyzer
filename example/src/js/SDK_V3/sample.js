import { Flagship } from "@flagship.io/js-sdk";

Flagship.start("your_env_id", "your_api_key");

const visitor = Flagship.newVisitor({
    visitorId: "your_visitor_id",
    context: { key: "value" },
});

visitor.on("ready", (error) => {
    if (error) {
        return;
    }

    const btnColorFlag = visitor.getFlag("btnColor", 'red').getValue();
    const backgroundColorFlag = visitor.getFlag("backgroundColor", 'green').getValue();
    const backgroundSize = visitor.getFlag("backgroundSize", 16).getValue();
    const showBackground = visitor.getFlag("showBackground", true).getValue();

    console.log('btnColorFlag : ', btnColorFlag)
    console.log('backgroundColor : ', backgroundColorFlag)
});

visitor.fetchFlags();


