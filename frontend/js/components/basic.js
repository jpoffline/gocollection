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

function divClass(cls){
    var d = div();
    d.setAttribute("class", cls);
    return(d)
}

function form(id){
    var frm = document.createElement("form");
    frm.setAttribute('id', id);
    return(frm)
}

function genlabel(text){
    var lbl = document.createElement("label");
    lbl.innerHTML = text;
    return(lbl);
}

function labelClass(text, cls){
    var lbl = genlabel(text);
    lbl.setAttribute("class", cls);
    return(lbl);
}

function link(lbl, href, cls) {
    var a = document.createElement("a");
    a.setAttribute("class", cls);
    a.setAttribute("href", href);
    a.innerHTML = lbl;
    return (a)
}

function genSpan(lbl){
    var spn = document.createElement("span");
    spn.innerHTML = spn;
    return(spn);
}