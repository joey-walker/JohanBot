package common

import log "github.com/Sirupsen/logrus"

// Our separate api channels
const (
	DISCORD string = "DISCORD"
	SLACK   string = "SLACK"
	COMMON  string = "COMMON"
	TEST    string = "TEST"
)

func init() {
	// Create our own custom format so the time stamp is more readable
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	//customFormatter.ForceColors = true
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(customFormatter)

}

// Log information to the console or wherever stdout is directed.  This should be used at the
// main package level or the common package.
//
// Example:
//	LogInfo("/common/utilities.go", "Removing spaces in url")
func LogInfo(fileSource string, info string) {
	LogInfoAPI(COMMON, fileSource, info)
}

// Log information related to any of the api streams such as Discord or Slack.  Use within their separate packages.
//
// Example:
//	LogApiInfo(common.SLACK, "apiSlack.go", "starting connection to slack...")
func LogInfoAPI(api string, fileSource string, info string) {
	log.WithFields(log.Fields{
		"api":         api,
		"info_Source": fileSource,
	}).Info(info)
}

// Log debug information of any type.
//
// Example:
//	LogDebugAPI("apiSlack.go", "The Bot's id", Bot.ID)
func LogDebug(fileSource string, info string, object interface{}) {
	LogDebugAPI(COMMON, fileSource, info, object)
}

// Log debug information of any type.
//
// Example:
//	LogDebugAPI("common.SLACK", "apiSlack.go", "The Bot's id", Bot.ID)
func LogDebugAPI(api string, fileSource string, info string, object interface{}) {
	log.WithFields(log.Fields{
		"api":             api,
		"info_Source":     fileSource,
		"additional_info": info,
	}).Debug(object)
}

// Anything unexpected, but that we easily recover from such as trying a different method to parse a url.
//
// Example:
//	LogWarning("utilities.go", "failed to parse url")
func LogWarning(fileSource string, warning string) {
	LogWarningAPI(COMMON, fileSource, warning)
}

// Something strange that happened in one of the streams, but we were able to provide a result anyways.
//
// Example:
//	LogWarningAPI(common.DISCORD,"main.go", "failed to parse url")
func LogWarningAPI(api string, fileSource string, warning string) {
	log.WithFields(log.Fields{
		"api":         api,
		"info_Source": fileSource,
	}).Warn(warning)
}

// This should log when we have failed to perform an action in a shared method.
// Such as stripping out whitespace in a url and replacing with %20s
//
// Example:
//	LogError("utilities.go", err)
func LogError(fileSource string, err error) {
	LogErrorAPI(COMMON, fileSource, err)
}

//  This is for when we attempted an action, but failed due to connectivity issues or failure to parse.
//  For example, failing to send message to the slack channel.
//
// Example:
//	LogErrorApi(common.SLACK, "main.go", err)
func LogErrorAPI(api string, fileSource string, err error) {
	log.WithError(err).WithFields(log.Fields{
		"api":          api,
		"error_source": fileSource,
	}).Error()
}

// Anything in the shared or main functions that cause us to shut down unexpectedly.
//
// Example:
//	LogErrorAPIFatal("common.go", err)
func LogErrorFatal(fileSource string, err error) {
	LogErrorAPIFatal(COMMON, fileSource, err)
}

// Anything that causes us to shut down unexpectedly that came from one of the APIs such as Slack.
//
// Example:
//	LogErrorApiFatal(common.DISCORD, "main.go",err)
func LogErrorAPIFatal(api string, fileSource string, err error) {
	log.WithError(err).WithFields(log.Fields{
		"api":          api,
		"error_source": fileSource,
	}).Fatal()
}
