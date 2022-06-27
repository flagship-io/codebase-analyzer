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

    console.log('btnColorFlag : ', btnColorFlag);
    console.log('backgroundColorFlag : ', backgroundColorFlag);
});

visitor.fetchFlags();