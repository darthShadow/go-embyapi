/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaProfileConditionType string

// List of Dlna.ProfileConditionType
const (
	EQUALS_DlnaProfileConditionType             DlnaProfileConditionType = "Equals"
	NOT_EQUALS_DlnaProfileConditionType         DlnaProfileConditionType = "NotEquals"
	LESS_THAN_EQUAL_DlnaProfileConditionType    DlnaProfileConditionType = "LessThanEqual"
	GREATER_THAN_EQUAL_DlnaProfileConditionType DlnaProfileConditionType = "GreaterThanEqual"
	EQUALS_ANY_DlnaProfileConditionType         DlnaProfileConditionType = "EqualsAny"
)
