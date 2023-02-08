import { Flagship } from "@flagship.io/js-sdk";

Flagship.start("your_env_id", "your_api_key");

const visitor = Flagship.newVisitor({
    visitorId: "your_visitor_id",
    context: { isVip: true },
});

visitor.on("ready",  (error) => {
    if (error) {
        return;
    }

    const btnColorFlag = visitor.getFlag('btnColor', 'red');
    const backgroundColorFlag: string = visitor.getFlag("backgroundColor", 'green').getValue();
    const backgroundSizeFlag: number = visitor.getFlag("backgroundSize", 16).getValue();
    const showBackgroundFlag: boolean = visitor.getFlag("showBackground", true).getValue();

    console.log('btnColorFlag : ', btnColorFlag);
    console.log('backgroundColorFlag : ', backgroundColorFlag);
});

visitor.fetchFlags();