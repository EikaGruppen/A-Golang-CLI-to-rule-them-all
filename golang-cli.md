# A Golang CLI

### to rule them all


#### Tormod Alf Try Tufteland

# Intro




#### Tormod Alf Try Tufteland

- Developer at Eika

- Main contributor of the Eika CLI
- Working with the new mobile bank for Eika
- Creator of wsdl2kotlin - written in Rust






#### Eika Gruppen

- Bank group consisting of 50+ small to medium size banks



[github.com/EikaGruppen](https://github.com/EikaGruppen)



# Agenda
  
- Demo 10 LoC CLI
- Why Go?


#### Create a 0-dep CLI for listing and cloning repos
- Go stdlib
- Project structure
- Testing
- UI libraries
- Distribution


#### Rule your stack
- Simplify auth
- Safer with OS keychain
- Fetch info from multiple sources
- Leverage working directory
- Consider SDK's
- Examples





Β«pΒ»Code and slides here: [github.com/EikaGruppen](https://github.com/EikaGruppen)




# Hello JavaZone


```go
package main

func main() {
	println("Hello JavaZone!")
}
```

# Hello CLI

<!-- - [hello-cli](./hello/cli.go) -->
```go
package main

import "flag"

func main() {
	name := flag.String("name", "anonymous", "your name")
	flag.Parse()

	println("Hello", *name)
}
```


 
 
 

# Go

Β«pΒ»π
- Easy to learn
- Easy to write


- Compiles to machine code
- Cross compilation included
- Run C code directly


- Networking, IO, concurrency etc. inluded


- Many DevSecOps tools are written in Go (Docker, k8s, Jaeger, Jenkins X, ...)



Β«pΒ»βοΈ
Β«pΒ»Perfection sacrificed for simplicity?




# Demo gh


#### Today's goal: Official Github CLI


# stdlib

#### net/http, ioutil, fmt

[Github Server](./repo/pkg/github/server/server.go)

[pkg/github/client.go](./repo/pkg/github/client.go)	  [pkg/github/client_test.go](./repo/pkg/github/client.go)

<!-- Show client/server, test -->

# Structure

#### repo CLI
```js
βββ main.go
βββ cmd
β   βββ list.go
β   βββ list_test.go
β   βββ clone.go
β   βββ clone_test.go
βββ pkg
β   βββ github
β       βββ client.go
β       βββ client_test.go
βββ go.mod
βββ go.sum
```


- Scalability
- Isolation
- Avstraction layers
- Easy to find



Β«pΒ»_π‘ Tip: code in `internal` folders cannot be imported by other projects_
Β«pΒ»_π‘ Tip: initalize go.mod and go.sum with_ go mod init example.com/project


# Cmd


[main.go](./repo/main.go)
```go
import (
	"repo/cmd"
	"os"
)

func main() {

	switch os.Args[1] {
	case "list":
		cmd.ListRepos()
	case "clone":
		cmd.CloneRepo()
	default:
		println("Unknown subcommand")
		os.Exit(1)
	}
}
```
# Cmd

[cmd/list.go](./repo/cmd/list.go)
```go 
import (
	"fmt"
	"repo/pkg/github"
)

func ListRepos() {
	repos, err := github.GetRepos()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("%s - %s\n\t%s\n\n", green(repo.Project), green(repo.Name), repo.Description)
	}

}

func green(text string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", text)
}
```

# Testing CLI's

- Test it like a normal Go app!



#### Extra useful in CLI's

- Mocking
  - Eg. the pkg/github package



- Send input, capture output



# Testing CLI's
### Mocking

#### Use interfaces
```go
type repoStorage interface {
	GetRepos() ([]Repo, error)
}
```

[pkg/github/client.go](./repo/pkg/github/client.go)
```diff
+ type Client struct {}

- func GetRepos() ([]Repo, error) {
+ func (s *Client) GetRepos() ([]Repo, error) {
```

# Testing CLI's
#### Mock impl
```go
var alfa = cmd.Repo{
		Name:        "alfa",
		Project:     "ap",
		Description: "ad",
}

var beta = cmd.Repo{
		Name:        "beta",
		Project:     "bp",
		Description: "bd",
}

type reposMock struct{}

func (m *reposMock) GetRepos() ([]cmd.Repo, error) {
	return []cmd.Repo{alfa, beta}, nil
}
```

# Testing CLI's

[cmd/list.go](./repo/cmd/list.go)
```diff
-func ListRepos() { ... }
-	repos, err := github.GetRepos()
+func ListRepos(rs repoStorage) { ... }
+	repos, err := rs.GetRepos()
```

#### Use it
[main.go]()
```go
cmd.ListRepos(&github.Client{})
```
[cmd/clone_test.go]()
```go
cmd.ListRepos(&reposMock{})
```

# Testing CLI's

#### Capturing output
```go
func ListRepos(storage repoStorage, w io.WriteCloser) { ... }
```

```diff
- fmt.Printf("%s %s", repo.Project)
+ fmt.Fprintf(w, "%s %s", repo.Name)
```

#### Use it
[main.go]()
```go
cmd.ListRepos(&gitlab.Client{}, os.Stdout)
```
[cmd/list_test.go](./repo/cmd/list_test.go)
```go
var buffer = ...

cmd.ListRepos(&reposMock{}, buffer)

captured := buffer.String()
// TODO asserts
```

<!-- - Golden files? -->
# Demo


#### repo list






#### Find full code at
[github.com/EikaGruppen](github.com/EikaGruppen)

# UI


#### Colors
- github.com/fatih/color


#### Fancy prompts
- github.com/manifoldco/promptui
- demo `repo clone`

- github.com/c-bata/go-prompt
- demo `bit`


#### Terminal visualizations
- github.com/gizak/termui


# Distribute

#### homebrew, rpm, chocolatey etc.

```
brew tap your-company/homebrew-formulas
brew install your-cli

brew upgrade your-cli
```

- Cross compile vs build from source?

# Done!



#### right?
<!-- haven't touch upon the tings that make your cli rule the stack yet -->







Β«pΒ»_Haven't we just re-implemented an existing CLI?_


# Auth

```go
type Client struct {
	Username string
	Password func() (password string, err error)

	TokenFromStore     func(name string) (token string)
	UpdateTokenInStore func(name string, token string) (err error)
}
```

- Store tokens _safely_ till expiry
- Make auth more convenient _and_ safer with OS keychain
- Local machine and/or build server


Β«pΒ»βοΈ
Β«pΒ»Avoid Basic Auth


- MFA / 2FA and CLI?

Β«pΒ»_π€ Tip:_ MFA _and_ 2FA _stands for multi-/two-factor-authorizaion_


# MFA

[github.com/EikaGruppen/go-oauth-cli-client](github.com/EikaGruppen/go-oauth-cli-client)

Β«pΒ»ββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββ
Β«pΒ»β   ββββββββββββββββββββββββββββββββββββββββββββββββββββββ       β
Β«pΒ»β   β ποΈ |  http://my-oidc-server/oauth/authenticate     β       β
Β«pΒ»β   ββββββββββββββββββββββββββββββββββββββββββββββββββββββ       β
Β«pΒ»β                                                                β
Β«pΒ»β                                                                β
Β«pΒ»β                   ... MFA steps ...                            β
Β«pΒ»β                                                                β
Β«pΒ»ββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββ

Β«pΒ»ββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββ
Β«pΒ»β   ββββββββββββββββββββββββββββββββββββββββββββββββββββββ       β
Β«pΒ»β   β ποΈ |  http://my-oidc-server/oauth/callback         β       β
Β«pΒ»β   ββββββββββββββββββββββββββββββββββββββββββββββββββββββ       β
Β«pΒ»β                                                                β
Β«pΒ»β               Logged in successfully!                          β
Β«pΒ»β                                                                β
Β«pΒ»β        You may now close this browser window...                β
Β«pΒ»β                                                                β
Β«pΒ»ββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββ

Β«pΒ»_π€ Tip: the offical azure-cli does this exact thing (only without PKCE)_



# OS keychain

#### macOS
[github.com/EikaGruppen/go-macos-keychain](github.com/EikaGruppen/go-macos-keychain)

Β«pΒ»ββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββ
Β«pΒ»β                                                                β
Β«pΒ»β     my-CLI wants to access my-item in OS keychain              β
Β«pΒ»β   =====================================================        β
Β«pΒ»β                                                                β
Β«pΒ»β                                                                β
Β«pΒ»β        | Always allow |      | Deny |       | Allow |          β
Β«pΒ»β                                                                β
Β«pΒ»ββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββββ



Β«pΒ»βοΈ
Β«pΒ»Avoid macOS "security" command for OS keychain


# SDK

[github.com/google/go-github](github.com/google/go-github)

```go 
import "github.com/google/go-github/v40/github"

client := github.NewClient(nil)

// list public repositories for org "github"
opt := &github.RepositoryListByOrgOptions{Type: "public"}
repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
```




Β«pΒ»_π€ Tip: SDK stands for_ Software Development Kit _, and is a client-code for a specific language and service_


# Examples

#### status


```
$ j javaz
[my-javazone-repo] $ ec status

Env			Version				Latest deployment
test        0.1.5 -> 1.0.0      2 seconds ago
qa			0.1.5				1 week ago
prod		0.1.5				1 week ago

Last build _FAILED_. For logs: ec build-logs
Latest version: 1.0.0

```
Β«pΒ»Integrations: build server, build system, container orchestrator, artifact manager, git


Β«pΒ»_π€ Tip: brew install autojump_

# Examples


#### Print progress

- Build:                DONE!
- Tests:                DONE!
- Scan:        [===>    ] 51%
- Analysis:    [=====>  ] 82%



# Examples
#### Follow logs


```go 
request := c.kubeClient.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)
logsStream, err := request.Stream(context.TODO())
if err != nil {
	return errors.New("error in opening stream")
}
defer logsStream.Close()
// read, and print logsStream...
```


```
$ j javaz
[my-javazone-repo] $ ec logs --env test --follow

12:09 [main] - INFO no.yourapp.RequestLogger - Got request!
12:09 [main] - INFO no.yourapp.ResponseLogger
  Headers: [a=2, b=4],
  Payload: {"isItWorking", "sure!"}
12:09 [main] - INFO no.yourapp.LogFilter

_following, hit CTRL-C to cancel_
```


# Q&A







                                                                           ********                               
                                                                        ****************                          
                                                                      ****** **** ********                        
                                                                      **** *********  ****                        
                                                                       ********************                       
                                                                          \\   //  ********                       
                                                                           \\////  ******                         
                                                                             \\\\////                             
                                                                              ||||//                              
                                                                              ||||                                
                                                          ,,,,,, ,,,, ,,,,,,,//||||,,,,,,,,, ,,,,,,,,             
                                                        ,;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;



Β«pΒ»ο u/tormodatt

# Thanks!


Β«pΒ»Share command ideas in the Slack-channel!








Β«pΒ»Code and slides here: [github.com/EikaGruppen](https://github.com/EikaGruppen)

























