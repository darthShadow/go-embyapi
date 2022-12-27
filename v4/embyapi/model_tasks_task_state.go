/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type TasksTaskState string

// List of Tasks.TaskState
const (
	IDLE_TasksTaskState       TasksTaskState = "Idle"
	CANCELLING_TasksTaskState TasksTaskState = "Cancelling"
	RUNNING_TasksTaskState    TasksTaskState = "Running"
)
