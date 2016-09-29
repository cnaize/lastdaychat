<!DOCTYPE html>
<html lang="en">
<head>
    <title>lastday chat</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.7/css/materialize.min.css">
    <link rel="stylesheet" href="/static/css/main.css" type="text/css">
    <script src="/static/js/jquery-2.1.1.min.js" type="text/javascript"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.7/js/materialize.min.js" type="text/javascript"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/rollups/aes.js" type="text/javascript"></script>
    <script src="/static/js/js.cookie.js" type="text/javascript"></script>
    <script src="/static/js/ws.js" type="text/javascript"></script>
    <script src="/static/js/util.js" type="text/javascript"></script>
    
    <script type="text/javascript">
        window.onload = runWs({{.}});
        window.onunload = clearAll;
    </script>
</head>
<body>
    <div class="row">
        <div class="col s12 m6 l6 offset-m3 offset-l3 hoverable z-depth-3">
            <nav class="green darken-4 z-depth-3">
                <div class="nav-wrapper">
                    <ul class="left">
                        <li>
                            <a onclick="zoomRoomId();">
                                room id: <span id="roomId" class="black-text">{{.roomId}}</span>
                            </a>
                        </li>
                    </ul>
                    <span id="people-counter" class="new badge right black" data-badge-caption="people">0</span>
                </div>
            </nav>
            <ul id="log-wrapper" class="collection hoverable z-depth-3">
            </ul>
            <form id="form-msg">
                <div row>
                    <div class="input-field col s12 m10 l10">
                        <input placeholder="secure message" type="text" id="msg"/>
                    </div>
                    <div class="input-field col m2 l2 hide-on-med-and-down">
                        <button type="submit" class="btn waves-effect green darken-4 z-depth-3">send</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</body>
</html>
