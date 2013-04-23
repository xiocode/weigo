package weigo

//通过id获取mid
func (api *APIClient) get_statuses_querymid(params map[string]interface{}, result interface{}) error {
	return api.get.call("emotions", params, result)
}
