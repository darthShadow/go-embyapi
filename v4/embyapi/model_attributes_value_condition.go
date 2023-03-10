/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type AttributesValueCondition string

// List of Attributes.ValueCondition
const (
	IS_EQUAL_AttributesValueCondition            AttributesValueCondition = "IsEqual"
	IS_NOT_EQUAL_AttributesValueCondition        AttributesValueCondition = "IsNotEqual"
	IS_GREATER_AttributesValueCondition          AttributesValueCondition = "IsGreater"
	IS_GREATER_OR_EQUAL_AttributesValueCondition AttributesValueCondition = "IsGreaterOrEqual"
	IS_LESS_AttributesValueCondition             AttributesValueCondition = "IsLess"
	IS_LESS_OR_EQUAL_AttributesValueCondition    AttributesValueCondition = "IsLessOrEqual"
)
