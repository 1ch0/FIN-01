DIFF_CMD   = @git diff --exit-code

tests = custom

.PHONY: all
all: $(tests)

custom: custom.go
	go run custom.go > custom1.cue
	$(DIFF_CMD) custom.cue custom1.cue




