package inmemory

import (
	"testing"

	"github.com/rcastillo9x/go-ms-boilerplate/app/internal/report/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateReportInMemory(t *testing.T) {

	// Arrange
	s := NewReportInMemory()
	rExpect := &model.Report{
		ID:         1,
		Definition: "test",
	}

	// Act
	rResult, err := s.Create(rExpect)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, rExpect, rResult)

}
func TestFindByIDReportInMemory(t *testing.T) {

	// Arrange
	s := NewReportInMemory()
	rAdd := &model.Report{
		ID:         1,
		Definition: "test",
	}
	rResult, addErr := s.Create(rAdd)
	assert.NoError(t, addErr)

	//
	rFound, findErr := s.FindByID(rResult.ID)

	// Assert
	assert.NoError(t, findErr)
	assert.Equal(t, rFound.ID, rResult.ID)

}

func TestUpdateReportInMemory(t *testing.T) {

	// Arrange
	s := NewReportInMemory()
	rAdd := &model.Report{
		ID:         1,
		Definition: "test",
	}
	rResult, addErr := s.Create(rAdd)
	assert.NoError(t, addErr)

	//Act
	rResult.Definition = "Test2"

	rFound2, err := s.UpdateByID(rAdd)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, rFound2.Definition, rResult.Definition)

}

func TestDeleteByIDReportInMemory(t *testing.T) {

	// Arrange
	s := NewReportInMemory()
	rAdd := &model.Report{
		ID:         1,
		Definition: "test",
	}
	rResult, addErr := s.Create(rAdd)
	assert.NoError(t, addErr)

	//Act
	err := s.DeleteByID(rResult.ID)
	r, errNotFound := s.FindByID(rAdd.ID)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, r)
	assert.Error(t, errNotFound)

}
