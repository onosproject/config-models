// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rbac_1_0_0

import (
	"fmt"
	"sort"
)

func NewRbac() *Rbac_Rbac {
	return &Rbac_Rbac{
		Group: make(map[string]*Rbac_Rbac_Group),
		Role:  make(map[string]*Rbac_Rbac_Role),
	}
}

// GetRolesByGroupIDs - get all Roles for a set of Groups from the RBAC cache
func (r *Rbac_Rbac) GetRolesByGroupIDs(groupIDs []string, permission E_RbacIdentities_PERMISSION) ([]*Rbac_Rbac_Role, error) {
	matchingRoles := make(map[string]*Rbac_Rbac_Role)
	for _, gID := range groupIDs {
		group, ok := r.Group[gID]
		if !ok {
			return nil, fmt.Errorf("group '%s' not found in RBAC cache", gID)
		}
		for groupRoleID := range group.Role {
			if _, alreadyHandled := matchingRoles[groupRoleID]; !alreadyHandled {
				role, ok := r.Role[groupRoleID]
				if !ok {
					return nil, fmt.Errorf("role '%s' not found in RBAC cache", groupRoleID)
				}
				if role.Permission.Type == RbacIdentities_NOUNTYPE_CONFIG &&
					ConfigPermissionIncludes(permission, role.Permission.Operation) {
					matchingRoles[groupRoleID] = role
				}
			}
		}
	}
	roleArray := make([]*Rbac_Rbac_Role, 0, len(matchingRoles))
	for _, role := range matchingRoles {
		roleArray = append(roleArray, role)
	}
	sort.Slice(roleArray, func(i, j int) bool {
		return *roleArray[i].Roleid < *roleArray[j].Roleid
	})
	return roleArray, nil
}

// ConsolidateNouns - a set of roles might have overlapping nouns
// Here we get unique entries and sort them alphabetically
func (r *Rbac_Rbac) ConsolidateNouns(roles []*Rbac_Rbac_Role) []string {
	consolidatedMap := make(map[string]interface{})
	for _, r := range roles {
		for _, n := range r.Permission.Noun {
			consolidatedMap[n] = struct{}{}
		}
	}
	consolidated := make([]string, 0, len(consolidatedMap))
	for c := range consolidatedMap {
		consolidated = append(consolidated, c)
	}
	sort.Strings(consolidated)
	return consolidated
}

// ConfigPermissionIncludes - compare the given 'perm' to what the `rolePerm` allows
// For config ALL and CREATE are the same thing, and also allows READ
// If role only allows READ and 'perm' is READ then it is included, otherwise not
func ConfigPermissionIncludes(perm E_RbacIdentities_PERMISSION, rolePerm E_RbacIdentities_PERMISSION) bool {
	switch rolePerm {
	case RbacIdentities_PERMISSION_ALL, RbacIdentities_PERMISSION_CREATE:
		return true
	case RbacIdentities_PERMISSION_READ:
		if perm == RbacIdentities_PERMISSION_READ {
			return true
		}
	}
	return false
}
