package messageHandling

import (
	"testing"
)

//TestInvalidMessageId tests the messages
func TestInvalidMessageId(t *testing.T) {
	//Send invalid message
	var msgDetails = GetMessageDetails(999)
	if msgDetails.code != NoMessageFound {
		t.Fail()
	}
}

//TestValidMessageId tests that the valid message is returned
func TestValidMessageId(t *testing.T) {
	//Send invalid message
	var msgDetails = GetMessageDetails(TestMessage)
	if msgDetails.code != TestMessage && msgDetails.message != "This is a test message" {
		t.Fail()
	}
}
