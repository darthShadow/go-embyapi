/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type AttributesSimpleCondition string

// List of Attributes.SimpleCondition
const (
	IS_TRUE_AttributesSimpleCondition              AttributesSimpleCondition = "IsTrue"
	IS_FALSE_AttributesSimpleCondition             AttributesSimpleCondition = "IsFalse"
	IS_NULL_AttributesSimpleCondition              AttributesSimpleCondition = "IsNull"
	IS_NOT_NULL_OR_EMPTY_AttributesSimpleCondition AttributesSimpleCondition = "IsNotNullOrEmpty"
)
