/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UsersForgotPasswordAction string

// List of Users.ForgotPasswordAction
const (
	CONTACT_ADMIN_UsersForgotPasswordAction       UsersForgotPasswordAction = "ContactAdmin"
	PIN_CODE_UsersForgotPasswordAction            UsersForgotPasswordAction = "PinCode"
	IN_NETWORK_REQUIRED_UsersForgotPasswordAction UsersForgotPasswordAction = "InNetworkRequired"
)
