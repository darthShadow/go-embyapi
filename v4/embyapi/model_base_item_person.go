/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type BaseItemPerson struct {
	Name            string      `json:"Name,omitempty"`
	Id              string      `json:"Id,omitempty"`
	Role            string      `json:"Role,omitempty"`
	Type_           *PersonType `json:"Type,omitempty"`
	PrimaryImageTag string      `json:"PrimaryImageTag,omitempty"`
}
