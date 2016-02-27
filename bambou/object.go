// Copyright (c) 2015, Alcatel-Lucent Inc.
// All rights reserved.
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
// * Neither the name of bambou nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package bambou

import (
	"fmt"
)

// Rootable is the interface that must be implemented by the root object of the API.
// A Rootable also implements the Identifiable interface.
type Rootable interface {
	Identifiable

	APIKey() string
	SetAPIKey(string)
}

// RemoteObject represents an object than contains information common to all objects.
// exposed by the server.
// This struct must be embedded into all objects that are available
// throught the ReST api.
type RemoteObject struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
	identity   Identity
}

// Identity returns the object's Identity.
func (o *RemoteObject) Identity() Identity {

	return o.identity
}

// SetIdentity sets the Identity of the object.
func (o *RemoteObject) SetIdentity(identity Identity) {

	o.identity = identity
}

// Identifier returns the object's ID.
func (o *RemoteObject) Identifier() string {

	return o.ID
}

// SetIdentifier sets the ID of the obje.
func (o *RemoteObject) SetIdentifier(ID string) {

	o.ID = ID
}

// String returns the string representation of the object.
func (o *RemoteObject) String() string {

	return fmt.Sprintf("<ExposedObject %s:%s>", o.identity.RESTName, o.ID)
}
