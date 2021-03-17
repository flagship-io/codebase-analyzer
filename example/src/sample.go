package sample

func main() {
  // Using the Decision API (default)
  fsClient, err := flagship.Start(environmentID, apiKey)
  
  // Using the Bucketing mode
  fsClient, err := flagship.Start(environmentID, apiKey, client.WithBucketing())
  
  // Create visitor context
  context := map[string]interface{}{
    "isVip": true,
    "age": 30,
    "name": "visitor",
    }
    // Create a visitor
    fsVisitor, err := fsClient.NewVisitor("visitor_id", context)
  
    // Update a single key
  fsVisitor.UpdateContextKey("vipUser", true)
  fsVisitor.UpdateContextLey("age", 30)
  
  // Update the whole context
  newContext := map[string]interface{}{
    "isVip": true,
    "age": 30,
    "name": "visitor",
  }
  fsVisitor.UpdateContext(newContext)
  
  discountName, err := fsVisitor.GetModificationString("btnColor", "VString", true);
  discountName, err := fsVisitor.GetModificationNumber("btnColor", "VNumber", true);
  discountName, err := fsVisitor.GetModificationBool("btnColor", "VBool", true);
  discountName, err := fsVisitor.GetModificationArray("btnColor", "VArray", true);
  discountName, err := fsVisitor.GetModificationObject("btnColor", "VObject", true);
  
  // If there is not error (and if there is, your value will still be set to defaut), you can use your modification value in your business logic
  discountValue := getDiscountFromDB(discountName)
}
