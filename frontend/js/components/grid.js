function tableElement(width, body) {
    var elem = document.createElement("div");
    elem.setAttribute("class", "col-" + width);
    elem.appendChild(body)
    return (elem)
}