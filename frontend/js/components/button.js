function button(label='Submit', id='form', onclick='', icon='plus'){
    if (typeof icon === 'undefined') { icon = 'plus'; }

    var btn = document.createElement("button");
    btn.setAttribute('type', 'button');
    btn.setAttribute('value', label)
    btn.setAttribute('onclick', onclick)
    btn.setAttribute('id', id);
    btn.setAttribute('class', 'btn btn-primary');
    
    //btn.innerHTML = label;
    var icn = genIcon(icon);
    btn.appendChild(icn);
    
    btn.insertAdjacentHTML('beforeend', '&nbsp;&nbsp;');
    btn.insertAdjacentHTML('beforeend', label);
    return(btn)
}