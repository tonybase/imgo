package im

import (
	//"flag"
	///"im-go2"
	"fmt"
	"im-go/im/model"
	"testing"
)

func TestConfig(t *testing.T) {
	/**
	  var tt []string;
	  tt=append(tt,"444");
	  fmt.Println(tt[0])
	*/
	var buddies []model.User
	group := model.Category{"id", "tt", buddies}

	user := model.User{"userid", "", "", "", ""}
	users := []model.User{user}

	group.Buddies = users
	fmt.Println(string(group.Encode()))

}
