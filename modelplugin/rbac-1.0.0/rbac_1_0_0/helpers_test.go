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
	assert "github.com/stretchr/testify/require"
	"testing"
)

const (
	role1ID = "role-1"
	role2ID = "role-2"
	role3ID = "role-3"

	group1ID = "group-1"
	group2ID = "group-2"
	group3ID = "group-3"
)

// 3 roles referred to in 3 groups
//         group-1    group-2     group-3
// role-1     ✔
// role-2     ✔          ✔
// role-3                ✔           ✔

func setupRbacCache() *Rbac_Rbac {
	cache := NewRbac()

	role1ID := role1ID
	role1Desc := fmt.Sprintf("Role %s", role1ID)
	role1 := Rbac_Rbac_Role{
		Description: &role1Desc,
		Permission: &Rbac_Rbac_Role_Permission{
			Noun: []string{
				"/a/b/c",
				"/d/e/f",
			},
			Operation: RbacIdentities_PERMISSION_READ,
			Type:      RbacIdentities_NOUNTYPE_CONFIG,
		},
		Roleid: &role1ID,
	}
	cache.Role[role1ID] = &role1

	role2ID := role2ID
	role2Desc := fmt.Sprintf("Role %s", role2ID)
	role2 := Rbac_Rbac_Role{
		Description: &role2Desc,
		Permission: &Rbac_Rbac_Role_Permission{
			Noun: []string{
				"/a/b/c",
				"/d/e/f",
				"/g/h/i",
			},
			Operation: RbacIdentities_PERMISSION_ALL,
			Type:      RbacIdentities_NOUNTYPE_CONFIG,
		},
		Roleid: &role2ID,
	}
	cache.Role[role2ID] = &role2

	role3ID := role3ID
	role3Desc := fmt.Sprintf("Role %s", role3ID)
	role3 := Rbac_Rbac_Role{
		Description: &role3Desc,
		Permission: &Rbac_Rbac_Role_Permission{
			Noun: []string{
				"/a/*/c",
				"/x/y/z",
			},
			Operation: RbacIdentities_PERMISSION_CREATE,
			Type:      RbacIdentities_NOUNTYPE_CONFIG,
		},
		Roleid: &role3ID,
	}
	cache.Role[role3ID] = &role3

	group1ID := group1ID
	group1Desc := fmt.Sprintf("Group %s", group1ID)
	group1 := Rbac_Rbac_Group{
		Description: &group1Desc,
		Groupid:     &group1ID,
		Role:        make(map[string]*Rbac_Rbac_Group_Role),
	}
	g1R1 := fmt.Sprintf("%s in %s", role1ID, group1ID)
	group1.Role[role1ID] = &Rbac_Rbac_Group_Role{
		Description: &g1R1,
		Roleid:      &role1ID,
	}
	g1R2 := fmt.Sprintf("%s in %s", role2ID, group1ID)
	group1.Role[role2ID] = &Rbac_Rbac_Group_Role{
		Description: &g1R2,
		Roleid:      &role2ID,
	}
	cache.Group[group1ID] = &group1

	group2ID := group2ID
	group2Desc := fmt.Sprintf("Group %s", group2ID)
	group2 := Rbac_Rbac_Group{
		Description: &group2Desc,
		Groupid:     &group2ID,
		Role:        make(map[string]*Rbac_Rbac_Group_Role),
	}
	g2R2 := fmt.Sprintf("%s in %s", role2ID, group2ID)
	group2.Role[role2ID] = &Rbac_Rbac_Group_Role{
		Description: &g2R2,
		Roleid:      &role2ID,
	}
	g2R3 := fmt.Sprintf("%s in %s", role3ID, group2ID)
	group2.Role[role3ID] = &Rbac_Rbac_Group_Role{
		Description: &g2R3,
		Roleid:      &role3ID,
	}
	cache.Group[group2ID] = &group2

	group3ID := group3ID
	group3Desc := fmt.Sprintf("Group %s", group3ID)
	group3 := Rbac_Rbac_Group{
		Description: &group3Desc,
		Groupid:     &group3ID,
		Role:        make(map[string]*Rbac_Rbac_Group_Role),
	}
	g3R3 := fmt.Sprintf("%s in %s", role3ID, group3ID)
	group3.Role[role3ID] = &Rbac_Rbac_Group_Role{
		Description: &g3R3,
		Roleid:      &role3ID,
	}
	cache.Group[group3ID] = &group3
	return cache
}


func Test_GetRolesByGroupIDs(t *testing.T) {
	cache := setupRbacCache()

	roles, err := cache.GetRolesByGroupIDs([]string{group1ID, group3ID}, RbacIdentities_PERMISSION_READ)
	assert.NoError(t, err, "unexpected error")
	assert.True(t, roles != nil)

	assert.Equal(t, 3, len(roles))

	for _, r := range roles {
		switch *r.Roleid {
		case role1ID:
			assert.Equal(t, "Role role-1", *r.Description)
			assert.EqualValues(t, []string{"/a/b/c", "/d/e/f"}, r.Permission.Noun)
		case role2ID:
			assert.Equal(t, "Role role-2", *r.Description)
			assert.EqualValues(t, []string{"/a/b/c", "/d/e/f", "/g/h/i"}, r.Permission.Noun)
		case role3ID:
			assert.Equal(t, "Role role-3", *r.Description)
			assert.EqualValues(t, []string{"/a/*/c", "/x/y/z"}, r.Permission.Noun)
		default:
			t.Errorf("unexpected role %s", *r.Roleid)
		}
	}

	consolidated := cache.ConsolidateNouns(roles)
	assert.Equal(t, 5, len(consolidated))
	assert.EqualValues(t, []string{"/a/*/c", "/a/b/c", "/d/e/f", "/g/h/i", "/x/y/z"}, consolidated)
}

func Test_ConfigPermissionIncludes(t *testing.T) {
	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_READ, RbacIdentities_PERMISSION_ALL))
	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_CREATE, RbacIdentities_PERMISSION_ALL))
	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_ALL, RbacIdentities_PERMISSION_ALL))

	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_READ, RbacIdentities_PERMISSION_CREATE))
	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_CREATE, RbacIdentities_PERMISSION_CREATE))
	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_ALL, RbacIdentities_PERMISSION_CREATE))

	assert.True(t, ConfigPermissionIncludes(RbacIdentities_PERMISSION_READ, RbacIdentities_PERMISSION_READ))
	assert.True(t, !ConfigPermissionIncludes(RbacIdentities_PERMISSION_CREATE, RbacIdentities_PERMISSION_READ))
	assert.True(t, !ConfigPermissionIncludes(RbacIdentities_PERMISSION_ALL, RbacIdentities_PERMISSION_READ))

	assert.True(t, !ConfigPermissionIncludes(RbacIdentities_PERMISSION_READ, RbacIdentities_PERMISSION_UNSET))
	assert.True(t, !ConfigPermissionIncludes(RbacIdentities_PERMISSION_CREATE, RbacIdentities_PERMISSION_UNSET))
	assert.True(t, !ConfigPermissionIncludes(RbacIdentities_PERMISSION_ALL, RbacIdentities_PERMISSION_UNSET))

}
