.PHONY: init
init:
ifeq ($(shell uname -s),Darwin)
	@grep -r -l go-mask * | xargs sed -i "" "s/go-lib-template/$$(basename `git rev-parse --show-toplevel`)/"
else
	@grep -r -l go-mask * | xargs sed -i "s/go-lib-template/$$(basename `git rev-parse --show-toplevel`)/"
endif
