#update:
#	hz new -idl ./../../idl/http/$(service).thrift -force

update:
	hz update -idl ./../../idl/http/$(service).thrift

user:
	make service=user update

interview:
	make service=interview update

question:
	make service=question update

storage:
	make service=storage update

all:
	make user
	make interview
	make question
	make storage