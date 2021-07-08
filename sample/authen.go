/*
 * Copyright (c) 2020 wellwell.work, LLC by Zoe
 *
 * Licensed under the Apache License 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sample

// 先快速实现，不考虑任何扩展性
// 方式：
// - OTP
// - password
// 物料：
// - uid
// - email
// - login
// 动作：
// - login<register>
// - register
// 内存实现：

type User interface {
	UID() string      // UID
	Username() string // 可选可设置
}

type Authen interface {
	// 登陆: 输入一些信息进行登陆，返回用户
	Login()

	// 唯一认证方法：输入
	Auth()
}

// OTP one time password
// flow for login
type OTP struct {
	// 准备(预)登陆请求，发送一次性密码和链接
	// 登陆请求(magic link 或 username/password)
	//
	// prelogin -> send notify -> waiting
	// login(magic link or passsword)
}
