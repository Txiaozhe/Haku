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

package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func IsValidObjectHex(id string) bool {
	return bson.IsObjectIdHex(id)
}

func IsValidObjectId(id bson.ObjectId) bool {
	return IsValidObjectHex(id.Hex())
}

func GetById(collection *mgo.Collection, id string, i interface{})  {
	collection.FindId(bson.ObjectIdHex(id)).One(i)
}

// 根据条件查找单条记录
func GetUniqueOne(collection *mgo.Collection, q interface{}, doc interface{}) error {
	return collection.Find(q).All(doc)
}

// 根据条件模糊查询多条记录
func GetByOneCondition(collection *mgo.Collection, q interface{}, doc interface{}) error {
	return collection.Find(q).All(doc)
}

// 插入记录
func Insert(collection *mgo.Collection, doc interface{}) error {
	return collection.Insert(doc)
}

// 修改记录
func UpdateByQueryField(collection *mgo.Collection, q interface{}, field string, value interface{}) error {
	_, err := collection.UpdateAll(q, bson.M{"$set": bson.M{field: value}})

	return err
}

// 通用更新
func Update(collection *mgo.Collection, query interface{}, i interface{}) error {
	return collection.Update(query, i)
}

// 删除记录
func Delete(collection *mgo.Collection, query interface{}) error {
	return collection.Remove(query)
}
