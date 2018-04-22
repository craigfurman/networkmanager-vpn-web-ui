function onReady(fn) {
  if (document.attachEvent ? document.readyState === "complete" : document.readyState !== "loading"){
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}

function connHtml(conn) {
  var status = conn.active ? 'ON' : 'OFF';
  return '<div class="conn">' +
    '<div class="conn-name">' + conn.name + '</div>' +
    '<button type="button" class="conn-status">' + status + '</button>' +
    '</div>';
}

function connsHtml(conns) {
  var s = '';
  for (var conn of conns) {
    s += connHtml(conn);
  }
  return s;
}

class Request {
  constructor(method, url) {
    this.method = method;
    this.url = url;
  }

  make(onSuccess) {
    var request = new XMLHttpRequest();
    request.open(this.method, this.url, true);

    request.onload = function() {
      if (request.status == 200) {
        onSuccess(request);
      } else {
        throw new Error('error making request: ' + request.responseText);
      }
    };

    request.onerror = function() {
      throw new Error('error attempting to make request');
    };

    request.send();
  }
}

function refreshConns() {
  var req = new Request('GET', '/api/connections');
  req.make(function(xhr) {
    var conns = JSON.parse(xhr.responseText);
    var connList = document.getElementById('conn-list');
    connList.innerHTML = connsHtml(conns);
  });
}

onReady(function() {
  refreshConns();
  window.setInterval(refreshConns, 5000);
});
