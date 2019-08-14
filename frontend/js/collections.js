


function ShowCollectionNames(names){
    var divContainer = document.getElementById("collectionNameBoxes");
    
    var divs = document.createElement("div");
    divs.setAttribute('class', 'container');
    
    names.forEach(collectionname => {
        divs.appendChild(CollectionNameBox(collectionname))
    })

    divContainer.innerHTML = "";
    divContainer.appendChild(divs);
}

function ShowSelectedCollectionMeta(meta){
    document.getElementById("selectedCollectionName").innerHTML = meta.label;
    document.getElementById("selectedCollectionName").setAttribute("name", meta.name);
}

function submitNewRecord(id){
    var elements = document.getElementById(id).childNodes;
    var obj ={};
    for(var i = 0 ; i < elements.length-1 ; i++){
        var item = elements.item(i);
        var field = item.querySelectorAll('[type="text"]')[0];
        obj[field.id] = field.value;
    }
    
    var collname = document.getElementById('selectedCollectionName').getAttribute("name");

    RouteRespondPost(
        '/collection/' + collname,
        {'data':obj, 'id' : 4},
        ShowCollectionTable,
        collname
    )
}

function RenderAddRowToCollection(data){
    
    fields = data.fields
    console.log(data)
    var formid = 'newrecordform';
    var divContainer = document.getElementById("formAddRowToCollection");
    var form = document.createElement("form");
    form.setAttribute('id', formid);
    form.setAttribute('class', 'container');
    fields.forEach(field => {
        form.appendChild(FieldNameArea(field));
    })
    form.appendChild(
        button(
            label='submit', 
            id='submitform',
            onclick='submitNewRecord("' + formid +'")'
        )
    );
    divContainer.innerHTML = "";
    divContainer.appendChild(form);
    
}


