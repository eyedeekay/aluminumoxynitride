aluminumoxynitride
==================

A Chrome-Wrapping configuration tool for browsing I2P. The very simplest way to
configure Chromium-based browsers to visit I2P sites, without interfering with
your main browser configuration.

This is **NOT FINISHED** software, but it is usable and probably pretty good at
what it's supposed to do.

It works by using the excellent [Lorca](https://github.com/zserge/lorca) library
created the same people who made the Webview bindings for Go as a means to wrap
several variants of the Chrome browser in order to configure them to work with
I2P. It does this using a different working area to your "main" Chrome browser,
which prevents it from interfering with other ways you browse.

It is not actually a Chromium fork. It's just a wrapper around Chromium, in the
same way that `i2p.firefox` is a wrapper around Firefox.

Works pretty much anyplace you can shoehorn a Chromium somewhere on the `$PATH`.