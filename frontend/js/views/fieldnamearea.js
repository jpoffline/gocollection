
function FieldNameArea(name, defaultValue=""){

    var divs = divClass('form-group row');

    var divlbl = labelClass(name.name, "col-sm-3 col-form-label");
    divlbl.setAttribute("for", name.id);


    var input = document.createElement("input");
    input.setAttribute('type', 'text');
    input.setAttribute('id',  name.id);
    input.setAttribute('name',  name.name);
    input.setAttribute('class', 'form-control')
    input.setAttribute("value", defaultValue)
    
    var idiv = divClass("col-sm-9");

    idiv.appendChild(input);

    divs.appendChild(divlbl);
    divs.appendChild(idiv);

    return(divs)
}