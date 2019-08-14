function button(label='submit', id='form', onclick=''){
    var btn = document.createElement("button");
    btn.setAttribute('type', 'button');
    btn.setAttribute('value', label)
    btn.setAttribute('onclick', onclick)
    btn.setAttribute('id', id);
    btn.innerHTML = label;
    console.log(btn)
    return(btn)
}