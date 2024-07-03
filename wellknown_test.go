package syncspecv1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWellknown_MarshalJSON(t *testing.T) {
	w := Wellknown{
		TokenEndpoint:            "https://www.example.com/sync/v1/token",
		ListDepartmentsEndpoint:  "https://www.example.com/sync/v1/depts",
		ListUsersInDeptEndpoint:  "https://www.example.com/sync/v1/users",
		SearchUserEndpoint:       "https://www.example.com/sync/v1/users:search",
		SearchDepartmentEndpoint: "https://www.example.com/sync/v1/depts:search",
	}

	b, err := json.Marshal(&w)
	assert.NoError(t, err)
	assert.Contains(t, string(b), `"spec":"v1"`)
}
