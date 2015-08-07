package models

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
)

var GlobalSessions *session.Manager

func init() {
	// GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	// GlobalSessions, _ = session.NewManager("file", `{"cookieName":"gosessionid","sessionsavepath":"./sessionpath/", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	GlobalSessions, _ = session.NewManager("mysql", `{"cookieName":"gosessionid","sessionsavepath":"./sessionpath/", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": "root:1234@3306(localhost)/fss?param1=session_key&param2=session_data&param3=session_expiry"}`)
	defer GlobalSessions.GC()

}
