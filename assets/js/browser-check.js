if (!window.WebSocket && window.location.toString().indexOf('/unsupported-browser') == -1){
     window.location = '/unsupported-browser';
}