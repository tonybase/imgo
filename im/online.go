package im
import "im-go/im/model"

var (
    Online = make(map[string]model.User)    // conn key, user data
)

func IsOnline(uid string) bool {
    for _, user := range Online {
        if user.Id == uid {
            return true
        }
    }
    return false
}