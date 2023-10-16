package test

import (
	"encoding/json"
	"github.com/duzhiyuan3353669/utils/http_utils"
	"io"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {

	resp, err := http_utils.Get("http://grcadmin-beta.zgzb.com/v1/api/ss-words/noauth/list?page=1&limit=9999999&scope=2", http.Header{
		"Token": []string{"aaaa"},
	})
	if err != nil {
		t.Error(err.Error())
	}
	rs, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	defer resp.Body.Close()

	t.Logf("%s", rs)

}

func Test_Post(t *testing.T) {
	da := map[string]interface{}{
		"aa": 1,
		"bb": 2,
	}

	data, _ := json.Marshal(da)

	resp, err := http_utils.Post("http://127.0.0.1:82/index/test/index", data, http.Header{
		"Token": []string{"aaaa"},
	})
	if err != nil {
		t.Error(err.Error())
	}
	rs, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	defer resp.Body.Close()

	t.Logf("%s", rs)

}
