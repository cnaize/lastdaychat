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
    <script src="/static/js/js.cookie.js" type="text/javascript"></script>
    <script src="/static/js/forms.js" type="text/javascript"></script>
    <script type="text/javascript">
        window.onload = prepareForms;
    </script>
</head>
<body>
    <div class="row">
        <div class="col s12 m4 l4 offset-m4 offset-l4">
            <h1 class="center green-text text-darken-4">lastday chat</h1>
            <h4 class="center">anonymous secure chat</h4>
            <div class="card-panel hoverable z-depth-3">
                <ul class="tabs">
                    <li class="tab col s6 m4 l4 active"><a href="#join">join</a></li>
                    <li class="tab col s6 m4 l4"><a href="#create">create</a></li>
                </ul>
                <div id="join">
                    <form id="joinForm" action="/rooms" method="GET">
                        <div class="input-field">
                            <input id="joinUsername" class="validate" placeholder="username" type="text" required="" aria-required="true" maxlength="16" />
                        </div>
                        <div class="input-field">
                            <input id="joinRoomId" class="validate" placeholder="room id" name="roomId" type="text" required="" aria-required="true" maxlength="52" />
                        </div>
                        <div class="input-field">
                            <input id="joinSecret" class="validate" placeholder="secret" type="password" required="" aria-required="true" maxlength="52" />
                        </div>
                        <button type="submit" class="btn waves-effect green darken-4 z-depth-3">join</button>
                    </form>
                </div>
                <div id="create">
                    <form id="createForm" action="/rooms" method="POST">
                        <div class="input-field">
                            <input id="createUsername" class="validate" type="text" placeholder="username" type="text" required="" aria-required="true" maxlength="16"/>
                        </div>
                        <div class="input-field">
                            <input id="createSecret" class="validate" placeholder="secret" type="password" required="" aria-required="true" maxlength="52"/>
                        </div>
                        <button type="submit" class="btn waves-effect green darken-4 z-depth-3">create</button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</body>
</html>