const BEROOT = "http://localhost:10000";



function urlRoot(){
    return(BEROOT)
}


function ShowCollectionTable(which='coffee'){
    RouteRespondGet('/collection/' + which, toTable)
    document.getElementById("selectedCollectionName").innerHTML = which
}

function CollectionNames(){
    RouteRespondGet('/collections', ShowCollectionNames)
}

function showAddRowToCollectionForm(which = 'coffee'){
    RouteRespondGet('/fields/' + which, RenderAddRowToCollection)
}