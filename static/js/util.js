function clearAll() {
	Cookies.remove("username");
	Cookies.remove("secret");
};

function zoomRoomId() {
	window.prompt("copy room id:", document.getElementById("roomId").innerText);
};