# h1spec

h1spec is a conformance testing tool for HTTP/1.1 implementation.
This tool is compliant with `RFC 7230`, `RFC 7231`, `RFC 7232`, `RFC 7233`, `RFC 7234` and `RFC 7235`.

This project is highly inspired by [h2spec](https://github.com/summerwind/h2spec).

## Controbution

### Project Structure

- `h1spec/`: The main package of the project.
- `h1spec/hispec.go`: The entry point of the project.
- `h1spec/cmd/`: The command line interface of h1spec.
- `h1spec/config/`: The configuration of server and test cases.
- `h1spec/spec/`: The test cases of RFC 7230, RFC 7231, RFC 7232, RFC 7233, RFC 7234 and RFC 7235.
- `h1spec/brenchmark/`: The golden standard of the test cases.
- `h1spec/log/`: Some utils for logging.
- `h1spec/experiment/`: Some utils for manual testing.
