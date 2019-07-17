package sns

import "encoding/json"

const (
	gcmKeyMessage         = "message"
	fcmAndroidKeyPriority = "priority"
	apnsKeyMessage        = "alert"
	apnsKeyTitle          = "title"
	apnsKeyBody           = "body"
	apnsKeySound          = "sound"
	apnsKeyCategory       = "category"
	apnsKeyBadge          = "badge"
	apnsKeyMutableContent = "mutable-content"
)

const fcmPriorityHigh = "high"

// make sns message for Google Cloud Messaging.
func composeMessageGCM(msg string, opt map[string]interface{}) (payload string, err error) {
	data := make(map[string]interface{})
	data[gcmKeyMessage] = msg
	for k, v := range opt {
		data[k] = v
	}

	message := make(map[string]interface{})
	message["data"] = data

	// set Android FCM priority, which is compatible to GCM
	message["android"] = map[string]string{fcmAndroidKeyPriority: fcmPriorityHigh}

	b, err := json.Marshal(message)
	return string(b), err
}

// make sns message for Apple Push Notification Service.
// https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html
func composeMessageAPNS(msg string, opt map[string]interface{}) (payload string, err error) {
	aps := make(map[string]interface{})

	if v, ok := opt[apnsKeyTitle]; ok {
		// The title and body keys provide the contents of the alert.
		aps[apnsKeyMessage] = map[string]interface{}{
			apnsKeyTitle: v,
			apnsKeyBody:  msg,
		}
	} else {
		aps[apnsKeyMessage] = msg
	}

	aps[apnsKeySound] = "default"
	if v, ok := opt[apnsKeySound]; ok {
		aps[apnsKeySound] = v
	}

	if v, ok := opt[apnsKeyCategory]; ok {
		aps[apnsKeyCategory] = v
	}

	if v, ok := opt[apnsKeyBadge]; ok {
		aps[apnsKeyBadge] = v
	}

	if v, ok := opt[apnsKeyMutableContent]; ok {
		aps[apnsKeyMutableContent] = v
	}

	message := make(map[string]interface{})
	message["aps"] = aps
	for k, v := range opt {
		switch k {
		case apnsKeySound:
			continue
		case apnsKeyCategory:
			continue
		case apnsKeyBadge:
			continue
		case apnsKeyMutableContent:
			continue
		default:
			message[k] = v
		}
	}

	b, err := json.Marshal(message)
	return string(b), err
}
