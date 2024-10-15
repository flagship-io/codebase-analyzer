////////////////////////////////
///////    Get Flag      ///////
////////////////////////////////

import 'package:flagship/flagship.dart';
import 'package:flagship/model/flag.dart';


// Get the current visitor
    var currentVisitor = Flagship.getCurrentVisitor();

// Fetch flags
    currentVisitor?.fetchFlags().whenComplete(() {
      // Ex: get flag for vip feature
      Flag flag = currentVisitor.getFlag("displayVipFeature",false);
      Flag flag = currentVisitor.getFlag("backgroundColor", "green");
      Flag flag = currentVisitor.getFlag("backgroundSize", 16);
    });

// The getFlag function can be directly called through the "currentVisitor" instance if you
// already fetched it elsewhere in the application
