function prepareForms() {
    var joinForm = document.getElementById("joinForm");
    joinForm.onsubmit = function() {
        var username = document.getElementById("joinUsername");
        var secret = document.getElementById("joinSecret");
        Cookies.set("username", username.value);
        Cookies.set("secret", secret.value);
    };
    var createForm = document.getElementById("createForm");
    createForm.onsubmit = function() {
        username = document.getElementById("createUsername");
        secret = document.getElementById("createSecret");
        Cookies.set("username", username.value);
        Cookies.set("secret", secret.value);
    };
};
