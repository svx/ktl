=======
Coding
=======

.. admonition:: About

    Info about the setup

Dependencies
============

- `dep <https://github.com/golang/dep/>`_

We use ``dep`` for dependency management and pining.

Getting Started
===============

Clone the repository:

.. code-block:: console

    cd $PATH_TO_YOUR_GOPATH/src/github.com
    git clone https://github.com/svx/ktl.git

Install dependencies with ``dep``:

.. code-block:: console

    dep ensure -vendor-only

Or run ``make setup`` form the /root of the repository.

Code Blueprint
==============

``ktl`` is based on `cobra <https://github.com/spf13/cobra/>`_.

We inject the version number, the build date and the git hash via ``ldflags`` during package build into version.go (cmd/version.go).

See the example below how we do that in the ``Makefile`` (here for local test installs).

.. code-block:: make

    .PHONY: install
    install: ## Creating test builds (binaries) for local testing
               go install -ldflags "$(LD_FLAGS) -X main.Version=${VERSION} \
                -X github.com/svx/ktl/cmd.BuildTime=$(BUILD_DATE)"
