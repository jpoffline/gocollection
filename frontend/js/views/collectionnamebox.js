
function CollectionNameBox(collection){
    var div = document.createElement("button");
    div.setAttribute('class', 'btn btn-info');
    div.setAttribute('id',  collection.name);
    div.setAttribute('onclick', 'ShowCollectionTable(this.id)');
    div.innerHTML = collection.label
    return(div)
}
