/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

/////////////////////////////////////////////// 读取接口 /////////////////////////////////////////////////

//获取某条微博的评论列表
func (api *APIClient) GET_comments_show(params map[string]interface{}, result *Comments) error {
	return api.get.call("comments/show", params, result)
}

//我发出的评论列表
func (api *APIClient) GET_comments_by_me(params map[string]interface{}, result *Comments) error {
	return api.get.call("comments/by_me", params, result)
}

//我收到的评论列表
func (api *APIClient) GET_comments_to_me(params map[string]interface{}, result *Comments) error {
	return api.get.call("comments/to_me", params, result)
}

//获取用户发送及收到的评论列表
func (api *APIClient) GET_comments_timeline(params map[string]interface{}, result *Comments) error {
	return api.get.call("comments/timeline", params, result)
}

//获取@到我的评论
func (api *APIClient) GET_comments_mentions(params map[string]interface{}, result *Comments) error {
	return api.get.call("comments/mentions", params, result)
}

//批量获取评论内容
func (api *APIClient) GET_comments_show_batch(params map[string]interface{}, result *Comments) error {
	return api.get.call("comments/show_batch", params, result)
}

/////////////////////////////////////////////// 写入接口 /////////////////////////////////////////////////

//评论一条微博
func (api *APIClient) POST_comments_create(params map[string]interface{}, result *Comment) error {
	return api.post.call("comments/create", params, result)
}

//删除一条评论
func (api *APIClient) POST_comments_destroy(params map[string]interface{}, result *Comment) error {
	return api.post.call("comments/destroy", params, result)
}

//批量删除评论
func (api *APIClient) POST_comments_destroy_batch(params map[string]interface{}, result *[]Comment) error {
	return api.post.call("comments/destroy_batch", params, result)
}

//回复一条评论
func (api *APIClient) POST_comments_reply(params map[string]interface{}, result *Comment) error {
	return api.post.call("comments/reply", params, result)
}
