package userTransport

import (
	"backend_autotest/component"
	"backend_autotest/modules/user/userBiz"
	"backend_autotest/modules/user/userModel"
	"backend_autotest/modules/user/userStorage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UserRegister(app component.AppContext) http.HandlerFunc {
	fb := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome teachers!")
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var data userModel.User
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := userBiz.NewCreateUserBiz(store)
		if err := biz.CreateNewUser(r.Context(), &data); err != nil {
			panic(err)
		}

		fmt.Fprint(w, "successful registration")
	}

	return http.HandlerFunc(fb)
}
