import Flagship = Flagship

Private Sub SurroundingSub()
    Dim client As [let] = FlagshipBuilder.Start("ENV_ID", "API_KEY")

    Dim context = New Dictionary(Of String, Object)()
    context.Add("key", "value")

    Dim visitor As [let] = client.NewVisitor("visitor_id", context)

    Dim btnColorFlag As [let] = visitor.GetModification("btnColor", "red", True)
    Dim btnSizeFlag As [let] = visitor.GetModification("btnSize", 13, True)
    Dim showBtnFlag As [let] = visitor.GetModification("showBtn", True, True)
End Sub