package test

import (
	"os"
	"strconv"
	"testing"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/requests"
)

func TestCreateNewCourse(t *testing.T) {
	token := os.Getenv("CANVAS_API_TOKEN")
	canvasURL := "atomicjolt.instructure.com"
	testAccountID := os.Getenv("CANVAS_ACCOUNT_ID")
	canvas := canvasapi.New(token, canvasURL)

	// Create the course
	createNewCourse := requests.CreateNewCourse{}
	createNewCourse.Path.AccountID = testAccountID
	createNewCourse.Form.Course.Name = "canvasapi test course"
	createNewCourse.Form.Course.CourseCode = "ctc101"
	createNewCourse.Form.Course.DefaultView = "assignments"
	course, err := createNewCourse.Do(&canvas)
	if err != nil {
		t.Errorf("CreateNewCourse failed: %v", err)
		return
	} else {
		t.Logf("CreateNewCourse returned: %v", course)
	}

	courseID := strconv.FormatInt(course.ID, 10)

	// Update the course
	updateCourse := requests.UpdateCourse{}
	updateCourse.Path.ID = courseID
	updateCourse.Form.Course.Name = "canvasapi test course updated"
	updateCourse.Form.Course.DefaultView = "assignments"
	uerr := updateCourse.Do(&canvas)
	if uerr != nil {
		t.Errorf("UpdateCourse failed: %v", uerr)
	} else {
		t.Logf("UpdateCourse returned: %v", course)
	}

	// Create an assignment in the course
	createAssignment := requests.CreateAssignment{}
	createAssignment.Path.CourseID = courseID
	createAssignment.Form.Assignment.Name = "a test assignment"
	createAssignment.Form.Assignment.SubmissionTypes = []string{"none"}
	createAssignment.Form.Assignment.GradingType = "pass_fail"
	assignment, aerr := createAssignment.Do(&canvas)
	if aerr != nil {
		t.Errorf("CreateAssignment failed: %v", aerr)
	} else {
		t.Logf("CreateAssignment returned: %v", assignment)
	}

	listAssignments := requests.ListAssignmentsAssignments{}
	listAssignments.Path.CourseID = courseID
	listAssignments.Query.Include = []string{"submission", "can_edit"}
	assignments, pager, laerr := listAssignments.Do(&canvas, nil)
	if laerr != nil {
		t.Errorf("ListAssignmentsAssignments failed: %v", laerr)
	} else {
		t.Logf("ListAssignmentsAssignments returned: %v", assignments[0])
	}
	if pager.Current.Page != 1 {
		t.Errorf("Expected pager to be on page 1")
	}
	if pager.Next != nil {
		moreAssignments, pager2, laerr2 := listAssignments.Do(&canvas, pager.Next.URL)
		if laerr2 != nil {
			t.Errorf("ListAssignmentsAssignments failed: %v", laerr)
		} else {
			t.Logf("ListAssignmentsAssignments returned: %v", moreAssignments[0])
		}
		if pager2.Next != nil {
			t.Errorf("Didn't expect so many pages")
		}
	}

	accessCode := "atestcode"
	createQuiz := requests.CreateQuiz{}
	createQuiz.Path.CourseID = courseID
	createQuiz.Form.Quiz.Title = "a test quiz"
	createQuiz.Form.Quiz.AccessCode = accessCode
	quiz, qerr := createQuiz.Do(&canvas)
	if qerr != nil {
		t.Errorf("CreateQuiz failed: %v", qerr)
		return
	} else {
		t.Logf("CreateQuiz returned: %v", course)
	}
	quizID := strconv.FormatInt(quiz.ID, 10)

	// Test validate access code
	validateCode := requests.ValidateQuizAccessCode{}
	validateCode.Path.CourseID = courseID
	validateCode.Path.ID = quizID
	validateCode.Form.AccessCode = accessCode
	valid, verr := validateCode.Do(&canvas)
	if verr != nil {
		t.Errorf("ValidateQuizAccessCode failed: %v", verr)
	} else {
		if !valid {
			t.Errorf("ValidateQuizAccessCode expected access code %v to be valid.", accessCode)
		}
	}

	// Delete the course
	deleteCourse := requests.DeleteConcludeCourse{}
	deleteCourse.Path.ID = courseID
	deleteCourse.Query.Event = "delete"
	derr := deleteCourse.Do(&canvas)
	if derr != nil {
		t.Errorf("DeleteConcludeCourse failed: %v", derr)
	} else {
		t.Logf("DeleteConcludeCourse returned: %v", course)
	}
}
