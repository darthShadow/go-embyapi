/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyWebGenericEditActionsPostbackAction struct {
	TargetEditorId             string `json:"TargetEditorId,omitempty"`
	PostbackCommandId          string `json:"PostbackCommandId,omitempty"`
	CommandParameterPropertyId string `json:"CommandParameterPropertyId,omitempty"`
}
