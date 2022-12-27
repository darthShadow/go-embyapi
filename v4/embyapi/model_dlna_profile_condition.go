/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaProfileCondition struct {
	Condition  *DlnaProfileConditionType  `json:"Condition,omitempty"`
	Property   *DlnaProfileConditionValue `json:"Property,omitempty"`
	Value      string                     `json:"Value,omitempty"`
	IsRequired bool                       `json:"IsRequired,omitempty"`
}
