
function CollectionNameBox(collection){
    var div = document.createElement("div");
    div.setAttribute('class', 'collectionnamebox');
    div.setAttribute('id',  collection.name);
    div.setAttribute('onclick', 'ShowCollectionTable(this.id)');
    div.innerHTML = collection.id + "//" + collection.label
    return(div)
}
