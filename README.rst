============
Kube Control
============

.. image:: docs/_static/logo-small.png

About
=====

A wrapper around `kubectl`.

I am lazy, I do not want to type `kubectl get pods --all-namespaces`.

`ktl` (Kube Control) is a tiny wrapper written in Go which calls `kubectl` in the background.

For example `ktl pods` is calling `kubectl get pods --all-namespaces`.

Why Go and not *just* some *bash aliases* ?

Because I want to learn Go and I switch all the time between mutiple Operating Systems.

Dependencies
============

- `kubectl`

Install
=======

Install it via `curl`:

::

    curl https://raw.githubusercontent.com/svx/ktl/master/install.sh | sh

Install via `go get`:

::

    go get github.com/svx/ktl

Documentation
=============

You can find the HTML version `online <https://ktl.leftxs.org>`_ and the source on `GitHub <https://github.com/svx/ktl/>`_. 

Contributing
============

Pull requests are welcome.

For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

- `Issue Tracker <https://github.com/svx/ktl/issues/>`_
- `Source Code <https://github.com/svx/ktl/>`_

License
=======

`MIT <https://choosealicense.com/licenses/mit/>`_
