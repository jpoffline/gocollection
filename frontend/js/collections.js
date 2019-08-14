


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

function submitNewRecord(id){
    var elements = document.getElementById(id).childNodes;
    var obj ={};
    for(var i = 0 ; i < elements.length-1 ; i++){
        var item = elements.item(i);
        var field = item.querySelectorAll('[type="text"]')[0];
        obj[field.id] = field.value;
    }
    console.log(obj)
    var collname = document.getElementById('selectedCollectionName').innerHTML;

    var request = new XMLHttpRequest();
    
    url = urlRoot() + '/collection/' + collname;
    
    console.log(url)
    request.open('POST', url, true);
    request.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');

    request.onreadystatechange = function () {
        if (this.readyState != 4) return;
    
        if (this.status == 200) {
            //var data = JSON.parse(this.responseText);
            ShowCollectionTable(collname)
            // we get the returned data
        }
    
        // end of state change: it can be after some time (async)
    };

    request.send(JSON.stringify({'data':obj, 'id' : 4}));
    


}

function RenderAddRowToCollection(names){
    
    var formid = 'thisform';
    var divContainer = document.getElementById("formAddRowToCollection");
    var form = document.createElement("form");
    form.setAttribute('id', formid);
    form.setAttribute('class', 'container');
    names.forEach(field => {
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


