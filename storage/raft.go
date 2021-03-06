// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Ben Darnell

package storage

import "github.com/cockroachdb/cockroach/proto"

// raft is the interface exposed by a raft implementation.
type raft interface {
	// propose a command to raft. If accepted by the consensus protocol it will
	// eventually appear in the committed channel, but this is not guaranteed
	// so callers may need to retry.
	propose(proto.InternalRaftCommand)

	// committed returns a channel that yields commands as they are
	// committed. Note that this includes commands proposed by this node
	// and others.
	committed() <-chan proto.InternalRaftCommand
}

// noopRaft is a trivial implementation of the raft interface for testing.
type noopRaft chan proto.InternalRaftCommand

func newNoopRaft() noopRaft {
	return make(noopRaft, 10)
}

func (nr noopRaft) propose(req proto.InternalRaftCommand) {
	nr <- req
}

func (nr noopRaft) committed() <-chan proto.InternalRaftCommand {
	return nr
}
