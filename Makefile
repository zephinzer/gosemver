tags:
	git tag dev0.0.0
	git tag dev0.0.1
	git tag dev0.0.10
	git tag dev0.1.1
	git tag dev0.10.1
	git tag dev1.1.1
	git tag dev10.1.1
	git tag dev11.1.1-alpha.1
	git tag dev11.1.1-alpha.2
	git tag dev11.1.1-alpha
	git tag dev11.1.1-beta.1
	git tag dev11.1.1-beta.2
	git tag dev11.1.1-beta
	git tag dev11.1.1-rc.1
	git tag dev11.1.1-rc.2
	git tag dev11.1.1-rc
	git tag dev11.1.1
tags.get:
	git tag --list
tags.clean:
	git tag --list | grep dev | xargs -I{} git tag -d {}
test:
	godev test