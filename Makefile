readme:
	git add README.md
	git commit -m README
	git push origin HEAD

makefile:
	git add Makefile
	git commit -m Makefile
	git push origin HEAD

push-all:
	git add .
	git commit -m "commit"
	git push origin HEAD