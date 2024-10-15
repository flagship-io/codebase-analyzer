import 'package:flagship/flagship.dart';

// Step 1 - Start the Flagship sdk with default configuration.
Flagship.start("_ENV_ID_", "_API_KEY_");

// Step 2 - Create visitor with context "isVip" : true
var visitor = Flagship.newVisitor(visitorId: "visitorId", hasConsented: true)
        .withContext({"isVip": true}).build();

// Step 3 - Fetch flags
    visitor.fetchFlags().whenComplete(() {
      // Step 4 - Get Flag key
      var flag = v.getFlag("displayVipFeature");
      // Step 5 - Read Flag value
      var value = flag.value(false);
      var value1 = v.getFlag("backgroundColor").value("red");
      var value1 = v.getFlag("backgroundSize").value(1);
    });