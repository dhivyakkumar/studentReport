package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Client(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	Stud := NewMockStudentRepoOperations(ctrl)

	t.Run("Client can list students resources", func(t *testing.T) {
		ExpectedRes := StudentsResponse()
		Stud.EXPECT().GetStudents().Return(ExpectedRes, nil)
		ctl := Controller{
			studentRepo: Stud,
		}

		request, _ := http.NewRequest("GET", "/api/students", nil)
		resp := httptest.NewRecorder()

		handler := http.HandlerFunc(ctl.getAllStudentList)
		handler.ServeHTTP(resp, request)

		assert.Equal(t, 200, resp.Code)
	})

	t.Run("Client can list student resource", func(t *testing.T) {
		ExpectedRes := StudentStruct()
		id := 0
		Stud.EXPECT().GetStudentInfo(id).Return(ExpectedRes, nil)
		ctl := Controller{
			studentRepo: Stud,
		}

		request, _ := http.NewRequest("GET", "/api/student/1", nil)
		resp := httptest.NewRecorder()

		handler := http.HandlerFunc(ctl.getStudent)
		handler.ServeHTTP(resp, request)

		assert.Equal(t, 200, resp.Code)
		//Erep, _:=json.Marshal(ExpectedRes)
		//assert.Equal(t, string(Erep), resp.Body.String())
	})

	t.Run("Client can create student resource", func(t *testing.T) {
		ExpectedRes := StudentStruct()
		studReq := StudentStruct()
		Stud.EXPECT().CreateStudentInfo(studReq).Return(ExpectedRes, nil)
		ctl := Controller{
			studentRepo: Stud,
		}

		payload, err := json.Marshal(studReq)
		if err != nil {
			fmt.Errorf("Failed to marshal %v", err)
		}

		request, _ := http.NewRequest("POST", "/api/student", bytes.NewBuffer(payload))
		resp := httptest.NewRecorder()

		handler := http.HandlerFunc(ctl.createStudent)
		handler.ServeHTTP(resp, request)

		assert.Equal(t, 200, resp.Code)
	})

	t.Run("Client can remove student resource", func(t *testing.T) {
		id := 0
		Stud.EXPECT().RemoveStudentInfo(id).Return(nil)
		ctl := Controller{
			studentRepo: Stud,
		}

		request, _ := http.NewRequest("DELETE", "/api/student/1", nil)
		resp := httptest.NewRecorder()

		handler := http.HandlerFunc(ctl.removeStudent)
		handler.ServeHTTP(resp, request)

		assert.Equal(t, 200, resp.Code)
	})

	t.Run("Client can update student resource", func(t *testing.T) {
		id := 0
		studReq := StudentStruct()
		ExpectedRes := StudentStruct()
		Stud.EXPECT().UpdateStudentInfo(id, studReq).Return(ExpectedRes, nil)
		ctl := Controller{
			studentRepo: Stud,
		}

		payload, err := json.Marshal(studReq)
		if err != nil {
			fmt.Errorf("Failed to marshal %v", err)
		}

		request, _ := http.NewRequest("DELETE", "/api/student/1", bytes.NewBuffer(payload))
		resp := httptest.NewRecorder()

		handler := http.HandlerFunc(ctl.updateStudent)
		handler.ServeHTTP(resp, request)

		assert.Equal(t, 200, resp.Code)
	})
}
