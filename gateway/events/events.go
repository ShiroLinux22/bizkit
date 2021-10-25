/*
	Struct(s) for event handling
    Copyright (C) 2021 Jack C <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package events

import (
	"github.com/chakernet/bizkit/common/handler"
	"github.com/chakernet/bizkit/common/util"
)

var (
	log = util.Logger{
		Name: "EventHandler",
	}
)

type Handler struct {
	handler.HandlerR
}
