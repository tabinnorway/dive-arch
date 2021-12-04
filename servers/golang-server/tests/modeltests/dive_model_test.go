package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/tabinnorway/golang-server/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllDives(t *testing.T) {

	err := refreshUserAndDiveTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}
	_, _, err = seedUsersAndDives()
	if err != nil {
		log.Fatalf("Error seeding user and post  table %v\n", err)
	}
	posts, err := diveInstance.FindAllDives(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the posts: %v\n", err)
		return
	}
	assert.Equal(t, len(*posts), 1)
}

func TestSaveDive(t *testing.T) {

	err := refreshUserAndDiveTable()
	if err != nil {
		log.Fatalf("Error user and post refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newDive := models.Dive{
		ID:      1,
		Title:   "This is the title",
		DiverID: user.ID,
	}
	savedDive, err := newDive.SaveDive(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the post: %v\n", err)
		return
	}
	assert.Equal(t, newDive.ID, savedDive.ID)
	assert.Equal(t, newDive.Title, savedDive.Title)
	assert.Equal(t, newDive.DiverID, savedDive.DiverID)

}

func TestGetDiveByID(t *testing.T) {

	err := refreshUserAndDiveTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	post, err := seedOneUserAndOneDive()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundDive, err := diveInstance.FindDiveByID(server.DB, post.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundDive.ID, post.ID)
	assert.Equal(t, foundDive.Title, post.Title)
}

func TestUpdateADive(t *testing.T) {

	err := refreshUserAndDiveTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	post, err := seedOneUserAndOneDive()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	postUpdate := models.Dive{
		ID:      1,
		Title:   "modiUpdate",
		DiverID: post.DiverID,
	}
	updatedDive, err := postUpdate.UpdateADive(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedDive.ID, postUpdate.ID)
	assert.Equal(t, updatedDive.Title, postUpdate.Title)
	assert.Equal(t, updatedDive.DiverID, postUpdate.DiverID)
}

func TestDeleteADive(t *testing.T) {
	err := refreshUserAndDiveTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	post, err := seedOneUserAndOneDive()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := diveInstance.DeleteADive(server.DB, post.ID, post.DiverID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
