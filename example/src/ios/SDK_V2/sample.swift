/// Update the context when basket value change
Flagship.sharedInstance.synchronizeModifications { (result) in

  if result == .Updated{

    // Update the UI for users that have basket over or equal 100
    if (Flagship.sharedInstance.getModification("freeDelivery", defaultBool: false, activate: true)){

      DispatchQueue.main.async {

        /// Show your message for free delivery

      }
    }
  }
}


// Retrieve modification and activate
let title = FlagShip.sharedInstance.getModification("btnColor", defaultString: "#000000", activate: true)

// Retrieve modification and activate
let title = FlagShip.sharedInstance.getModification("btnSize", defaultDouble: 13, activate: true)
let title = FlagShip.sharedInstance.getModification("btnSizeFloat", defaultFloat: 13.1, activate: true)
let title = FlagShip.sharedInstance.getModification("btnSizeInt", defaultInt: 13.00, activate: true)




// END
