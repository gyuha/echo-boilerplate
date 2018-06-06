$(function () {
    var loc = window.location;
    var protocol = loc.protocol == 'http:' ? 'ws:' : 'wss:';
    var url = protocol + '//' + loc.host + '/ws/channel/' + channel;
    var name = "Guest" + Math.floor(Math.random() * 1000);
    var ws = new WebSocket(url);

    ws.onmessage = function (msg) {
        data = JSON.parse(msg.data);
        side = data.name == name ? 'right' : 'left';
        messageAdd(side, data.text);
    };

    sendWsMessage = function(text) {
        text = JSON.stringify({name: name, text: text});
        ws.send(text);
    };
});