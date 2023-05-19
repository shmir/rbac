package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupRBACModule() *RBACModule {
	rbacModule := &RBACModule{
		Users: []User{
			{
				ID:            1,
				Name:          "John",
				Organizations: []int{1, 2},
			},
		},
		Organizations: []Organization{
			{
				ID:   1,
				Name: "Org 1",
				Users: []int{
					1,
				},
			},
		},
		Projects: []Project{
			{
				ID:           1,
				Name:         "Project 1",
				Organization: 1,
			},
		},
	}

	return rbacModule
}

func teardownRBACModule(rbacModule *RBACModule) {
	// Clean up any resources if needed
}

func TestCreateOrganization(t *testing.T) {
	// Creating a sample RBAC module instance
	rbacModule := &RBACModule{}

	// Creating an organization
	organization := rbacModule.CreateOrganization("My Organization", nil, nil)

	// Checking if the organization is created correctly
	assert.Equal(t, "My Organization", organization.Name)
	assert.Equal(t, 0, organization.ID)
	assert.Equal(t, 0, len(organization.Users))
	assert.Equal(t, 0, len(organization.Projects))

	// Checking if the organization is in the RBAC module
	assert.Equal(t, 1, len(rbacModule.Organizations))

}

func TestAddUserToOrganization(t *testing.T) {
	// Creating a sample RBAC module instance
	rbacModule := &RBACModule{}

	// Creating an organization
	organization := rbacModule.CreateOrganization("My Organization", nil, nil)
	user := rbacModule.CreateUser("John Doe", nil)

	organization.AddUser(user)

	// Check if the user is added to the organization's users.
	assert.Equal(t, 1, len(organization.Users))
	assert.Equal(t, 1, len(user.Organizations))
	assert.Equal(t, organization.ID, user.Organizations[0])
	assert.Equal(t, user.ID, organization.Users[0])

}

func TestAddProjectToOrganization(t *testing.T) {
	// Creating a sample RBAC module instance
	rbacModule := &RBACModule{}

	// Creating an organization
	organization := rbacModule.CreateOrganization("My Organization", nil, nil)
	project := rbacModule.CreateProject("My Project", nil)

	organization.AddProject(project)

	// Check if the project is added to the organization's projects.
	assert.Equal(t, 1, len(organization.Projects))
	assert.Equal(t, organization.ID, project.Organization)
	assert.Equal(t, project.ID, organization.Projects[0])

}

func TestCanUserAccessProjectWithIDs(t *testing.T) {
	rbacModule := setupRBACModule()
	defer teardownRBACModule(rbacModule)

	// Test user ID and project ID
	userID := 1
	projectID := 1

	// Verify user can access the project
	assert.True(t, rbacModule.CanUserAccessProject(userID, projectID))
}

func TestCanUserAccessProjectWithObjects(t *testing.T) {
	rbacModule := setupRBACModule()
	defer teardownRBACModule(rbacModule)

	// Test user object and project object
	user := User{
		ID:            1,
		Name:          "John",
		Organizations: []int{1, 2},
	}
	project := Project{
		ID:           1,
		Name:         "Project 1",
		Organization: 1,
	}

	// Verify user can access the project
	assert.True(t, rbacModule.CanUserAccessProject(user, project))
}
