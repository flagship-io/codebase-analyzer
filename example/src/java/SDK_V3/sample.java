Visitor visitor = Flagship.newVisitor("visitor_unique_id").build();
visitor.updateContext("isVip", true);
visitor.fetchFlags().whenComplete((instance, error) -> {
    Flag<String> btnColorFlag = visitor.getFlag("btnColor", "red");
    Flag<Number> backgroundColorFlag = visitor.getFlag("backgroundSize", 13);
    Flag<Boolean> backgroundColorFlag = visitor.getFlag("showBackground", true);
    Flag<String> backgroundColorFlag = visitor.getFlag("backgroundColor", "green");
});
