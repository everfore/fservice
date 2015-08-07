package models

import (
	"github.com/astaxie/beego/session"
)

var GlobalSessions *session.Manager

func init() {
	// GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	GlobalSessions, _ = session.NewManager("file", `{"cookieName":"gosessionid","sessionsavepath":"./sessionpath/", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	defer GlobalSessions.GC()

}
