.. sphinx_press_theme documentation master file, created by
   sphinx-quickstart on Mon Jul 30 11:44:38 2018.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

============
Kube Control
============

.. image:: _static/logo-small.png
   :alt: ktl (kube control) logo

About
=====

A wrapper around `kubectl`.

I am lazy, I do not want to type ``kubectl get pods --all-namespaces``.

``ktl`` (Kube Control) is a tiny wrapper written in Go which calls ``kubectl`` in the background.

For example ``ktl pods`` is calling ``kubectl get pods --all-namespaces``.

Why Go and not *just* some *bash aliases* ?

Because I want to learn Go and I switch all the time between multiple Operating Systems.

.. toctree::
   :hidden:
   :maxdepth: 2
   :caption: Overview

   install
   usage

.. toctree::
   :hidden:
   :maxdepth: 2
   :caption: Development

   release
   coding
