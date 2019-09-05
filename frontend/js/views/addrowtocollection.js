function RenderAddRowToCollection(data) {

    fields = data.fields

    var formid = 'newrecordform';

    var frm = CreateFieldsInputForm(formid, fields);

    frm.appendChild(
        button(
            label = 'Add',
            id = 'submitform',
            onclick = 'submitNewRecord("' + formid + '")'
        )
    );

    var divContainer = document.getElementById("formAddRowToCollection");
    divContainer.innerHTML = ""
    divContainer.appendChild(h4("Add row"));
    divContainer.appendChild(frm);

    RenderFieldsEditArea(fields)


}
