package controller

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"studentReports/src/model"
	"testing"
)

func Test_Control(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	Stud := NewMockStudentRepoOperations(ctrl)

	t.Run("successfully get students info", func(t *testing.T) {
		ExpectedRes := StudentsResponse()
		Stud.EXPECT().GetStudents().Return(ExpectedRes, nil)
		controller := Controller{
			studentRepo: Stud,
		}

		ActualResp, err := controller.studentRepo.GetStudents()

		require.Error(t, err)
		require.Equal(t, ExpectedRes, ActualResp)
	})

	t.Run("successfully get student info", func(t *testing.T) {
		ExpectedRes := StudentStruct()
		id := 1
		Stud.EXPECT().GetStudentInfo(id).Return(ExpectedRes, nil)
		controller := Controller{
			studentRepo: Stud,
		}

		ActualResp, err := controller.studentRepo.GetStudentInfo(id)

		require.NoError(t, err)
		require.Equal(t, ExpectedRes, ActualResp)
	})

	t.Run("successfully create student info", func(t *testing.T) {
		ExpectedRes := StudentStruct()
		studReq := StudentStruct()
		Stud.EXPECT().CreateStudentInfo(studReq).Return(ExpectedRes, nil)
		controller := Controller{
			studentRepo: Stud,
		}

		ActualResp, err := controller.studentRepo.CreateStudentInfo(studReq)

		require.NoError(t, err)
		require.Equal(t, ExpectedRes, ActualResp)
	})

	t.Run("successfully remove student info", func(t *testing.T) {
		id := 1
		Stud.EXPECT().RemoveStudentInfo(id).Return(nil)
		controller := Controller{
			studentRepo: Stud,
		}

		err := controller.studentRepo.RemoveStudentInfo(id)

		require.NoError(t, err)
	})

	t.Run("successfully update student info", func(t *testing.T) {
		id := 1
		studReq := StudentStruct()
		ExpectedRes := StudentStruct()
		Stud.EXPECT().UpdateStudentInfo(id, studReq).Return(ExpectedRes, nil)
		controller := Controller{
			studentRepo: Stud,
		}

		ActualResp, err := controller.studentRepo.UpdateStudentInfo(id, studReq)

		require.Equal(t, ExpectedRes, ActualResp)
		require.NoError(t, err)
	})
}

func StudentsResponse() []model.Student {
	return []model.Student{
		{
			ID:       1,
			Name:     "test",
			Subject1: 90,
			Subject2: 78,
			Total:    123,
			Avg:      80.0,
			Rank:     1,
		},
	}
}

func StudentStruct() model.Student {
	return model.Student{
		ID:       1,
		Name:     "test",
		Subject1: 90,
		Subject2: 78,
		Total:    123,
		Avg:      80.0,
		Rank:     1,
	}
}
