# jenkins-job-generator

Experimental Jenkins job generator written in Go.

## Installation

	go get github.com/suzuken/jenkins-job-generator

## Usage

prepare shell file.

```sh
#!/bin/bash
echo "Hello world!";
```

run,

```
$ jenkins-job-generator -file examples/echo.sh -project hello
```

and then job file (config.xml) is generated.

```xml
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
      <command>#!/bin/bash&#xA;echo &#34;Hello world!&#34;;&#xA;</command>
    </hudson.tasks.Shell>
  </builders>
  <publishers></publishers>
  <buildWrappers></buildWrappers>
</project>
```

## LICENSE

MIT

## Author

Kenta Suzuki (a.k.a. suzuken)
