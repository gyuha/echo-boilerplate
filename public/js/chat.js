$(function () {
    var url = 'ws://' + window.location.host + '/ws/channel/' + channel;
    var name = "Guest" + Math.floor(Math.random() * 1000);
    var ws = new WebSocket(url);

    ws.onmessage = function (msg) {
        data = JSON.parse(msg.data);
        if (data.name == name) return;

        messageAdd('left', data.text);
    };

    sendWsMessage = function(text) {
        text = JSON.stringify({name: name, text: text});
        ws.send(text);
    };
});