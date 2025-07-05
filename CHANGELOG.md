# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](https://semver.org/).

---

## [v1.2.0] - 2025-07-06

### Added
- `(*Err).WithStack()`: Adds the call site stack trace to the error.
- `(*Err).Stack()`: Retrieves a formatted stack trace if available; supports optional custom formatting via `StackWithFormat` `StackOption`.

---

## [v1.1.0]  - 2025-07-05

### Added
- `(*Err).Unwrap()`: Implements the standard `Unwrap() error` interface, allowing integration with `errors.Unwrap()`. Supports nested unwrapping of single and multi-error types of length 2.

---

## [v1.0.0] - 2025-07-03

### Added
- Initial release
