
function FieldNameArea(name){
    var divs = document.createElement("div");
    divs.setAttribute('class', 'container');

    var label = document.createElement("label")
    label.setAttribute('for', name.name);
    label.setAttribute('class', 'label')
    label.innerHTML = name.name;

    var input = document.createElement("input");
    input.setAttribute('type', 'text');
    input.setAttribute('id',  name.name);
  
    
    divs.appendChild(label);
    divs.appendChild(input);

    return(divs)
}