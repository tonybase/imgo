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
	var buddies []model.IMUser
	group := model.Group{"id", "tt", buddies}

	user := model.IMUser{"userid", "", "", "", ""}
	users := []model.IMUser{user}

	group.Buddies = users
	fmt.Println(string(group.Encode()))

}
