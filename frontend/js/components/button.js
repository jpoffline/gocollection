function button(label='Submit', id='form', onclick=''){
    var btn = document.createElement("button");
    btn.setAttribute('type', 'button');
    btn.setAttribute('value', label)
    btn.setAttribute('onclick', onclick)
    btn.setAttribute('id', id);
    btn.setAttribute('class', 'btn btn-primary')
    btn.innerHTML = label;
    
    return(btn)
}