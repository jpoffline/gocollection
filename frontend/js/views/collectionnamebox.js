
function CollectionNameBox(collection) {

    var ni = document.createElement("a");
    ni.setAttribute("class", "nav-item nav-link");
    ni.setAttribute('id', collection.name);
    ni.setAttribute('onclick', 'ShowCollectionTable(this.id)');
    ni.innerHTML = collection.label
    return (ni)

}

function MakeMenuItemActive(which) {
    var nv = document.getElementById("navbaritems").children;
    for (var chld = 0; chld < nv.length; chld++) {
        thisid = nv[chld].getAttribute("id");

        if (thisid == which) {
            nv[chld].classList.add("active")
        } else {
            nv[chld].classList.remove("active")
        }

    }

}