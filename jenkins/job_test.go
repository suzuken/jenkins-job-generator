package jenkins

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestEchoJob(t *testing.T) {
	job := `<?xml version="1.0" encoding="UTF-8"?>
<project>
  <actions></actions>
  <description>hello</description>
  <keepDependencies>false</keepDependencies>
  <properties></properties>
  <scm class="hudson.scm.NUllSCM"></scm>
  <canRoam>true</canRoam>
  <disabled>false</disabled>
  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
  <triggers></triggers>
  <concurrentBuild>false</concurrentBuild>
  <builders>
    <hudson.tasks.Shell>
      <command>echo &#34;Hello World!&#34;</command>
    </hudson.tasks.Shell>
  </builders>
  <publishers></publishers>
  <buildWrappers></buildWrappers>
</project>`
	j := NewProject("hello", "echo \"Hello World!\"")
	var buf bytes.Buffer
	buf.Write([]byte(xml.Header))
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")
	if err := enc.Encode(j); err != nil {
		t.Fatalf("encoding job failed: err %s", err)
	}
	if buf.String() != job {
		t.Fatalf("job is not equals, \nexpected\n%s\nactual\n%s\n", job, buf.String())
	}
}
