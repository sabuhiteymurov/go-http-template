# Contributing to go-net_http-template

First off, thank you for considering contributing to `go-http-template`!

## Code of Conduct

By participating in this project, you are expected to uphold our Code of Conduct:

- Use welcoming and inclusive language
- Focus on what is best for the community

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible.

### Suggestions

If you have a suggestion for the template, I'd love to hear it. Enhancement suggestions are tracked as GitHub issues.

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests

## Development Process

1. Clone the repository

```bash
git clone https://github.com/sabuhiteymurov/go-net_http-template.git
```

2. Create a new branch

```bash
git checkout -b feat/your-feature-name
```

3. Make your changes

4. Push to your fork and submit a pull request

### Go Code Style

- Use `gofmt` to format your code
- Keep functions focused and small
- Use meaningful variable names

### Git Commit Messages

- Follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)

### Documentation

- Keep `README.md` and other documentation up to date
- Document all public functions and types

## Project Structure

Please maintain the existing project structure:

```
.
├── docs/                  # Swagger documentation
├── internal/             # Internal packages
├── db/                   # Database related files
├── routes/               # HTTP route definitions
├── models/               # Public models
└── utils/               # Utility functions
```

## License

By contributing, you agree that your contributions will be licensed under the `MIT License`.
