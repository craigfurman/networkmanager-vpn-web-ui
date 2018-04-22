function onReady(fn) {
  if (document.attachEvent ? document.readyState === "complete" : document.readyState !== "loading"){
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}

onReady(function() {
  var request = new XMLHttpRequest();
  request.open('GET', '/api/connections', true);

  request.onload = function() {
    if (request.status == 200) {
      var conns = JSON.parse(request.responseText);
      var connList = document.getElementById('conn-list');
      connList.innerHTML = connsHtml(conns);
    } else {
      throw new Error('error getting connections: ' + request.responseText);
    }
  };

  request.onerror = function() {
    // TODO
  };

  request.send();
});

function connHtml(conn) {
  var status = conn.active ? 'ON' : 'OFF';
  return '<div class="conn">' +
    '<div class="conn-name">' + conn.name + '</div>' +
    '<div class="conn-status">' + status + '</div>' +
    '</div>';
}

function connsHtml(conns) {
  var s = '';
  for (var conn of conns) {
    s += connHtml(conn);
  }
  return s;
}
