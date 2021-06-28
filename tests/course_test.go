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
	testAccountID := "578"
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
