package messageHandling

//Message defined as a constant

//TestMessage indicates the configuration file is missing
const TestMessage = 1000

//ConfigurationFileNotSpecified indicates the configuration file is missing
const ConfigurationFileNotSpecified = 1001

//NoMessageFound indicates that the message being sent for handling doesn't exist
const NoMessageFound = 1002

//UnableToLoadConfiguration indicates that the configuration couldn't be loaded
const UnableToLoadConfiguration = 1003

//UnableToStartListener indicates that the listener couldn't be started
const UnableToStartListener = 1004

//SeverityDebug represents debug messages
const SeverityDebug = 1

//SeverityInfo represents Info messages
const SeverityInfo = 2

//SeverityWarning represents warning messages
const SeverityWarning = 3

//SeverityError represents error messages
const SeverityError = 4

//SeverityFatal represents fatal messages
const SeverityFatal = 5

//MessageDetails gives details of the message being printed
type MessageDetails struct {
	code     int
	severity int
	message  string
}

var messageMap = map[int]MessageDetails{
	TestMessage: MessageDetails{
		code:     TestMessage,
		severity: SeverityInfo,
		message:  "This is a test message",
	},
	ConfigurationFileNotSpecified: MessageDetails{
		code:     ConfigurationFileNotSpecified,
		severity: SeverityFatal,
		message:  "Configuration file was not received in input",
	},
	NoMessageFound: MessageDetails{
		code:     NoMessageFound,
		severity: SeverityError,
		message:  "No message found for message id provided",
	},
	UnableToLoadConfiguration: MessageDetails{
		code:     UnableToLoadConfiguration,
		severity: SeverityFatal,
		message:  "Unable to load configuration",
	},
	UnableToStartListener: MessageDetails{
		code:     UnableToStartListener,
		severity: SeverityFatal,
		message:  "Unable to start listener",
	},
}

/*GetMessageDetails Handles messages by
- Throwing panic for fatal messages. Expectation for fatal messages is that it would cause system crash
- For rest, it prints in the log and returns the messageDetails
*/
func GetMessageDetails(messageCode int) MessageDetails {
	var messageDetails MessageDetails
	var ok bool
	if messageDetails, ok = messageMap[messageCode]; ok {
		return messageDetails
	}
	return messageMap[NoMessageFound]
}
