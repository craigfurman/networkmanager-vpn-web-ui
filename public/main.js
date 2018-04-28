function onReady(fn) {
  if (document.attachEvent ? document.readyState === "complete" : document.readyState !== "loading") {
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}

function refreshConns() {
  const req = new Request('GET', '/api/connections');
  req.make(function(xhr) {
    const conns = JSON.parse(xhr.responseText);
    const connList = document.getElementById('conn-list');
    connList.innerHTML = connsHtml(conns);

    connList.querySelectorAll('button.conn-toggle').forEach(function(el) {
      const currentlyActive = el.textContent == 'ON'
      el.addEventListener('click', function() {
        el.textContent = 'CHANGING...'
        setConnState({
          name: el.id,
          active: !currentlyActive
        });
      });
    });

    refreshIP();
  });
}

function refreshIP() {
  const req = new Request('GET', '/api/address');
  req.make(function(xhr) {
    const ip = JSON.parse(xhr.responseText);
    document.getElementById('ip-address').textContent = `IP Address: ${ip.ip}`;
  });
}

function setConnState(conn) {
  const req = new Request('PUT', `/api/connections/${encodeURI(conn.name)}?active=${encodeURI(conn.active)}`);
  req.make(refreshConns);
}

function connsHtml(conns) {
  return conns.map(function(conn) {
    return connHtml(conn);
  }).reduce(function(htmlForConns, htmlForConn) {
    return htmlForConns + htmlForConn;
  }, '');
}

function connHtml(conn) {
  const status = conn.active ? 'ON' : 'OFF';
  return '<div class="conn">' +
    '<div class="conn-name">' + conn.name + '</div>' +
    '<button class="conn-toggle" type="button" id="'+ conn.name +
    '" class="conn-status">' + status + '</button>' +
    '</div>';
}

class Request {
  constructor(method, url) {
    this.method = method;
    this.url = url;
  }

  make(onSuccess) {
    const request = new XMLHttpRequest();
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

onReady(function() {
  refreshConns();
  window.setInterval(refreshConns, 5000);
});
