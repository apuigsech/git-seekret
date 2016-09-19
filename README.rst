|Build Status|

===========
git-seekret
===========

Git module to prevent from committing sensitive information into the repository.

Description
===========

``git-seekret`` inspect on commits and/or staged (and uncommited yet) files, to 
prevent adding sensitive information into the git repository. It can be easily
integrated with git hooks, forcing it to analise all staged files before they are
included into a commit.


Installing git-seekret
======================

It's important to have the following tools and libraries to make it work properly:

	* glide
	* pkg-config
	* golang >= 1.6
	* libgit >= 2.24


``git-seekret`` uses ``Glide`` to maintain its dependencies.

Follow the instructions for installing Glide: https://github.com/Masterminds/glide#install

::

	glide install
	go build

Usage
=====

::

	General Options:

	NAME:
	   git-seekret - prevent from committing sensitive information into git repository

	USAGE:
	   git-seekret [global options] command [command options] [arguments...]

	VERSION:
	   0.0.1

	AUTHOR(S):
	   Albert Puigsech Galicia <albert@puigsech.com>

	COMMANDS:
	     config   manage configuration seetings
	     rules    manage rules
	     check    inspect git repository
	     hook     manage git hooks
	     help, h  Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --global
	   --help, -h     show help
	   --version, -v  print the version


``--global``


Rules and Exceptions
====================

The definition of rules and exceptions for ``git-seekret`` are defined by the seekret go library. A proper documentatio for this library can be found here;

	https://github.com/apuigsech/seekret




Hands-On
========

The repository seekret-secrets is prepare to test ``git-seekret`, and can be used to perform the following hands-on examples:

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
