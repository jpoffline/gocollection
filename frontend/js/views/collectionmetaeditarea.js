function RenderFieldsEditArea(fields) {
    var area = document.getElementById("selectedCollectionOptions")
    var formid = "titleedit";
    var frm = CreateFieldsInputForm(formid, fields, "names");
    frm.appendChild(
        button(
            label = 'Update names',
            id = 'submittitleedit',
            onclick = 'submitFieldNameUpdates("' + formid + '")',
            icon = "pencil"
        )
    );
    area.innerHTML = ""
    area.appendChild(frm);

    renderAddFieldToCollectionArea(formid);
}

function renderAddFieldToCollectionArea(id) {
    var frm = CreateFieldsInputForm(
        'newfieldname', [
            {
                'id': 'newfieldnamevalue',
                'name': 'New field'
            }
        ]
    );
    document.getElementById('newFieldNameArea').innerHTML = ""
    document.getElementById('newFieldNameArea').appendChild(frm);
    document.getElementById('newFieldNameArea').appendChild(
        button(
            label = 'Add',
            id = 'submitaddfield',
            onclick = 'submitNewFieldToCollection("newfieldname")',
            icon = "plus"
        )
    );

}

function submitNewFieldToCollection(formid) {
    var newFieldName = readFormComplete(formid)[0];
    console.log(newFieldName.value);
    var collname = selectedCollectionName();
    data = { 'newfield': newFieldName.value }
    RouteRespondPost('/addfield/' + collname, data, ShowCollectionTable, collname)
}

function submitFieldNameUpdates(formid) {
    var collname = selectedCollectionName();
    var fieldNames = readFormComplete(formid);
    var data = []
    fieldNames.forEach(field => {
        data.push({
            'id': field.id,
            'name': field.value
        });
    })

    RouteRespondPost('/fields/' + collname, data, ShowCollectionTable, collname)

    console.log(collname);

}