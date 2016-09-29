function runWs(params) {
    return function() {
        var conn;
        var audio = new Audio("/static/sounds/om.mp3");
        var msg = document.getElementById("msg");
        var log = document.getElementById("log-wrapper");
        var people = document.getElementById("people-counter");
        var username = Cookies.get("username");
        var secret = Cookies.get("secret");
        clearAll();
        function appendLog(item) {
            audio.play();
            var doScroll = log.scrollTop === log.scrollHeight - log.clientHeight;
            log.appendChild(item);
            if (doScroll) {
                log.scrollTop = log.scrollHeight - log.clientHeight;
            }
        }
        document.getElementById("form-msg").onsubmit = function () {
            if (!conn) {
                return false;
            }
            if (!msg.value) {
                return false;
            }
            if (!username || !secret) {
                conn.close();
                return false;
            }
            var message = CryptoJS.enc.Utf8.parse(username + ": " + msg.value);
            var enc = CryptoJS.AES.encrypt(message, secret);
            conn.send(enc.toString());
            msg.value = "";
            return false;
        };
        if (window["WebSocket"]) {
            conn = new WebSocket("wss://" + params.host + "/ws?roomId=" + params.roomId);
            conn.onclose = function (evt) {
                var item = document.createElement("li");
                item.className = "collection-item";
                item.innerHTML = "<b>Connection closed</b>";
                appendLog(item);
            };
            conn.onmessage = function (evt) {
                var messages = evt.data.split('\n');
                for (var i = 0; i < messages.length; i++) {
                    if (messages[i].length === 0) {
                        return;
                    }
                    var dec = CryptoJS.AES.decrypt(messages[i], secret);
                    var message = dec.toString(CryptoJS.enc.Utf8);
                    var item = document.createElement("li");
                    item.className = "collection-item";
                    if (message.length === 0) {
                        item.innerHTML = "<b>invalid message</b>";
                    } else {
                        item.innerText = message;
                    }
                    appendLog(item);
                }
            };
        } else {
            var item = document.createElement("li");
            item.className = "collection-item";
            item.innerHTML = "<b>Your browser does not support WebSockets</b>";
            appendLog(item);
        }
        var getUserCount = function() {
            $.ajax({
                url: "https://" + params.host + "/users?roomId=" + params.roomId,
                cache: false,
                success: function(data) {
                    if (data !== "not found") {
                        people.innerText = data;
                    }
                    setTimeout(getUserCount, 5000);
                }
            });
        };
        getUserCount();
    };
};