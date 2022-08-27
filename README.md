Common Go library for OpenSourceCorp
====================================

<!-- badges: start -->
![Github Actions](https://github.com/opensourcecorp/go-common/actions/workflows/main.yaml/badge.svg)

[![Support OpenSourceCorp on Ko-Fi!](https://img.shields.io/badge/Ko--fi-F16061?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ko-fi.com/ryapric)
<!-- badges: end -->

This repository houses common functionality to be used across Go packages
authored by OpenSourceCorp.

Suppressing log messages in your own program's tests
----------------------------------------------------

If you want or need to suppress log messages from this repo's `logging` package,
set the env var `OSC_IS_TESTING=true` for your test runs. This will suppress all
but `logging.FatalLog()` messages.