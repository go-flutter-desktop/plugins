package file_picker

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockDialog struct {
	mock.Mock
}

type FilePickerPluginTestSuite struct {
	suite.Suite
	directoryPicker func(arguments interface{}) (reply interface{}, err error)
	filePicker func(arguments interface{}) (reply interface{}, err error)
	mockDirectoryDialog *MockDialog
	mockFileDialog *MockDialog
}



func (suite *FilePickerPluginTestSuite) SetupTest() {
	filePickerPlugin := FilePickerPlugin{}
	suite.mockDirectoryDialog = new(MockDialog)
	suite.mockFileDialog = new(MockDialog)
	suite.directoryPicker = filePickerPlugin.filePicker(suite.mockDirectoryDialog, true)
	suite.filePicker = filePickerPlugin.filePicker(suite.mockFileDialog, false)

}

func (m MockDialog) File(title string, filter string, directory bool) (string, bool, error) {
	args := m.Called(title, filter, directory)
	return args.String(0), args.Bool(1), args.Error(2)
}

func (suite *FilePickerPluginTestSuite) TestFilePickerShouldOnlyAcceptArgsInJSONFormat() {
	directoryPicker := suite.directoryPicker
	_, err := directoryPicker("im not a json value")
	assert.Equal(suite.T(), "arguments must be encoded in JSON format", err.Error())

	filePicker := suite.filePicker
	_, err = filePicker("im not a json value")
	assert.Equal(suite.T(), "arguments must be encoded in JSON format", err.Error())
}

func (suite *FilePickerPluginTestSuite) TestFilePickerShouldRequireATitleArgument() {
	directoryPicker := suite.directoryPicker
	_, err := directoryPicker(map[interface{}]interface{}{})
	assert.Equal(suite.T(), "arguments requires a title parameter with type string", err.Error())

	filePicker := suite.filePicker
	_, err = filePicker(map[interface{}]interface{}{})
	assert.Equal(suite.T(), "arguments requires a title parameter with type string", err.Error())
}

func (suite *FilePickerPluginTestSuite) TestFilePickerShouldGetTheChosenDirectory() {
	directoryPicker := suite.directoryPicker
	mockDirectoryDialog := suite.mockDirectoryDialog
	const title = "Open a Directory"
	const chosenDirectory = "/home/flutter-user"
	mockDirectoryDialog.On("File", title, "*", true).Return(chosenDirectory, true, nil)

	actualChosenDirectory, err := directoryPicker(map[interface{}]interface{}{
		"title": title,
	})
	assert.Equal(suite.T(), chosenDirectory, actualChosenDirectory)
	assert.Nil(suite.T(), err)
}

func (suite *FilePickerPluginTestSuite) TestFilePickerShouldGetTheChosenFile() {
	filePicker := suite.filePicker
	mockFileDialog := suite.mockFileDialog
	const title = "Open a File"
	const chosenFile = "/home/flutter-user/go/src/file_picker/main.go"
	mockFileDialog.On("File", title, "*", false).Return(chosenFile, true, nil)

	actualChosenFile, err := filePicker(map[interface{}]interface{}{
		"title": title,
	})
	assert.Equal(suite.T(), chosenFile, actualChosenFile)
	assert.Nil(suite.T(), err)
}

func (suite *FilePickerPluginTestSuite) TestFilePickerShouldHandleDialogPickerFailure() {
	directoryPicker := suite.directoryPicker
	mockDirectoryDialog := suite.mockDirectoryDialog
	const title = "Open a Directory"
	mockDirectoryDialog.On("File", title, "*", true).Return("", false, errors.New(""))

	_, err := directoryPicker(map[interface{}]interface{}{
		"title": title,
	})
	assert.Equal(suite.T(), "failed to open dialog picker: ", err.Error())

	filePicker := suite.filePicker
	mockFileDialog := suite.mockFileDialog
	mockFileDialog.On("File", title, "*", false).Return("", false, errors.New(""))

	_, err = filePicker(map[interface{}]interface{}{
		"title": title,
	})
	assert.Equal(suite.T(), "failed to open dialog picker: ", err.Error())
}

func TestFilePickerPluginTestSuite(t *testing.T) {
	suite.Run(t, new(FilePickerPluginTestSuite))
}
