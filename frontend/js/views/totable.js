
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