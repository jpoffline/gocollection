


function ShowCollectionNames(names) {
    var divContainer = document.getElementById("collectionsNav");

    var hldr = div();
    hldr.setAttribute("class", "navbar-nav")
    hldr.setAttribute("id", "navbaritems");
    names.forEach(collectionname => {
        hldr.appendChild(CollectionNameBox(collectionname));
    })
    var nv = div();
    nv.setAttribute("class", "collapse navbar-collapse");
    nv.setAttribute("id", "navbarNavAltMarkup");
    nv.appendChild(hldr);
    divContainer.innerHTML = '';
    divContainer.appendChild(nv)

    
}

function ShowSelectedCollectionMeta(meta) {
    var area = document.getElementById("selectedCollectionName");

    area.innerHTML = ""
    area.appendChild(h4(meta.label))
    area.setAttribute("name", meta.name);

    var colloptsdiv = document.createElement("div");
    colloptsdiv.innerHTML = "Edit collection properties";
    console.log(meta)




}

function readFormValues(formid) {
    var elements = document.getElementById(formid).childNodes;
    var obj = {};
    for (var i = 0; i < elements.length - 1; i++) {
        var item = elements.item(i);
        var field = item.querySelectorAll('[type="text"]')[0];
        obj[field.name] = field.value;
    }
    return (obj)
}

function isUndefined(value) {
    // Obtain `undefined` value that's
    // guaranteed to not have been re-assigned
    var undefined = void (0);
    return value === undefined;
}

function readFormComplete(formid) {
    var elements = document.getElementById(formid).childNodes;

    var obj = [];
    for (var i = 0; i < elements.length; i++) {
        var item = elements.item(i);
        var field = item.querySelectorAll('[type="text"]')[0];
        if (!isUndefined(field)) {
            var newrec = {};
            newrec['name'] = field.name;
            newrec['id'] = field.id;
            newrec['value'] = field.value
            obj.push(newrec);
        }
    }
    return (obj)
}

function selectedCollectionName() {
    return (
        document.getElementById('selectedCollectionName').getAttribute("name")
    )
}

function submitNewRecord(id) {
    var obj = readFormValues(id)

    var collname = selectedCollectionName();

    RouteRespondPost(
        '/collection/' + collname,
        { 'data': obj },
        ShowCollectionTable,
        collname
    )
}

function CreateFieldsInputForm(id, fields, deft = "") {
    var frm = form(id);

    fields.forEach(
        field => {
            if (deft != "") {
                frm.appendChild(FieldNameArea(field, defaultValue = field.name));
            } else {
                frm.appendChild(FieldNameArea(field));
            }
        }
    )

    return (frm)
}






function createNewCollection(id){
    var newV = document.getElementById(id);
    
    

    var bdy = {
        "name":newV.value.toLowerCase().replace(/\s/g, ''),
        "label":newV.value
    }
    console.log(bdy)
    RouteRespondPost('/addcollection', bdy, CollectionNames)
    
}