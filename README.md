# Jukebox

Golang WebSocket server.  I'm completely improvising this, so this sentence
will disappear in a future commit.

depends: labix mgo, go.net websocket

S -> C
S <- C

<-: (connect)
->: ready()
<-: session(id, secret) referring to a previous session or the lack thereof
->: session.response(newId, newSecret) updating the session information of the
client
->: user(...) updating the user information of the client; users can be 'fake'
i.e. they don't have a password, only a secret and auto-genned nick, but they
can still interact just like normal users; they're simply more ephemeral.
(possibly)
<-: login(nick, password)
->: login.response(ok|nok, reason)
->: user(...) the logged in user if successful

Apart from those administrative interactions, imagine application interactions
and stuff here.

->: error() at any time, this message indicates the server encountered an error
and will close the connection
