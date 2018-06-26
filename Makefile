dev:
	@echo $$AWS_ACCESS_KEY_ID
	jq '.profile |= "uneet-dev" |.stages.production |= (.domain = "e.dev.unee-t.com" | .zone = "dev.unee-t.com")| .actions[0].emails |= ["kai.hendry+edev@unee-t.com"]' up.json.in > up.json
	up deploy production

demo:
	@echo $$AWS_ACCESS_KEY_ID
	jq '.profile |= "uneet-demo" |.stages.production |= (.domain = "e.demo.unee-t.com" | .zone = "demo.unee-t.com") | .actions[0].emails |= ["kai.hendry+edemo@unee-t.com"]' up.json.in > up.json
	up deploy production

prod:
	@echo $$AWS_ACCESS_KEY_ID
	jq '.profile |= "uneet-prod" |.stages.production |= (.domain = "e.unee-t.com" | .zone = "unee-t.com")| .actions[0].emails |= ["kai.hendry+eprod@unee-t.com"]' up.json.in > up.json
	up deploy production

testdev:
	curl https://e.dev.unee-t.com/version

testdemo:
	curl https://e.demo.unee-t.com/version

testprod:
	curl https://e.unee-t.com/version

.PHONY: dev demo prod
