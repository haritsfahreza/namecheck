workflow "Test and build on push" {
  on = "push"
  resolves = ["Golang Action"]
}

action "Golang Action" {
  uses = "cedrickring/golang-action@1.3.0"
  args = "go get github.com/mitchellh/gox && go get -t -v ./... && diff -u <(echo -n) <(gofmt -d .) && go vet $(go list ./... | grep -v /vendor/) && ./go-test.sh"
}
