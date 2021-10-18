/*
	Struct(s) for event handling
    Copyright (C) 2021 jacany <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package events

import (
	"github.com/chakernet/ryuko/common/handler"
	"github.com/chakernet/ryuko/common/util"
)

var (
	log = util.Logger {
		Name: "EventHandler",
	}
)

type Handler struct {
	handler.EventHandlerR;
}
