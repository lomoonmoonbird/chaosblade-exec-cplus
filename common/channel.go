/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
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

package common

import (
	"context"

	"github.com/lomoonmoonbird/chaosblade-spec-go/channel"
	"github.com/lomoonmoonbird/chaosblade-spec-go/spec"
)

type AsyncChannel struct {
	//localChannel channel.OsChannel
	localChannel channel.Channel
}

func NewAsyncChannel() spec.Channel {
	return &AsyncChannel{
		localChannel: channel.NewLocalChannel(),
	}
}

func (a *AsyncChannel) Run(ctx context.Context, script, args string) *spec.Response {
	go a.localChannel.Run(ctx, script, args)
	// less judgement
	return spec.ReturnSuccess("success")
}

func (a *AsyncChannel) GetScriptPath() string {
	return a.localChannel.GetScriptPath()
}
