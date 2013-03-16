/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

/////////////////////////////////////////////// 读取接口 /////////////////////////////////////////////////

//获取用户信息
func (api *APIClient) GET_users_show(params map[string]interface{}, result *User) error {
	return api.get.call("users/show", params, result)
}

//获取用户信息
func (api *APIClient) GET_users_domain_show(params map[string]interface{}, result *User) error {
	return api.get.call("users/domain_show", params, result)
}

//获取用户信息
func (api *APIClient) GET_users_counts(params map[string]interface{}, result *[]UserCounts) error {
	return api.get.call("users/counts", params, result)
}

//获取用户信息
func (api *APIClient) GET_users_show_rank(params map[string]interface{}, result *UserRank) error {
	return api.get.call("users/show_rank", params, result)
}
