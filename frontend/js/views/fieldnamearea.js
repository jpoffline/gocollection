
function FieldNameArea(name, defaultValue=""){
    console.log(name)
    var divs = document.createElement("div");
    divs.setAttribute('class', 'input-group');

    var divlbl = document.createElement("div");
    divlbl.setAttribute("class", "input-group-prepend")

    var label = document.createElement("span")
    label.setAttribute("class", "input-group-text")
    //label.setAttribute('for', name.name);
    
    label.innerHTML = name.name;
    divlbl.appendChild(label)
    var input = document.createElement("input");
    input.setAttribute('type', 'text');
    input.setAttribute('id',  name.id);
    input.setAttribute('name',  name.name);
    input.setAttribute('class', 'form-control')
    input.setAttribute("value", defaultValue)
    
    divs.appendChild(label);
    divs.appendChild(input);

    return(divs)
}