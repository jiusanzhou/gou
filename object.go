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

package gou

// UserBuilder is an factory to create a user from form.
type UserBuilder struct {
	Fields []FieldBuilder `json:"fields,omitempty" yaml:"fields"`
}

// FieldBuilder contains the field name and value type,
// and how to generate the value.
type FieldBuilder struct {
	Name     string    `json:"name,omitempty" yaml:"name"`
	Type     FieldType `json:"type,omitempty" yaml:"type"`         // type can auto generate from ValueFrom function
	Required bool      `json:"required,omitempty" yaml:"required"` // useful for web form builder
	Pattern  *string   `json:"pattern,omitempty" yaml:"pattern"`   // useful for web form builder
	Default  string    `json:"default,omitempty" yaml:"default"`   // default value set value from string
}

// FieldType is a type to define field type
// we can use it to generate the default value
// even more the value generator
// without nested
type FieldType string

// NewUserBuilder return a user builder
func NewUserBuilder(fields ...FieldBuilder) UserBuilder {
	// default fields
	return UserBuilder{
		[]FieldBuilder{},
	}
}
