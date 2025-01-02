.PHONY: run commit
.SILENT:
.ONESHELL:

run:
	cls
	cd 02
	go run .

commit:
	git add .
	git commit -m "chore: commit everything"
	git push origin main