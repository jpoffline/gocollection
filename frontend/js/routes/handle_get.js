function RouteRespondGet(route, callback){
    
    var request = new XMLHttpRequest();
    request.open('GET', urlRoot() + route, true);
    
    request.onload = function () {
        // Begin accessing JSON data here
        callback(JSON.parse(this.response))
    }

    // Send request
    request.send();
}