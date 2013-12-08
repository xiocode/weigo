/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

func (api *APIClient) GET_account_privacy(params map[string]interface{}, result *Config) error {
	return api.get.call("account/get_privacy", params, result)
}

func (api *APIClient) GET_account_rate_limit_status(params map[string]interface{}, result *LimitStatus) error {
	return api.get.call("account/rate_limit_status", params, result)
}

func (api *APIClient) GET_account_get_uid(params map[string]interface{}, uid *UserID) error {
	return api.get.call("account/get_uid", params, uid)
}

func (api *APIClient) GET_account_get_email(params map[string]interface{}, email *Email) error {
	return api.get.call("account/profile/email", params, email)
}
