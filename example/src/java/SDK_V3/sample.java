Visitor visitor = Flagship.newVisitor("visitor_unique_id").build();
visitor.updateContext("isVip", true);
visitor.fetchFlags().whenComplete((instance, error) -> {
    Flag<Boolean> btnColorFlag = visitor.getFlag("btnColor", "red");
    Flag<Boolean> backgroundColorFlag = visitor.getFlag("backgroundColor", "green");
});
