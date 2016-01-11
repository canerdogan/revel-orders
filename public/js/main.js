$(function() {

	var conn;

	// $("#orders").on("click", "a", function(e) {
	// 	if (!conn) {
	// 		return false;
	// 	}
	// 	var $this = $(this);
	//
	// 	var resp = {
	// 		"Timestamp": Math.floor(Date.now() / 1000),
	// 		"Request": {
	// 			"RequestsId": $this.attr('id')
	// 		}
	// 	};
	//
	// 	conn.send(JSON.stringify(resp));
	// 	$this.remove();
	// 	return false;
	// });

	if (window["WebSocket"]) {
		conn = new WebSocket('ws://' + window.location.host + '/ws');

		conn.onmessage = function(event) {
			response = JSON.parse(event.data);
			console.log(response);
			$('#orders').append(tmpl('message_tmpl', {
				event: response.Request
			}));
		}
	} else {
		console.log('Browswer websocket desteklemiyor!');
	}
});
