/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UsersForgotPasswordAction string

// List of Users.ForgotPasswordAction
const (
	CONTACT_ADMIN_UsersForgotPasswordAction UsersForgotPasswordAction = "ContactAdmin"
	PIN_CODE_UsersForgotPasswordAction      UsersForgotPasswordAction = "PinCode"
)
