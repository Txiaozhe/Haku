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
 *     Initial: 2017/08/15     Tang Xiaoji
 */

package cockroach

import (
	"errors"
	"container/ring"
	"sync"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"Haku/orm"
)

const (
	poolMaxSize = 200
	dialect = "postgres"
)

var (
	ErrNoConnection = errors.New("CockroachDB Connection expired")
)

type Pool struct {
	lock		sync.Mutex
	pool		*ring.Ring
	size		int				// 实际长度
}

func NewPool(db string, size int) *Pool {
	var (
		err		error
		conn	*ring.Ring
	)

	if size > poolMaxSize {
		size = poolMaxSize
	}

	pool := &Pool{}

	pool.pool = ring.New(1)

	for i := 0; i < size; i++ {
		conn = ring.New(1)
		conn.Value, err = gorm.Open(dialect, db)

		if err != nil {
			continue
		}

		pool.pool.Link(conn)
	}

	pool.size = pool.pool.Len()
	if pool.size != size {
		log.Debug("New pool not enough!")
	}

	return pool
}

func (this *Pool) GetConnection() (orm.Connection, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.size == 0 {
		return nil, ErrNoConnection
	}

	this.size -= 1

	conn := this.pool.Unlink(1)
	return conn.Value.(orm.Connection), nil
}

func (this *Pool) ReleaseConnection(v orm.Connection) {
	conn := ring.New(1)
	conn.Value = v

	this.lock.Lock()
	defer this.lock.Unlock()

	this.size += 1
	this.pool.Prev().Link(conn)
}

func (this *Pool) Close() {
	f := func (v interface{}) {
		if v == nil {
			return
		}

		conn := v.(*gorm.DB)
		conn.Close()
	}

	this.lock.Lock()
	defer this.lock.Unlock()

	this.size = 0
	this.pool.Do(f)
	this.pool = nil			// 释放内存
}
