import Flagship = Flagship

Private Sub SurroundingSub()
    Dim client As [let] = FlagshipBuilder.Start("ENV_ID", "API_KEY")

    Dim context = New Dictionary(Of String, Object)()
    context.Add("key", "value")

    Dim visitor As [let] = client.NewVisitor("visitor_id", context)
    visitor.FetchFlags();

    Dim btnColorFlag As [let] = visitor.GetFlag("btnColor", "red");
    Dim btnSizeFlag As [let] = visitor.GetFlag("btnSize", 13);
    Dim showBtnFlag As [let] = visitor.GetFlag("showBtn", True);
End Sub