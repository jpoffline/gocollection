const BEROOT = "http://localhost:10000";



function urlRoot(){
    return(BEROOT)
}

function toTable(data){
    var table = document.createElement("table");
    var tr = table.insertRow(-1); 
    for (var i = 0; i < data.Header.length; i++) {
        var th = document.createElement("th");      // TABLE HEADER.
        th.innerHTML = data.Header[i];
        tr.appendChild(th);
    }

    for (var i = 0; i < data.Rows.length; i++) {

        tr = table.insertRow(-1);

        for (var j = 0; j < data.Header.length; j++) {
            var tabCell = tr.insertCell(-1);
            tabCell.innerHTML = data.Rows[i][j];
        }
    }
    var divContainer = document.getElementById("collectionDataTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}

function CollectionNameBox(collection){
    var div = document.createElement("div");
    div.setAttribute('class', 'collectionnamebox');
    div.setAttribute('id',  collection.name);
    div.setAttribute('onclick', 'ShowCollectionTable(this.id)');
    div.innerHTML = collection.id + "//" + collection.label
    return(div)
}

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

function RouteRespond(route, callback){
    
        var request = new XMLHttpRequest();
        request.open('GET', urlRoot() + route, true);
        
        request.onload = function () {
            // Begin accessing JSON data here
            callback(JSON.parse(this.response))
        }

        // Send request
        request.send();
}

function ShowCollectionTable(which='coffee'){
    RouteRespond('/collection/' + which, toTable)
    document.getElementById("selectedCollectionName").innerHTML = which
}

function CollectionNames(){
    RouteRespond('/collections', ShowCollectionNames)
}