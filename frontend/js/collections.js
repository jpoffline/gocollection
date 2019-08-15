


function ShowCollectionNames(names){
    var divContainer = document.getElementById("collectionNameBoxes");
    
    var divs = div();
    var row = div();
    row.setAttribute("class", "row")
    divs.setAttribute('class', 'container');
    
    names.forEach(collectionname => {
        row.appendChild(tableElement("2", CollectionNameBox(collectionname)));
    })

    divs.appendChild(row)

    
    divContainer.appendChild(divs);
}

function ShowSelectedCollectionMeta(meta){
    var area = document.getElementById("selectedCollectionName");
    
    area.innerHTML = ""
    area.appendChild(h4(meta.label))
    area.setAttribute("name", meta.name);

    var colloptsdiv = document.createElement("div");
    colloptsdiv.innerHTML = "Edit collection properties";
    console.log(meta)
    
    


}

function readFormValues(formid){
    var elements = document.getElementById(formid).childNodes;
    var obj ={};
    for(var i = 0 ; i < elements.length-1 ; i++){
        var item = elements.item(i);
        var field = item.querySelectorAll('[type="text"]')[0];
        obj[field.name] = field.value;
    }
    return(obj)
}

function readFormComplete(formid){
    var elements = document.getElementById(formid).childNodes;
    var obj =[];
    for(var i = 0 ; i < elements.length-1 ; i++){
        var item = elements.item(i);
        var field = item.querySelectorAll('[type="text"]')[0];
        //obj[field.name] = field.value;
        var newrec = {};
        newrec['name'] = field.name;
        newrec['id'] = field.id;
        newrec['value'] = field.value
        obj.push(newrec);
    }
    return(obj)
}

function selectedCollectionName(){
    return(
        document.getElementById('selectedCollectionName').getAttribute("name")
    )
}

function submitNewRecord(id){
    var obj = readFormValues(id)
    
    var collname = selectedCollectionName();

    RouteRespondPost(
        '/collection/' + collname,
        {'data':obj},
        ShowCollectionTable,
        collname
    )
}

function CreateFieldsInputForm(id, fields, deft=""){
    var frm = form(id);

    fields.forEach(
        field => {
            if(deft != "") {
                frm.appendChild(FieldNameArea(field, defaultValue = field.name));
            } else {
                frm.appendChild(FieldNameArea(field));
            }
        }
    )

    return(frm)
}

function RenderAddRowToCollection(data){
    
    fields = data.fields
    
    var formid = 'newrecordform';
    
    var frm = CreateFieldsInputForm(formid, fields);

    frm.appendChild(
        button(
            label='Add', 
            id='submitform',
            onclick='submitNewRecord("' + formid +'")'
        )
    );

    var divContainer = document.getElementById("formAddRowToCollection");
    divContainer.innerHTML = ""
    divContainer.appendChild(h4("Add row"));
    divContainer.appendChild(frm);
    
    RenderFieldsEditArea(fields)
    

}

function RenderFieldsEditArea(fields){
    var area = document.getElementById("selectedCollectionOptions")
    var formid = "titleedit";
    var frm  = CreateFieldsInputForm(formid, fields, "names");
    frm.appendChild(
        button(
            label='Update names', 
            id='submittitleedit',
            onclick='submitFieldNameUpdates("' + formid +'")'
        )
    );
    area.innerHTML = ""

    

    area.appendChild(frm);
}


function submitFieldNameUpdates(formid){
    var collname = selectedCollectionName();
    var fieldNames = readFormComplete(formid);
    var data = []
    fieldNames.forEach(field=>{
        data.push({
            'id': field.id,
            'name':field.value
        });
    })

    RouteRespondPost('/fields/' + collname, data, ShowCollectionTable, collname)
    
    console.log(collname);
    
}