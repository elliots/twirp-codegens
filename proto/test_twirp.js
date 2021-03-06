// Code generated by protoc-gen-twirp_browserjs v5.0.0, DO NOT EDIT.
// source: test.proto

// _request takes the HTTP method, URL path (this will usually contain the domain name)
// json body that will be sent to the server, a callback on successful requests and a
// callback for requests that error out.
var _request = function(method, path, body, onSuccess, onError) {
  var xhr = new XMLHttpRequest();
  xhr.open(method, path, true);
  xhr.setRequestHeader("Accept","application/json");
  xhr.setRequestHeader("Content-Type","application/json");

  xhr.onreadystatechange = function (e) {
    if (xhr.readyState == 4) {
      if (xhr.status == 204 || xhr.status == 205) {
        onSuccess();
      } else if (xhr.status == 200) {
        var value = JSON.parse(xhr.responseText);
        onSuccess(value);
      } else {
        var value = JSON.parse(xhr.responseText);
        onError(value);
      }
    }
  };

  if (body != null) {
    xhr.send(JSON.stringify(body));
  } else {
    xhr.send(null);
  }
};

// methods for HelloWorldClient

var HelloWorld_speak = function(server_address, words, onSuccess, onError) {
  var full_method = server_address + "/twirp/" + "us.xeserv.api.HelloWorld" + "/" + "Speak";
  _request("POST", full_method, words, onSuccess, onError);
};
