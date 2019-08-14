function RouteRespondPost(route, data, callback, args){
    
    var request = new XMLHttpRequest();
    request.open('POST', urlRoot() + route, true);
    request.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    request.onreadystatechange = function () {
        if (this.readyState != 4) return;
        if (this.status == 200) {
            callback(args)
        }
    
    };

    // Send request
    
    request.send(JSON.stringify(data));
}