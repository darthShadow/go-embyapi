/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationAccessSchedule struct {
	DayOfWeek *ConfigurationDynamicDayOfWeek `json:"DayOfWeek,omitempty"`
	StartHour float64                        `json:"StartHour,omitempty"`
	EndHour   float64                        `json:"EndHour,omitempty"`
}
