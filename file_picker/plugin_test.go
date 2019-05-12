package file_picker

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FilePickerPluginTestSuite struct {
	suite.Suite
	FilePickerPlugin FilePickerPlugin
}

type MockDialog struct {
	mock.Mock
}

func (m MockDialog) File(title string, filter string, directory bool) (string, bool, error) {
	args := m.Called(title, filter, directory)
	return args.String(0), args.Bool(1), args.Error(2)
}

func (suite *FilePickerPluginTestSuite) TestDirectoryPickerShouldOnlyAcceptArgsInJSONFormat() {
	filePickerPlugin := suite.FilePickerPlugin
	mockDialog := new(MockDialog)
	directoryPicker := filePickerPlugin.directoryPicker(mockDialog)

	_, err := directoryPicker("im not a json value")
	assert.Equal(suite.T(), "arguments must be encoded in JSON format", err.Error())
}

func (suite *FilePickerPluginTestSuite) TestDirectoryPickerShouldRequireATitleArgument() {
	filePickerPlugin := suite.FilePickerPlugin
	mockDialog := new(MockDialog)
	directoryPicker := filePickerPlugin.directoryPicker(mockDialog)

	_, err := directoryPicker(map[interface{}]interface{}{})
	assert.Equal(suite.T(), "arguments requires a title parameter with type string", err.Error())
}

func (suite *FilePickerPluginTestSuite) TestDirectoryPickerShouldGetTheChosenDirectory() {
	filePickerPlugin := suite.FilePickerPlugin
	mockDialog := new(MockDialog)
	const title = "Open a Directory"
	const chosenDirectory = "/home/flutter-user"
	mockDialog.On("File", title, "*", true).Return(chosenDirectory, true, nil)
	directoryPicker := filePickerPlugin.directoryPicker(mockDialog)

	actualChosenDirectory, err := directoryPicker(map[interface{}]interface{}{
		"title": title,
	})
	assert.Equal(suite.T(), chosenDirectory, actualChosenDirectory)
	assert.Nil(suite.T(), err)
}

func (suite *FilePickerPluginTestSuite) TestDirectoryPickerShouldHandleDialogPickerFailure() {
	filePickerPlugin := suite.FilePickerPlugin
	mockDialog := new(MockDialog)
	const title = "Open a Directory"
	mockDialog.On("File", title, "*", true).Return("", false, errors.New(""))
	directoryPicker := filePickerPlugin.directoryPicker(mockDialog)

	_, err := directoryPicker(map[interface{}]interface{}{
		"title": title,
	})
	assert.Equal(suite.T(), "failed to open dialog picker: ", err.Error())
}

func TestFilePickerPluginTestSuite(t *testing.T) {
	suite.Run(t, new(FilePickerPluginTestSuite))
}