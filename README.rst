|Build Status| |Documentation Status|

===========
git-seekret
===========

Git module to prevent commit secrets.

NOTE: It's not ready to be used.


Interactive session
===================

::
	$ git clone https://github.com/apuigsech/seekret-secrets

	$ cd seekret-secrets

	$ git seekret config --init
	Config:
		version = 1
		rulespath = /Users/apuigsech/Develop//.go/src/github.com/apuigsech/seekret/rules
		rulesenabled =
		exceptionsfile =

	$ git seekret rules
	List of rules:
		[ ] aws.secret_key
		[ ] aws.access_key
		[ ] certs.rsa
		[ ] certs.generic
		[ ] certs.pgp
		[ ] password.pass
		[ ] password.cred
		[ ] password.password
		[ ] password.pwd
		[ ] unix.passwd

	$ git seekret rules --enable password.password
	List of rules:
		[ ] aws.secret_key
		[ ] aws.access_key
		[ ] certs.generic
		[ ] certs.pgp
		[ ] certs.rsa
		[x] password.password
		[ ] password.pwd
		[ ] password.pass
		[ ] password.cred
		[ ] unix.passwd

	$ git seekret check -c 1   # Check on last commit.
	Found Secrets: 9
		secret_6:2
			- Metadata:
			  commit: 442d574a5e233d9cec7d245f7c85177cd1a827e4
			  uniq-id: e4ac21ceef17fff49d2f0d1fdd46f0abe7d0f62c
			- Rule:
			  password.password
			- Content:
			  password = 's3cr3t'
		secret_8:5
			- Metadata:
			  uniq-id: 373978394eb25268890ebee17966024300f3997b
			  commit: 442d574a5e233d9cec7d245f7c85177cd1a827e4
			- Rule:
			  password.password
			- Content:
			  password = 'thisISnotSECRET'

		... 

	$ git seekret check -s     # Check on staged files.
	Found Secrets: 0

	$ echo "password = 'this is super secret'" > new_file

	$ git add new_file

	$ git seekret check -s
	Found Secrets: 1
		new_file:1
			- Metadata:
			  status: test
			- Rule:
			  password.password
			- Content:
			  password = 'this is super secret'



.. |Build Status| image:: https://travis-ci.org/apuigsech/git-seekret.svg
   :target: https://travis-ci.org/apuigsech/seekret
   :width: 88px
   :height: 20px
