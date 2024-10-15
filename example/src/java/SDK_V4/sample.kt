val flagRank = visitor.getFlag("btnColor")
val flagRankValue = flagRank.value("red")

val flagRankValue2 = visitor.getFlag("backgroundSize").value(1)

val flagRank1 = visitor.getFlag("showBackground")
val flagRankValue = flagRank1.value(true)