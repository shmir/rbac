package model

import "fmt"

type Organization struct {
	ID       int
	Name     string
	Users    []int
	Projects []int
}

type User struct {
	ID            int
	Name          string
	Organizations []int
}

type Project struct {
	ID           int
	Name         string
	Organization int
}

type RBACModule struct {
	Users         []User
	Organizations []Organization
	Projects      []Project
}

func (r *RBACModule) CreateOrganization(name string, user *User, project *Project) *Organization {
	organization := &Organization{
		ID:   len(r.Organizations),
		Name: name,
	}

	organization.AddUser(user)
	organization.AddProject(project)

	r.Organizations = append(r.Organizations, *organization)
	return organization
}

func (o *Organization) AddUser(user *User) {
	if user != nil {
		o.Users = append(o.Users, user.ID)
		user.Organizations = append(user.Organizations, o.ID)
	}
}

func (o *Organization) AddProject(project *Project) {
	if project != nil {
		o.Projects = append(o.Projects, project.ID)
		project.Organization = o.ID
	}
}

func (r *RBACModule) CreateUser(name string, organization *Organization) *User {
	user := &User{
		ID:   len(r.Users),
		Name: name,
	}

	if organization != nil {
		user.Organizations = append(user.Organizations, organization.ID)
		organization.Users = append(organization.Users, user.ID)
	}

	r.Users = append(r.Users, *user)

	return user
}

func (r *RBACModule) CreateProject(name string, organization *Organization) *Project {
	project := &Project{
		ID:   len(r.Projects),
		Name: name,
	}

	if organization != nil {
		project.Organization = organization.ID
		organization.Projects = append(organization.Projects, project.ID)
	}

	r.Projects = append(r.Projects, *project)

	return project
}

func (r *RBACModule) CanUserAccessProject(userParam interface{}, projectParam interface{}) bool {
	var user User
	var project Project

	// Determine the type of userParam and projectParam
	switch user_ := userParam.(type) {
	case int:
		foundUser, err := r.getUserByID(user_)
		if err != nil {
			// User not found, cannot access project
			return false
		}
		user = *foundUser
	case User:
		user = userParam.(User)
	default:
		// Invalid userParam type
		return false
	}

	switch project_ := projectParam.(type) {
	case int:
		foundProject, err := r.getProjectByID(project_)
		if err != nil {
			// Project not found, cannot access project
			return false
		}
		project = *foundProject
	case Project:
		project = projectParam.(Project)
	default:
		// Invalid projectParam type
		return false
	}

	// Check if the user's organization matches the project's organization
	for _, orgID := range user.Organizations {
		if orgID == project.Organization {
			return true
		}
	}

	return false
}

func (r *RBACModule) getUserByID(userID int) (*User, error) {
	for i, user := range r.Users {
		if user.ID == userID {
			return &r.Users[i], nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (r *RBACModule) getProjectByID(projectID int) (*Project, error) {
	for i, project := range r.Projects {
		if project.ID == projectID {
			return &r.Projects[i], nil
		}
	}

	return nil, fmt.Errorf("project not found")
}
