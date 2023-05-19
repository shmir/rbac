package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	assert.Equal(t, organization.ID, user.Organizations[0].ID)
	assert.Equal(t, user.ID, organization.Users[0].ID)

}

func TestAddProjectToOrganization(t *testing.T) {
	// Creating a sample RBAC module instance
	rbacModule := &RBACModule{}

	// Creating an organization
	organization := rbacModule.CreateOrganization("My Organization", nil, nil)
	project := rbacModule.CreateProject("My Project", nil, nil)

	organization.AddProject(project)

	// Check if the project is added to the organization's projects.
	assert.Equal(t, 1, len(organization.Projects))
	assert.Equal(t, 1, len(project.Organization.Projects))
	assert.Equal(t, organization.ID, project.Organization.ID)
	assert.Equal(t, project.ID, organization.Projects[0].ID)

}
