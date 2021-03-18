const flagship = require("@flagship.io/js-sdk"); // ES5

const fsInstance = flagship.start("YOUR_ENV_ID", "YOUR_API_KEY", {
  /* sdk settings */
});

const fsVisitorInstance = fsInstance.newVisitor("YOUR_VISITOR_ID", {
    //...
    some: "VISITOR_CONTEXT",
    //...
  });

const getModificationsOutput = fsVisitorInstance.getModifications(
  [
    {
      key: "btnColor", // required
      defaultValue: "#ff0000", // required
      activate: true, // optional ("false" by default)
    },
    {
      key: "customLabel", // required
      defaultValue: "Flagship is awesome", // required
    },
    {
      key: "key", // required
      defaultValue: "Flagship is awesome", // required
    },
  ] /* ActivateAllModifications */
);
