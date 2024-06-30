module stack_min

go 1.22.4

replace stack => /home/nkane/dev/cracking-the-code-interview/chapter-3-stack-and-queues/stack

require (
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8
	gotest.tools v2.2.0+incompatible
	stack v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)
