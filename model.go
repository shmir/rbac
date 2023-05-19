package main

type Organization struct {
	ID       int
	Name     string
	Users    []User
	Projects []Project
}

type User struct {
	ID            int
	Name          string
	Organizations []Organization
}

type Project struct {
	ID           int
	Name         string
	Organization Organization
	Environments []Environment
}

type Environment struct {
	ID      int
	Name    string
	Project Project
}

type RBACModule struct {
	Users         []User
	Organizations []Organization
	Projects      []Project
	Environments  []Environment
}

func (m *RBACModule) CreateOrganization(name string, user *User, project *Project) *Organization {
	organization := &Organization{
		ID:       len(m.Organizations),
		Name:     name,
		Users:    []User{},
		Projects: []Project{},
	}

	organization.AddUser(user)
	organization.AddProject(project)

	m.Organizations = append(m.Organizations, *organization)
	return organization
}

func (o *Organization) AddUser(user *User) {
	if user != nil {
		o.Users = append(o.Users, *user)
		user.Organizations = append(user.Organizations, *o)
	}
}

func (o *Organization) AddProject(project *Project) {
	if project != nil {
		o.Projects = append(o.Projects, *project)
		project.Organization = *o
	}
}

func (m *RBACModule) CreateUser(name string, organization *Organization) *User {
	user := &User{
		ID:   len(m.Users),
		Name: name,
	}

	if organization != nil {
		user.Organizations = append(user.Organizations, *organization)
		organization.Users = append(organization.Users, *user)
	}

	m.Users = append(m.Users, *user)

	return user
}

func (m *RBACModule) CreateProject(name string, organization *Organization, environment *Environment) *Project {
	project := &Project{
		ID:           len(m.Projects),
		Name:         name,
		Organization: Organization{},
		Environments: []Environment{},
	}

	if organization != nil {
		project.Organization = *organization
		organization.Projects = append(organization.Projects, *project)
	}

	if environment != nil {
		project.Environments = append(project.Environments, *environment)
		environment.Project = *project
	}

	m.Projects = append(m.Projects, *project)

	return project
}

func (m *RBACModule) CreateEnvironment(name string, project *Project) *Environment {
	environment := &Environment{
		ID:      len(m.Environments),
		Name:    name,
		Project: Project{},
	}

	if project != nil {
		environment.Project = *project
		project.Environments = append(project.Environments, *environment)
	}

	m.Environments = append(m.Environments, *environment)

	return environment
}
