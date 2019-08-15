function title(level, text){
    var h = document.createElement(level);
    h.innerHTML = text;
    return(h)
}

function h4(text){
    return(title("h4", text))
}

function h2(text){
    return(title("h2", text))
}

function div(){
    return(document.createElement("div"))
}

function form(id){
    var frm = document.createElement("form");
    frm.setAttribute('id', id);
    return(frm)
}