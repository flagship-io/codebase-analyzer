import flagship from "@flagship.io/js-sdk";

const fsInstance = flagship.start("YOUR_ENV_ID", "YOUR_API_KEY", {});

const fsVisitorInstance = fsInstance.newVisitor("YOUR_VISITOR_ID", {
    some: "VISITOR_CONTEXT",
});

const modificationsOutput = fsVisitorInstance.getModifications(
    [
        {
            key: "btnColor", // required
            defaultValue: "#ff0000", // required
            activate: true, // optional ("false" by default)
        },
        {
            key: "backgroundColor", // required
            defaultValue: "green", // required
        },
        {
            key: "backgroundSize", // required
            defaultValue: 16, // required
        },
        {
            key: "showbackground", // required
            defaultValue: true, // required
        },
    ]
);

console.log('Flags : ', modificationsOutput)

