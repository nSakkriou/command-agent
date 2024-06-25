package endpoint

// Mandatory endpoint
const (
	AVAILABLE_ENDPOINT = "/available"
	HEALTH_ENDPOINT    = "/health"
	UI_ENDPOINT        = "/ui"
)

var MANDATORY_ENDPOINT = []string{
	AVAILABLE_ENDPOINT, HEALTH_ENDPOINT, UI_ENDPOINT,
}
