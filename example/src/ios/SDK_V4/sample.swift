import Flagship

// Step 1 - Start the Flagship sdk with default configuration.
Flagship.sharedInstance.start(envId: "_ENV_ID_", apiKey: "_API_KEY_")
        
// Step 2 - Create visitor with context "isVip" : true
let visitor = Flagship.sharedInstance.newVisitor(visitorId: "visitorId", hasConsented: true)
              .withContext(context: ["isVip": true])
              .build()

// Step 3 - Fetch flags
visitor.fetchFlags {
  
    // Fetch completed
  
    // Step 4 - Get Flag key
    let flag = visitor.getFlag(key: "btnColor")

    // Step 5 - Read Flag value
    let value = flag.value(defaultValue: "red")

    let value = visitor.getFlag(key: "displayVipFeature").value(defaultValue: false)

    // Step 4 - Get Flag key
    let flag2 = visitor.getFlag(key: "vipFeature")

    // Step 5 - Read Flag value
    let value = flag2.value(defaultValue: 16)

}