/* walter: a deployment pipeline template
* Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
* (see CONTRIBUTORS.md)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
package services

import "container/list"

type LocalClient struct{}

func (self *LocalClient) RegisterResult(result Result) error {
	return nil
}

func (self *LocalClient) GetCommits(update Update) (*list.List, error) {
	return list.New(), nil
}

func (self *LocalClient) GetUpdateFilePath() string {
	return ""
}
