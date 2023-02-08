val visitor1 = Flagship.newVisitor("visitor_1")
            .context(hashMapOf("age" to "32", "isVIP" to true))
            .hasConsented(true)
            .isAuthenticated(true)
            .build()

val visitor1 = Flagship.newVisitor("visitor_1").build()
visitor1.updateContext("isVip", true)
visitor1.fetchFlags().invokeOnCompletion {
    val btnColorFlag = visitor1.getFlag("btnColor", "#000000")
    val btnSizeFlag = visitor1.getFlag("btnSize", 12)
    val showBackgroundFlag = visitor1.getFlag("showBackground", false)
}
