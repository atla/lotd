new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        room: null,
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {

            var msg = JSON.parse(e.data);

            switch (msg.type) {
                case 'displayRoom':

                    self.room = msg.room;

                    break;

                    
                default:

                if (msg.username != undefined || msg.username == '') {
                    self.chatContent += '<li>' + msg.username + '$  ' + msg.message + '<li/>';
                } else {
                    self.chatContent += '<li><pre>' + msg.message + '</pre><li/>';
                }
    
                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
                    break;
            }

           
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function (event) {
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        }
    }
});