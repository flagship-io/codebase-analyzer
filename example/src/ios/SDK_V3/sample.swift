import Flagship


// Create the visitor
let visitor1 = Flagship.sharedInstance.newVisitor("visitor_1").build()

// Fetch flags
visitor1.fetchFlags {
    let flag = visitor1.getFlag(key: "btnColor", defaultValue: 'red')
}