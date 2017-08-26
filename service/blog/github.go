/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co,Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the 'Software'), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/24     Tang Xiaoji
 */

package blog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

type Label struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GitHubBlog struct {
	Id         int32   `json:"id"`
	Number     int32   `json:"number"`
	Title      string  `json:"title"`
	Labels     []Label `json:"labels"`
	Updated_at string  `json:"updated_at"`
}

type GitHubBlogDetail struct {
	Id         int32   `json:"id"`
	Number     int32   `json:"number"`
	Title      string  `json:"title"`
	Labels     []Label `json:"labels"`
	Body       string  `json:"body"`
	Updated_at string  `json:"updated_at"`
}

func (b *blogServiceProvider) GetLabelFromGitHub() ([]Label, error) {
	var (
		err    error
		labels []Label
	)

	url := "https://api.github.com/repos/Txiaozhe/docs/labels"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &labels)
	return labels, nil
}

func (b *blogServiceProvider) GetListFromGitHub() ([]GitHubBlog, error) {
	var (
		err   error
		blogs []GitHubBlog
	)
	url := "https://api.github.com/repos/Txiaozhe/docs/issues"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &blogs)
	return blogs, nil
}

func (b *blogServiceProvider) GetDetailFromGitHub(number string) (GitHubBlogDetail, error) {
	var (
		err   error
		blog  GitHubBlogDetail
	)
	url := "https://api.github.com/repos/Txiaozhe/docs/issues/" + number
	fmt.Println(url)

	resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}

	json.Unmarshal([]byte(body), &blog)
	return blog, err
}
