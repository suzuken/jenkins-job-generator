package jenkins

import (
	"encoding/xml"
)

type Project struct {
	XMLName                          xml.Name      `xml:"project"`
	Actions                          string        `xml:"actions"`
	Description                      string        `xml:"description"`
	KeepDependencies                 bool          `xml:"keepDependencies"`
	Properties                       JobProperties `xml:"properties"`
	Scm                              Scm           `xml:"scm"`
	CanRoam                          bool          `xml:"canRoam"`
	Disabled                         bool          `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding bool          `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   bool          `xml:"blockBuildWhenUpstreamBuilding"`
	Triggers                         Triggers      `xml:"triggers"`
	ConcurrentBuild                  bool          `xml:"concurrentBuild"`
	Builders                         Builders      `xml:"builders"`
	Publishers                       Publishers    `xml:"publishers"`
	BuildWrappers                    BuildWrappers `xml:"buildWrappers"`
}

func DefaultProject() *Project {
	scm := Scm{Class: "hudson.scm.NUllSCM"}
	return &Project{
		CanRoam: true,
		Scm:     scm,
	}
}

func NewProject(desc, cmd string) *Project {
	b := Builders{cmd}
	p := DefaultProject()
	p.Builders = b
	p.Description = desc
	return p
}

type Scm struct {
	Class string `xml:"class,attr"`
}

type JobProperties struct {
}

type Triggers struct {
}

type Builders struct {
	Command string `xml:"hudson.tasks.Shell>command"`
}

type Publishers struct {
}

type BuildWrappers struct {
}
